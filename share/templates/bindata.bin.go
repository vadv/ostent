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

var _defines_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\xdf\x73\xda\x3e\x12\x7f\xcf\x5f\xa1\xf1\xe4\x7b\xd3\xcc\x14\xe7\xda\x69\xef\x6e\x7a\xc4\x33\x34\x90\xd6\xd3\x90\x30\xfc\xc8\xcd\x3d\x75\x0c\x16\x41\x17\x63\xf9\x64\x99\x96\x63\xf8\xdf\x4f\x2b\x59\xfe\x01\x36\xd8\x86\xf4\xfb\xd2\x3c\x04\x24\x7d\xbc\xbb\x5a\xed\x67\x57\x92\xd9\x6c\x5c\x3c\x27\x3e\x46\x86\xfa\x0c\x3f\x7d\x52\x5f\xbe\xcf\x3c\x12\x18\xdb\xed\x66\x73\x79\x8b\x3e\xdd\x20\xca\x90\x79\x7b\xdf\x19\x8d\x1e\x3a\xfd\x1e\x32\x66\x9e\x13\x86\x62\xb8\x4d\xfc\x20\xe2\x68\xb3\x31\xed\x6e\x87\x73\xb6\xdd\x22\x04\x8f\x6c\xb7\x37\xc6\x8c\x7a\x9e\x13\x84\xb8\x35\x5b\xe0\xd9\xcb\x94\xfe\x34\x10\x5f\x07\x58\x0c\x24\x6d\x87\x11\xa7\xb5\x20\xae\x8b\xfd\x1b\x83\xb3\x08\x1b\x48\xb5\xac\xf6\xb5\x94\x6c\xb5\x3d\x67\x8a\x3d\x50\x70\x47\x59\xac\x21\x51\x00\x26\xc2\x50\xff\x5f\x23\xbe\xf6\xb0\x1a\xb6\x44\xc7\x18\xff\xe4\xc2\xb8\x6b\xf9\xb0\x75\x71\xb1\xd9\x60\xdf\xdd\x6e\x2f\x0e\x4c\x37\x88\x18\xfd\x11\xca\x19\xff\x20\x7c\x81\xcc\xc7\xa7\xde\x70\x68\x77\x7b\xd0\x63\xc2\x3f\xec\x85\xf8\xa8\x43\x36\x1b\xe6\xf8\xcf\x18\x5d\x92\xb7\xe8\x72\x46\x19\x06\xac\xd9\x75\xb8\x63\xde\x0e\x26\xe6\x3d\x09\xc1\x30\xce\xd0\x0b\x5e\xdf\x08\x59\x00\x31\x1f\x84\xd5\x6d\xee\x22\x29\x44\x38\x42\x58\xdf\x62\xe4\x79\xc1\x91\x4f\x7f\x30\x27\x30\xac\x0c\xb2\x7d\xcd\xdd\x62\xb4\x61\xb5\xc3\xc0\xf1\xb5\x7f\xf4\x33\x93\x10\xb3\x5b\xc0\x4a\xdf\xa4\x7d\x20\x0a\xf0\x56\x6d\x89\xa3\x75\x18\x0b\x44\x56\xa6\xaf\xb9\x40\xdb\xf5\xf0\xae\x89\xd0\xb7\x23\xf1\x9a\x33\x2b\x5e\xca\xfc\xc7\xe1\x85\xe5\xce\xd4\xc3\x10\xac\xf2\x4b\x62\x14\x34\xde\xa1\x90\x33\x12\xe0\x96\xc2\x08\xa3\x17\xd8\x01\xdb\x19\x7c\x05\x8d\x0b\xf8\x72\x60\x65\xc0\x93\x6a\x4e\x31\x26\xf2\x89\x98\xe6\x1f\xa9\xe1\xc7\x24\x08\xcf\x9d\x26\x00\x3c\x75\x54\x02\xf8\x0e\xbe\xa9\xe9\x4d\xa9\xbb\x16\xae\xe4\x78\x19\x78\x0e\x3f\xc0\x07\x64\xca\x88\x93\x78\xf1\x09\x5e\xb2\x8e\x3b\xdd\x9d\x4f\xd7\x1c\x87\xdf\x2b\x53\xea\x4e\x53\xea\xeb\xb8\x7f\x7f\xf7\x38\x44\xc6\x9c\xb2\xe3\xe9\x27\x61\x9b\x4b\xc2\x97\x94\x69\xdd\x3b\xa9\x3e\x65\xdb\x66\x23\xe8\x86\x84\x28\x23\xb6\xac\x25\x2c\x9b\xae\x5b\x2e\x61\xbe\xb3\x14\x59\x47\x0a\x30\xbb\x84\x3d\x88\x66\x9e\x8d\x29\x05\x0f\x78\x4b\x66\xa2\x37\xf0\x81\xc4\x5c\x94\xa2\x96\xd2\xe4\xe2\x55\x4e\x07\x5e\x81\x0e\xf4\xee\xfd\xd5\x3e\x97\x4f\xd6\x55\x38\x1f\xf4\xb7\x02\x55\x59\x4e\x0a\x2f\x4b\x7c\x67\xe5\x10\xaf\x1a\x54\x84\xbd\x08\x81\xbf\xf8\xd3\x30\xf8\x67\x3b\x8c\x82\x3d\x5e\x6b\xd4\x00\xb3\x19\xf6\x79\x86\xdd\x3b\x23\xdb\x6d\x1a\xa8\x52\x50\x05\xed\x63\xca\x9d\xc4\xd0\x26\x59\x41\x07\xa8\x4e\x0d\xc7\xca\x5c\xad\xcc\x11\x87\xfb\xe5\x7d\x1a\x91\xf7\xc4\x7f\x81\xfc\x98\xf2\x18\xd0\x98\x19\x16\x4a\xff\xda\x0e\x5a\x30\x3c\x07\xef\xdd\x8b\x18\xbe\x1b\x99\x5f\x45\x13\x4a\x2a\xca\x38\x36\x1e\xd2\x09\x18\x21\x4b\x84\x14\x99\xc5\x19\xa0\x00\xe8\x30\xcc\x33\xe8\xc4\xd9\xce\x5e\x6e\x39\x6e\x53\x7f\x50\x6a\x93\x18\xca\x68\xe9\xd3\xc8\xe7\xd8\xdd\x8d\x0a\x0d\xac\x6b\x13\xca\x95\x90\xbc\x4d\x9d\xa7\x8e\x7d\xaf\xcd\xca\xab\x52\x43\x3a\xf6\x64\x78\xc7\x6e\x2a\x06\x66\xcc\x3a\xcd\xa6\xc9\xa8\xd7\x4d\x3c\x95\x57\x25\x87\x92\xf2\x09\x3c\xd2\x26\x15\x02\xb3\x9e\x3a\xcd\xa6\xf1\xe3\xb8\x53\xe2\x27\x35\xa4\x67\x2e\xd9\x55\xec\xa7\x18\x78\xc8\x4f\x31\xf9\xea\x16\x9c\x5c\xc9\x68\x5a\x75\x88\x4f\xdd\x3f\xb3\xec\x28\xfd\xc5\x75\x47\x8d\xfd\x92\xc2\x13\xab\xfa\x35\x95\xa7\x44\x59\xcd\xd2\x63\xcf\x19\xc6\x15\xa1\x51\xa5\xda\x03\xb0\x92\xe2\x93\x19\x6a\x52\x7d\x6c\x39\xe3\xd3\xca\x4f\x1c\xa9\x8d\xb7\xa6\xbb\x29\x3b\xae\x01\x25\x09\x5d\x67\xe3\xe3\x39\x43\xe5\xc8\xe3\x38\x99\xb7\x8e\xc3\x54\x2a\x69\xb2\x01\xcd\x93\xb9\x59\x42\x20\x35\xb7\xa1\xfa\xe8\x36\x4f\x39\x6d\xef\x6e\x25\xf5\xc1\x8d\xcc\x4d\x08\xf0\x6f\x78\xad\x58\x6b\xa5\x7d\x90\x4c\x8e\x86\xb2\xc0\x76\xb1\xc7\x1d\xdb\xaf\x0c\x7d\x8c\x78\x15\x6c\x35\x89\x19\x61\x4d\x22\x98\xec\x6c\xa0\xea\x9f\xad\x6c\x11\x91\x6c\xee\xec\x07\xed\xfe\x09\x07\x71\xc2\x3d\x7c\x63\x7c\xb6\xc7\x23\x14\x88\x10\x0b\xf1\x8c\xfa\xae\x38\xf9\xf8\x05\xe7\x9e\x36\xb1\xa6\xed\x6b\x62\x05\x61\xe5\x23\x54\xa9\x02\xe1\xa5\xf3\x6a\xe0\x40\x08\xf4\xf9\xdf\xe3\xde\x08\x2d\xa9\x1b\x79\x14\x7d\xf8\x52\x3e\x91\xcf\xa0\xe6\x8f\x0f\x5f\xce\xa5\xa7\x74\x3e\x85\x8a\x6a\x52\x96\x9c\x5e\xc2\xc9\x1c\x33\x46\xd9\x89\x94\x55\x32\x7e\x73\xb6\xcc\xb9\xbf\x94\xb4\x79\xbe\xea\xaa\xbd\x17\x83\x0d\xb8\xb4\x43\xd4\x73\x8a\xce\x51\xb3\x54\x70\x13\x62\xe6\xc9\x58\x59\x74\x6d\x2a\x66\x79\xd4\x94\x8b\x81\x33\x7b\xc1\xfc\x44\x32\xc6\x42\x7e\xb3\xb1\xd4\xbd\xbf\xe9\x58\xad\x9a\xbd\x1e\x29\xf7\xe4\xbf\x2a\x35\x73\xb4\x6a\xc6\xcd\x25\x5e\xd6\x25\xe5\x5b\x74\x29\x9e\x4a\x99\xd9\xef\xf5\xf7\x39\x29\x10\xe6\x37\x02\xca\x35\x23\xd3\x9e\x23\x8c\x00\xe0\x5d\x95\x43\x1c\x00\xf7\xae\x0f\xd3\xfe\xf8\x60\xa6\x53\x40\x85\x43\x19\x3c\x78\xf2\x8d\xa0\x10\xf2\x1a\xef\x09\x0c\x0b\x7c\x72\x68\xbc\xe8\x24\x75\x9e\x23\x94\x0e\x92\x66\x21\x16\x84\x35\x23\x2c\x60\x74\x96\x86\xd7\x60\x24\x15\xed\x87\x18\xc0\xcc\x81\xdd\x2d\x7b\xe7\x65\x58\x28\x87\x3a\xb4\xf8\x09\x72\x92\x45\x66\xef\x2f\x53\x44\x88\xd9\xf1\xb2\x92\xaa\x66\x84\x32\xc2\xd7\x15\xe1\x0f\xe2\xf0\x5d\x11\x3a\x22\xff\xab\x0a\x1d\xe2\x90\xb8\xf2\x92\xa2\x10\x0e\x34\x81\xd3\xbd\x86\x8f\xc9\xb2\x40\xb2\xbe\xd9\xc9\x79\x63\xb7\xc8\x36\x61\x4c\x10\x36\xbf\x3d\x7f\x7f\x96\xdb\xf3\xdc\x55\x43\xfe\xaa\x7a\x30\x12\xc1\x53\x78\x03\xab\x46\x92\x2b\x55\xd1\x48\xee\x5f\x8b\x70\x27\xdc\xbf\xee\x5a\x34\x29\xb5\x68\x92\xb5\x68\x52\x6e\xd1\xa4\xbe\x45\xb9\xdb\xfc\x3d\x8b\x46\xbd\x61\xd1\x8d\x70\x3c\xa2\xef\xcd\xa0\x51\x78\x49\xad\x71\xcd\xef\xcd\xf7\x56\x6d\x68\x97\xad\x9a\x18\x49\x57\x2d\xb5\xa7\x10\x77\xc6\x55\x7b\xb0\x6f\x7b\xc5\x3e\x92\x23\x7a\xda\x0f\x76\xa9\x45\x0a\x77\x3e\x1f\x3d\xd9\xc3\x71\xb1\x45\x72\x44\x2b\x81\x46\xc9\xaa\x29\xdc\xf9\x2c\x1a\xf6\x46\x25\xab\x06\x23\xc9\x42\x88\x46\x59\x64\x4b\x5c\x93\x55\xd3\x49\x70\xc7\xa2\xb1\xdd\x2f\x59\x35\x39\x92\xbc\xea\x10\x8d\x12\x1f\x29\x5c\x2d\x1f\x1d\xe4\x1a\xe4\xc4\x92\x38\xea\x64\x2c\xba\x7d\xec\xf7\x3b\x0f\xdd\xfd\x17\x67\x1a\xf7\x0a\xef\x5e\xe2\x52\xdf\x6c\x9f\xc0\x29\xf5\x38\x09\xf4\xde\x49\x19\x94\xdd\x35\x67\x01\xc8\x15\xf9\xbc\x25\xcc\x98\xe1\xa5\x58\xb8\x1b\xc3\x89\x38\x45\x1e\x9e\xf3\x78\x48\x1c\x1f\xb8\x1c\x80\x5f\xfc\x44\x9e\xd8\xd2\x41\x7d\x33\x47\x0b\xca\x78\xfa\x6b\x90\x12\xf1\x2e\xe5\x21\x92\x9b\xf4\x46\x7a\x4c\xd3\xbc\xb8\x48\xdf\x06\xc0\xc7\xf1\xe9\x47\x21\x0e\xd4\x8e\xb5\x42\x21\xcc\x2e\x2a\x68\x8e\xd7\x51\x4e\xf1\xc9\xf1\x22\x9c\xbe\x8e\x38\xae\x79\xf5\x2c\xcf\xf8\x46\xba\xbb\x32\x74\x97\xaa\xff\xd4\x03\x49\x37\xc6\x47\x29\x5f\x56\xd2\x27\xe7\x59\xec\xd6\x78\x0f\x50\x05\xb5\xff\xa0\xb6\x64\x3b\x78\xf2\x6b\xb3\xa5\x33\x5b\xa4\xd5\x3d\xb6\xa9\x2f\x3a\x41\x63\xd1\xdb\xb3\x95\x82\xc4\x2f\xcf\xa2\x88\xb8\x86\x92\x62\x4e\x26\x7a\x27\x69\x25\x1b\xbe\x5a\x6f\xb0\xb4\xe8\x5d\xa1\xe8\xef\x57\xfb\x3b\x4a\x38\x6c\x00\x40\xbd\xac\x2b\x1b\x1d\x30\xba\x12\xfb\x36\x56\x84\x68\x64\x9b\xd8\x68\x71\xac\x8d\x1b\x41\x03\xfd\xa3\xd0\xba\x46\xd2\xe7\xc4\xc3\x81\xc3\x17\x5a\x41\xbc\x1e\xd0\xfd\x1d\xfa\xd1\xc7\xbf\x5e\x95\xef\x13\x0f\x86\x4c\x28\xb6\x72\xb3\x45\x95\x33\x04\x99\x23\xfc\x5f\x14\xff\x06\xf0\xcd\x7f\x42\xea\xa3\xfd\x88\x75\xaf\x00\x5a\x3e\x49\x1d\xfe\x28\x27\xfa\x10\x3e\xc9\x7b\x75\x37\xbf\xab\xe7\x9d\xd3\x22\x71\x81\x7e\xaa\xb3\xce\xd1\x91\x24\x87\x3f\x0b\xde\x9f\x26\x8d\x20\x0e\xa2\xa4\x43\x06\x41\xd2\x72\x09\xc3\x33\x4e\xd9\xba\xd1\xc9\x30\x59\x99\x9a\x29\x5f\x54\x2f\x86\x43\x58\xd0\xb6\x4b\x56\x7a\x9a\xf2\x97\x9a\x2d\xe1\xcb\x28\x40\x99\xef\xad\x70\x89\xe2\x07\x54\x87\x91\x4f\xdc\x59\xa8\xe3\xba\xd4\x37\xac\xa1\x82\xeb\x1c\xac\x7e\x5c\x1a\xc3\x45\x8a\x59\xca\xb4\xcd\xa8\x97\xc8\x95\x08\xfd\xbb\x52\xd8\x0f\x18\x48\x26\xfd\x05\xf5\x84\xef\x64\x8e\xed\xe2\xb9\x13\x79\x3c\x0e\x32\x31\x1d\x1d\x5d\xaa\xff\x4a\xb7\x07\x58\x1c\xf9\x20\xc2\x0c\xb4\x82\x54\x2c\x1f\x56\x9d\x49\x40\x88\x19\x5c\x8b\x89\x27\x2e\xfa\x7f\x00\x00\x00\xff\xff\x46\xa6\x5c\x3b\x3f\x2b\x00\x00")

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

	info := bindata_file_info{name: "defines.html", size: 11071, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3c\x6d\x6f\xdb\xb6\xd6\xdf\xfb\x2b\xf8\xa8\xdd\x83\x15\x98\xe5\xdb\xde\xed\x6e\xe8\x1c\x01\x6e\xec\xb4\xc6\xf2\x62\xc4\x4e\x77\xf7\x69\x90\x2d\xda\x66\x23\x4b\x9a\x44\x39\xcd\x82\xfc\xf7\x7b\x0e\x29\xea\xcd\xa2\x2c\xcb\x4e\xb7\x0e\x2d\x50\x5b\x22\x8f\xce\x39\x3c\xef\x87\xa2\xf3\xf0\xe0\xd0\x05\xf3\x28\x31\xe4\x77\xf4\xe6\x8d\xbc\xf8\x7d\xee\xb2\xc0\x78\x7c\x7c\x78\x78\x71\x4a\xde\x9c\x10\x3f\x24\xe6\xe9\x79\x7f\x32\xb9\xec\x5f\x0c\x89\x31\x77\xed\x28\x82\xe9\x1e\xf3\x82\x98\x93\x87\x07\x73\x34\xe8\x73\x1e\x3e\x3e\x12\x82\x8f\x3c\x3e\x9e\x18\x73\xdf\x75\xed\x20\xa2\x9d\xf9\x8a\xce\x6f\x67\xfe\x27\x83\xf0\xfb\x80\xc2\x44\x7a\x6f\x87\xcc\xee\xac\x98\xe3\x50\xef\xc4\xe0\x61\x4c\x0d\x22\xef\xac\x5e\x57\x60\xb6\x7a\xae\x3d\xa3\x2e\x12\x38\xf3\xc3\x84\x42\x4a\x00\x59\xc4\xa9\x8b\x5f\x27\xfc\xde\xa5\x72\xda\x82\x81\x29\xfd\xc4\x81\xb9\xae\x78\xd8\x7a\xf6\xec\xe1\x81\x7a\xce\xe3\xe3\xb3\x9a\xe5\x06\x71\xe8\xdf\x45\x62\xc5\x77\x8c\xaf\x88\x79\xf5\x61\x78\x7d\x3d\x1a\x0c\x71\xc4\xc4\x0f\xea\x46\x74\xa7\x40\x1e\x1e\x42\xdb\x5b\x52\xf2\x82\x7d\x47\x5e\xcc\xfd\x90\x22\xac\x39\xb0\xb9\x6d\x9e\x8e\x6f\xcc\x73\x16\x21\x63\x3c\x24\xb7\xf4\xfe\x04\x70\x21\x88\x79\x09\x5c\xf7\xb8\x43\x04\x12\x10\x04\x70\xdf\x09\xd9\x72\xc5\x89\xe7\xdf\x85\x76\x60\x58\x39\xc8\x5e\x97\x3b\xd5\xd0\x86\xd5\x8b\x02\xdb\x53\xf2\x51\xcf\xdc\x44\x34\x3c\x45\x58\x21\x9b\x6c\x0c\x51\x21\xbc\xb5\x37\xc6\xc9\x7d\x94\x20\x24\x56\x6e\xac\x3d\xc2\x91\xe3\xd2\x32\x8b\x38\x56\xc2\xd8\xe5\xa1\x95\xa8\xb2\xf8\x55\xaf\x58\x6e\xcf\x5c\x8a\xc6\x2a\x2e\x52\xa6\xf0\xe6\x15\x89\x78\xc8\x02\xda\x91\x30\xc0\xf4\x8a\xda\xc8\x7b\x88\x97\x48\x71\x85\x17\x35\x9a\x41\x49\xca\x35\x25\x30\xb1\xc7\x60\x99\xdf\x64\x8c\xef\xc2\x00\x92\x3b\x0c\x01\x4a\x6a\x27\x06\x94\x1d\x5e\xc9\xe5\xcd\x7c\xe7\x1e\x44\xc9\xe9\x3a\x70\x6d\x5e\xe3\x0f\xc4\x14\x16\x27\xe0\xe1\x1b\xa5\x64\xed\x16\xba\xb3\x98\xdd\x73\x1a\xfd\xde\xd8\xa5\xce\x94\x4b\xbd\x9f\x5e\x9c\x9f\x5d\x5d\x13\x63\xe1\x87\xbb\xc3\x4f\xea\x6d\x0e\x8b\x6e\x33\x4f\x1b\x9c\x09\xf2\x99\xb7\x3d\x3c\x80\xbb\x11\x40\x65\x24\x9c\x75\x80\xb3\xd9\x7d\xc7\x61\xa1\x67\xaf\x21\xea\x08\x04\xe6\x80\x85\x97\x70\x5b\xf4\xc6\xcc\x05\x6b\xa4\x25\x22\xd1\xb7\xf8\x45\x60\x2d\x92\x50\x47\x52\x72\xe8\xa6\x40\x83\x6e\x90\x06\x79\xf5\xfa\xe5\xb6\x2f\x1f\x4c\xab\x72\x3d\xe4\x3f\x15\xa4\xf2\x3e\x09\x52\x16\xf0\xfd\x8d\xcd\xdc\x66\xa0\x60\xf6\x60\x02\xff\xef\xcd\xa2\xe0\xe7\x5e\x14\x07\x5b\x7e\xad\xa0\xc6\x34\x9c\x53\x8f\xe7\xbc\xbb\x34\xf3\xf8\x98\x19\xaa\x40\xd4\x80\xfa\xd4\xe7\x76\xca\x68\x9b\xa8\xa0\x0c\x54\x85\x86\x5d\x69\x6e\xaf\xc8\x91\x98\xfb\x8b\xf3\xcc\x22\xcf\x99\x77\x8b\xf1\x31\xf3\x63\x84\xa6\xa1\x61\x91\xec\x5f\xcf\x26\xab\x90\x2e\x50\x7a\xe7\x60\xc3\x67\x13\xf3\x3d\xdc\x62\x4a\x25\x39\xc1\x26\x53\x2a\x00\x13\x62\x81\x49\xb1\x79\x12\x01\x2a\x00\xed\x90\xf2\x1c\x74\x2a\x6c\x7b\x2b\xb6\xec\xe6\xe9\x62\xac\xe5\x09\xa6\x72\x54\x2e\xfc\xd8\xe3\xd4\x29\x5b\x85\x02\xdc\x97\x27\x52\x48\x21\x45\x9e\xfa\x1f\xfa\xa3\x73\xc5\x56\x91\x94\x9c\x52\xb6\x27\xcc\x3b\x11\x53\x35\x60\x8e\xad\xc3\x78\xba\x99\x0c\x07\xa9\xa4\x8a\xa4\xc4\x54\x9a\x3e\xd1\x8f\x14\x4b\x95\x80\x79\x49\x1d\xc6\xd3\xf4\x6a\xda\xd7\xc8\x49\x4e\xa9\x95\x0b\xef\xaa\x96\x53\x02\x58\x27\xa7\xc4\xf9\xf6\x4d\x38\x85\x94\xd1\x36\xeb\x30\xcf\x77\xfe\xca\xb4\x23\xe9\x57\xe7\x1d\x39\xf7\x59\x12\x4f\x42\xea\xf3\x64\x1e\x0d\xb1\x3d\x53\xcf\x68\x11\x52\xda\x10\x34\x6e\x94\x7b\x10\x4c\x93\x7c\x72\x53\x6d\xb2\xcf\x48\xac\xf8\xb0\xf4\x93\x58\x6a\xeb\xd2\xb4\x1c\xb2\x93\x1c\xa0\x09\xe8\x2a\x1a\xef\x8e\x19\x32\x46\xee\x86\x13\x71\x6b\x37\x98\x0c\x25\x6d\x0a\xd0\xa2\x33\xb7\x0b\x08\x6c\xcf\x32\x54\xb5\x6e\x8b\xcc\xa7\x47\xe5\x52\x52\x35\x6e\x6c\x61\xa2\x81\xff\x42\xef\xa5\xd7\x5a\xd9\x18\x06\x93\x9d\xa6\x0c\xb0\x03\xea\x72\x7b\xe4\x35\x06\xbd\x8a\x79\x13\xd8\x66\x18\x73\xc8\xda\x58\x30\x2b\x15\x50\xfb\xf7\x56\x23\xb0\xc8\x70\x61\x6f\x1b\xed\x76\x87\x43\x38\xe3\x2e\x3d\x31\xde\x8e\xa6\x13\x12\x80\x89\x45\x74\xee\x7b\x0e\x74\x3e\x5e\x45\xdf\xd3\x63\xd6\xac\xd7\x65\x56\x10\x35\x6e\xa1\xb4\x04\x40\x4a\xc7\xa5\xc0\xd1\x21\xc8\xdb\xdf\xa6\xc3\x09\x59\xfb\x4e\xec\xfa\xe4\xfb\x77\xfa\x85\xbc\x45\x32\xdf\x7c\xff\xee\x58\x74\xb4\xeb\xa9\x24\xb4\xa7\xcb\xb2\xc3\x53\x38\x5b\xd0\x30\xf4\xc3\x03\x5d\x56\xe2\xf8\xea\xb3\x3a\xe1\x7e\x56\xa7\x2d\xfa\xab\xca\xda\x5b\x36\xd8\xc2\x97\x4a\x8e\x7a\x4c\xd4\x05\xd7\xd4\x22\x6e\xe3\x98\x45\x67\x6c\x8c\x7a\x6f\x57\xcc\xfb\x51\x5b\x5f\x0c\xec\xf9\x2d\xe5\x07\x3a\x63\x82\xe4\xab\x37\x6a\xc5\xfb\xd5\x1d\x9b\x65\xb3\xa7\x73\xca\x2d\xfc\x4f\xea\x9a\x05\xb7\x6a\xe7\x9b\x6b\xba\xde\xd7\x29\xbf\x23\x2f\xe0\xa9\xcc\x33\x2f\x86\x17\xdb\x3e\x09\x10\xe6\x2f\x0c\x89\x2b\x8f\xcc\x46\x76\x78\x04\x02\x9e\x35\x69\xe2\x10\x70\x6b\xfb\x30\x1b\x4f\x1a\x33\x15\x02\x1a\x34\x65\xf8\xe0\xc1\x3b\x82\x80\xe4\x29\xde\x13\x18\x16\xca\xa4\x6e\xbe\xaa\x93\x3a\x4e\x0b\xa5\x8c\xa4\x9d\x89\x05\xd1\x9e\x16\x16\x84\xfe\x3c\x33\xaf\xf1\x44\x10\xda\x36\x31\x04\x33\xc7\xa3\x81\xee\x9d\x97\x61\x91\x02\x54\x9d\xf2\x53\xc8\x9b\x3c\x64\x7e\xff\x32\x83\x88\x68\xb8\x3b\xad\x64\xa4\x43\xe6\x87\x8c\xdf\x37\x04\xbf\x84\xe6\xbb\x21\xe8\x84\xfd\xd9\x14\xf4\x9a\x46\xcc\x11\x9b\x14\x95\xe0\xe8\x26\xd8\xdd\x2b\xf0\x29\x5b\x57\x60\x56\x3b\x3b\x05\x69\x94\x93\x6c\x1b\x8f\x09\xa2\xf6\xbb\xe7\xaf\x8f\xb2\x7b\x5e\xd8\x6a\x28\x6e\x55\x8f\x27\x60\x3c\x95\x3b\xb0\x72\x26\xdd\x52\x85\x9b\x74\xff\xb5\x0a\xee\x80\xfd\xd7\x32\x47\x37\x5a\x8e\x6e\xf2\x1c\xdd\xe8\x39\xba\xd9\x9f\xa3\xc2\x6e\xfe\x16\x47\x93\xe1\x75\xd5\x8e\x70\x32\xa3\xf6\xcd\xf0\xa6\x72\x93\x5a\xc1\xb5\xdf\x37\xdf\xd2\xda\xf5\x48\xa7\x35\x98\xc9\xb4\x96\xf1\x53\x09\x77\x44\xad\x5d\x8e\x4e\x87\xd5\x32\x12\x33\x6a\xd9\x97\x23\x2d\x47\x12\xee\x78\x32\xfa\x30\xba\x9e\x56\x73\x24\x66\x14\x11\xbc\xd1\x68\x4d\xc2\x1d\x8f\xa3\xeb\xe1\x44\xa3\x35\x9c\x49\x15\x01\x37\x3a\xcb\x16\x70\x6d\xb4\xa6\x82\x60\x89\xa3\xe9\xe8\x42\xa3\x35\x31\x93\xbe\xea\x80\x1b\x8d\x8c\x24\xdc\x5e\x32\xaa\xf5\x35\x8c\x89\x1a\x3b\xea\xe7\x38\x3a\xbd\xba\xb8\xe8\x5f\x0e\xb6\x5f\x9c\x29\xb8\x27\x78\xf7\x92\xa4\xfa\x76\x75\x02\xf7\x7d\x97\xb3\x40\xd5\x4e\x92\xa1\x7c\xd5\x9c\x07\x20\x0e\xc4\xf3\x0e\xb0\x31\xa7\x6b\x50\xdc\x89\x61\xc7\xdc\x27\x2e\x5d\xf0\x64\x0a\xda\x07\x2e\x26\xf0\xc4\x4f\xec\x42\x49\x87\xf9\xcd\x9c\xac\xfc\x90\x67\xa7\x41\x34\xe8\x1d\x9f\x47\x44\x14\xe9\xad\xe8\x98\xa6\xf9\xec\x59\xf6\x36\x00\xbf\x76\x2f\x3f\x8e\x68\x20\x2b\xd6\x06\x89\x30\xaf\x54\xa4\x9c\xe8\x51\x2c\xf1\x83\xed\xc6\x34\x7b\x1d\xb1\x9b\xf2\x66\x29\x7a\x7c\x23\xab\xae\x0c\x35\x24\xf3\xbf\xef\x22\xa6\x13\xe3\x07\x81\x5f\x64\xd2\x0f\xf6\x12\xaa\x35\x3e\x44\xa8\x8a\xdc\x5f\x4b\x2d\x2d\x07\x0f\x7e\x6d\xb6\xb6\xe7\xab\x2c\xbb\x27\x3c\x5d\xc0\x20\x52\xac\x7a\x7b\xb6\x91\x20\xc9\xcb\xb3\x38\x66\x8e\x21\xb1\x98\x37\x37\xaa\x92\xb4\xd2\x82\x6f\xaf\x37\x58\x0a\x75\x19\x29\xf9\xf1\xe5\x76\x45\x89\xcd\x06\x02\xc8\x97\x75\xba\xd9\x71\xe8\x6f\xa0\x6e\x0b\xab\x20\x5a\xf1\x06\x85\x16\xa7\x8a\xb9\x09\xde\x90\x9f\x2a\xb9\x6b\x85\x7d\xc1\x5c\x1a\xd8\x7c\xa5\x08\x24\xfa\xc0\xe1\xdf\x71\x9c\xfc\xf0\xaf\x97\xfa\x3a\xb1\xd6\x64\x22\x28\xe5\xe6\xab\x26\x3d\x04\x5b\x10\xfa\x07\x49\xce\x00\x7e\xfb\x31\xf2\x3d\xb2\x6d\xb1\xce\x4b\x04\xd5\x2f\x52\x99\x3f\x29\xa0\xae\x83\x4f\xe3\xde\xbe\xc5\xef\x66\x59\xea\x16\x99\x83\xee\x27\x07\xf7\x69\x1d\x59\xda\xfc\x59\xf8\xfe\x34\xbd\x09\x12\x23\x4a\x07\x84\x11\xa4\x77\x0e\x0b\xe9\x9c\xfb\xe1\x7d\xab\xce\x30\xd5\xcc\x9e\x21\x1f\xb2\x57\x48\x23\x54\x68\xcf\x61\x1b\xb5\x4c\x71\x52\xb3\x03\xb2\x8c\x03\x92\xbb\xee\x44\x6b\x92\x3c\x20\x07\x8c\x62\xe0\xce\x83\xda\x8e\xe3\x7b\x86\x75\x2d\xc1\x55\x0c\x96\x87\x4b\x13\x70\x08\x31\x6b\x11\xb6\x43\xdf\x4d\xf1\x0a\x08\x75\xae\x14\xeb\x01\x83\x88\xa0\xbf\xf2\x5d\x90\x9d\x88\xb1\x03\xba\xb0\x63\x97\x27\x46\x06\xcb\x51\xd6\x25\xc7\x5f\xaa\xfb\x31\x85\x96\x0f\x2d\xcc\x20\x1b\x0c\xc5\xe2\x61\x39\x98\x1a\x04\xac\xa0\x0b\x0b\x4f\x45\xd4\xfb\xbf\xc1\xd5\xe9\xf4\xb7\xf1\x90\xac\xf8\xda\xb5\x7a\xc9\xa7\xd0\xc2\x9a\x72\x9b\xcc\x57\x76\x18\x51\x48\x33\x31\x5f\x74\x7e\x32\x92\xd1\x15\xe7\x41\x87\xfe\x11\xb3\xcd\x89\xf1\xdf\xce\x4d\xbf\x73\xea\xaf\xc1\xd3\x98\xb4\x1d\x95\x99\x46\xc3\x13\xea\x2c\xa9\x7a\x08\xad\x03\x4c\x8c\xd1\xbb\x00\x72\x61\x0e\xee\x8e\x39\x7c\x75\xe2\x88\x17\xc0\x1d\x71\xf3\x1d\xa8\x81\x71\x66\xbb\x9d\x68\x6e\xbb\xf4\xe4\x15\x5a\x1d\xee\x74\xa5\x79\xe0\xbd\x1f\x71\x4f\x04\x31\x02\x57\x80\x06\x94\x2f\x00\x7a\x2e\xf4\x5a\x20\x5e\x17\xf4\x03\x14\x94\x68\xd9\xda\x5e\xd2\x6e\xe0\x2d\x0d\x59\xd3\x18\xdd\x85\xbd\x41\x00\x13\xc7\xf2\x4f\x45\x78\x12\x37\x5a\x51\x5a\x50\x4b\x77\x0e\x19\x40\x3d\x0a\xd7\x5d\xe6\x39\xf4\x93\x89\xa3\x60\x14\x73\xf0\x0f\x9e\x87\xfe\x68\x6f\x6c\x39\x6a\x94\x45\x48\xa2\x70\x0e\x38\x3e\x46\xdd\x10\x25\x18\x52\xb8\x7a\x6d\xbe\x32\x5f\xfd\xa8\x06\xcc\x35\xf3\xcc\x8f\x51\x92\xe8\xd7\x36\xf3\x24\xbc\xd0\xbf\x39\xed\xbf\x7b\x37\x1c\xcc\x98\xf7\xf8\x08\x70\x09\x1f\xf2\x09\x15\x32\x80\x42\x77\xcd\xdc\xdb\x64\x52\x4c\x28\xed\x4b\xae\xe0\x42\x2a\x59\x3a\x0e\xfa\x02\xfa\x3f\x44\x57\xac\x03\x12\x23\xe9\x79\x76\xea\x22\x70\x39\xb3\x43\x22\xbf\x3a\x8e\x34\x3d\x75\xbb\x60\x9f\xa8\xd3\xe1\x3e\x04\x68\xb0\x6d\x2a\xa0\xd9\x12\xec\x01\x5d\x22\xef\x68\xa8\x72\x58\x0e\x85\x47\x5c\x4c\x57\x85\xc9\x04\x99\xaa\x44\x2b\xa6\x66\x10\x48\x1d\x51\x29\x23\xb3\xab\xc4\x04\x52\xb5\x6c\x17\x4d\x33\x9f\x73\x7f\x9d\x8c\x73\x7f\xb9\x14\xdb\xbc\x7e\xe0\x6f\x80\x42\x32\x0a\x1d\xc1\x12\x3d\x6d\x85\x83\x64\xe1\xcf\xe3\x28\x57\x62\x09\x6e\x4f\x8c\xe7\x0e\xc3\x20\x38\xeb\x14\x1f\x4e\x4d\x38\x8c\x3d\x8f\x79\x4b\xa2\xcc\x73\x30\x9a\x4c\xaf\x47\x6f\x93\xca\xa8\x64\xb1\xb2\xda\x15\x12\x16\xe1\x04\x17\x53\xc6\x9f\x55\xc6\x02\x4e\x29\x28\x11\x84\x3a\x99\x6e\x64\x82\x95\x03\xa4\x0c\x90\x3d\x1a\x07\x9c\xad\x69\x27\x80\xf2\xdb\xe3\x5b\x0f\x1a\x96\x9c\xcf\x6f\x1a\x67\x4f\x65\xab\xb8\x11\xf7\xb9\x63\xcd\x82\xbd\xd8\xcd\x69\x4a\xf1\x00\x5f\xc2\xb1\x04\x1e\x4c\xde\x45\x65\x8a\x60\x97\xe2\x1d\x8d\xc5\x71\x77\x96\x3e\xe0\xda\xf5\x0f\x9c\xf7\xd5\x03\xdd\xd8\xd5\x73\xa0\x2e\xd5\x89\x32\xc0\x5e\x85\x35\xf7\xbc\x0b\x9a\x80\xb0\xec\x62\xe2\x28\xad\x24\xed\xd1\x8c\xe7\x68\xec\x16\x7e\x4a\x65\x4a\xc6\x73\xf3\x01\xc4\x85\xe7\xf8\xa2\xa0\x72\x76\x03\x21\xe7\x79\x52\xbe\x64\x20\x62\x21\x4a\x06\x09\x37\xf6\x9c\xb3\x0d\xcd\xba\x43\x03\x43\x6f\xf4\xa6\xdb\xbd\xbb\xbb\x33\xc1\xa2\x42\xf8\x6f\xce\xfd\x75\x57\x46\x41\x08\x21\x2e\xb5\x23\x1a\x75\x31\x6f\x46\x3c\x7d\x41\x00\x36\x15\x81\x3b\xa6\x06\x0a\x65\xcc\x64\x74\x75\x89\x06\x7a\x35\x99\x0e\x2f\xa7\x65\x36\x84\x5e\x93\x4f\x58\x7f\x66\x47\xc2\xcd\xab\xfd\x99\x04\x90\x06\xc1\x0d\x30\x14\x44\x74\x03\xfc\xdc\x43\xae\x2d\x78\x32\xd4\x2b\xc5\x81\x99\xeb\xcf\x6f\xb1\xd4\xef\xac\x9d\xce\x6b\x75\xe1\x2f\x16\x10\x32\x3b\xaf\x8a\xb0\x60\x71\xd4\xc5\x31\xf9\x9b\x8a\x92\x09\x77\xe4\xa8\x00\x12\x51\x04\x1d\x72\xc6\xbd\x8e\x20\x61\x80\x6b\x83\x23\x3b\xf1\x7a\x7d\xbf\xa6\x6b\xe0\x7c\xc1\x96\xd9\x0f\x39\x60\xb2\x13\x61\xf8\x7f\x9e\x4e\x0a\xa1\x2b\x2e\x01\x8d\x0c\x4c\x12\x5b\x5a\xef\x2d\x6c\x37\x2a\x15\x7c\xa7\x22\x88\x9a\xef\xa1\xf8\x91\x88\x2e\x86\x17\x90\x95\x89\x54\xa5\x0a\xc5\xca\x14\x80\x9c\x61\x5d\xd0\xb5\xa8\x84\x6c\x55\x32\x88\x90\xa0\x61\xb5\xbc\xec\x76\xbf\x4e\x91\xbf\x2c\x29\x56\x28\x87\xa1\x2c\x04\xaa\x4c\x8a\x0a\x79\x4d\x89\xac\x91\x58\xca\x8e\xa4\x90\x25\x31\xac\xa4\xd2\x8d\x0b\x3f\x64\x7f\xa2\x19\xba\x1d\x31\x3c\x83\x2a\x1b\xb5\x24\x22\xbf\x18\x2a\xda\x90\xa8\xc2\x64\xcd\x86\xd6\x4a\x4b\x49\x08\x15\x2d\x0b\xc1\xf4\x0a\xca\xc0\x52\x06\x99\xc5\x80\xdc\x8b\xca\x86\x08\x4f\x88\xa7\x54\x76\x2c\x1a\x66\xc9\xc6\x0c\x25\xfb\x92\x84\xad\xf7\x62\xb1\xa9\x82\xea\x28\xa8\x82\x70\x87\x15\x4e\x7e\xed\x8f\xf5\x06\x08\xc2\xa5\xd1\x1d\x6e\xb0\xeb\x19\xa2\x04\x21\x70\x6f\x21\xe1\x2a\x1f\x1e\xea\x85\xab\x1a\xa7\x3c\x53\x49\x95\x0c\x5a\xd6\xb4\x37\xaa\x4c\xcf\x35\x36\x8a\x24\x92\xd8\xa2\x5f\xb0\xbb\x3d\x2d\x6e\x87\xad\xa5\x99\x10\x30\xab\x0e\x68\xc7\x3b\xab\xa4\x87\x32\xcb\x59\x72\xfb\x53\x17\x07\xff\xfd\x34\x71\x8f\x2d\xf4\x61\x4f\xcd\x1d\x2d\xea\x8d\xce\xf4\x36\xc7\x16\x59\x22\x4f\x1e\x9c\xda\xb3\xd1\x99\x39\xc5\x6c\x95\x54\x48\xa5\x48\xb8\xcd\xfc\x17\x11\x08\x53\xb9\xb6\x89\x83\x42\x86\x6d\xc2\xe0\x5f\x16\xf3\x80\x47\x26\x1a\xaa\xe3\x04\xb7\x7a\x63\x1b\x7e\x02\xa3\x77\xd0\xdb\xa4\xb1\x41\x11\x2d\x76\x51\xab\xcc\x4d\xc3\x50\xc9\x0a\x25\x46\xac\x09\x47\x67\xd0\x1e\x1f\x39\xe0\x21\xce\x76\xf1\x6e\xab\xbc\xc5\x50\x14\x15\xea\x59\xb6\xe8\xa8\x3d\x11\xa9\x38\x7b\x86\xf6\x97\x16\xd8\x67\xd3\xfe\xdb\x89\xc9\xce\xc6\xfd\xd3\x5f\x86\xd3\x89\x79\xc3\xf0\xbd\x6c\xae\xb2\x7c\x9e\x1e\xec\x30\xac\xb1\xbc\xc8\x97\xaf\xfb\xd2\x19\x5e\x5f\x5f\x5d\x57\x93\x91\x47\xbb\x0c\x4b\xec\x89\xd5\x13\x49\x02\xc8\x0e\x5a\xe2\x98\x68\x25\x29\x71\xa0\xd3\xb0\xde\xe2\x57\x4d\x89\x5b\x72\x59\x25\x87\x1c\x3b\x40\xba\xbc\xff\xa2\x15\xeb\xcb\x2a\x63\x95\x01\x4e\x4c\x83\xa9\x6a\x9c\x3a\x59\xa6\xac\xd9\xb5\xaa\x4b\xa2\xe4\xbe\x8a\x4e\x53\x59\xba\xc0\x26\x09\xad\x7c\xcc\x6b\x3b\xaf\x65\x72\x4b\x14\xbb\x87\xd8\x72\x56\xf2\x54\x52\x4b\x48\x34\x13\x5a\xd9\x6a\x73\x32\x93\x8b\x6b\x26\xb2\xc2\x39\xd5\x3a\x89\x49\xfb\xdc\x43\x60\x99\xa9\x3f\x95\xbc\x24\x85\x66\xe2\x2a\x39\x5e\x4e\x5a\xf2\xf7\x9c\x8d\x84\x95\x3f\x88\x7f\x48\xd5\xf4\xfa\x69\xaa\xa6\x79\x10\xeb\xcb\xa6\x74\xf2\x68\x75\xd3\xe9\xf8\x46\x5f\x38\x01\x39\xc3\x02\x88\xaa\x02\xa9\x82\xcf\x2f\xa2\x42\xca\x44\xd8\xa6\x44\x92\xe2\xfa\xe7\xb6\x8a\x42\xe3\x9f\xb9\x9a\x4a\x4c\x50\x57\x4e\xd5\xb1\xa4\xad\xa7\x00\xe7\xf1\x0b\x2a\x81\xf4\xc8\x1d\x24\xae\x6e\x3f\x43\xdc\x61\x82\x69\x50\x04\xcc\x4d\x02\x62\xfa\x07\x1f\xfe\x86\x1d\xa4\x53\xd3\x41\x3a\xc7\xee\x20\x07\x35\x1d\xa4\x53\xd9\x41\x0e\x6a\x3b\xc8\x6d\xe6\xbf\x88\xf8\xe8\x1c\xd2\x41\x0e\xbe\x76\x90\xcd\x62\xde\xa0\xb6\x83\x74\x5a\x74\x90\x83\x27\xe8\x20\x07\x4f\xd9\x41\x3a\xf5\x9d\xdd\x40\x16\x7d\xce\xd9\xe8\xf2\x6a\x50\xd9\x6e\xa9\x9f\xbc\xe2\xcf\x1b\xf0\xbb\xaa\xb3\x73\x1a\x75\x76\x29\x2d\x5d\x67\xe7\xb4\xea\xec\x52\x06\x73\xdc\x68\x0a\xee\xaa\xd5\xea\x2a\xee\x41\xbb\x8a\xbb\x4c\xa2\xb6\xe4\xd6\x4a\x3f\x7b\x85\xa8\x7e\xa1\xdf\x20\xc7\x94\x7e\xbf\xad\xef\x50\x9c\x52\x87\xd2\x40\x60\xbb\x3b\x94\x43\xe5\xd5\xa0\x43\xd1\x19\x50\x4e\x5a\x8d\x3b\x14\x67\xef\x0e\x45\xe3\xeb\xb5\xaf\xbf\xeb\xdf\x8c\xfd\x54\x7a\x33\xf6\x44\xbd\x4e\x10\xe9\xf3\xbb\x9a\x3b\x5a\x7e\x1f\x4f\xf4\xf9\x1d\xdf\x9f\x8e\x43\x7f\x4e\xa3\x28\x71\xef\x52\x32\xdf\xe6\xf4\x8b\x48\xe6\xa9\x10\xdb\x24\x73\x21\xb0\x46\xc9\x5c\xbe\xbf\xd6\xb5\x39\xc5\x3f\xb2\xd2\x3a\x19\x8d\x27\x7b\x27\xa3\xbf\x5b\x77\x85\x66\xd6\xa2\xd0\x70\xc1\x2c\x89\xae\xda\xd0\x2a\x72\x3c\xf1\x7c\x3e\xa0\xf3\x90\xda\xa2\xca\xa8\x2b\x38\x82\x08\x69\x68\x99\xeb\x54\xf3\xb5\xc6\xbf\xd7\xd7\x8e\xaf\xac\x12\xaa\x67\x0b\x49\x34\xad\x85\xc6\x93\xc0\x8d\x23\xf9\xb7\x0c\x35\x95\xd0\xae\x6e\x2c\x88\xf6\xf4\x94\x7a\x1f\x49\xc3\x7f\xd0\x28\xf2\xab\x5f\x08\xfd\xa3\x63\xfe\x66\xa9\x8f\xf9\x6a\xee\x68\x31\xff\xc3\x3b\x7d\xcc\xc7\x53\x31\xc9\x69\x59\xb2\x74\xfd\x19\x1e\xff\xe3\x36\x8f\x2b\xe3\xff\x36\xd7\x5f\x44\xfc\x4f\x05\xda\x26\xfe\x0b\xe1\xfd\x7d\xe2\xff\x87\x77\x5f\x7c\xfc\xdf\xec\x6a\x34\x5b\x86\xad\xcd\xbe\x0a\xae\x57\x6d\x1a\xb6\x36\xcb\x26\x61\x2b\x3d\xc6\xbd\x67\xd8\xaa\x3d\xc7\x6a\x3d\xdb\xd8\x21\x41\xbe\xc9\x89\x3a\x46\xf6\xf8\xf8\xf3\xb3\xdc\x91\xd2\xe4\x14\xb6\x38\x44\xfc\xbf\x00\x00\x00\xff\xff\x95\x44\xee\x28\x99\x57\x00\x00")

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

	info := bindata_file_info{name: "index.html", size: 22425, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
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

