// +build bin

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _defines_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5f\x6f\xdb\x36\x10\x7f\xef\xa7\x20\x84\x74\x58\x81\x46\xc1\x86\xee\xa5\x73\x04\xb8\xb1\xd3\x0a\x8d\x93\xc0\x7f\x02\xec\xa9\x50\x2c\x3a\xe6\x22\x8b\x1a\x49\xb9\xf3\x02\x7f\xf7\xf1\x28\x51\x7f\x2c\xd1\x96\x64\x67\xd8\x43\xf3\x50\x8b\xe4\x4f\xc7\xe3\xdd\xfd\x8e\x47\xaa\x2f\x2f\x3e\x5e\x90\x10\x23\x2b\xf9\xe5\x1f\x3f\x26\x0f\xdf\xe6\x51\xcc\xe8\x77\x6e\x6d\xb7\x2f\x2f\xdf\x89\x58\x22\xfb\xee\x61\x38\x1e\xbb\x83\x21\xf4\xd8\xf0\x0f\x0e\x38\x86\xdf\xb3\x2b\xf4\xf1\x12\x51\x86\xec\xab\x9b\xfe\x64\x72\xdb\x1f\x0d\x91\x35\x0f\x3c\x9e\xbc\xcd\xbc\xf0\x09\xa3\x33\xf2\x1e\x9d\xcd\x29\xc3\x80\xb5\x07\x9e\xf0\xec\xab\xfb\x99\x7d\x43\xb8\xd8\x6e\x7b\x82\xa1\x67\xbc\xb9\x94\xb2\x00\x62\xdf\x6e\xb7\x4e\x4f\xf8\x48\x09\xb9\xb4\x04\xfe\x5b\x9c\x33\xf2\xb4\x14\x28\xa4\xdf\x99\x17\x59\x4e\x01\xd9\xbb\x10\x7e\x3d\xda\x72\x7a\x3c\xf2\x42\x04\x2a\x6e\xb7\x99\xf4\x19\xc7\xec\x0a\xb0\x72\x96\x62\x1f\x88\x02\xbc\xd3\x5a\xe2\x64\xc3\x53\x81\xc8\x29\xf4\x75\x17\xe8\xfa\x01\xde\x55\x11\xfa\x76\x24\x5e\x08\x26\x87\x71\xe8\x2b\x77\x14\x7e\xde\xec\x75\xac\xf0\x1e\x03\x6c\x81\xd9\xe1\x21\x53\x0a\x1a\xbf\x20\x2e\x18\x89\xf0\x79\x82\x91\x4a\x2f\xb1\x07\xba\x33\x78\x84\x19\x97\xf0\xb0\xc7\x33\x60\xc9\x64\x4d\x29\x26\x0e\x89\x5c\xe6\xdb\x5c\xf1\x43\x12\xa4\xe5\x8e\x13\x00\x96\x3a\x28\x01\x6c\x07\x4f\xc9\xf2\x1e\xa9\xbf\x91\xa6\x14\x78\x15\x05\x9e\xd8\xc3\x07\x64\xab\x88\x53\x78\xf9\x0b\x56\x72\x0e\x1b\xdd\x5f\x3c\x6e\x04\xe6\xdf\x4e\x4b\x29\x9f\xf0\xe7\x9c\x4e\x83\x6b\x35\x47\x95\x52\x00\xb3\x07\x84\xdd\x7a\x2b\xfc\x15\x6f\xca\xdc\x2a\x10\x2a\xc5\xf1\x67\x00\x7e\x99\x8e\x6e\xaa\xdc\xaa\xa2\x99\x19\x5c\x8c\x72\xfd\x42\x7f\xed\x91\xa0\x19\x54\x06\x92\x34\xea\x4f\xe1\x23\x8f\x7e\xef\xf1\x38\xaa\x30\x45\xa3\xee\x31\x9b\xe3\x50\x14\xf8\xb2\x33\xb2\xdd\xe6\xae\x57\x82\x1a\xcc\x3e\xa5\xc2\xcb\x14\xed\xc2\x33\xed\x72\x4d\xb6\x03\x4e\x6d\xc7\xc5\x34\x80\xce\x6e\x72\xf7\xdf\x90\xf0\x19\x32\x4e\xce\x0c\x40\x63\x66\x39\x28\xff\xeb\x79\x68\xc9\xf0\x02\xac\x77\x23\x03\xe6\x7a\x62\x7f\x91\x4d\x99\xb6\xe4\x5f\x6e\xd8\x74\x48\xa7\x34\x84\x9c\x01\x5e\x93\x79\xca\xa9\x1a\xa0\xc7\xb0\x28\xa0\x33\x63\x7b\x15\xb6\x1e\xd6\x69\x74\x6f\xd4\x49\x0e\x15\x66\x19\xd1\x38\x14\xd8\xdf\x8d\x0a\x0d\x6c\xab\x13\x2a\x25\xe5\xb2\x4e\xfd\x87\xbe\x7b\xa3\xd5\x2a\x4f\x95\x0c\xe9\xd8\x53\xe1\x9d\x9a\xa9\x1e\x58\x50\xeb\x38\x9d\x66\x93\xe1\x20\xb3\x54\x79\x2a\x35\x94\x6d\x48\xc0\x23\xad\x52\x2d\xb0\x68\xa9\xe3\x74\x9a\xde\x4d\xfb\x06\x3b\x25\x43\x7a\xe5\x8a\x5d\xf5\x76\x4a\x81\xfb\xec\x94\x92\xaf\x6d\x0a\x2f\x25\xe1\xae\x79\x9c\x84\xd4\x7f\xf5\x44\x9e\x4c\xf2\xff\xce\xe4\xee\x82\x61\xdc\x10\x1a\x37\x4a\xe5\x00\x33\xe4\xf2\xc2\x50\x97\x64\xee\x2a\x7b\x1e\x97\xcd\x53\xc7\x77\xae\x9d\x76\x33\x60\x9a\x52\x0d\xf9\x51\x27\xb7\xc3\x14\x4c\x52\xce\x61\x9c\x4a\x03\x87\x61\x09\x33\xbb\x54\x48\x65\x6e\x74\xe3\x17\x69\x59\x27\xe9\xb3\xc5\x22\x67\x8f\x6b\x2a\x83\xc8\xc2\x2e\xf1\xc6\xc9\xfb\x1a\x45\xbd\xc4\x0e\x70\x20\x3c\x37\x6c\x0c\xbd\x8b\x45\x13\x6c\x33\x89\x05\x61\x5d\x22\x98\xec\xd4\x23\xed\x8b\x7f\x57\x46\x24\x5b\x78\xd5\xa0\xad\x96\xe0\x48\x10\x11\xe0\x4b\xeb\x93\x3b\x9d\xa0\x48\x86\x18\xc7\x73\x1a\xfa\xb2\x34\x0f\x6b\x0a\xf3\x1e\x71\x1e\x7b\x17\xc4\x89\x78\xe3\x1a\xdf\x38\x81\xb4\xd2\x69\x67\x10\x40\x08\xf4\xe9\x8f\xe9\x70\x82\x56\xd4\x8f\x03\x8a\x3e\x7c\x36\x2f\xe4\x13\x4c\xf3\xf6\xc3\xe7\x53\xcd\x63\x5c\x4f\xed\x44\x2d\x29\x4b\x8e\xdf\x11\xc9\x02\x33\x46\xd9\x91\x94\x4d\x64\xfc\xe0\xac\xc9\xb8\xff\x29\x69\xcb\x7c\xd5\xbb\x76\x25\x06\x3b\x70\x69\x87\xa8\xa7\x14\x5d\xa2\xa6\x51\x70\x17\x62\x96\xc9\xd8\x58\x74\x6b\x2a\x16\x79\xd4\x95\x8b\x91\x37\x7f\xc6\xe2\x48\x32\xa6\x42\x7e\xb0\xd1\x68\xde\x1f\x74\x6c\xb6\x9b\xbd\x1e\x29\x2b\xf2\x5f\x95\x9a\x25\x5a\x75\xe3\xe6\x0a\xaf\xda\x92\xf2\x3d\x3a\x93\x6f\xe5\xcc\x1c\x0d\x47\x55\x4e\x4a\x84\xfd\x95\xc0\xe4\x9a\x91\x79\xcf\x01\x46\x00\xf0\xba\xc9\x21\x0e\x80\x95\xdb\xb8\xbc\x3f\x3d\x98\xe9\x14\xd0\xe0\x50\x06\x2f\x1e\x7d\xc1\x26\x85\xbc\xc6\x45\xb6\xe5\x80\x4d\xf6\x8d\xd7\x9d\xa4\x4e\x73\x84\xd2\x41\xd2\x2d\xc4\x22\xde\x32\xc2\x22\x46\xe7\x79\x78\xdd\x4f\xd4\x44\xd5\x10\x03\x98\x7d\xef\x0e\x4c\x1f\x65\x2c\x07\x95\x50\xfb\x9c\x9f\x21\x67\x45\x64\xf1\x3a\x30\x47\x70\xcc\x0e\x6f\x2b\xf9\xd4\x8c\x50\x46\xc4\xa6\x21\xfc\x56\x1e\xbe\x1b\x42\x27\xe4\x9f\xa6\xd0\x31\xe6\xc4\x57\x97\x14\xb5\x70\xa0\x09\x9c\xee\x35\x7c\x4a\x56\x35\x92\xf5\xd5\x4c\xc9\x1a\xbb\x9b\x6c\x17\xc6\x44\xbc\xfb\x65\xf4\xaf\x27\xb9\x8c\x2e\x5d\x35\x94\x6f\x7e\xef\x27\x32\x78\x6a\x2f\x34\x93\x91\xec\x86\x52\x36\xb2\xeb\xcc\x3a\xdc\x11\xd7\x99\xbb\x1a\xcd\x8c\x1a\xcd\x8a\x1a\xcd\xcc\x1a\xcd\xda\x6b\x54\xba\x1c\xaf\x68\x34\x19\x8e\xeb\x2e\x58\xd3\x11\x7d\x6f\x06\x8d\xda\x3b\x5f\x8d\xeb\x7e\x0d\x5d\xf1\xda\xd8\x35\x79\x4d\x8e\xe4\x5e\xcb\xf5\xa9\xc5\x9d\xd0\x6b\xb7\xee\xd5\xb0\xde\x46\x6a\x44\x2f\xfb\xd6\x35\x6a\x94\xe0\x4e\x67\xa3\x07\x77\x3c\xad\xd7\x48\x8d\xe8\x49\xa0\x61\xf0\x5a\x82\x3b\x9d\x46\xe3\xe1\xc4\xe0\x35\x18\xc9\x1c\x21\x1b\xa6\xc8\x56\xb8\x2e\x5e\xd3\x49\x70\x47\xa3\xa9\x3b\x32\x78\x4d\x8d\x64\x5f\x0e\x64\xc3\x60\xa3\x04\xd7\xca\x46\x7b\xb9\x06\x39\xd1\x10\x47\xfd\x82\x46\x57\x77\xa3\x51\xff\x76\x50\xfd\x0e\xa5\x71\xaf\xf0\x29\x23\xdd\xea\xbb\xd5\x09\x82\xd2\x40\x90\x48\xd7\x4e\x89\x42\xc5\xaa\xb9\x08\x40\xbe\xcc\xe7\xe7\xf2\x8c\x20\xa4\xdb\x2e\xad\xeb\x38\x08\xd2\x3e\xa9\xda\x1c\xaf\x54\xaf\x17\x0b\x8a\x02\xbc\x50\xd5\x9d\x3d\x59\x52\x26\xf2\xff\xa9\x60\x90\xec\x53\xc1\x91\xaa\xcf\xdb\x4e\x61\xdb\xf6\x9b\x37\xf9\x37\x00\xf8\x39\xbc\xe8\x98\xe3\x28\xa9\x53\x1b\x6c\x7f\x45\x57\x5a\x72\x45\xa9\xf7\xd4\xea\x1e\xbc\x20\xc6\xf9\x47\x88\xc3\x33\xaf\x9f\xd4\xc9\xde\xca\x6b\x2a\x4b\x77\x25\xbb\x3e\x0d\x40\xd2\xa5\xf5\x9b\x92\xaf\xf6\xcf\x07\xef\x49\xd6\x68\x62\x08\xa8\x9a\x1d\x7f\xef\x6c\x59\x11\x98\x56\x79\x2b\x6f\xbe\xcc\x77\xe6\x54\xf2\x48\x76\xc2\x7b\x35\x07\x0a\x39\x62\xcf\x66\xba\xdc\x73\xf2\xaa\x2c\x1b\x29\x57\x65\xbb\x08\xa8\x54\xcc\xa3\xf7\x8c\xae\x65\x89\xc4\xcc\x08\x59\x86\x8a\xc3\x1f\xad\x14\x34\x5d\xcc\x82\x04\x92\x12\x9e\x58\xee\x2f\x90\xf6\x5a\x8d\xcb\x1a\x66\xbe\x6c\x52\x3c\x93\x05\xc2\x7f\x21\x4b\xb0\x58\xb2\xe3\xe7\x3f\x39\x0d\x51\xd5\x69\xfe\x3b\x80\x9a\x09\xac\x23\x00\x95\x44\xef\xc3\x67\x84\x6f\x5b\xf5\xad\x9f\x76\x8e\x49\xc4\x87\x08\x4c\x3a\xdb\x9c\x99\x48\x76\xea\x71\x42\xe9\xe3\xac\x11\xa5\x2e\xcd\x3a\x38\x78\x30\x6b\xf9\x84\xe1\xb9\xa0\x6c\xd3\xe9\x48\x94\x79\xa6\x65\xae\x93\x69\x9b\x61\x0e\x0e\xed\xf9\x64\xad\x97\x49\xc2\x28\x16\xe7\xd2\x96\x71\x84\x0a\xcf\xe7\x7c\x85\xd2\x17\x92\x0e\xab\x9c\xb6\x8a\x50\xcf\xf7\x69\x68\x39\xe3\x04\xae\xd3\x90\x42\x68\xf8\x82\xb2\x95\xca\x67\x8c\x06\x99\x5c\x85\xb0\x90\xd8\x44\x38\x39\x0d\x58\x48\xa5\xb7\x25\x0d\xa4\xed\x54\x9a\x19\xe0\x85\x17\x07\x22\x0d\x32\xb9\x1c\x1d\x5d\x49\xff\x3b\xdd\x96\x07\x6e\x42\x21\xc2\x2c\xb4\x86\x6c\xa4\x5e\x4e\x3a\xb3\x80\x90\x2b\xb8\x90\x0b\xcf\x4c\xf4\x6f\x00\x00\x00\xff\xff\x9e\x3a\x0a\x51\xea\x27\x00\x00")

