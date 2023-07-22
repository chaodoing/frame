// Code generated by go-bindata. DO NOT EDIT.
// sources:
// vscode/index.html
// service/systemd.service
package assets

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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _vscodeIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xdd\x6e\xe3\x36\x13\xbd\xb6\x9f\x82\x9f\xbe\x02\x92\x0b\x8b\xb2\xb7\x8b\xed\x42\xb6\x8c\xb6\xdb\x14\xcd\x22\xbb\x29\x9a\x04\x6d\xd1\x16\x08\x2d\x8e\x25\x26\xd4\x50\x4b\x52\x72\x1c\xc3\xef\x5e\x50\x92\x63\x2b\x3f\x68\x81\xe8\x86\xd2\xcc\xf0\xf0\x70\xe6\xcc\xd8\xf3\xff\xfd\x78\xfe\xe1\xf2\x8f\x5f\x4e\x48\x6e\x0b\xb9\x18\xce\xdd\x42\x24\xc3\x2c\xf1\xee\xf3\xf0\x67\x86\xc6\x73\x56\x60\x7c\x31\x1c\xcc\x0b\xb0\x8c\xa4\x39\xd3\x06\x6c\xe2\x5d\x5d\xfe\x14\xbe\xf7\x9c\xdd\x0a\x2b\x61\xb1\xdd\xd2\x4b\xf7\xb2\xdb\xcd\xa3\xd6\xb2\xdf\x92\x5b\x5b\x86\xf0\xa5\x12\x75\xe2\xfd\x1e\x5e\x7d\x1f\x7e\x50\x45\xc9\xac\x58\x4a\xf0\x48\xaa\xd0\x02\xda\xc4\x3b\x3d\x49\x80\x67\xe0\x3d\x6c\x43\x56\x40\xe2\xd5\x02\xd6\xa5\xd2\xf6\x28\x72\x2d\xb8\xcd\x13\x0e\xb5\x48\x21\x6c\x3e\xc6\x44\xa0\xb0\x82\xc9\xd0\xa4\x4c\x42\x32\xa5\x93\x06\x47\x0a\xbc\x25\x1a\x64\xe2\x19\xbb\x91\x60\x72\x00\xeb\x91\x5c\xc3\x2a\xf1\x1c\x2b\x13\x47\x51\x85\xe5\x6d\x46\x53\x55\x44\x92\x6d\x2a\xf1\xdd\x1b\xfa\x8e\xbe\x8f\xb8\x30\x36\x4a\x8d\x69\x8d\x34\x35\xa6\x01\x34\xa9\x16\xa5\x5d\x0c\x49\xf7\xd4\x4c\x13\xc3\x8a\x52\xc2\x47\xa3\x90\x24\xe4\xe3\xc5\xf9\x67\x5a\xba\x14\x05\xdb\x2d\x75\xc6\xdd\x6e\x34\xeb\xc5\x73\x66\xd9\x3e\xd2\x58\x2d\x30\x13\xab\x4d\x70\x40\x19\x13\xac\xa4\x1c\x13\xef\x2f\xeb\x8d\x66\xc3\xc1\x3c\xda\x9f\x3a\x98\x37\xd7\x38\x1c\xbf\x54\x7c\x43\xb6\x0f\x9f\xee\x29\x19\xe7\x02\xb3\x98\x4c\x66\x3d\x7b\xc1\x74\x26\xf0\x89\x39\x07\x91\xe5\x36\x26\xd3\xc9\xa4\xce\xfb\xae\x26\xb3\xad\x67\xdd\xf7\xa8\x1a\xf4\x4a\xaa\x75\x4c\x72\xc1\x39\xe0\xc1\xbb\x7b\x78\xfb\x7f\x93\xb8\x30\x55\x1c\x1e\x31\x7c\x19\xf7\x05\x32\xbb\x26\x07\xed\xcd\xe7\x51\xab\xc6\xb9\xbb\xfa\x62\x38\xe7\xa2\x26\x82\x27\xde\xe1\x34\x6f\x31\x8f\xb8\xa8\x5d\x68\x17\xd3\xa6\x8f\x18\x9d\x1e\xaa\x9e\x72\xa4\x4b\xa5\xac\x5b\x11\x6c\xc4\x6e\xd8\x5d\x24\xc5\xd2\x44\xa9\x14\xe5\x52\x31\xcd\xe9\x8d\x89\xde\xd0\x09\x9d\x4e\x8f\x6c\x85\x40\x7a\x63\xdc\x19\xfb\xa2\x3c\x0b\xff\x54\x54\xdf\xd2\x77\xad\xa8\x5a\x41\xfd\x07\x8c\x97\x29\x16\x0a\x59\xaa\x42\xe0\xc2\x2a\x1d\x4d\xe8\x37\x6f\xe9\x34\x2a\x04\x46\xb5\x89\xa4\x62\x1c\xf4\xcb\x3c\x5b\xf1\x68\xd7\x8e\x1a\x68\xaa\x70\x25\xb2\xe0\x50\xa0\x92\xd9\xdc\xc4\x64\x4b\x6a\x13\x13\xff\x75\x54\x7c\xb2\x1b\x3f\x00\xfb\xb5\x89\x50\x1a\x3f\x26\x5b\x56\x33\x21\xd9\x52\xc2\x19\xc3\xac\x62\x19\xb8\x03\xfd\xaf\xfd\xd8\xbf\xcf\xc3\x14\xfd\x5d\xb7\x6d\xdf\x38\x6b\x81\x5c\xad\xe9\xa7\xe6\xa8\x13\xac\x85\x56\x58\x00\x5a\x92\x1c\x29\x2b\x03\xfb\x9b\xd2\xb7\xa0\xaf\xb4\x8c\xc9\xaa\xc2\xd4\x0a\x85\x24\x58\x37\xc6\x53\x3e\x26\x92\x2d\x41\x8e\x1e\x89\x51\x83\xad\x34\x92\x6b\xd7\x94\xb1\x85\x3b\x1b\xdd\xb0\x9a\xb5\xa9\x9a\xed\x67\x5d\x65\x57\xe1\xfb\xf1\x57\x5b\x40\x27\xb0\xab\x5f\x4f\xdd\x00\x53\x08\x68\x83\xeb\x1e\x9a\x01\xb9\xfa\x17\x9e\x5d\xe3\x32\x03\x0d\xd1\x57\xa4\xd8\xef\x61\xee\xfa\xad\x24\x0a\x37\x34\x2f\x9a\x7b\x98\xe0\x95\x85\x74\x6c\xa3\x36\x91\xdd\xf2\x89\x09\xec\x34\xe6\x8f\x66\xd7\xa3\xdd\xf5\x51\xbb\x76\xd5\x9b\x1d\x2b\x2d\xf8\xd3\x09\xa0\x43\x6f\x17\x5a\x30\x81\xfe\xdf\xe3\xa3\x62\x1d\x57\xa7\xab\x7a\xcb\x8b\x76\x5b\x52\x0d\xcc\x42\xc0\x55\x5a\xb9\xcc\xd2\x2f\x15\xe8\xcd\x05\x48\x48\xad\xd2\x81\x77\x34\x74\xbc\xd1\xf8\x51\xd6\x6b\x26\x2b\x88\x9b\xf1\x3b\xee\x39\x64\x27\xc3\x98\xf8\x37\x46\xa1\xdf\xf7\xda\x1c\x0a\x88\x89\xfb\x79\x73\x6f\x4e\x9c\x83\xc1\x60\xa5\xd0\x5e\x88\x7b\xb7\x67\xfa\xb6\xbc\x7b\xb4\x47\x0a\x84\xcf\x55\xb1\x04\xed\xba\xe8\x09\xa4\x56\x15\x72\xe0\x2d\x6f\xa1\x30\x26\x2b\x26\x0d\xf4\x83\x4c\xaa\x95\x94\x3f\xc0\x46\x21\x3f\x63\xc6\x9e\x09\x84\x67\x03\x35\x30\x7e\x8e\x72\xf3\xac\xd3\xb2\x65\x4b\xf3\x6d\xbf\xa7\x1e\xf5\x16\xe3\xfc\xa4\x06\xb4\x67\xc2\x58\x40\xd0\x81\xaf\xc1\x88\x7b\xf0\xc7\x04\xea\x46\xc2\x8b\xa7\xa5\x91\x2a\x65\x8e\x3e\xd5\xe0\xa6\x4e\x30\xea\x90\x87\x47\x43\x27\x6a\xfe\x59\xfc\x13\x00\x00\xff\xff\x5a\x22\x43\x46\x69\x08\x00\x00")

func vscodeIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_vscodeIndexHtml,
		"vscode/index.html",
	)
}

func vscodeIndexHtml() (*asset, error) {
	bytes, err := vscodeIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vscode/index.html", size: 2153, mode: os.FileMode(420), modTime: time.Unix(1687670566, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serviceSystemdService = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\xcf\x8e\xd3\x30\x10\x87\xef\x7e\x8a\x91\x72\xa6\x5a\x09\x95\x5b\x0e\x40\x03\x54\xf4\x0f\x6a\xbb\xdd\x43\xd5\x83\x49\x26\xed\x68\x1d\x3b\xcc\x4c\xd2\xb5\x56\x7e\x77\xe4\x6e\x2b\x01\x07\x4e\x9e\xf9\x3e\x8f\x25\xcf\xef\xf0\xe8\x49\x8f\x66\x86\x52\x33\xf5\x4a\xc1\x97\xaf\xaf\x93\x3f\xda\x94\xcc\xc7\x56\x91\x4b\x89\xe2\xc2\x69\xa2\x96\x4f\xa8\x37\xe6\x51\x2f\x81\x9f\xff\x86\x5d\x94\x5f\xae\x99\x08\xf2\x48\x35\x02\x63\x43\x72\xef\x8c\x39\x6c\xdf\xaa\xa3\x29\x60\x19\x1a\x6a\x23\xe8\x19\x05\x41\x2f\x01\x46\xeb\x06\x14\xb0\xbe\x81\xc1\xd7\xa1\xeb\xd0\x6b\xd6\x1d\x50\x0b\x31\x0c\x70\xb6\x63\x7e\xb1\x0f\x02\x17\xd2\x33\xb8\xa0\x02\xa1\x85\x96\xdc\x6d\xee\x84\x0a\xd6\xc3\xb7\xdd\xee\x07\x20\x73\x60\x98\x3e\x3c\xc0\x4f\xac\xed\x20\x98\xaf\xea\xd9\xaa\x29\x8a\xc2\x14\xb0\xa0\x8e\x74\x59\x2d\x17\xeb\xcf\xdf\x4b\xf2\x2d\x79\xd2\x78\xe7\xab\xf5\x97\xf9\xa2\x2a\x3f\x4c\xa7\xef\xa7\x66\x17\x7b\x2c\x85\xba\xde\xa1\x79\x14\xe4\xbc\xa4\x7c\x7a\xdb\x61\x4a\xe6\x2b\x87\xa1\xcf\xec\x5a\xa4\x64\x9e\x02\x3f\x93\x3f\xcd\x88\xb1\xd6\xc0\xf1\xba\xd4\x7b\x93\x92\xa9\xfc\x48\x1c\x7c\xfe\x5f\x39\x9b\x6f\xfe\xa7\xab\xd5\x3e\xeb\xca\x8f\xff\x88\x7d\xb5\xd9\xce\xd7\xab\x2c\xf7\xc8\xf2\x96\x55\xf5\x82\xf5\x56\x2d\xeb\x75\xe6\x05\xeb\xe1\x16\xe2\x06\xe5\x8a\xad\xbb\xd8\x28\xc6\x1c\xe6\x5e\xd4\x3a\x77\x34\x4f\xd6\x2b\x36\x9f\x62\xd9\x0d\x4e\xe9\xdd\x20\xc8\xb7\x40\x7f\x07\x00\x00\xff\xff\xd1\xa9\x95\xa7\x1d\x02\x00\x00")

func serviceSystemdServiceBytes() ([]byte, error) {
	return bindataRead(
		_serviceSystemdService,
		"service/systemd.service",
	)
}

func serviceSystemdService() (*asset, error) {
	bytes, err := serviceSystemdServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service/systemd.service", size: 541, mode: os.FileMode(420), modTime: time.Unix(1689127439, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	"vscode/index.html": vscodeIndexHtml,
	"service/systemd.service": serviceSystemdService,
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
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"service": &bintree{nil, map[string]*bintree{
		"systemd.service": &bintree{serviceSystemdService, map[string]*bintree{}},
	}},
	"vscode": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{vscodeIndexHtml, map[string]*bintree{}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

