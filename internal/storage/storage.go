package storage

import (
	"fmt"
	"io"
	"strings"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/internal/storage/local"
	"github.com/laojianzi/mdclubgo/log"
)

var instance Interface

// Interface storage methods
type Interface interface {
	Read(path string, thumbs map[string][2]int) map[string]string
	Write(path string, reader io.Reader, thumbs map[string][2]int) error
	Delete(path string, thumbs map[string][2]int) error
	io.Closer
}

// Init for storage
func Init() {
	typ := strings.ToLower(strings.TrimSpace(conf.Storage.Type))

	switch typ {
	case Local:
		instance = local.New()
	default:
		log.Fatal(fmt.Errorf("unrecognized dialect: %s", typ).Error())
	}
}

// Read get images form storage
func Read(path string, thumbs map[string][2]int) map[string]string {
	return instance.Read(path, thumbs)
}

// Write save image to storage
func Write(path string, reader io.Reader, thumbs map[string][2]int) error {
	return instance.Write(path, reader, thumbs)
}

// Delete remove image form storage
func Delete(path string, thumbs map[string][2]int) error {
	return instance.Delete(path, thumbs)
}

func Close() {
	_ = instance.Close()
}
