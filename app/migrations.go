// Code generated for package app by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/201911252319_init.down.sql
// migrations/201911252319_init.up.sql
// migrations/202001170654_add_patchset_announcements.down.sql
// migrations/202001170654_add_patchset_announcements.up.sql
package app

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

// Mode return file modify time
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

var __201911252319_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x49\x4d\xcc\x8d\x4f\xce\xcf\x4b\xcb\x4c\x2f\xb6\xe6\x02\x4b\x78\xfa\xb9\xb8\x46\x28\xe4\x24\x16\x97\xc4\x17\xa5\x16\xe4\x17\x95\xc4\x67\xa6\x54\x40\xe5\x20\x9a\xd2\x53\x8b\x8a\x32\x4b\xe2\x4b\x8b\x53\x8b\x50\x35\x95\x16\xa4\x24\x96\xa4\xa6\xc4\x27\x62\xea\xc9\xcc\xcb\xc9\xcc\x4b\x8d\x4f\xce\xcf\xcd\x4d\xcd\x2b\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\xbd\x25\xed\x8b\x85\x00\x00\x00")

func _201911252319_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__201911252319_initDownSql,
		"201911252319_init.down.sql",
	)
}

func _201911252319_initDownSql() (*asset, error) {
	bytes, err := _201911252319_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "201911252319_init.down.sql", size: 133, mode: os.FileMode(420), modTime: time.Unix(1574794364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __201911252319_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd1\xc1\x4e\x84\x40\x0c\x06\xe0\x33\xf3\x14\x3d\xee\x26\xfb\x06\x9e\x50\xe7\xb0\x11\x59\x43\x30\x71\x4f\x4d\x03\x75\x9d\xc8\x0c\x64\x28\x66\x7d\x7b\xe3\x2e\x6a\x65\xb8\x15\xda\xf9\xf3\x35\xbd\xab\x6c\x5e\x5b\xa8\xf3\xdb\xc2\xc2\x89\x63\x74\x82\xd3\xc8\x71\x84\x8d\xc9\xd4\x77\x20\xcf\x20\x7c\x16\x28\x0f\x35\x94\xcf\x45\xb1\x33\x59\xf3\x46\x82\xae\x4d\xfe\x77\x34\x0a\x46\x1e\xfa\x28\x20\xce\xf3\x28\xe4\x87\x9d\xc9\x9e\xaa\xfd\x63\x5e\x1d\xe1\xc1\x1e\x61\x03\xcb\xf4\xad\xd9\xde\x18\x33\x83\xf6\xe5\xbd\x7d\x01\x15\x84\xae\x3d\xc3\xa1\x5c\x18\xf5\x04\xa8\xe7\xd7\x7d\x5c\xe8\x5c\x60\x6c\x7a\xef\x39\xc8\x65\xa5\xb9\x5e\x53\x4f\x43\x4b\xc2\x2d\x92\x42\xeb\xfe\x7f\xbd\x0a\x5a\x81\xff\x65\xfd\xb8\x13\x8b\x9a\x49\xe5\xc2\xe4\xb1\xe9\xc3\xab\x3b\xcd\xec\xef\x12\xdf\xf9\x33\x3d\xc2\xb5\xf5\x41\xdd\x94\x5e\x68\x69\xfe\x4d\xb9\x98\xbf\x02\x00\x00\xff\xff\x69\x41\xee\x7c\xfd\x01\x00\x00")

func _201911252319_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__201911252319_initUpSql,
		"201911252319_init.up.sql",
	)
}

func _201911252319_initUpSql() (*asset, error) {
	bytes, err := _201911252319_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "201911252319_init.up.sql", size: 509, mode: os.FileMode(420), modTime: time.Unix(1574794364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __202001170654_add_patchset_announcementsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcc\x41\x0a\xc2\x30\x10\x85\xe1\x7d\x4f\x91\x0b\xf4\x04\x5d\x29\x69\xa1\x20\x55\xac\x0b\x77\xc3\x30\x0e\x36\xd2\xcc\xc4\x64\x02\x1e\x5f\xdc\x14\x17\x82\xcb\xc7\xfb\xf8\xdb\xd6\x89\x06\x29\x89\xc9\x82\x8a\x9b\x9f\xeb\xa4\x1e\x0d\x67\xad\x99\x78\xdc\x9e\x41\xf3\x10\x56\x6e\x1a\x7f\x3e\x9e\xdc\x38\xf9\xfe\xea\x12\x1a\x2d\x85\x0d\x50\x44\xab\x10\x47\x16\x83\x94\xf5\xc1\x64\x20\x18\x19\x68\x41\xb9\x33\x48\x8d\xb0\xe9\xcf\x08\xb7\x57\xf7\x3f\x65\xe5\x0b\x5e\x76\xfb\x43\xff\x1b\x96\xae\x79\x07\x00\x00\xff\xff\x48\x9d\xb1\xf1\xc9\x00\x00\x00")

func _202001170654_add_patchset_announcementsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__202001170654_add_patchset_announcementsDownSql,
		"202001170654_add_patchset_announcements.down.sql",
	)
}

