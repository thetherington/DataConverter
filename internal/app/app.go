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
	converter ConverterInterface
}

type ConverterInterface interface {
	Convert(doc string) (string, error)
	IncrementErrors()
	GetTemplate(name string) string
	GetSourceIndexName() string
	GetDestinationIndexName() string
	SendToCountChan(count int)
	GetCountChan() chan int
	ReturnErrors() int
}

func New(c ConverterInterface) *App {
	return &App{
		converter: c,
	}
}

func (app *App) Run(arg string) error {
	sm, spinners := helpers.CreateSpinGroups()

	// start the spinners
	sm.Start()

	srcIndex := app.converter.GetSourceIndexName()
	dstIndex := app.converter.GetDestinationIndexName()

	err := app.CreateWorkingDirExtract(arg, dstIndex)
	if err != nil {
		spinners["setup"].Error()
		sm.Stop()
		return err
	}

	spinners["setup"].UpdateMessage("Setup/Extraction...Complete")
	spinners["setup"].Complete()

	app.CreateTemplates(filepath.Join(dstIndex, "v11"))

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
		filepath.Join(dstIndex, "v10.3", fmt.Sprintf("%s-data.json", srcIndex)),
		filepath.Join(dstIndex, "v11", fmt.Sprintf("%s-data.json", dstIndex)),
	)

	done <- true

	if app.converter.ReturnErrors() == 0 {
		spinners["upgrade"].UpdateMessage("Upgrading Data...Complete")
	} else {
		spinners["upgrade"].UpdateMessagef("Upgrading Data...Errors: %d", app.converter.ReturnErrors())
		spinners["upgrade"].Error()
	}

	spinners["upgrade"].Complete()

	err = app.ArchiveCleanup(dstIndex)
	if err != nil {
		spinners["cleanup"].Error()
		sm.Stop()
		return err
	}

	spinners["cleanup"].UpdateMessage("Archiving/Cleanup...Complete")
	spinners["cleanup"].Complete()

	sm.Stop()

	fmt.Println()
	fmt.Println("New Archive: ", filepath.Join(dstIndex, fmt.Sprintf("%s.tar.gz", dstIndex)))

	return nil
}

func (app *App) CreateWorkingDirExtract(f, dir string) error {
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

func (app *App) ArchiveCleanup(index string) error {
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

func (app *App) CreateTemplates(path string) {
	for _, tmpl := range []string{"alias", "mapping", "settings"} {
		filename := fmt.Sprintf("%s-%s.json", app.converter.GetDestinationIndexName(), tmpl)

		w, err := os.OpenFile(filepath.Join(path, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer w.Close()

		f := app.converter.GetTemplate(tmpl)

		w.WriteString(f)
	}
}

func (app *App) ScanAndConvert(src, dst string) {
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
		j, err := app.converter.Convert(scanner.Text())
		if err != nil {
			app.converter.IncrementErrors()
			continue
		}

		_, err = w.WriteString(j)
		if err != nil {
			app.converter.IncrementErrors()
		}

		count++
		app.converter.SendToCountChan(count)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
