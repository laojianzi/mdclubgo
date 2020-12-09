package local

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/spf13/afero"

	"github.com/laojianzi/mdclubgo/internal/storage/util"
)

// Local storage
type Local struct {
	pathPrefix string
	fs         *afero.Afero
}

func (l Local) applyPathPrefix(path string) string {
	return fmt.Sprintf("%s%s", l.pathPrefix, strings.Trim(path, "\\/"))
}

// New return a *Local
func New() *Local {
	return &Local{
		pathPrefix: "../../public/upload/",
		fs:         &afero.Afero{Fs: afero.NewOsFs()},
	}
}

// Read data from local
func (l *Local) Read(path string, thumbs map[string][2]int) map[string]string {
	data := make(map[string]string)
	data["original"] = path

	for size := range thumbs {
		data[size] = util.GetThumbLocation(path, size)
	}

	return data
}

// Write data to local
func (l *Local) Write(path string, reader io.Reader, thumbs map[string][2]int) error {
	location := l.applyPathPrefix(path)
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("ioutil read all: %w", err)
	}

	if err := l.fs.WriteReader(location, ioutil.NopCloser(bytes.NewBuffer(buf))); err != nil {
		return fmt.Errorf("afero write reader: %w", err)
	}

	return util.CROP(ioutil.NopCloser(bytes.NewBuffer(buf)), location, thumbs, func(path string,
		reader io.Reader) error {
		return l.fs.WriteReader(path, reader)
	})
}

// Delete data in local
func (l *Local) Delete(path string, thumbs map[string][2]int) error {
	location := l.applyPathPrefix(path)
	err := l.fs.Remove(location)
	if err != nil {
		return err
	}

	for size := range thumbs {
		err = l.fs.Remove(util.GetThumbLocation(location, size))
		if err != nil {
			return err
		}
	}

	return nil
}

// Close file system object
func (l *Local) Close() error {
	l.fs = nil
	return nil
}
