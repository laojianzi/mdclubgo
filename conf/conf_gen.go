// Code generated by go-bindata. DO NOT EDIT.
// sources:
// app.ini (1.53kB)

package conf

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _confAppIni = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\x5d\x6f\xda\xe6\x1b\xc6\xcf\x9f\x4f\x71\x87\xff\xd9\x5f\x5a\x92\x2e\x52\xdb\x0d\x71\x40\x80\xa5\x68\x80\x19\xa6\xea\xa2\x68\x42\xbc\x38\x04\xd5\x08\x62\x9b\x4e\x39\x83\x2c\x49\x9b\x09\x37\x6c\x61\x2d\x24\xa4\x24\x5a\x93\xd2\x76\xc5\x6c\x07\xe1\x7d\x7c\x98\x3e\xcf\x63\x73\x94\xaf\x30\xd9\x4f\x40\x44\xed\xa2\x89\x13\x30\xd7\x75\xbf\x5c\xf7\x4f\xfe\x1f\xcc\xcd\xcd\x01\x6d\x75\x8c\x51\x0d\x77\x54\xe3\xa2\x80\x47\x4d\x5a\xee\x1a\xa3\x2a\xee\xbc\xa3\x2f\x9e\xe2\xfe\xa5\x29\x41\x4c\x68\x68\x6d\x52\x6b\xe8\xc7\x1d\xd2\x2b\xeb\x47\x3b\xc6\xd3\x77\xa4\x79\x84\xbb\xfb\xe3\x5d\x55\x1f\x36\x99\x1c\x77\x3e\x18\xa3\x63\xe3\xac\xc8\x2a\x59\x6e\x64\x07\x63\x54\xa5\x15\x4d\x7f\x5f\xd5\xb7\xbb\xa4\xa4\xea\x6f\x5a\x28\xe0\xf4\x7b\xc0\x01\x7e\xb7\x4b\xcc\xc5\x56\x32\xc8\x0e\x6e\x21\x96\x4b\x02\x6d\x9c\x91\xc1\xc1\xd5\xa0\x38\xee\x57\x8c\xe6\x6b\x32\xc8\x93\x92\x86\xdc\x9e\xe5\x87\x2b\xe0\x00\x45\xca\x09\x08\xad\xc9\x82\xf4\x44\x90\x7e\x40\x76\xa0\x15\x8d\x94\x2e\x98\x0a\x36\x14\x25\x2b\xa3\x07\xe1\x70\x90\x8f\x78\x02\xce\x65\x9f\xd9\x61\x3d\x2a\xca\x02\xb2\x83\xf9\x18\x78\xcb\x08\xfa\xd1\x8e\x7e\xfc\x0b\x29\xfd\x41\x6a\x2d\x72\x92\xb7\x2c\x11\xa7\xdb\x1d\x02\x07\x2c\xce\x5b\x9f\x7f\x73\xe8\xef\x35\x72\xf0\x3b\x73\x04\xb9\x50\x18\x1c\x70\x7f\xf1\xfe\x44\xce\x4f\xf4\x86\x56\xc0\xdd\x0b\x16\x8a\xd1\xd6\xc8\xdf\x3b\x57\x83\x22\xee\xef\xce\x4e\x0a\xf4\xe5\x25\xdd\x57\xf1\xe0\x08\x0f\x47\x7a\xb9\x81\x5c\x9e\x50\x38\xf2\x8d\xd7\x1a\x3b\x9e\x93\x95\x4c\x7a\xc1\x12\x2e\xc4\x05\x49\x99\xcf\x0a\x69\xf4\xad\x67\xf5\xb3\x82\xc7\xc2\x96\xf5\xff\xf5\xd0\x46\xbb\x41\xea\x75\x76\x98\x69\x94\xb4\xa2\xc1\xff\x81\xee\xe7\x69\x6d\x9f\xd4\xeb\xa4\xa4\x8e\x7f\x1a\x92\x03\x0d\xf7\xcf\x99\x1e\x39\x5d\x2e\x0f\xcf\x47\x5c\x5c\x20\x1c\xe2\x7c\x11\xa7\xcf\xc7\x3d\x8a\x70\x21\xef\x8a\x37\x00\x0e\x84\xd6\xc4\x4c\xd2\xcc\x5c\xcc\x24\x81\x6d\x46\x3e\x54\xc8\x76\x63\xba\x1f\xeb\xa3\xbf\xed\xe1\x8e\x4a\xf6\xaa\xc8\x0e\xe4\x62\x9b\xbe\xaa\x99\xb7\x3f\xee\x10\xad\x3b\x55\x92\xd6\x1e\xee\x9f\xc3\x8f\x19\xe9\x31\x64\xa3\xca\x06\xe0\x4e\xcf\xa4\x69\xa8\xa2\x10\xc7\x85\x23\x41\x67\xf8\x81\xd5\x32\x11\x55\xa2\xb1\xa8\x2c\x58\xb7\xfe\xad\x45\xd5\x26\xe9\x1d\xd2\xb2\x46\x8b\x05\xb0\x65\x33\xb2\x92\x94\x04\xd9\xf6\x31\x5f\xb0\xa5\xb7\xe4\x4d\xd1\xfa\x26\x6f\x8a\x29\x45\x58\x62\x4f\x65\xf3\x29\xd0\x67\x2f\x8c\xfc\x2e\xd8\x94\x54\x22\x66\x33\xe7\x1a\x1e\xea\x3b\x97\x78\x78\x8a\xbb\x75\x16\x01\x3b\x01\xf8\x57\xf9\xef\x7c\x40\xd4\x03\xa3\xd9\xa4\x27\xe7\xc6\xe8\x15\x7d\x7e\x0e\x61\xaf\x7b\x19\x85\x57\x83\x66\xee\x56\x9f\xd9\x69\x98\x66\x42\x12\xc7\x9b\x48\xdc\xf9\xf2\x9e\x85\xd1\x9d\xaf\x97\x96\x16\xef\xce\xaa\x6f\xe2\x9f\x4e\xc4\xc5\x5c\x2c\x99\xf9\xb4\x9e\x5e\x6e\xd0\x67\x6d\x52\x52\xd1\x43\xde\x13\xba\x55\x4a\xb4\x3d\xfd\xb4\x80\x82\x4e\x9e\x7f\xc4\x85\xdc\xe0\x40\x76\xc0\xfd\x5d\xbd\xdc\xc0\xbd\xe7\x33\x29\x5d\x0d\x8a\x93\xe4\x12\x29\x39\x1a\x13\x05\x2b\x22\x49\xd8\xcc\xa5\x24\xc1\x06\xe4\xd7\x22\xd8\x9e\x08\x52\x6a\x7d\xeb\x8b\xf5\x9c\x28\xda\x10\xcf\xfb\x22\x7e\xce\x6d\x4e\x7a\xed\xb8\x59\x7a\x92\x34\x30\x10\xd8\x5c\x0c\x8d\xab\x41\xd1\xd0\xda\xfa\x59\x13\x8f\x4e\x58\xb6\x7a\xff\x64\x8a\xc0\xc7\xfc\x36\x62\x37\x06\xf3\xc2\x0b\x93\xe5\xe6\x13\xb1\x4f\xf7\xa3\x7f\x9e\xd2\x5a\x9e\xbc\x7e\xc3\xea\xe0\xce\xcf\xe3\x6a\x09\xf9\x9d\xdf\x47\xb8\xa0\x27\x60\xe2\x1a\xe0\xc1\x01\x4b\x8b\xb7\x58\xf5\xb7\xbd\xf1\xcb\xbf\x66\xac\x5e\xb7\xcf\x33\x6b\x45\x6b\xf1\x68\x7c\xc3\xe2\x4c\x1f\x1c\x9a\xfb\x58\x2b\x91\x6a\x63\x92\x99\x24\x24\x52\xd7\xa8\x09\x69\x4b\x9c\xb0\x7e\x65\x13\x19\xdb\x04\x0e\x4b\x33\x2d\x71\x3b\x19\x77\x97\xee\x7d\x65\xa2\xc8\xfa\xd4\xeb\xec\x05\xc1\x4a\xc0\xf5\x1b\xc1\xc4\x84\x0f\x3a\x5d\x66\xed\xcf\xac\x77\x13\x12\xc6\xd4\x7f\xe1\xe3\x9f\x00\x00\x00\xff\xff\xa0\x0e\x8a\xb6\xfa\x05\x00\x00"

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

	info := bindataFileInfo{name: "conf/app.ini", size: 1530, mode: os.FileMode(0644), modTime: time.Unix(1604568256, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x59, 0x72, 0x98, 0xf8, 0xeb, 0x17, 0x46, 0xee, 0x60, 0xaf, 0xed, 0x95, 0x6a, 0x64, 0xda, 0xe3, 0xa, 0x9b, 0xa6, 0xe7, 0x3e, 0xc9, 0x51, 0x87, 0x3, 0x2a, 0x64, 0x73, 0xfe, 0xdd, 0xa1, 0xa4}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"conf": {nil, map[string]*bintree{
		"app.ini": {confAppIni, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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