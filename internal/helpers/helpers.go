package helpers

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chelnak/ysmrr"
)

type Converter interface {
	Convert(i string) (string, error)
	IncrementErrors()
	GetTemplate(tmpl string) string
	GetIndexName() string
}

func ValidateArchiveGetIndexName(f string) (string, error) {
	// Check if the input file exists
	r, err := os.Open(f)
	if err != nil {
		return "", errors.New("archive file provided does not exist")
	}
	r.Close()

	// parse the file name to get the index name and date from the filename
	re := regexp.MustCompile(`(\.\/)*(.*)(\d{4}\.\d{2}\.\d{2})`)

	for _, match := range re.FindAllStringSubmatch(filepath.Base(r.Name()), -1) {
		index := match[2] + match[3]
		return index, nil
	}

	return "", errors.New("couldn't get index name from archive file")
}

func CreateWorkingDirExtract(f, dir string) error {
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

	err = Untar(filepath.Join(dir, "v10.3"), r)
	if err != nil {
		return err
	}

	return nil
}

func ArchiveCleanup(index string) error {
	w, err := os.Create(filepath.Join(index, fmt.Sprintf("%s.tar.gz", index)))
	if err != nil {
		return err
	}

	err = Tar(filepath.Join(index, "v11"), w)
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

func CreateTemplates(path string, c Converter) {
	for _, tmpl := range []string{"alias", "mapping", "settings"} {
		filename := fmt.Sprintf("%s-%s.json", c.GetIndexName(), tmpl)

		w, err := os.OpenFile(filepath.Join(path, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer w.Close()

		f := c.GetTemplate(tmpl)

		w.WriteString(f)
	}
}

func ScanAndConvert(src, dst string, c Converter) {
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		j, err := c.Convert(scanner.Text())
		if err != nil {
			c.IncrementErrors()
			continue
		}

		_, err = w.WriteString(j)
		if err != nil {
			c.IncrementErrors()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func ReadJSON(data string, dst any) error {
	stringReader := strings.NewReader(data)

	dec := json.NewDecoder(stringReader)
	// dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		// If there is an error during decoding, start the triage..
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmashalError *json.InvalidUnmarshalError

		switch {
		// Use the errors.As() function to check whether the error has the type
		// *json.SyntaxError. If it does, then return a plain-english error message
		// which includes the location of the problem.
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		// In some circumstances Decode() may also return an io.ErrUnexpectedEOF error
		// for syntax errors in the JSON. So we check for this using errors.Is() and
		// return a generic error message. There is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		// Likewise, catch any *json.UnmarshalTypeError errors. These occur when the
		// JSON value is the wrong type for the target destination. If the error relates
		// to a specific field, then we include that in our error message to make it
		// easier for the client to debug.
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// An io.EOF error will be returned by Decode() if the request body is empty. We
		// check for this with errors.Is() and return a plain-english error message
		// instead.
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// A json.InvalidUnmarshalError error will be returned if we pass something
		// that is not a non-nil pointer to Decode(). We catch this and panic,
		// rather than returning an error to caller.
		case errors.As(err, &invalidUnmashalError):
			panic(err)

		// For anything else, return the error message as-is.
		default:
			return err

		}

	}

	return nil
}

func WriteJSON(data any) (string, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)

	encoder.SetEscapeHTML(false)
	// encoder.SetIndent("", "    ")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func Untar(dst string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {
		// if no more file are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happen)
		case header == nil:
			continue

		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}

	}
}

// Tar takes a source and variable writers and walks 'source' writing each file
// found to the tar writer; the purpose for accepting multiple writers is to allow
// for multiple outputs (for example a file, or md5 hash)
func Tar(src string, writers ...io.Writer) error {

	// ensure the src actually exists before trying to tar it
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("unable to tar files - %v", err.Error())
	}

	mw := io.MultiWriter(writers...)

	gzw := gzip.NewWriter(mw)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	// walk path
	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {

		// return on any error
		if err != nil {
			return err
		}

		// return on non-regular files (thanks to [kumo](https://medium.com/@komuw/just-like-you-did-fbdd7df829d3) for this suggested update)
		if !fi.Mode().IsRegular() {
			return nil
		}

		// create a new dir/file header
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		// update the name to correctly reflect the desired destination when untaring
		header.Name = strings.TrimPrefix(strings.Replace(file, src, "", -1), string(filepath.Separator))

		// write the header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// open files for taring
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		// copy file data into tar writer
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		// manually close here after each file operation; defering would cause each file close
		// to wait until all operations have completed.
		f.Close()

		return nil
	})
}

func CreateSpinGroups() (ysmrr.SpinnerManager, map[string]*ysmrr.Spinner) {
	// Create a new spin manager
	sm := ysmrr.NewSpinnerManager()

	spinners := make(map[string]*ysmrr.Spinner)

	// Setup the spinners
	spinners["setup"] = sm.AddSpinner("Setup/Extraction...")
	spinners["upgrade"] = sm.AddSpinner("Upgrading Data...")
	spinners["cleanup"] = sm.AddSpinner("Archiving/Cleanup...")

	return sm, spinners
}
