package storage

import (
	"net/http"
	"testing"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/internal/storage/local"
)

func TestInit(t *testing.T) {
	t.Run("uses local new a storage", func(t *testing.T) {
		conf.Storage.Type = Local
		Init()

		_, ok := instance.(*local.Local)
		if !ok {
			t.Fatalf("instance type want: *local.Local; got: %T", instance)
		}
	})
}

func TestLocal(t *testing.T) {
	conf.Storage.Type = Local
	Init()

	path := "tmp/storage_test_local.png"
	resp, err := http.Get("https://raw.githubusercontent.com/laojianzi/mdavatar/main/mdavatar.png")
	if err != nil {
		t.Fatal(err)
	}

	thumbs := map[string][2]int{
		"small":  {64, 64},
		"middle": {128, 128},
		"large":  {256, 256},
	}

	err = Write(path, resp.Body, thumbs)
	_ = resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	result := Read(path, thumbs)
	check := func(key string, wantV string) {
		v, ok := result[key]
		if !ok {
			t.Fatalf("Read() = %x; want include key '%s'", v, key)
		}

		if v != wantV {
			t.Fatalf("%s want: %s, got: %s", key, wantV, v)
		}
	}

	tests := map[string]string{
		"original": path,
		"small":    "tmp/storage_test_local_small.png",
		"middle":   "tmp/storage_test_local_middle.png",
		"large":    "tmp/storage_test_local_large.png",
	}

	for key, wantV := range tests {
		check(key, wantV)
	}

	if err = Delete(path, thumbs); err != nil {
		t.Fatalf("Write() error: %s", err.Error())
	}

	Close()
}
