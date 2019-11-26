// Code generated by go-bindata. DO NOT EDIT.
// sources:
// html/index.html (6.547kB)

package main

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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x71\x6f\xda\x4a\x12\xff\x9f\x4f\x31\xc7\x53\x65\xa3\x33\x8b\x0d\x0e\x0d\x0e\x70\xd7\xb4\xcd\x6b\xa4\x17\x35\x4a\xaa\xa7\x7b\x17\x45\xd1\x62\xaf\xf1\xe6\xec\x5d\x6b\x77\x21\xa1\x27\xbe\xfb\x69\x6d\x03\xb6\x31\x24\xb9\xd3\x3b\x3d\x55\xcf\x6a\x63\x76\x66\x76\x66\x76\x66\x7e\xe3\xc1\x8c\xff\xf2\xe9\xeb\xc7\x6f\xbf\x5d\x7f\x86\x48\x25\xf1\xb4\x35\xce\x6f\xad\x71\x44\x70\x30\x6d\x01\x00\x8c\x13\xa2\x30\xf8\x11\x16\x92\xa8\x49\x7b\xa1\xc2\xee\x69\xbb\x60\x49\x5f\xd0\x54\x81\x14\xfe\xa4\x1d\x29\x95\x4a\xaf\xd7\xf3\x03\x86\x1e\x65\x40\x62\xba\x14\x88\x11\xd5\x63\x69\xd2\x23\x7a\xbf\x92\x7f\x77\x91\x8b\xec\x5e\x40\xa5\xda\x90\x50\x42\xb5\x7c\x7b\x3a\xee\xe5\xda\xa6\xad\x71\x2f\xb7\xde\x6a\x8d\x67\x3c\x58\x15\xb6\x66\x62\xaa\xff\xe7\x8b\x80\x2e\x81\x06\x93\xb6\xa2\x2a\x26\x85\x37\x19\x23\x72\x40\xaa\x55\x4c\x26\x6d\x45\x9e\x55\x17\xc7\x74\xce\x3c\xf0\x09\x53\x44\xb4\xa7\xbf\xd0\x25\x81\x5b\x25\x08\x4e\xe0\x8a\x33\xaa\xb8\x18\xf7\x22\xa7\x50\xda\x0b\xe8\x72\x67\xac\x6a\x48\x10\x1c\x3f\x28\x9a\x90\xf6\x46\x7f\x44\xe8\x3c\x52\x5e\xff\xc4\x4e\x9f\xcf\x20\xa0\x32\x8d\xf1\xca\x83\xb9\xa0\xc1\x59\xf6\xb7\xab\x48\x92\xc6\x58\x91\xae\xcf\xe3\x45\xc2\xa4\x07\xa7\xef\xc0\xb5\xdf\x81\x9b\xdd\xce\xca\x7e\x6b\xd3\x25\x07\x9a\x6d\x3f\x84\xa9\xdc\xda\x9f\x61\xff\x5f\x73\xc1\x17\x2c\xd0\xfa\xb9\xf0\xe0\x27\x32\x0b\x1d\x32\x3b\x6b\x37\x69\x7a\x85\xf6\x19\x55\x02\x2b\xf2\x56\x0b\xb5\xb8\x35\x27\x2a\x4c\xe5\x43\x44\xa5\xe2\x62\x55\x8f\xe0\xc0\xd6\x11\xdc\xaa\x1c\xf7\xf2\xa4\xb7\x36\xd5\xa5\x56\x69\x91\xce\xde\x23\x5e\xe2\x9c\x5a\xc4\x6e\x89\x05\x3c\x7f\x78\xa6\xf2\x13\x56\x18\x26\x70\x77\x7f\xb6\xa5\x47\x54\x5e\x38\x4d\x34\x7b\x9f\xf8\x09\x2b\x52\x50\xb7\x64\xa1\xf6\xb6\x6b\x52\x7d\xb7\x50\xe7\x37\x7b\xa4\xba\x3e\x9e\x2a\xca\x19\x4c\xe0\xdf\xdb\x04\x28\xce\x63\x45\x53\xaf\x44\xca\xc8\x82\xce\xe7\x44\x78\x60\xe0\x67\x2a\x0d\xab\xc2\x4c\xb9\xa4\x5a\x91\x07\xe1\x82\xf9\x99\x4a\x33\x55\x9d\x9a\x0a\x7d\x09\xa2\x16\x82\xc1\x5d\xaa\xee\xec\x7b\x0b\x0c\xc7\x7e\x67\x14\x1e\x6e\xae\xf5\x76\xb5\xde\x59\xc9\x00\x55\xf7\x29\x26\xa1\xf2\xc0\xc8\x51\x54\x73\x49\xe7\xc5\x03\xe3\xe2\xfa\x16\xbe\xe4\xf9\x2d\x09\x94\x15\x73\x1e\xcf\xf8\x73\x5d\x75\x48\xb0\x5a\x88\x3d\x8b\xfa\x0a\xb0\xc2\xff\xe4\x3c\x69\xe2\xe9\x6b\xa5\xf3\x7e\xc9\x02\xf2\xec\x81\xc1\x38\x23\xc6\x9e\xd8\xda\x6a\x08\x8c\x76\x52\x1b\x6c\x60\x4a\xbc\x24\x1f\xe4\x65\x82\xe7\x99\xc0\xcb\xe1\xca\x6a\x6f\x2f\x85\xab\x94\xe8\x70\x61\x45\xe6\xd5\x78\xe8\x6b\xa6\x01\x85\xc5\xea\x67\x9c\x7a\x10\xe2\x58\x92\x2a\x5f\x1f\xdb\xdb\xd4\x64\x93\xcd\xd5\x11\x9b\x4b\x1c\x2f\xc8\x31\x83\x77\x76\x56\x0c\xba\x1a\x9a\x74\xef\x62\x7e\xd7\xa8\x9f\x32\x49\x83\xba\x01\xa9\xb0\x50\x1e\x8c\xec\x2a\x99\xb0\xc0\x03\xc7\xb6\x4b\x76\x6a\x4e\x17\x1b\x9b\xf7\x55\x89\x11\x66\x41\x4c\x2e\x7d\x5d\xfb\xc6\x95\x63\xa3\xf7\x96\xe3\xa0\xd1\xb2\xeb\xa0\xc1\x97\x11\x1a\x2c\x1d\x34\xf0\xbb\x2e\x1a\x59\x36\x1a\x74\x4f\xd1\xa9\xe5\x22\x37\xbb\x8f\x90\xeb\xdb\xd6\x89\x35\x40\x23\x6b\x84\x1c\xab\xa0\xe9\x0d\x91\x83\x06\x99\x06\xdf\x45\xa3\xae\x8d\x06\x9a\xd9\x75\x91\x9b\xdd\x47\xc8\xfd\xe8\x8c\xd0\x89\xe5\x0c\xd1\xc0\x72\x4e\xd0\xd0\x72\xfa\xa8\x6f\x6d\x8d\x7f\x87\x2b\x67\x80\x06\x56\xdf\x45\xee\x97\x21\x7a\xff\x6b\x7f\x10\x0d\xd1\xf0\x57\xbd\xde\xf0\x9c\x11\x1a\x6a\x9e\x36\xe3\x66\x5c\x4d\xf9\x6e\x34\x9d\xee\x96\x7e\xd7\x31\x3e\xb5\xdf\x35\xb3\x75\xd7\x6c\x02\x43\xd1\x9c\x8d\x9f\xc2\x30\x34\x1a\x8a\x3a\xc2\x01\x7f\x3a\x8f\x17\xc2\x83\xc1\x21\xf6\xc7\x42\x87\x98\xcf\xb0\x69\x5b\x50\xfc\x43\xc3\xce\x41\x8d\x5f\xc3\x50\x12\xf5\x0f\x0f\xfa\xc7\x25\x7e\xf3\xa0\x7f\x08\x49\xf7\xbb\xad\x92\x08\x4a\xa4\x07\x77\x15\xd9\xfd\xd3\x32\x9c\x90\xbc\xe1\x3c\x38\xf5\xfe\x08\xbb\x4a\x9d\xe1\x7a\xa7\xca\x8c\x24\x9c\xab\xc8\x03\x25\x16\xa4\x81\xbb\x4a\x66\x3c\xde\x74\x93\xa6\xf6\x90\xa4\x31\x65\x73\xdd\x9c\x97\x44\xe0\x79\x93\x10\x55\x24\x39\x98\x2a\x28\xa5\x4b\xcc\x67\x66\xff\xe4\xc4\x82\xf7\xb6\x05\xce\xc0\xe9\xbc\xaa\x7f\x61\x41\xf0\x6b\xd4\x33\xf2\x04\x9b\x11\x6b\x2e\x70\x1a\x51\x1f\xfd\x42\x19\xc1\xe2\x67\x81\x03\x4a\x98\xda\xa5\xd9\xb1\xea\x70\x2f\x5f\x3c\xcb\xe2\x1e\x4c\x8f\x1e\xc9\x39\x39\xb5\x60\x78\xda\x70\x24\x68\x6a\x04\x4d\xe6\x9c\x37\x98\x3b\x1c\xc1\xcc\xdc\x7d\xe7\x35\x91\xdd\x36\xde\x0b\xa7\x75\x44\xf4\x78\x45\xda\x47\x4a\x32\xa6\xcd\x55\xf5\x07\xab\x49\xd7\xcd\x03\xda\xef\xf7\x9b\x4b\xf2\x58\xe4\xec\x43\xa1\xcb\x1f\x37\xeb\xca\x3c\x74\xf3\xed\xa2\x3a\x12\x35\x0d\x1f\xc5\x80\x71\x43\x70\x0c\xdf\x68\x42\xe0\xe2\xfa\xd6\x38\x34\x61\xfc\xf1\x06\xaa\xdf\x6d\x42\xc8\x87\xcc\xdf\x73\x40\xa8\x3d\xe5\xd3\x98\x2a\xdd\x41\x9a\x6a\x48\x46\xfc\xa9\xf0\xf5\xe5\x90\x6c\x3b\x7d\x55\xcf\xb1\xc6\x7e\xb0\xa9\x6b\xc3\xb7\x05\x44\x1a\x42\x15\xf1\x25\x11\x1f\x18\x4d\x70\x91\xe5\x7d\x91\xa3\xd8\x78\x65\xa7\x59\x37\x27\xa8\xd4\x48\x4a\x12\x87\x4f\xbd\xd7\x3c\x0e\x37\x8e\xff\xe3\xb9\x07\x83\x37\x1f\xbb\x34\xf3\x35\xe1\x3e\xfb\xc2\xf4\x56\xd8\x9f\xe7\xdf\x4e\xff\x84\xfe\x8f\x0b\xfd\x4d\x8a\x7f\x04\xec\x9f\xdf\x1c\xc2\xc0\x12\x0b\x08\x53\x59\x7c\x5b\x86\xc9\x76\x4c\xa3\x8c\x2a\x33\xe0\xfe\x22\x21\x4c\xa1\x39\x51\x9f\x63\xa2\x3f\x9e\xaf\x2e\x03\xd3\x28\xbd\x41\x31\x3a\x9d\xb3\xb2\x2a\x0d\x92\x0c\x23\xaf\xd5\x55\x79\xa3\x54\xd1\x36\x13\xff\x83\xb2\xe2\x05\x52\xa6\x70\xab\x31\xc6\x52\x5d\x73\xca\x14\x4c\xa0\x6d\xb7\x4b\xef\x57\x08\x0b\x88\x28\x5e\xdc\xec\x30\xa8\x03\x58\x46\xa1\x5e\xa3\x90\x8b\xcf\xd8\x8f\xcc\x9d\x18\xae\x23\x55\xab\x64\xfc\x09\x26\xd9\xe8\xab\x01\x62\x62\xa4\xdd\x92\x0a\x27\xe9\x43\x22\x3b\x7b\xe2\x2a\x91\x4a\xc0\x04\xee\x18\x7f\xd2\xc7\xfa\xc2\x17\x42\x9a\x1d\x0b\x8a\xf5\x15\x65\x0b\x45\xca\x94\x5b\xe2\x73\x16\x48\xb3\x73\x8f\x1e\x39\x65\xa6\xe1\x19\x9d\xea\x3c\x94\x4d\x41\x28\x5d\xc8\xc8\xc4\x08\x2f\xe7\x3a\xc2\x0f\x4e\xcd\x78\x3e\x2a\xed\x49\xd9\xfb\x62\xfa\x1c\xb9\x5c\xe6\x6c\xcd\x98\xee\xb5\x2f\xd9\xca\xfa\xf1\x8b\xa6\x74\xc1\x6e\x84\x8a\x34\xd6\x05\x8e\xbb\x42\x43\x30\x33\x77\x62\xc2\xe6\x2a\x82\x29\x0c\xed\xc6\x66\xaa\x65\x64\x44\x43\x65\xee\x4f\xe5\xb9\xaf\x87\xb9\x99\x0b\x87\xd9\xe7\x37\x8d\xcc\x75\xd5\xd3\x72\x41\x56\x2b\x64\x87\xd8\x4e\x2b\xdf\x53\x6c\x0d\x89\xf2\x23\xd3\xe8\x2d\x9d\x5e\xaa\x77\xca\xbf\x69\x25\x0f\xd9\xe7\x89\x01\x7f\xdd\xe9\xdc\x59\x46\x2a\x22\xcc\x14\x44\xa6\x9c\x49\x02\x93\x29\x6c\x3e\xa3\x47\xc9\x99\xd9\xa9\x8b\x06\x19\x12\xa6\xb5\x98\xed\x50\x92\x23\xa3\xc2\x2c\x61\x1f\x49\xa2\xbe\x66\x4f\x57\x73\x3b\x5c\x77\xaa\x0f\xa9\x1d\xb6\xf7\x84\xcf\x6f\x6a\xb2\xbb\x06\xb5\x27\x5b\x92\x5c\x6f\x80\x2e\x89\xba\x64\x8a\x88\x25\x8e\x4b\x18\x2d\xe7\xff\xbf\x09\xe1\x1b\xc3\xf8\x52\x28\x5f\x0c\xe7\xdb\x43\xfa\xd6\xb0\xae\x8b\xf5\xda\x02\xc7\xb6\xed\x4e\x6b\xf7\xab\x47\x6b\xdc\xcb\x7e\x7c\xf9\x4f\x00\x00\x00\xff\xff\xce\xbb\x45\x4a\x93\x19\x00\x00")

func htmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexHtml,
		"html/index.html",
	)
}

func htmlIndexHtml() (*asset, error) {
	bytes, err := htmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.html", size: 6547, mode: os.FileMode(0644), modTime: time.Unix(1574740149, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0xe6, 0xa6, 0x66, 0xb1, 0xb0, 0xa4, 0x6, 0x7, 0xf1, 0x38, 0x35, 0xc0, 0x5a, 0x71, 0x59, 0x4, 0x9d, 0x54, 0x83, 0xbe, 0x3a, 0xbf, 0x39, 0xbe, 0x31, 0xeb, 0xc6, 0xd4, 0x25, 0xc1, 0xf6}}
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
	"html/index.html": htmlIndexHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{htmlIndexHtml, map[string]*bintree{}},
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