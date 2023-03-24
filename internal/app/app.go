package app

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chelnak/ysmrr"
	"github.com/thetherington/DataConverter/internal/helpers"
)

type App struct {
	index     string
	converter ConverterInterface
}

type ConverterInterface interface {
	Convert(doc string) (string, error)
	IncrementErrors()
	GetTemplate(name string) string
	GetIndexName() string
	SendToCountChan(count int)
	GetCountChan() chan int
	ReturnErrors() int
}

func New(c ConverterInterface, i string) *App {
	return &App{
		index:     i,
		converter: c,
	}
}

func (app *App) Run(arg string) error {
	sm, spinners := helpers.CreateSpinGroups()

	// start the spinners
	sm.Start()

	err := app.CreateWorkingDirExtract(arg, app.index)
	if err != nil {
		spinners["setup"].Error()
		sm.Stop()
		return err
	}

	spinners["setup"].UpdateMessage("Setup/Extraction...Complete")
	spinners["setup"].Complete()

	app.CreateTemplates(filepath.Join(app.index, "v11"))

	done := make(chan bool)

	go func(spinner *ysmrr.Spinner, ch chan int, done chan bool) {
		var c int
		for {
			select {
			case count := <-ch:
				if count > c+10000 {
					spinner.UpdateMessagef("Upgrading Data...%d", count)
					c = count
				}
			case <-done:
				close(ch)
				close(done)
				return
			}
		}
	}(spinners["upgrade"], app.converter.GetCountChan(), done)

	app.ScanAndConvert(
		filepath.Join(app.index, "v10.3", fmt.Sprintf("%s-data.json", app.index)),
		filepath.Join(app.index, "v11", fmt.Sprintf("%s-data.json", app.index)),
	)

	done <- true

	if app.converter.ReturnErrors() == 0 {
		spinners["upgrade"].UpdateMessage("Upgrading Data...Complete")
	} else {
		spinners["upgrade"].UpdateMessagef("Upgrading Data...Errors: %d", app.converter.ReturnErrors())
		spinners["upgrade"].Error()
	}

	spinners["upgrade"].Complete()

	err = app.ArchiveCleanup(app.index)
	if err != nil {
		spinners["cleanup"].Error()
		sm.Stop()
		return err
	}

	spinners["cleanup"].UpdateMessage("Archiving/Cleanup...Complete")
	spinners["cleanup"].Complete()

	sm.Stop()

	fmt.Println()
	fmt.Println("New Archive: ", filepath.Join(app.index, fmt.Sprintf("%s.tar.gz", app.index)))

	return nil
}

func (a *App) CreateWorkingDirExtract(f, dir string) error {
	err := os.MkdirAll(filepath.Join(dir, "v11"), 0755)
	if err != nil {
		return errors.New("failed to create working directory")
	}

	err = os.MkdirAll(filepath.Join(dir, "v10.3"), 0755)
	if err != nil {
		return errors.New("failed to create working directory")
	}

	r, err := os.Open(f)
	if err != nil {
		return errors.New("archive file provided does not exist")
	}

	err = helpers.Untar(filepath.Join(dir, "v10.3"), r)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) ArchiveCleanup(index string) error {
	w, err := os.Create(filepath.Join(index, fmt.Sprintf("%s.tar.gz", index)))
	if err != nil {
		return err
	}

	err = helpers.Tar(filepath.Join(index, "v11"), w)
	if err != nil {
		return err
	}

	err = os.RemoveAll(filepath.Join(index, "v11"))
	if err != nil {
		return err
	}

	err = os.RemoveAll(filepath.Join(index, "v10.3"))
	if err != nil {
		return err
	}

	return nil
}

func (a *App) CreateTemplates(path string) {
	for _, tmpl := range []string{"alias", "mapping", "settings"} {
		filename := fmt.Sprintf("%s-%s.json", a.converter.GetIndexName(), tmpl)

		w, err := os.OpenFile(filepath.Join(path, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer w.Close()

		f := a.converter.GetTemplate(tmpl)

		w.WriteString(f)
	}
}

func (a *App) ScanAndConvert(src, dst string) {
	f, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	w, err := os.OpenFile(dst, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer w.Close()

	var count int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		j, err := a.converter.Convert(scanner.Text())
		if err != nil {
			a.converter.IncrementErrors()
			continue
		}

		_, err = w.WriteString(j)
		if err != nil {
			a.converter.IncrementErrors()
		}

		count++
		a.converter.SendToCountChan(count)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
