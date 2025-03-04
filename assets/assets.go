// Code generated by go-bindata.
// sources:
// plugins/codeamp/graphql/schema.graphql
// plugins/codeamp/graphql/static/index.html
// DO NOT EDIT!

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

var _pluginsCodeampGraphqlSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xcd\x72\x23\xb9\x0d\xbe\xeb\x29\xe8\xda\x8b\xb6\xca\x4f\xa0\x5b\x66\xec\x8c\x9d\xcc\x24\x8e\xb5\x73\x48\x4d\xf9\x40\xb7\x60\x89\x71\x37\xd9\x4b\xb2\x3d\xe3\x4a\xe5\xdd\x53\xfc\x6d\x80\x64\xcb\x96\x76\x52\x95\x8b\xad\x46\x93\x20\x00\x02\x1f\x7e\xda\x74\xbc\xe7\x9a\xfd\x26\x06\x58\xc5\xdf\x7f\xd9\xfe\xfd\x6f\xab\x95\xe9\x0e\x30\x70\xf6\xef\x15\x63\xbf\x4f\xa0\x5f\x37\xec\x1f\xee\xdf\x8a\xb1\x61\xb2\xdc\x0a\x25\x37\xec\x4b\xfc\xb5\xfa\xcf\x6a\xf5\x4b\x7c\x6f\x5f\x47\x08\x3f\xfd\xde\x5f\xd8\x57\x03\x7a\xc5\xd8\x64\x40\xaf\xc5\x6e\xc3\x6e\xaf\x7e\xdd\x24\x62\x78\x6b\xe2\x6b\xb3\xfe\x75\xc3\xbe\x39\xca\xc3\x85\x7f\x79\xa7\xd5\xbf\xa0\xb3\x2b\xc6\xc6\xf0\x2b\x32\xb8\x64\xa6\x9f\xf6\x1b\xb6\xb5\x5a\xc8\xfd\x25\x93\x7c\x80\xf9\x09\xe4\x8b\xd0\x4a\x0e\x20\xed\xed\x55\x22\xff\xba\x41\xdc\x32\x67\x33\xb3\x36\xeb\xf8\x63\x0b\x5c\x77\x87\xbc\x3c\x3c\xde\xca\x71\xb2\x97\x6c\xe4\x9a\x0f\x66\xc3\xee\xf8\x5e\x48\x6e\x95\xf6\xf4\x99\xf7\x67\x61\x6c\x10\xfd\xcf\xc0\xed\xa4\xc1\x1d\xf0\x14\x7f\xae\x17\x77\xc7\xc5\xf3\xee\x2d\xe8\x17\xd1\xf9\xdd\x26\xfe\x5c\xde\x1d\x17\xcf\xbb\xaf\x7f\x8c\x4a\x5b\xc6\xfb\x3e\xef\x66\xdf\x85\x3d\x08\xc9\xf6\xe2\x05\x64\x52\xf9\xf6\x8a\x71\xb9\xa3\xf6\x5a\x31\x06\x7e\xfb\xb6\x3c\xf7\x9a\x90\xf3\xe1\xde\xba\x17\x8c\x61\xb9\x99\x19\xa1\x43\xc2\x6f\xdd\xa3\xbf\xdc\xed\x4c\x88\x77\x7c\x0f\x3d\x70\xe3\x55\xd5\xf1\xe7\xb2\xaa\x71\x31\x52\x75\x96\xdd\x71\x40\xaa\xcc\xf7\x89\x5c\xc5\x89\x80\xb6\x3c\x54\x4c\xd8\x0b\xd7\x82\x3f\xf6\xd1\xf4\x9d\x06\x7b\xd4\xf2\x6e\x41\xd3\xf0\xd0\xe2\x79\xce\x2d\x50\x11\xae\x31\x75\xe1\x0e\x84\xdc\xf7\x10\x65\xcb\x5a\xf8\xc8\xc9\x46\xc8\x2f\x73\x24\x5c\xff\xb0\x20\x8d\x50\xd2\xdf\x9d\x3f\x3f\x12\xcc\x7a\x29\xa0\xbe\xe5\x4d\x34\x5e\x33\x19\x85\xd7\x4c\xf3\x7e\x50\xae\xa4\xce\x80\x0e\x9f\xdd\xa2\xe0\x70\x5f\x50\x93\x08\xa0\x07\x61\xf2\xe1\xf3\x93\xdb\xe4\x90\xed\x22\x80\x55\x86\x2e\x8f\x57\xe9\x29\x42\xd6\x47\x0d\xdc\x42\x12\x7d\xc5\x58\xe7\x09\x51\xe8\xe4\x58\x39\xea\x0b\x10\x08\xb8\x36\xee\x28\x8b\xc9\x13\x4e\x61\x11\xa5\x88\xea\x67\x29\xa2\xe2\xeb\x48\xcf\x41\x51\xc4\x48\xf0\x05\xab\x46\xc4\xc0\x58\x35\xa6\xed\x01\x49\x2f\x8a\x0d\xf1\xcc\x18\xb7\xf9\xcc\x18\xb6\xeb\x48\xcf\x98\x53\x40\x10\xd6\x7c\x26\x05\xcd\x4f\x61\x71\x05\x3d\x10\x29\x76\x9e\x70\x0a\x8b\xdb\xc1\x87\xe2\xc0\xe5\xeb\x0c\x82\xdc\x32\x25\xfd\x02\x31\x10\x8c\x4b\x2b\x36\x71\x5f\x89\x72\x09\xb7\x2e\xa2\x97\x51\x33\xa5\x90\x21\xb6\x72\x10\xb7\x46\xf8\x97\xc5\x73\x0f\x85\xc4\xdb\xb0\x3f\xdb\xae\xe0\x4b\x0c\x78\x16\x5f\x6a\xd0\xc4\x97\x58\xf5\x2c\xbe\xd1\x0e\x08\x20\xb2\x19\x10\xa8\x62\x00\xd9\x60\xb4\x4d\x6c\xaf\xc9\xfe\x6c\x06\xca\x36\x58\xe1\x8f\xb0\x8d\x56\xa0\x6c\x83\x11\xfe\x08\xdb\xda\x08\x19\xf1\x91\x53\x78\x10\x0e\x58\x9c\xf0\x97\xa6\x91\x05\xcd\x31\xaf\xe4\x08\xef\xe3\x55\xab\x9b\x79\x31\x74\xfb\xef\x63\x86\xe3\xa9\x9d\xdb\xea\xe0\x0a\xa9\x2b\x66\xd1\x39\xb4\x48\xee\xfa\x16\x9e\x8b\xc0\x82\x32\x1b\x45\x9f\x4a\xe4\x75\x5e\xe0\x52\x62\xfc\x99\xef\x27\x11\x88\x45\x4b\x8e\xd1\x9d\xce\xe0\x98\xec\x5a\x72\x8c\x9e\x74\x06\xc7\x52\xeb\x32\xe3\xcc\x3c\xcb\x6c\xba\xa9\x72\x6e\x91\x49\x8e\x1b\xa3\xcc\x4b\x3f\xed\x20\x64\xa3\x48\x0b\xd6\xf9\x1f\x29\xe4\xba\x07\x9c\xec\xb3\x5e\xae\x99\x40\x25\xc1\x7a\xa2\xcf\xa1\x11\x41\x84\xd9\x2d\x43\x4d\x15\xdd\x92\x66\x73\x52\x66\x56\x26\x6c\x94\xa0\x98\x36\x2b\x82\x88\xf9\x58\x44\x4c\x67\x7f\x50\xea\x79\xe0\xfa\x19\xd5\x12\x8f\x91\x74\x47\xba\x22\x97\xcb\x3f\x28\xd5\x03\x97\x61\xe7\x27\xb0\xec\x93\xb0\xec\xa3\x1a\x06\xe1\x25\xdd\x83\xfd\x24\x6c\x7c\x5e\xe7\x02\xd4\xef\xae\x1a\x27\x4f\x93\xf0\x3d\x73\xc5\xfc\x7d\x11\x95\x0b\xe2\x95\x90\x16\xf4\x13\xef\x60\xa6\xf9\x5a\xaa\x53\x93\x43\xd1\x5b\x69\xe3\x16\x54\xc1\x87\xd2\x0b\x11\x1c\x70\xf4\xe0\x0d\x72\x84\x8d\xab\xf2\xad\x16\x2e\x59\xa7\x22\xf0\x21\x32\x9f\xeb\xf1\xc0\x7b\x7e\x3e\x9d\x75\xd8\x3b\x73\xce\x3d\x56\x62\x9d\x09\xe7\xf0\xf6\x9b\x13\x73\xd4\xfe\x05\xe6\x88\x70\x3a\xf3\xb8\x39\x31\x47\x9d\x69\x60\x8e\x08\xa7\x33\x8f\x9b\x13\x73\xdf\xc6\x7b\xae\xee\x97\xdf\x19\x7d\xd1\xed\x1a\xb8\xe8\x73\x77\x42\xab\xf1\x22\xc2\x02\xda\xed\x36\x7e\x0e\x41\xad\x42\x2c\x52\x9c\x50\x56\xd1\x8e\x36\x80\x31\x7c\x0f\xf8\x5c\x17\xf4\xf8\xf9\xc0\xcd\x81\xc8\xc5\x35\x48\x7b\x53\x50\x35\x3c\xe1\xc7\x96\x88\xa9\xe0\xc4\x1e\xf1\x0e\x11\x3b\x35\x0c\x5c\xee\x30\x77\x3c\xc1\xb8\xa0\x8d\x33\xa9\xba\x2e\xaa\xbb\x71\x39\xd5\x99\xd3\xf5\x36\x0f\x17\xb4\x07\x26\xb5\x8a\x7b\xe7\x04\x3d\xa2\x94\x43\xe9\xb1\x57\xaf\x6e\xf9\xd6\x6a\x6e\x61\xff\x1a\xba\xa6\x15\x63\xbd\x6b\x5a\xc1\x98\x3b\xad\x1e\x21\x53\x35\xf0\x9d\xa8\xc9\xa3\x06\xd7\x7b\xdc\x28\xf5\x9c\xce\x0b\x26\xc3\xc5\x93\x37\x1b\x6e\xbe\xa9\xe9\x4a\x9b\x3c\xc3\x2b\x7e\x14\xe6\x0a\x9e\xf8\xd4\x5b\x82\x7a\x9d\xea\x95\x3e\xaa\x62\x9a\xf7\xd4\xde\xdc\x9a\x03\x60\x20\x29\xe4\x2b\xe4\x79\xe1\xfd\x44\xef\xb0\x53\xd4\xda\x2d\x5f\x08\xbe\xe9\xa2\xa7\x75\x3f\x2f\xa0\x73\xb8\x24\x3c\x3a\x7e\xc3\x95\xba\xc2\x6c\x63\x6d\x47\xb1\x1b\x8f\x6a\x88\x03\x3b\x37\x7b\xe3\x26\xba\x71\xba\x87\xdf\x27\x30\xb6\xa0\x7e\x16\x83\x20\xb4\x01\x06\xa5\x5f\x1b\x8b\xc3\x8b\x6a\xbd\x75\x10\x21\x7d\x2f\xfe\x49\xf3\x0e\xee\x40\x0b\xb5\x6b\x44\x46\x8e\x8a\x05\xa5\x6b\xdf\xc0\xe9\x87\xa4\x9e\x77\x44\x2c\xbd\x25\xae\xad\x78\xe2\xde\x85\xc2\x44\x81\xb1\x03\xf0\x5d\xc4\xa8\x3c\xce\xf3\xfa\x70\xd1\xb7\xe8\xc6\x72\x0b\x14\x6d\x8a\x21\xc7\xd2\x88\xc3\xef\xfc\x52\x83\xdc\x49\x4e\x61\x2c\xd7\x84\xf0\x24\xa4\x30\x07\x6a\xc2\x7b\xd5\xf7\x8f\xbc\x7b\xae\xb2\x7e\xac\x43\x70\x36\x79\xc3\x61\xf0\xb8\x36\x68\x3b\x2a\x23\xac\xd2\xaf\xf4\x6a\x63\x13\x92\x29\x7b\x61\xbf\xea\xbe\xa0\xdc\x69\x65\x55\xa7\x08\x59\x1b\x7e\xa7\xc5\x0b\xb7\xf0\x57\x1a\x95\xee\xc5\xf4\xd8\x8b\xae\xa0\xe7\xa1\xac\x39\xa8\xef\x57\x1e\xf5\x9c\xf6\x51\xd3\x23\x93\xde\x62\x56\xdb\x4d\xda\x25\x8f\xfb\x62\x18\x73\xce\x34\xf3\x8d\x49\xef\x25\x33\x7e\x14\x8d\x14\xa9\x87\xbf\xc7\x47\x96\x4b\x2c\xf0\x14\x13\xb0\x07\x36\xc7\x74\x7b\x61\x3f\x68\x2e\x3b\x92\x2d\x3b\x25\xad\x90\x93\x9a\x4c\x30\x26\x01\x65\x20\xf5\x6f\x5d\xe4\xa6\x7a\x16\xdd\xc0\x52\xca\x2d\xe6\x94\x21\x85\x64\xda\x1b\xb0\xa5\x86\x51\x49\x1f\x20\x08\x71\xca\x94\xc8\xbb\x03\x38\xe4\x27\xa2\x14\x58\x7f\x34\xd8\x94\x7c\x12\xfb\x19\x1a\x5a\x5a\x54\xbd\x0c\x8e\xa5\x25\x75\x5a\xc0\xd4\xea\x31\x17\x20\xaa\x92\x6b\x32\x56\x0d\x1f\x0b\x6a\x05\x4c\x3f\x01\x6f\x30\xf2\xa2\xde\x13\x63\xf0\x92\xce\xe5\x90\xb3\xd0\xb9\xb4\x18\x0e\xa3\xad\xd8\xcb\x08\xbc\x25\xc4\x2c\xbc\x2b\x55\x2f\x5d\x63\xc9\x14\x0d\x63\x53\x80\xf5\x93\xf9\x1a\x62\x8f\x78\x06\x13\x2e\x5e\x57\xfe\x2f\x99\x11\x67\xfb\xc4\xa2\x6a\x11\x11\x9b\xd0\x49\xe3\xb6\x1d\x79\xf4\x6e\xe7\x81\xff\xb1\x20\x27\xb2\x87\x2f\x66\x2d\x0d\xd0\xb7\x34\xaf\x47\x9d\x06\x96\xc0\x00\xb7\x9b\xce\x4b\x08\x73\x82\x71\x9e\xf1\xe8\x6f\xe8\xd6\x0f\xe8\xfa\x50\x63\xb8\x07\xe2\x88\x98\x05\x1e\xa1\xd7\x16\x46\xc9\x7d\x36\x07\x8a\x48\x4a\x6c\x1a\xcf\xa7\x1c\xa5\x3b\xb8\x87\xc7\x49\xf4\x95\x6a\xa9\x1a\xc3\x42\xe1\xe9\x76\x2d\x54\xf3\xec\x77\x74\x16\xed\xee\x21\x9e\x75\xa7\x74\x70\xb3\x8b\x87\x86\xff\x2f\x6a\xd6\xea\x1a\xae\x2a\x5a\x60\x5c\x77\x12\x37\xc0\x7b\x7b\xf0\x0f\x7e\x49\xa3\xab\x68\x2c\x59\xec\x30\xd2\xec\x3e\x8e\x21\x89\x49\x1b\xd3\x7d\x6f\xd9\x84\x19\xff\xfc\xd3\x97\xcf\x81\xd7\xb9\xd7\x4c\x45\x08\x1f\xf1\x88\x08\x8d\xcf\xa8\xc1\x65\x4f\x3d\xe4\xa3\x92\x96\x0b\x09\x9a\x55\x67\x94\xb7\x19\x0e\x50\x1a\xdd\x7b\x46\x8c\x34\x9a\x09\x3b\x17\x6e\xcd\x33\x28\xdd\x61\xe0\x3f\xb6\x93\x8e\x81\x16\x09\x5f\x25\x7f\xe1\xa2\x0f\xf9\x93\xb2\x2e\xef\xb0\xe2\xe9\xbb\x02\x7b\xa0\xe5\x7e\xe1\xd1\x58\x0f\xdf\x62\x1d\x60\xc0\x0c\x46\x6e\x31\xbc\x09\x29\xac\xe0\xfd\x15\xf4\xfc\x75\x0b\x9d\x92\x3b\x93\xb6\x8e\xbe\xb3\x28\x88\x56\x0c\xa0\x26\x5b\x50\xcd\xd4\x75\x60\xcc\x6f\x07\x0d\xe6\xa0\x5c\xf0\x06\xfa\x13\x17\xfd\xa4\xa1\xa2\x1f\xac\x1d\x6f\x80\xef\x40\xbb\xd0\x42\x7a\xdf\xe4\x17\x29\xc8\x5a\xd6\x29\x56\x79\x3b\x95\x51\x5c\x34\x9b\x55\x47\xd7\x72\x87\xfc\x21\xa7\x06\x93\xff\xa7\x06\x6f\xb1\x77\xc3\xed\x39\x89\xa8\xe2\x13\xcd\xdb\xea\x9d\x33\x49\x58\x9c\x10\x14\xa6\xce\xdf\x4f\x6a\x31\xde\x1a\x18\x54\xf5\xc6\xc2\x00\x81\xe4\xe4\x45\x50\x5e\x6a\xfc\xdb\x3a\x2c\x63\xe5\xfc\xb9\x26\x42\xa5\x27\xfc\x24\xa4\x6c\x0a\x73\x04\x35\x0b\x61\x4e\x3f\x8f\xb4\x0f\xc5\x11\xf8\x83\xc3\xc9\x4e\xd4\x6c\x2e\x16\x2f\xa7\xdd\x65\x94\x25\x3a\x75\x88\x66\xe7\xd0\x2a\xb3\x8e\x6a\x72\xd9\x34\xdb\x25\xae\xab\x0b\xfa\xfb\x1a\x87\x63\x56\xa7\x1f\x59\x88\xc8\xad\xef\x2f\x5e\xe2\xc9\x80\x2e\xaa\x2d\x32\x3e\xa6\x1b\x23\x9e\x9e\x7a\xde\x7c\x5c\x15\x8c\x7b\xcd\x65\x15\x3b\xf5\xa7\x9b\xa6\xfd\x5b\x80\xb4\xe8\x0a\xef\x3c\x28\x98\x66\xf1\xa0\xd9\x72\xad\x98\xa0\xa6\x5b\x10\x33\x98\xef\xbf\x01\x00\x00\xff\xff\x3d\x35\xd0\x46\x11\x28\x00\x00")

func pluginsCodeampGraphqlSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsCodeampGraphqlSchemaGraphql,
		"plugins/codeamp/graphql/schema.graphql",
	)
}

func pluginsCodeampGraphqlSchemaGraphql() (*asset, error) {
	bytes, err := pluginsCodeampGraphqlSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/codeamp/graphql/schema.graphql", size: 10257, mode: os.FileMode(420), modTime: time.Unix(1549399964, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsCodeampGraphqlStaticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xeb\x6f\xdb\x36\x10\xff\xde\xbf\xe2\xe6\x6c\x90\x5d\xd8\x94\xd3\xf5\x01\xa8\x76\x86\xb6\x49\xbb\x16\x69\xd3\xc6\x19\x8a\x7d\x2b\x4d\x9e\x2c\xa6\x14\xa9\x1e\x29\x3b\x6a\x91\xff\x7d\xa0\x64\xc9\x8a\x91\x0c\xd9\xb0\x41\x1f\x4c\xde\xe3\x77\x0f\xde\xc3\xb3\x9f\x8e\xcf\x5e\x5d\xfc\xf9\xf1\x04\x32\x9f\xeb\xa3\x07\xb3\xe6\x07\x60\x96\x21\x97\xe1\x00\x30\x73\xbe\xd2\xd8\x9c\x01\x96\x56\x56\xf0\x63\x7b\x01\xc8\x50\xad\x32\x9f\xc0\xe1\x74\xfa\xcb\xf3\x8e\x9a\x73\x5a\x29\x93\xc0\x74\x47\xda\x28\xe9\xb3\x7d\x39\xbb\x46\x4a\xb5\xdd\x24\x90\x29\x29\xd1\xb4\x9c\xeb\xed\xef\xc1\x8a\x78\x91\xa9\x6f\xfa\x76\x8b\xeb\x6c\x5f\x81\x5d\x6e\xfc\xc4\xdb\xaf\x68\x7a\x1a\x4b\x2e\xbe\xae\xc8\x96\x46\x26\xa0\x95\x41\x4e\x93\x15\x71\xa9\xd0\xf8\xe1\x41\xfa\x2c\x7c\x63\x38\xc0\x47\xe1\x1b\xed\x9c\x5b\x5a\x92\x48\x93\xa5\xf5\xde\xe6\x09\x1c\x16\x57\xe0\xac\x56\x12\x0e\xe4\x34\x7c\x3b\xc9\xd4\x1a\x3f\x49\x79\xae\x74\x95\x80\xab\x9c\xc7\x7c\x0c\x13\x5e\x14\x1a\x27\xed\x35\x5a\x70\x03\xaf\x89\x1b\xa1\x9c\xb0\xd1\x18\x22\xb6\x78\xfd\x61\x71\xac\x5c\xa1\x79\x35\x39\xc7\x55\xa9\x39\x05\xfa\x02\x57\x16\xe1\x8f\xb7\xd1\x18\xea\x63\x47\xfa\xfc\x31\xb0\x7f\x47\xbd\x46\xaf\x04\x87\x0f\x58\x62\x34\x86\xac\x25\x8c\x21\x3a\x2d\x85\x92\x1c\xde\x10\x37\x32\xf0\x38\x29\xae\xc7\xe0\xb8\x71\x13\x87\xa4\xd2\x9d\xd3\x05\x97\x52\x99\x55\x02\xcf\x8a\x2b\x38\x7c\x5c\x5c\xc1\xd3\xe2\x6a\x2f\x26\xa7\xbe\x63\x52\x33\x6f\x26\x7a\x16\xf7\x6a\x62\xa6\x95\xf9\x0a\x84\x7a\x3e\xa8\xa9\x2e\x43\xf4\x03\xc8\x08\xd3\xf9\x20\xf3\xbe\x70\x49\x1c\x0b\x69\x2e\x1d\x13\xda\x96\x32\xd5\x9c\x90\x09\x9b\xc7\xfc\x92\x5f\xc5\x5a\x2d\x5d\xdc\xbe\x73\x3c\x65\x87\x53\xf6\xa8\xbb\x33\xe1\xdc\x00\xe2\xdb\x0a\x31\x7e\x08\x67\x6b\x24\x52\x12\x1d\x3c\x8c\xdb\x02\x68\x35\x27\xc2\x1a\xcf\x95\x41\x02\xb6\x0e\x69\x58\x6a\x9c\xa0\x54\xde\xd2\x2d\xc5\xf4\xf4\xe9\xdf\x87\xe8\x04\xa9\xc2\x83\x23\x71\xef\x90\x52\xf4\x22\x8b\x1f\xb1\x29\xfb\xb5\x39\xb3\x5c\x19\x76\xe9\x06\x47\xb3\xb8\x81\xfb\xf7\xd8\x84\x5c\xf8\xf8\xf0\x09\x7b\xc2\x1e\x37\x97\xff\x15\x7c\x22\x6d\xfe\x1f\x1a\xb8\xf3\xb1\xf7\xe1\x67\x71\x3b\x85\x66\x61\xec\x6c\x2d\x4a\xb5\x06\xa1\xb9\x73\xf3\x41\xd7\xed\x83\xa3\x77\x9f\x2f\xe0\xa2\x6e\xfc\x99\x32\x45\xe9\x41\xc9\x3e\x1f\x0a\xcd\x05\x66\x56\x4b\xa4\xf9\x60\x27\xbc\xb2\xe8\x20\x43\xc2\x60\x59\xaa\x75\xcf\x46\x00\x68\x5d\x1b\x1c\x9d\x5a\x1e\xda\x85\x31\xd6\x97\xeb\xa7\x62\xcd\x09\x1c\x72\x12\x19\xcc\x61\xa3\x8c\xb4\x1b\xa6\xad\xe0\x5e\x59\xc3\x1a\xc6\xf3\x4e\xb0\xe0\xc4\x73\xf4\x48\x0e\xe6\xf0\xe3\xba\x61\x48\x2b\xca\x1c\x8d\x67\x2b\xf4\x27\x1a\xc3\xf1\x65\xf5\x56\x0e\xa3\x2e\x8e\x68\xc4\xd6\x5c\x97\x08\x73\x08\xd0\x7a\xe1\x2d\xf1\x15\x06\x85\xb7\x1e\xf3\x61\xd4\x3a\x9c\x5c\x6e\xfc\x45\xa3\xf1\xfc\x41\x0d\x9e\x96\x46\x04\x57\xa0\x16\xf9\x74\xfa\x3a\xd4\x24\xd2\x70\x7b\xfd\x18\x1c\x72\xa3\xae\x37\x84\x35\xce\x43\x8b\x02\xf3\x7f\xe0\x5c\xdb\x48\x1a\x3d\x84\x07\xdc\x06\xd9\x35\x5d\xf4\x42\x08\x2c\x7c\x94\x40\x14\x66\xa4\x6a\x52\x14\x5f\x3a\x6b\xa2\xf1\x4e\xea\x95\x35\x1e\x8d\x9f\x5c\x54\x05\xde\x2a\xdb\xf6\x6b\x6b\x4f\xa5\x30\x6c\x1d\x1e\xf5\xec\xdd\xc8\x94\xbb\x3b\x53\xe3\x2e\xdc\xde\x16\xb8\x2d\x80\xfb\x86\x70\x9f\x20\x6e\x48\xbf\x28\x7d\x66\x49\x7d\xaf\xf9\x51\xb2\x4b\xff\x6f\xf0\xe5\x25\x72\x42\x82\x9f\x7f\xb4\xc4\xeb\x2f\x90\x80\x29\xb5\xee\x10\xae\xf7\xd7\x21\xa1\x2f\xc9\x40\x3d\x7e\x86\x51\xfc\xad\x44\xaa\xa2\x71\x2f\x92\x1c\x7d\x66\x65\x02\x51\x61\x9d\xef\xf9\xb2\x8d\x7a\xdc\x5b\x86\xb2\x4a\xe0\xdd\xe2\xec\x03\x73\x9e\x94\x59\xa9\xb4\xda\x2b\x9d\x9d\xb0\x20\x94\x68\xbc\xe2\xda\x25\x10\x29\x23\x74\x19\xb6\x51\xeb\xdc\x88\xf9\x0c\xcd\xb0\x2b\xc8\x21\xa1\x2b\xac\x71\xd8\x7f\xb4\xad\xeb\x2d\x8b\x79\xbc\xf2\xc3\xee\x5d\xee\xc6\x78\x69\x65\xd5\xc7\xf1\x54\xdd\x78\xb9\x2d\x6e\x1d\x49\xc1\xc9\xe1\x4d\xcd\xdd\xc3\x5f\x83\xe0\x5e\x64\x30\x44\x22\x4b\xa3\xdb\x40\xfa\x9a\x3d\xc5\xce\xc7\x86\xd6\xdc\xe3\x18\xce\xd1\x48\x24\x98\xbd\xa9\x2b\xef\xd3\x29\xc4\x47\xa0\x8c\xb7\xe0\x33\xac\x13\xcc\x5a\xc9\x05\x62\x4d\x3c\x3f\x79\x71\xfc\xfe\x04\x94\xa9\x6f\xde\x16\xa0\x71\x8d\x1a\x6c\x0a\x3e\x53\x0e\x72\x2b\x4b\x1d\x18\xa0\x91\x93\x81\xdc\x12\x02\x5f\xda\xd2\xb7\x48\x99\xdd\x40\x65\x4b\x10\xdc\x80\x28\x9d\xb7\xb9\xfa\x8e\xd0\x79\xb0\xac\xa0\x20\xbb\x56\x61\xb4\x81\x54\x69\x8a\x84\xc6\x43\xdd\xc6\x0e\x2c\xb5\x30\xe1\xbf\x42\xc8\x33\xd7\x20\x32\xa5\x25\x60\x33\x01\x5c\xe3\xf2\x79\xd8\x12\xc7\x67\xef\x19\xd5\x21\x0e\xb7\x19\xa8\xc9\x4c\x10\x72\x8f\xdb\x91\x31\x6c\x4d\xf7\xab\x30\x6d\x46\x51\xb2\x37\x9a\x76\x05\xd3\x9e\xee\x9c\x41\x6d\x33\x47\xa3\x5a\x72\x9b\xf9\x9b\xbb\xa4\x59\x21\xb3\xb8\xf9\x8b\xfb\x57\x00\x00\x00\xff\xff\x45\xab\x31\x54\xfa\x0a\x00\x00")

func pluginsCodeampGraphqlStaticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsCodeampGraphqlStaticIndexHtml,
		"plugins/codeamp/graphql/static/index.html",
	)
}

func pluginsCodeampGraphqlStaticIndexHtml() (*asset, error) {
	bytes, err := pluginsCodeampGraphqlStaticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/codeamp/graphql/static/index.html", size: 2810, mode: os.FileMode(420), modTime: time.Unix(1533663930, 0)}
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
	"plugins/codeamp/graphql/schema.graphql": pluginsCodeampGraphqlSchemaGraphql,
	"plugins/codeamp/graphql/static/index.html": pluginsCodeampGraphqlStaticIndexHtml,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"codeamp": &bintree{nil, map[string]*bintree{
			"graphql": &bintree{nil, map[string]*bintree{
				"schema.graphql": &bintree{pluginsCodeampGraphqlSchemaGraphql, map[string]*bintree{}},
				"static": &bintree{nil, map[string]*bintree{
					"index.html": &bintree{pluginsCodeampGraphqlStaticIndexHtml, map[string]*bintree{}},
				}},
			}},
		}},
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

