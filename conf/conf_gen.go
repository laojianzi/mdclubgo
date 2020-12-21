// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package conf generated by go-bindata.// sources:
// app.ini
package conf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _confAppIni = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x51\x73\x12\x57\x14\x7e\xbf\xbf\xe2\x48\xdf\x3a\x36\x89\xcd\x54\xad\x19\x66\x4a\x00\x95\x29\x09\x94\x5d\xc7\x3a\x19\x87\x59\x60\x43\x76\x5c\xdc\x75\x77\xb1\x93\x37\x92\x26\xd1\x58\x56\xb0\xa6\x1a\x12\x22\x66\x6a\x22\x6a\x65\x63\x1f\x12\x12\x48\xf9\x31\xee\xdd\xbb\x3c\xe5\x2f\x74\xee\x5e\x40\xac\xa9\xd3\xe9\xe4\xe5\xe6\xee\x77\xbe\x73\xce\x77\xbe\x7b\xf8\x02\xce\x9c\x39\x03\xce\x5e\xd3\xed\x54\xed\xa6\xe9\xee\x2e\xd8\x9d\x86\xb3\x76\xe8\x76\x2a\x76\xf3\xb5\xf3\xe4\x9e\xdd\xda\xa7\x10\xc4\x80\xae\x75\x80\xab\x75\xb2\xd9\xc4\x47\x6b\x64\x63\xc9\xbd\xf7\x1a\x37\x36\xec\xc3\xd5\xee\xb2\x49\x8e\x1b\x0c\x6e\x37\xdf\xba\x9d\x4d\x77\xbb\xc8\x98\xbc\x68\x34\x01\x6e\xa7\xe2\xac\x5b\xe4\x4d\x85\x2c\x1e\xe2\xb2\x49\x5e\xee\xa1\xe9\xc0\x54\x18\xfc\x30\x15\x0a\xca\xf9\xd4\x15\x05\x4d\x40\x48\x4c\xe5\xb3\xe0\xd4\xb7\x71\xbb\x74\xd2\x2e\x76\x5b\xeb\x6e\xe3\x05\x6e\x17\x70\xd9\x42\xa1\xf0\xe4\xb5\x2b\xe0\x07\x43\xcb\x8b\x68\x02\x22\x71\xc0\xd5\x3d\xbc\x55\x20\x1b\x4b\xf8\xed\x3a\x5e\xac\xb3\xf4\xee\x81\x85\xff\x5a\x42\x91\x78\x68\x32\x19\x0f\xf0\x57\xc1\x0f\xe9\xbc\x6e\x28\xb9\xd1\xb4\x72\x7b\x76\x34\x97\x49\xcb\xf9\x54\x56\x19\x91\xd4\x4c\x0a\xa1\x19\x5d\xd4\xee\x8a\xda\x4d\x34\x01\xce\xba\x85\xcb\xbb\x2c\x19\xcc\x19\x86\xaa\xa3\xab\x3c\x1f\xe7\x92\xe1\xe9\xc0\x64\x94\x16\x3a\x2b\xc8\x3a\x4d\x4d\xaf\x81\xf3\x02\x81\x6c\x2c\x91\xcd\x47\xb8\xfc\x07\x2b\xc6\x0b\x49\x06\x42\xa1\x04\xf8\x61\x6c\xc4\xfb\xfb\xb7\x08\xf2\xc6\xc2\xa5\xdf\x59\x44\x3c\x96\xe0\xc1\x0f\x17\xc7\x2e\xf6\xe1\x5c\x1f\xef\x5a\x0b\xf6\xe1\xee\x70\x73\x27\xed\xa2\xdd\x5a\x1e\xae\x14\x9c\xa7\xfb\xce\xaa\x69\xb7\x37\xec\xe3\x0e\x59\xab\xa3\x60\x38\xc1\x27\x2f\x47\xbc\xb2\x7b\xed\x7b\xc0\xd1\xb4\xa8\x19\x23\xaa\x98\x43\xdf\x87\x6f\x9c\x0a\xb8\x25\xce\x7b\xdf\x7b\x45\xbb\x07\x75\x5c\xab\xb1\xf9\x0e\x26\xe2\xac\x5b\xf0\x25\x38\xab\x05\xa7\xba\x8a\x6b\x35\x5c\x36\xbb\x3f\x1f\xe3\x92\x65\xb7\x76\x18\x1e\x05\x82\xc1\x30\xc7\x25\x83\xb1\x69\x3e\x11\x8b\x26\x03\xd1\x68\xec\x7a\x32\x96\x88\x5c\x89\x4c\x83\x1f\x4d\xc0\x75\x31\x05\xba\x64\x88\xa0\x1b\x82\x21\xa5\x21\xaf\xc9\xc0\xb8\xc9\xab\x23\x4a\x3f\x8a\xb8\x08\x1f\x4e\x72\x7c\x80\x8f\x04\x93\xd7\x12\x51\xf0\x23\x34\x23\x2b\x59\x3a\x2a\x59\xc9\x02\x13\x84\x4d\x7e\x20\xcb\x80\xc2\x6e\x9a\x78\xa5\x82\x26\x00\xef\x2e\x3a\xcf\xaa\xd4\x79\x9b\x4d\x6c\x1d\x0e\x90\x78\x6f\xc5\x6e\xed\xc0\x4f\x8a\x76\x0b\x54\xc1\x98\x03\xbb\x79\x44\xbd\x7c\x6c\xa2\x44\x2c\xc6\xf7\xac\x83\xd0\x4c\x46\x30\x84\x94\xa0\x8b\x9e\x45\x7e\xdb\x73\xcc\x06\x3e\x7a\xec\xac\x59\x4e\x71\x01\x7c\xaa\xa2\x1b\x59\x4d\xd4\x7d\xef\x0b\x0b\xbe\xdc\xbc\x7e\x47\xf6\x4e\xfa\x1d\x59\x32\xc4\x71\x76\xab\xd3\x5b\x70\xee\x3f\x71\x0b\xcb\xe0\x33\xa4\x4c\xca\x47\xeb\x3a\x7e\x4c\x96\xf6\xed\xe3\xe7\xf6\x61\x8d\x29\xc7\x26\x07\x53\x37\xb8\x1f\xa2\x80\xcd\x92\xdb\x68\x38\x5b\x3b\x6e\xe7\x99\xf3\x70\x07\xf8\x48\x68\x12\xf1\x37\xe2\x74\x5c\x5e\x9e\xe1\x6a\x18\xa6\x6f\xc0\x18\x47\x9d\x74\xee\xeb\x0b\x9e\xfb\xce\x5d\x1a\x1f\x1f\x3b\x3f\x8c\xfe\xf8\xf1\xf5\xdf\xc3\xa7\x7c\x64\xad\xee\xdc\x3f\xc0\x65\x13\x5d\xe3\xc2\x89\xcf\x42\xb1\xb5\x42\x9e\x2f\xa0\x78\x80\xe3\xae\xc7\x12\x21\x6f\xc2\x76\x6b\x99\xac\xd5\xed\xa3\x87\x43\x2a\x9d\xb4\x8b\x7d\xe5\x32\x92\x2e\xa4\x64\xd1\x93\x48\x13\xef\xe4\x25\x4d\xf4\x01\xfe\xb5\x08\xbe\xbb\xa2\x26\xcd\xce\x7f\x35\x9b\x97\x65\x1f\xe2\xb8\x68\x72\x2a\x16\xa2\x95\xf6\x22\x3e\xa6\xee\x2b\x0d\xbd\x15\xe0\xd5\xc5\xac\x71\xd2\x2e\xba\xd6\x01\xd9\x6e\xd8\x9d\x2d\xa6\x2d\x69\x6d\x0d\x2c\xf0\xbe\xb0\x88\x7a\xeb\x81\x4e\xf8\xc3\x5e\xc8\xa4\x3e\xed\xcf\x79\xf7\xdc\xa9\x16\xf0\x8b\x97\x8c\xc7\x6e\x3e\xe8\x56\xca\x68\x2a\xf0\x63\x32\x16\x0f\x4f\x53\x97\x4f\x73\xe0\x87\xf1\xb1\xcf\x84\x92\x57\x47\xdd\xa7\x7f\x0e\x85\x46\x42\xd1\xf0\x70\x28\x9a\x49\x0b\xe9\x39\xcf\x67\xa4\xfd\x98\xf6\xe3\xb5\x84\x2b\xf5\xbe\x66\x9a\x98\x91\x7a\x56\x13\x73\x1e\x38\xe3\xfd\xa7\x66\x14\x5f\xdf\x1c\x1e\x66\x40\xf1\x79\x67\x9c\x1f\xbf\xf0\x2d\xb5\x22\xcb\x53\xab\xb1\xbd\xc2\x28\xa0\xb7\x48\xa8\x4d\xb8\x78\x20\x48\xb9\x4f\x69\xef\x63\x93\x30\x4f\xfd\x17\x7f\xa0\x19\x31\x27\x48\x32\x6d\xd6\x3b\x00\x2e\x3d\xea\x16\x16\x9c\xaa\x89\x1f\x6c\xe3\x4a\x9d\xbc\x6b\xe1\x67\xbf\xd0\x77\x5a\xb2\xba\x85\xd5\x19\x3d\x67\xa8\x67\x41\x55\xd4\xf1\xb3\x20\xe5\x04\xf5\xe6\x49\xbb\x48\x36\x1b\x78\xd5\xa4\x15\xd3\xaf\x80\x4b\x16\x2d\xb8\xa7\x03\xbd\x1a\x90\x0f\x68\x4f\xd3\x83\x22\x47\x28\x2c\x9b\xbf\x3d\xa2\x68\xd9\x4b\xdf\x5c\xbc\xf0\xcf\xb2\x9e\xee\x93\x8d\x25\xb8\x9c\x88\x4d\xf5\x7e\x74\x90\x77\xf6\xc3\x6d\x45\x13\x55\x79\xfe\xbb\x81\x7d\x64\x25\x2d\xc8\x73\x8a\x6e\x7c\x9a\x9c\x54\x5a\xdd\xca\x0a\x53\x6c\x48\xae\xff\xc1\x71\x9a\x9e\xba\xa1\x68\x42\x96\xad\xa9\xa1\xd5\x38\xd0\xb2\x2f\x8c\x47\xfe\x01\xcf\x92\xdd\xec\xbf\x05\x35\x9f\x92\xa5\xf4\x68\x5e\x95\x15\x21\x33\x8a\xfe\x0e\x00\x00\xff\xff\xd9\x9b\x11\x41\x1f\x08\x00\x00"

func confAppIniBytes() ([]byte, error) {
	return bindataRead(
		_confAppIni,
		"conf/app.ini",
	)
}

func confAppIni() (*asset, error) {
	bytes, err := confAppIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.ini", size: 2079, mode: os.FileMode(420), modTime: time.Unix(1608538558, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"conf/app.ini": confAppIni,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"conf": &bintree{nil, map[string]*bintree{
		"app.ini": &bintree{confAppIni, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