func defines_html_bytes() ([]byte, error) {
	return bindata_read(
		_defines_html,
		"defines.html",
	)
}

func defines_html() (*asset, error) {
	bytes, err := defines_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "defines.html", size: 10218, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\x7b\x6f\xdb\xc6\xb2\xff\x3f\x9f\x62\x2f\x93\x16\xb7\x40\x45\x21\xb9\xbd\xb8\x45\x6a\x0b\x50\x2c\x39\x11\xea\x87\x60\xc9\xbe\xa7\x7f\x15\x94\xb8\x92\x36\xa6\x48\x96\xbb\x92\xe3\x06\xfa\xee\x67\x66\x1f\x7c\x89\x2b\x51\xb4\xdc\xe3\x16\x09\x10\x93\xdc\x9d\x9d\x9d\x9d\xc7\x6f\x67\x96\xb4\xbf\x7e\xf5\xe9\x8c\x85\x94\x38\xea\xca\xdf\xbf\x57\x37\xbf\x4f\xe3\x55\x12\x3d\x70\x67\xb3\xf9\xfa\xf5\x81\x89\x05\x71\xaf\xef\xfa\x37\x37\x83\x5e\x1f\x5b\x5c\xfc\x41\x03\x4e\xf1\xfa\xe6\x8c\xbc\x3f\x25\x51\x42\xdc\xb3\x8b\xee\x68\x74\xd5\xbd\xec\x13\x67\x1a\x78\x5c\x8d\x4e\xbc\x70\x4e\xc9\x1b\xf6\x23\x79\x33\x8d\x12\x8a\xb4\x6e\xcf\x13\x9e\x7b\x36\xbc\x75\x2f\x18\x17\x9b\xcd\x89\x48\xc8\x3d\x7d\x3c\x05\x5e\x48\xe2\x5e\x6d\x36\x9d\x13\xe1\x13\xc9\xe4\xd4\x11\xf4\x8b\x68\x25\x6c\xbe\x10\x24\x8c\x1e\x12\x2f\x76\x3a\x39\xca\x93\xb6\xf0\xab\xa9\x9d\xce\x09\x8f\xbd\x90\xa0\x88\x9b\x4d\xca\xfd\x96\xd3\xe4\x0c\x69\x61\x96\x7c\x1b\xb2\x42\xfa\xce\xc1\x1c\x47\x8f\x5c\x33\x24\x9d\x5c\x5b\x73\x86\x03\x3f\xa0\x65\x11\xb1\xad\xc4\xb1\x2d\x12\xe8\xa6\xa1\x2f\xcd\x91\xbb\xbc\xda\x69\x58\xe1\x4d\x02\xea\xa0\xda\xf1\x26\x15\x0a\x1f\xde\x12\x2e\x12\x16\xd3\x96\xa2\x01\xa1\x17\xd4\x43\xd9\x13\xbc\xc5\x19\x17\x78\xb3\xc3\x32\xa8\x49\xb5\x26\x4d\xb3\x0a\x19\x2c\xf3\xbb\x4c\xf0\x7d\x1c\x40\x73\x4f\x63\x80\x9a\xda\xcb\x01\x75\x87\x77\x6a\x79\x93\xc8\x7f\x04\x55\x0a\xba\x8c\x03\x4f\xec\x88\x07\xe2\x4a\x8f\x93\xf4\x70\x45\x2d\x75\xf6\x2b\xdd\x9f\x4d\x1e\x05\xe5\xbf\x1f\x37\xa4\x7c\xc6\xef\xb3\x70\xea\x9d\xcb\x39\xb6\x43\x0a\xc9\xdc\x1e\x4b\xae\xbc\x25\xfd\x95\x3e\x16\x63\x2b\x17\x50\x9a\x8e\xdf\x23\xe1\xa7\xf1\xe5\xc5\x76\x6c\x6d\x53\x27\x76\xe2\xbc\x97\x9b\x01\xdd\xb5\xc7\x82\x7a\xa4\xe0\x48\xa0\xd4\xef\xc3\x09\x8f\x7f\x39\xe1\xab\x78\x2b\x52\x0c\xd5\x90\x26\x53\x1a\x8a\x5c\xbc\x94\x7a\x36\x9b\xcc\xf4\x92\x51\x8d\xd9\xc7\x91\xf0\x52\x41\x9b\xc4\x99\x31\xb9\x09\xb6\x3d\x46\x3d\x2c\x16\xb5\x03\xbd\xb9\xc8\xcc\x7f\xc1\xc2\x7b\x44\x9c\x2c\x32\x90\x9a\x26\x4e\x87\x64\xff\x4e\x3c\xb2\x48\xe8\x0c\xb5\x77\x01\x0e\x73\x3e\x72\x3f\xc1\x23\xc0\x16\xfc\xcb\x14\xab\xbb\x0c\xa4\x11\xd2\xe9\xd1\x35\x9b\xea\x98\xaa\x20\xf4\x12\x2a\x72\xd4\xa9\xb2\xbd\xad\x68\xdd\x2f\xd3\xe5\xd0\x2a\x13\x74\xe5\x66\xb9\x8c\x56\xa1\xa0\x7e\xd9\x2b\x0c\xe1\xa1\x32\x91\x02\x28\x17\x65\xea\xde\x75\x07\x17\x46\xac\xe2\x54\xaa\xcb\xf8\x9e\x74\x6f\xad\xa6\x6a\xc2\x9c\x58\x4f\x93\xe9\x76\xd4\xef\xa5\x9a\x2a\x4e\x25\xbb\xd2\x0d\x09\xe3\xc8\x88\x54\x49\x98\xd7\xd4\xd3\x64\x1a\x5f\x8f\xbb\x16\x3d\xa9\x2e\xb3\x72\x19\x5d\xd5\x7a\xd2\x84\xbb\xf4\xa4\x83\xef\x50\x08\x2f\x80\x70\x53\x1c\x67\x61\xe4\x3f\x3b\x90\xab\x49\x5e\x36\x92\x0f\x66\x09\xa5\x35\x49\x57\xb5\xa0\x1c\xc9\x2c\x58\x9e\xeb\x6a\x02\xe6\x03\xa9\xcf\xa7\xa1\xb9\x36\x7c\xe3\xdc\xa9\x8c\x80\x1a\x52\x2d\xf8\x68\xc0\x6d\x7f\x08\x2a\xc8\xd9\x4f\x27\x61\x60\x3f\x99\x8a\xcc\x26\x19\x52\x31\x36\x9a\xc5\x17\x3b\x30\x4f\x32\xb5\xc5\x2c\x8b\x9e\x81\x2d\x0d\x62\x33\xb7\x10\x37\x9d\xac\xad\x96\xd7\x03\x6d\x8f\x06\xc2\x1b\x84\xb5\x49\xaf\x57\xa2\x0e\x6d\x3d\x8e\x39\x66\x4d\x3c\x98\x95\xf2\x91\xc3\x93\xff\x01\x78\x64\x32\xf3\xb6\x9d\x76\x3b\x05\x27\x82\x89\x80\x9e\x3a\x1f\x06\xe3\x11\x89\xc1\xc5\x38\x9d\x46\xa1\x0f\xa9\x79\x58\x91\x98\x9f\xb0\xce\xe4\xa4\xcd\x3a\x31\xaf\x9d\xe3\x5b\x27\x00\x2d\x1d\x77\x06\x81\x01\x41\x3e\xfc\x36\xee\x8f\xc8\x32\xf2\x57\x41\x44\x7e\xfa\x68\x5f\xc8\x07\x9c\xe6\xbb\x9f\x3e\x1e\x6b\x1e\xeb\x7a\x2a\x27\x3a\x30\x64\xd9\xd3\x77\x44\x36\xa3\x49\x12\x25\x4f\x0c\x59\xc5\xe3\x5b\xcc\xda\x94\xfb\x97\x06\x6d\x31\x5e\xcd\xae\xbd\xe5\x83\x0d\x62\xa9\x14\xa8\xc7\x64\x5d\x08\x4d\x2b\xe3\x26\x81\x59\x0c\xc6\xda\xac\x0f\x0e\xc5\x7c\x1c\x35\x8d\xc5\xd8\x9b\xde\x53\xf1\xc4\x60\xd4\x4c\xbe\x45\xa3\x55\xbd\xdf\xc2\xb1\xde\x6e\xf6\x7c\x41\xb9\xc5\xff\x59\x43\xb3\x10\x56\xcd\x62\x73\x49\x97\x87\x06\xe5\x8f\xe4\x0d\x8c\xca\x22\xf3\xb2\x7f\xb9\x1d\x93\x40\xe1\xfe\xca\x70\x72\x13\x91\x59\xcb\x9e\x88\x40\xc2\xf3\x3a\x45\x1c\x12\x6e\x9d\xc6\x65\xed\xba\x30\x33\x10\x50\xa3\x28\xc3\x81\x4f\x3e\x60\x03\x26\xcf\x71\x90\xed\x74\x50\x27\xbb\xfa\xab\x2a\xa9\xe3\x94\x50\xc6\x49\x9a\xb9\x58\xcc\x0f\xf4\xb0\x38\x89\xa6\x99\x7b\x0d\x47\x72\xa2\x6d\x17\x43\x32\x77\x38\xe8\xd9\x5e\xca\x38\x1d\x52\xa0\xda\x65\xfc\x94\xf2\x36\x4f\x99\x3f\x0e\xcc\x28\x38\x4d\xf6\x6f\x2b\xd9\xd4\x09\x8b\x12\x26\x1e\x6b\x92\x5f\x41\xf1\x5d\x93\x74\xc4\xfe\xac\x4b\x7a\x43\x39\xf3\xe5\x21\x45\x25\x39\x86\x09\x56\xf7\x86\x7c\xcc\x96\x15\x9c\xcd\xd1\x4c\x41\x1b\xe5\x4d\xb6\x49\xc4\xc4\xbc\xf9\x61\xf4\xbb\xa3\x1c\x46\x17\x8e\x1a\x8a\x27\xbf\xc3\x11\x38\x4f\xe5\x81\xa6\xea\x49\x4f\x28\xe1\x21\x3d\xce\xac\xa2\x7b\xc2\x71\x66\x59\xa2\x5b\xab\x44\xb7\x79\x89\x6e\xed\x12\xdd\x1e\x2e\x51\xe1\x70\x7c\x4b\xa2\x51\xff\xa6\xea\x80\x55\xf7\x98\x73\x33\x7c\xa8\x3c\xf3\x35\x74\xcd\x8f\xa1\xb7\xac\x76\x33\xb0\x59\x0d\x7a\x32\xab\x65\xf2\x54\xd2\x1d\xd1\x6a\x57\x83\xb3\x7e\xb5\x8e\x64\x8f\x59\xf6\xd5\xc0\x2a\x91\xa2\x3b\x9e\x8e\xee\x06\x37\xe3\x6a\x89\x64\x8f\x99\x04\x1f\x2c\x56\x53\x74\xc7\x93\xe8\xa6\x3f\xb2\x58\x0d\x7b\x52\x43\xc0\x83\xcd\xb3\x25\x5d\x13\xab\x19\x10\x2c\x49\x34\x1e\x5c\x5a\xac\x26\x7b\xd2\x37\x07\xf0\x60\xd1\x91\xa2\x3b\x48\x47\x3b\x63\x0d\x31\xd1\xe2\x47\xdd\x9c\x44\x67\xd7\x97\x97\xdd\xab\xde\xf6\x7b\x28\x43\xf7\x0c\xaf\x32\xf4\x56\xdf\x2c\x4f\x10\x51\x14\x08\x16\x9b\xdc\x49\x09\x94\xcf\x9a\xf3\x04\xc4\x07\x3c\x6f\x41\x8d\x20\xc0\x6c\xa7\xce\xf9\x2a\x08\x74\x1b\x88\x36\xa5\x4b\xd9\xea\xad\x44\x44\x02\x3a\x93\xd9\x9d\x3b\x5a\x44\x89\xc8\xbe\x54\xb0\x70\xf6\x23\xc1\x89\xcc\xcf\x0f\x9d\xc2\x75\xdd\x57\xaf\xb2\x77\x00\x78\xd9\xbf\xe8\x15\xa7\xb1\xca\x53\x6b\x6c\x7f\x79\x53\x3a\xb0\x22\x6d\x3d\xb9\xba\x3b\x2f\x58\xd1\xec\x25\xc4\xfe\x99\xd7\x73\x59\xd9\x3b\x59\x4e\xe5\x98\x26\xb5\xeb\x47\x01\x72\x3a\x75\xfe\x57\xf2\x97\xfb\xe7\x9d\x37\x87\x1c\x4d\xf4\x91\xaa\x62\xc7\xdf\x39\x5b\x9a\x04\xea\x2c\x6f\xe9\x4d\x17\xd9\xce\xac\x39\x5f\x42\x23\x8e\xab\x28\x28\xa0\xc7\xbd\xbd\x35\xe9\x5e\x27\xcb\xca\xd2\x9e\x62\x56\x56\xa6\xc0\x4c\xc5\xde\x3b\x4c\xa2\x35\xa4\x48\x89\x9d\x02\xd2\x50\xb1\xff\xa5\x95\x24\xd5\x8b\x99\xb1\x00\x42\xc2\x13\x8b\xdd\x09\xd2\x4e\xad\x71\xc8\x61\xa6\x8b\x3a\xc9\x33\x9b\x11\xfa\x07\x71\x44\xb2\x82\xe8\xf8\xef\xcf\x3c\x0a\xc9\xb6\xd1\xfc\x1f\x90\xd4\x1e\xc0\xc6\x03\x48\x81\xf5\x2e\xfa\x34\xe0\x0f\xcd\xfa\xd6\xf3\x52\x99\xc4\x7c\xf4\x40\xd5\x78\x48\xcd\xc4\xd2\xaa\xa7\x13\x82\x8d\xd3\x87\x58\x9b\x34\x6d\xe0\x68\xc1\xf4\xc9\x67\x09\x9d\x8a\x28\x79\x6c\x54\x12\xa5\x96\x39\x10\xeb\x00\xb6\x13\xca\xd1\xa0\x27\x3e\x5b\x9b\x65\xb2\x30\x5e\x89\x16\xe8\x72\x15\x93\xdc\x7d\x8b\x2f\x89\x1e\xa0\x1a\x9c\x22\x6c\xe5\x49\x3d\xdf\x8f\x42\xa7\x73\xa3\xc8\x0d\x0c\x49\x0a\x43\x3e\x8b\x92\xa5\xc4\xb3\x24\x0a\x52\xbe\x92\xc2\x21\xe2\x31\xa6\xaa\x1a\x70\x88\x84\xb7\x45\x14\x80\xee\x24\xcc\xf4\xe8\xcc\x5b\x05\x42\x3b\x19\x2c\xc7\x78\x97\x6a\xff\xc1\x3c\x43\xc1\xcd\x22\xf4\x30\x87\xac\x11\x8d\xe4\x60\xd5\x98\x3a\x04\xac\xa0\x0d\x0b\x4f\x55\x74\xf2\x5f\xbd\xeb\xb3\xf1\x6f\xc3\x3e\x59\x88\x65\xd0\x39\xd1\x3f\xa5\x15\x96\x54\x78\x64\xba\xf0\x12\x4e\x01\x69\x57\x62\xd6\xfa\xd9\xd1\xad\x0b\x21\xe2\x16\xfd\x63\xc5\xd6\xa7\xce\xbf\x5a\xb7\xdd\xd6\x59\xb4\x84\x48\x63\xca\x77\x0c\x64\x0f\xfa\xa7\xd4\x9f\x53\x33\x08\xbd\x03\x5c\x8c\xd1\x87\x18\x76\x82\x1c\xdd\x03\xf3\xc5\xe2\xd4\x97\x6f\x3e\x5b\xf2\xe1\x47\x30\x03\x13\xcc\x0b\x5a\x7c\xea\x05\xf4\xf4\x2d\x7a\x1d\x1e\xf1\xa4\x50\xf8\x29\xe2\x22\x94\x90\x42\xe0\x0e\xd8\x80\xf1\x25\xc1\x49\x00\x45\x06\xa8\x37\x00\xfb\xc0\x0c\x46\xb5\x6c\xe9\xcd\x69\x3b\x0e\xe7\x8e\xda\xcc\x9d\xf6\xcc\x5b\x23\x81\x8b\x6d\xf9\x51\x5c\x3c\x06\x60\x19\x4a\x0b\x66\x69\x4f\x61\x0f\x30\x43\xe1\xbe\xcd\x42\x9f\x7e\x71\xb1\x15\x9c\x62\x0a\xf1\x21\xf2\xd4\x9f\xbd\xb5\xa7\x5a\x9d\xb2\x0a\x09\x4f\xa6\xc0\xe3\x33\x6f\x27\xa8\xc1\x84\xc2\xdd\x3b\xf7\xad\xfb\xf6\xff\x4c\x83\xbb\x64\xa1\xfb\x99\xeb\xdd\x6e\xe9\xb1\x50\xd1\x4b\xfb\xbb\xe3\xee\xc7\x8f\xfd\xde\x84\x85\x9b\x0d\xd0\x69\x39\xd4\x08\x03\x19\x30\x43\x7b\xc9\x82\x7b\xdd\x29\x3b\x8c\xf5\x95\x54\x70\xa3\x8c\xac\x02\x07\x63\x01\xe3\x7f\x1a\x30\xdc\x0a\xb5\x93\x9c\x84\x5e\x1a\x22\x70\x3b\xf1\x12\xa2\x2e\x2d\x5f\xb9\x9e\x79\x9c\xb1\x2f\xd4\x6f\x89\x28\x76\x08\xf8\x36\x95\xd4\x6c\x0e\xfe\x80\x21\x91\x0f\x34\x34\x39\x2c\x87\xc2\x90\x60\xc5\xfc\x62\xa7\x66\x66\x52\xb0\x8a\xae\x09\x00\xa9\x2f\x53\x44\x14\x76\xa1\x5d\x20\x35\xcb\x76\x7a\x30\x89\x84\x88\x96\xba\x5d\x44\xf3\xb9\x3c\xdf\x8c\xe2\x68\x0d\x33\xe8\x56\x48\x85\xe7\x18\x69\x0b\x6c\x24\xb3\x68\xba\xe2\xb9\xdc\x43\x4a\x7b\xea\xbc\xf6\x19\x82\xe0\xa4\x55\x1c\x9c\xba\x70\xb2\x0a\x43\x16\xce\x89\x71\xcf\xde\x60\x34\xbe\x19\x7c\xd0\xc9\x41\xc9\x63\x55\x9a\x27\x35\x2c\xe1\x04\x17\x53\xe6\x9f\xa5\x84\x92\xce\x18\x48\x2b\x02\x92\x83\xc0\x8b\x79\x06\xd4\xa6\x81\x94\x09\xb2\xa1\xab\x58\xb0\x25\x6d\xc5\x90\x77\x86\x62\x6b\xa0\xd3\x51\xfd\xf9\xd3\xd2\x6c\x54\xb6\x8a\x5b\xf9\x9c\xfb\xe0\x54\x8a\xb7\x0a\x72\x96\x32\x32\xc0\x45\x06\x96\xe4\xc3\x62\xa7\x64\x4c\x09\x76\x29\xdf\xc1\x10\x79\x06\x2c\x1d\x10\x78\xbb\x07\x5c\x74\xcd\x80\xf6\x2a\xb0\x4b\x60\x6e\xcd\x97\x49\xc0\xbd\x8a\x6b\x6e\x7c\x00\x96\x00\x58\x0e\x70\xe3\x28\xad\x24\x2d\x4e\x9c\xd7\xe8\xec\x1d\xfc\xa9\x8c\xa9\x04\xcf\xf5\xc7\x80\x0b\xaf\xf1\x84\xbc\xb2\x77\x0d\x90\xf3\x7a\xad\x12\x83\x8c\x44\x2e\xc4\xe8\x40\x4b\xe3\x4d\x05\x5b\xd3\xac\x2c\x72\x10\x7a\xf9\xfb\x76\xfb\xe1\xe1\xc1\x05\x8f\x4a\xe0\xbf\x3b\x8d\x96\x6d\x85\x82\x00\x21\x01\xf5\x38\xe5\x6d\xdc\x37\xb9\x48\x4f\xc6\xc1\xa7\x38\x84\x63\xea\xa0\x90\xc6\x8c\x06\xd7\x57\xe8\xa0\xd7\xa3\x71\xff\x6a\x5c\x16\x43\xda\x35\x73\xbe\x9d\x01\x8c\xf6\x02\x98\x4e\x44\xd9\xbb\x20\x3b\x21\x25\x57\x34\xeb\x42\x72\xa2\x06\xf9\x98\x91\x26\xe6\x81\xf1\x25\xe3\x5c\xa7\x1a\x93\x15\x04\x70\xba\xe1\x4e\x83\x08\x7d\x5e\xe1\xac\xea\xd2\x81\xa8\x47\x69\xbe\x0e\xf1\x12\xe6\xb5\x16\xcc\xf7\x29\x60\xa7\xcc\xc9\x3a\xdf\xa3\xe3\xf2\x5f\x4e\xda\x6a\x60\x2e\xf6\xd4\xc4\xd0\xc9\x3d\xdc\xad\xba\xf8\x58\xf4\xef\xc2\x4f\xf0\x86\x6c\xd9\x12\xf4\xaa\x95\x43\x62\x48\x0a\x00\x14\x10\x18\x39\x5d\x83\x5e\x1e\x21\xf3\x28\x68\x01\xf4\x53\x6c\x98\x04\xd1\xf4\x1e\x55\xd6\x5a\xfa\xad\x77\xe6\x26\x9a\xcd\x60\x03\x69\xbd\x2d\xd2\x82\x7c\x34\xc0\xb6\xc0\x9b\xd0\xa0\x1c\xd0\x2d\xd5\x2a\x89\x24\xa6\x22\x3c\x4d\x44\xd8\x92\x53\x38\x00\x74\x00\x6b\xfe\x6a\xb9\x7c\x5c\xd2\x25\x48\x3e\x63\xf3\xe9\x82\x4e\xef\x27\xd1\x17\xd9\xd9\xe2\xb8\x19\xbe\x4e\x3b\xa5\x0b\x1a\x29\x81\x8d\x82\x69\xc5\x2d\xcd\x7e\x67\x5e\xc0\x4b\xe9\xef\x99\xdc\x52\xdc\x4f\x90\x0a\x2a\x46\x97\xfd\x4b\xc8\x51\x88\x72\x6c\xb3\x31\x99\xc0\x80\xe9\x9c\xce\x25\x5d\xca\xbc\xd0\x33\x09\x94\x04\x48\x8b\xa8\xe5\x65\x67\x3d\xca\x4d\xb2\xe7\x6d\x9f\x20\xea\x09\xfd\x1e\x95\x55\xca\xd7\x9e\xc6\xb2\x00\xdb\x99\x16\x0d\xf3\x1d\x05\x83\x45\x63\xa9\x38\x6a\x86\x6c\x4b\xc7\xbc\x32\x3d\xbf\x88\x12\xf6\x27\xba\x61\xd0\x92\xcd\x13\xa8\x39\xd0\x4a\x72\x1f\x94\x4d\x45\x1f\x92\x39\xa9\xca\x60\xd1\x5b\x69\x69\x4b\x46\x43\xab\xb4\x38\xbd\x83\xa4\xb8\xb4\x9f\xaa\x78\xe2\x65\x47\x84\x11\x72\x94\xc9\x15\x8a\x8e\x59\xf2\x31\xc7\xe8\xbe\xa4\xe1\xce\x27\xb9\xd8\xd4\x40\xbb\x66\x30\xe9\xf1\x1e\x2f\x1c\xfd\x7f\x77\x68\x77\x40\x50\x2e\xe5\x0f\x58\x4d\xda\x05\xa2\x04\x29\xf0\xb0\x41\x4b\x65\x01\xcb\x0a\xe5\x9a\x32\x32\x2f\x94\xae\x19\xc0\xca\x96\x62\xcf\x14\x2d\xb9\x32\xcf\x4c\x89\x53\x6c\xcf\xaf\x5d\x6e\x2b\x38\x6a\xc6\xa9\x8e\x50\x96\xf3\xb1\x14\x2b\x81\xad\xa9\x03\xf7\xbc\xb2\xd2\x95\xa4\x5b\xce\x15\x2c\xca\x2a\x20\xdf\xff\x3c\x0f\xd2\xb1\x99\x1d\xe8\x4c\xdf\xd1\x70\x6e\x70\x6e\xf7\x32\x36\xcb\x12\x19\x3d\x70\xec\x4d\x06\xe7\xee\x18\x77\x6b\x9d\x21\x96\xb0\x6f\x5b\xf8\xbf\x05\xf4\xa5\x7a\x6d\x82\x7c\x52\x87\x4d\x80\xef\x3f\x86\x72\x20\x23\x93\x05\xe5\x71\xe0\x6c\xb7\xb3\xf5\xbf\x80\xd3\xfb\x18\x67\xca\xd9\x20\x0b\x92\x67\xa8\x55\xee\x66\x11\xa8\xe4\x85\x8a\x23\xe6\xc4\x83\xf3\xcd\xe6\xd8\x10\x87\x3c\x9b\x21\xdc\x56\x7a\x8f\x20\xc4\x0b\xf9\x3c\x9b\xb5\xcc\x99\x90\x32\x9c\x37\x41\xff\x4b\x0b\x8c\xf3\x71\xf7\xc3\xc8\x65\xe7\xc3\xee\xd9\xaf\xfd\xf1\xc8\xbd\x65\xf8\x42\x36\x97\x59\xbf\x4e\xbf\xe8\x70\x3a\x43\x75\x93\x4f\xdf\x0f\x9d\xa7\x7f\x73\x73\x7d\x53\x3d\x8d\xfa\xa6\xcb\xe9\xc8\x33\xc1\xdd\x93\x68\x00\xd9\x33\x97\xfc\x3e\xb4\x72\x2a\xf9\x25\xa7\xd3\xf9\x80\x97\x3d\x29\xbe\x8a\x56\xa3\x82\x9c\x24\x30\x2b\x29\x6f\x21\x79\x6f\xac\x54\xed\x0f\x55\x0e\xab\x40\x4e\x76\x17\x36\x18\xbd\x3a\x55\xaa\x58\x2d\xa6\xc1\xf1\x50\xfb\xa6\x7b\x57\xba\xb8\x3a\x3b\x58\xf9\xb3\xae\xed\x8d\x2c\xd3\x99\xb6\x67\x13\x95\xe5\xbc\xe4\x88\x1a\xd3\x5c\xeb\x29\xac\xec\xa8\x39\x7d\xa9\x85\xd5\x53\x57\xe1\x9b\xd4\x5d\xda\x52\x2e\xd9\x44\x59\x99\x9b\x1f\x51\x57\x8a\x69\x3d\x55\x95\xe2\x2c\xa7\x29\xb9\xa8\x7a\x8a\xca\x7f\x70\xdf\x2c\x3d\x7a\xf7\x3c\xe9\xd1\x34\x5e\xd9\xf3\xa3\xb4\xf3\x68\x09\xd2\xd9\xf0\x16\xec\x54\x9d\x20\xc1\x6c\x4e\x07\x08\xaa\x12\xa1\x0a\x31\x5f\x70\x26\xa4\xce\x55\x53\xe5\x35\xc9\x82\x94\xa2\xfe\xb9\xf5\x9f\x34\xf6\x5f\x9c\x30\x29\x9d\x5a\x33\xa6\x5d\x22\x59\x53\x26\xe0\x79\xfc\x9c\x49\x32\x7d\x4a\x59\x98\xf3\xc1\xa6\x05\xa1\x56\x56\x65\x41\x08\x6c\xeb\xc0\x5e\xfa\xf7\x05\x5e\x54\x41\xe8\xef\x28\x08\xfd\x63\x17\x84\xbd\x73\x2b\xdc\xf9\x95\xf5\x60\x6f\x67\x3d\xb8\x2d\xfb\x0b\x46\xc1\xcc\x0d\xfd\xa7\xd4\x83\xbd\x6f\xf5\x60\x3d\x78\xeb\xed\xac\x07\xfd\x06\xf5\x60\xef\x19\xea\xc1\xde\x73\xd6\x83\xfe\xee\x3a\xad\xa7\x72\x3a\xff\x7c\x70\x75\xdd\xab\x2c\x9e\xcc\x6f\xae\xe2\x6f\x29\xe0\xb5\xaa\x4e\xf3\x6b\xd5\x69\xe9\x5c\xb6\x3a\xcd\x3f\xb4\x4e\x4b\x65\xcb\x09\xb2\x2f\x8d\xae\x5a\xb1\x2d\x8f\xee\xd5\xce\xa3\xcb\x5c\x77\x26\xd2\x56\xa5\x67\xef\x41\xf5\xc2\xea\xec\x29\xa5\xdf\xbe\xb6\xd7\x1c\x7e\xa9\xe6\x38\x44\x59\xfb\x6b\x8e\x06\xba\xaa\x51\x73\xd8\x7c\x26\xa7\xa9\xda\x35\x87\x7f\x40\xcd\x71\xd0\x8b\xbf\x5a\xaf\xb1\x7e\x2e\xbd\xc6\x7a\xa6\xea\x25\xe6\xf6\xbd\xdc\xf4\x1d\x6d\x2f\x1f\x8e\xac\x7b\x39\xbe\xf9\x1d\x26\xd1\x94\x72\xae\x43\xb9\xb4\x71\x6f\x0b\xfa\x82\x37\x6e\x29\x76\xaa\xbe\x26\x7b\xb6\x54\x55\xad\x3d\x5b\xbd\x74\xb7\x15\x2e\xc5\xbf\x30\xd2\x78\xcf\x19\x8e\x0e\xde\x73\x5e\x5a\xbd\x84\x1e\xd6\x20\x9f\x08\xc0\x23\x89\x2d\xa9\xb0\x1a\x72\x38\x0a\x23\xd1\xa3\xd3\x84\x7a\x32\x99\xd8\x95\x57\xc4\x1c\xe7\xb0\x0a\xd7\xaa\x96\x6b\x89\x7f\xfe\xad\x99\x5c\x59\xc2\xb3\x5b\x2c\x9c\xa2\x6e\xca\x33\x1c\xc5\xc1\x8a\x8f\xc1\xd9\xac\x09\x4f\x65\x7d\x15\xf3\xa6\xe5\x95\x8c\x90\xea\xea\x2a\xae\x85\xef\xe6\x37\x78\xfe\xa1\xc8\xbe\x9e\xdb\x91\xdd\xf4\x1d\x0d\xd9\xef\x3e\x5a\x91\x1d\xbf\xda\xd1\x5f\xf3\x92\x79\x10\x4d\xf0\xf3\x44\xe1\x89\x55\x25\xca\x6f\x0b\xfd\xd2\x51\x3e\x55\x65\x13\x94\x97\x6a\x7b\x39\x28\x7f\xf7\xf1\x6f\x8f\xf2\xeb\x7d\x55\xe3\x21\xe0\xb4\x9e\x37\x05\x27\x69\xd8\x6a\x70\x5a\xcf\xeb\x80\x53\xfa\x4d\x79\x6d\x70\xda\xf9\x39\x6d\xe7\xd5\xda\x4b\x08\x4a\x49\x4e\xcd\xd7\x6c\x9b\xcd\x2f\xaf\x72\x5f\xb6\xea\x8f\xc1\xe5\xb7\xcc\xff\x0e\x00\x00\xff\xff\xbc\x30\x51\xf5\xcb\x54\x00\x00")

func index_html_bytes() ([]byte, error) {
	return bindata_read(
		_index_html,
		"index.html",
	)
}

func index_html() (*asset, error) {
	bytes, err := index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "index.html", size: 21707, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"defines.html": defines_html,
	"index.html": index_html,
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

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"defines.html": &_bintree_t{defines_html, map[string]*_bintree_t{
	}},
	"index.html": &_bintree_t{index_html, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