func _202001170654_add_patchset_announcementsDownSql() (*asset, error) {
	bytes, err := _202001170654_add_patchset_announcementsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "202001170654_add_patchset_announcements.down.sql", size: 201, mode: os.FileMode(420), modTime: time.Unix(1574794364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __202001170654_add_patchset_announcementsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcf\x6a\xc2\x40\x10\xc6\xef\x79\x8a\x39\x2a\xc4\x27\xf0\x94\xd6\xb1\x2c\x4d\x56\x49\xb6\xa0\xa7\x65\x59\x07\x93\x92\xcc\xa6\xd9\x09\xf4\xf1\x4b\x5b\xb4\x11\xab\xed\x5c\xbf\xe1\xfb\xf3\x5b\x2c\x80\x43\xc3\xb1\x27\x2f\x4d\x60\xa8\xde\x5a\x1d\x56\x4e\x5c\x15\xc6\xc1\x93\x3a\x2b\xeb\x30\xac\x9b\x96\x92\xe4\xb1\xc4\xcc\x20\x98\xec\x21\x47\xe8\x9d\xf8\x3a\x92\x58\xc7\x1c\x46\xf6\xd4\x11\x4b\x84\x59\x02\xdf\xc7\x63\x07\x4a\x1b\x7c\xc2\x12\xf4\xc6\x80\x7e\xc9\xf3\xf4\x24\xf6\x43\x78\x25\x2f\x96\x5d\x47\x60\x70\x67\xae\x5f\x7c\xed\xf8\x48\xf6\xbe\xcd\xa9\xc3\xdd\xaf\x8e\x62\x74\x47\xb2\xb5\xe3\x43\x7b\x2b\x4e\x22\x18\x55\x60\x65\xb2\x62\x7b\xad\x6e\x4b\x55\x64\xe5\x1e\x9e\x71\x0f\xb3\xaf\x69\xf3\x64\xbe\x3c\x13\x51\x7a\x85\xbb\xdf\x89\xd8\xe9\x54\xfb\x33\xca\x4e\xbb\xdb\xe6\xf0\x0e\x1b\x7d\x93\xe9\x05\xaf\x74\x82\x26\xbd\x44\x30\x5f\xfe\xa7\x91\xc4\x3f\xf3\x24\x7e\x7a\x7d\x04\x00\x00\xff\xff\x63\xd0\x34\x07\x22\x02\x00\x00")

func _202001170654_add_patchset_announcementsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__202001170654_add_patchset_announcementsUpSql,
		"202001170654_add_patchset_announcements.up.sql",
	)
}

func _202001170654_add_patchset_announcementsUpSql() (*asset, error) {
	bytes, err := _202001170654_add_patchset_announcementsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "202001170654_add_patchset_announcements.up.sql", size: 546, mode: os.FileMode(420), modTime: time.Unix(1574794364, 0)}
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
	"201911252319_init.down.sql":                       _201911252319_initDownSql,
	"201911252319_init.up.sql":                         _201911252319_initUpSql,
	"202001170654_add_patchset_announcements.down.sql": _202001170654_add_patchset_announcementsDownSql,
	"202001170654_add_patchset_announcements.up.sql":   _202001170654_add_patchset_announcementsUpSql,
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
	"201911252319_init.down.sql":                       &bintree{_201911252319_initDownSql, map[string]*bintree{}},
	"201911252319_init.up.sql":                         &bintree{_201911252319_initUpSql, map[string]*bintree{}},
	"202001170654_add_patchset_announcements.down.sql": &bintree{_202001170654_add_patchset_announcementsDownSql, map[string]*bintree{}},
	"202001170654_add_patchset_announcements.up.sql":   &bintree{_202001170654_add_patchset_announcementsUpSql, map[string]*bintree{}},
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
