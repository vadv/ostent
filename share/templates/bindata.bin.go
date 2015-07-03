// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

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
	name string
	size int64
	mode os.FileMode
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\x7b\x73\xdb\x36\xb6\xff\x3f\x9f\x02\x97\x6d\xef\xdc\x76\x2a\xea\x26\xdb\xee\x66\x52\xdb\x33\x8e\xa5\xa4\x9a\xc6\x8f\xb5\xe5\xec\xf6\xaf\x0e\x25\x42\x12\x62\x8a\x64\x09\x50\x8e\xab\xf1\x77\xdf\x73\x00\xf0\xfd\x10\x29\x52\x8e\x67\x36\x9d\x69\xcc\x07\x70\x5e\x38\xf8\x9d\x03\xf0\x08\x47\xff\x33\xba\x3c\x9b\xfe\x7e\x35\x26\x2b\xb1\x76\x5e\x9c\x1c\xc9\x3f\x84\xc0\x05\xb5\x6c\xb8\xc0\xcb\x35\x15\x16\x99\xaf\xac\x80\x53\x71\x6c\x84\x62\x31\x78\x6d\xa4\x5f\xad\x84\xf0\x07\xf4\xcf\x90\x6d\x8e\x8d\x7f\x0f\x6e\x4f\x07\x67\xde\xda\xb7\x04\x9b\x39\xd4\x20\x73\xcf\x15\xd4\x85\x7e\x93\xf1\x31\xb5\x97\x34\xd3\xd3\xb5\xd6\xf4\xd8\xd8\x30\x7a\xef\x7b\x81\x48\x35\xbe\x67\xb6\x58\x1d\xdb\x74\xc3\xe6\x74\x20\x6f\x7e\x24\xcc\x65\x82\x59\xce\x80\xcf\x2d\x87\x1e\xbf\x8c\x08\x09\x26\x1c\x2a\xaf\xe1\x6e\xbb\x35\x47\x96\xb0\xcc\x5f\x3d\x2e\x90\xf8\xe3\x23\x81\x2b\x20\x79\x34\x4c\xda\x9d\x1c\x39\xcc\xbd\x23\x01\x75\x8e\x0d\x06\x2c\x0d\x22\x1e\x7c\x90\x83\xad\xad\x25\x1d\xfa\xee\xd2\x20\xab\x80\x2e\x8e\x8d\xe1\xc2\xda\x60\x03\x13\x9f\x15\xba\x72\xf1\xe0\x50\xbe\xa2\x54\x44\x04\x04\xfd\x2c\x86\x73\xce\xe3\xfe\x70\x3d\x64\xae\x4d\x3f\x9b\xf8\x54\x53\xe0\xf3\x80\xf9\x22\xdd\xe5\x93\xb5\xb1\xd4\x53\x23\x6f\x67\xc2\x83\x39\x10\xfa\xc4\x87\x01\x5a\x38\xa0\x70\xf5\xca\x7c\x69\xbe\x7c\x1d\x3d\x30\xd7\xcc\x35\x3f\x01\x4f\x1b\x14\x1f\xac\x2d\xe6\xaa\xf6\xdb\x2d\x5b\x10\x73\x7a\xfa\xfe\xfd\x78\x34\x63\xee\xe3\x23\xb4\xd3\xc2\xa8\x1e\xdb\x2d\x75\x38\x58\x08\x38\x0c\xd7\xcc\xb9\xd3\x2f\xe5\x0b\xd7\x7e\x7c\x34\x22\xa3\x1e\x0d\x95\x70\x5a\xfe\xa1\x76\x8d\x93\xa3\x99\x67\x3f\xe8\x87\xae\xb5\x21\x73\xc7\xe2\xfc\xd8\x80\xcb\x99\x15\x10\xf5\x67\x60\xd3\x85\x15\x3a\x22\xba\x5d\xb0\xcf\xd4\x1e\x08\xcf\x37\x48\xe0\xc1\x30\x62\x6b\xb6\x04\x67\x81\x61\x88\xf9\xd9\x2c\x26\x86\x1e\x01\x2a\x51\xe8\xe9\x84\xcc\x8e\xda\xe4\x5a\xd9\x81\xe7\xdb\xde\xbd\x1b\x71\x41\x09\x69\x90\x34\x96\xc2\x86\x42\x78\x6e\x56\x48\x10\x64\xb9\x74\x28\xf8\x9d\xe3\x58\x3e\xa7\x76\x34\x92\xaa\xb1\xb6\xa9\x6a\x84\xb2\xa8\x56\xd1\x63\x2b\x58\xe2\x40\x7d\xa3\x69\xc5\xaf\x53\x6c\xe5\x78\xfb\x56\xcc\x96\x07\x03\xcf\x75\x1e\xb2\x4d\xa0\xd1\x54\xc9\x91\x18\x03\x6c\x0e\xdd\x6a\x28\xa1\x5f\x0e\x80\x6d\x81\xd4\xb3\xe9\x39\x54\x46\xcc\x0c\x82\x45\x98\x9d\x0c\xd7\x2c\xb0\x5c\x30\x79\x6e\x10\xa3\x41\xd1\x66\xd5\x8d\x32\x23\x11\x35\x35\x88\x15\x30\x6b\x40\x3f\x03\x77\x9b\x02\x65\x11\x84\x34\x9e\x7d\xf9\x81\x40\x87\x41\xf6\x2b\x0d\x0e\x05\x35\x8a\xe8\x71\x34\x84\x4e\x79\xbd\xac\x8c\x4a\xa1\x53\x50\x60\x4d\xdd\x30\xf2\x6f\x75\x2d\xc5\x74\xac\x19\x75\x1c\x6a\xcf\x1e\xf2\x26\xc8\x71\x70\x58\x81\xa4\xf6\x68\x4d\xd4\x0f\x28\x07\x50\xcb\x4c\x9b\xb8\x7b\x10\xba\x2e\x73\x97\x47\x43\x87\x15\xe9\x36\xe8\x0f\x83\x54\x18\x11\x40\xa9\x7c\xbb\x94\xbd\x46\x93\x9b\xe9\xf5\xe4\x2d\x9a\xcb\x2a\xba\x46\x41\x8c\x61\xe8\x64\x2c\x98\xb1\x71\x32\x4c\xf9\x69\x95\x40\x82\x7a\x40\x6a\xe6\x5d\x7a\x58\x10\x9a\x74\x53\xf8\x53\x62\x6b\x64\xc6\x7c\x23\x0f\x0d\x25\x3a\xc7\x1a\x4f\xae\x50\xd9\x32\x03\xe7\x4c\x57\xb4\x6e\xa4\x5e\xd4\xc2\xb1\xca\x38\x93\xdc\x6c\x68\x35\x01\x8a\x23\x15\x73\x05\x6e\x85\xb7\x29\xbd\x3e\x9c\x96\xf9\x7c\xc9\x30\x15\xcd\xbc\x8f\xf7\x97\x49\xd3\xd1\xfd\x25\x89\xd0\x17\x6c\x4d\x0b\x03\x14\x93\x6f\x48\x46\xc3\x95\xa2\x56\x82\x53\xa5\xb3\x22\x63\xcf\x5b\xd9\xb5\x6c\x62\x68\x9b\x96\x4a\x98\x9d\x20\xe5\x2d\x8b\xd3\xa8\xca\xe3\xa3\xcb\x80\x2d\x57\xa2\x1a\x6a\xea\xdc\x3e\x45\xdb\x61\x5c\x0c\x98\x0b\x69\x10\xad\x9c\x55\x11\xe9\x0a\x9b\x2a\x78\xfe\x06\x93\x80\x72\xeb\xe1\xab\x76\x16\xdb\xc5\xcb\xe7\x55\xac\x7c\xde\x2f\xa7\xcd\xb2\x8a\xd3\xc6\x5a\x02\xd8\x8b\x43\xb8\x42\x71\xc6\x10\x6b\x2e\xd8\xa6\x18\xe2\x72\x11\x78\x43\x03\x8e\xde\x5f\x11\x83\xf7\x09\xbb\xb8\x1a\xe0\x6f\x86\xc3\xfb\xfb\x7b\x13\xe2\x68\x00\xff\x9b\x73\x6f\x3d\x54\x79\x38\x64\xad\x0e\xb5\x38\xe5\x43\xc7\x12\x94\x97\x45\x95\xcb\x9b\xe9\xf8\x62\x5a\x16\x4a\xba\x81\x4d\xa4\xeb\x41\x10\x47\x13\xef\x03\x72\xfa\x31\xa3\xa4\x16\xc1\xd0\xc7\xf1\xf5\xcd\xe4\xf2\xe2\x69\x70\x28\x17\xce\x53\xb7\x70\x03\x58\xa1\x57\x0c\x51\x44\x92\x6b\x81\xf2\x6c\x9f\xf8\x96\x6d\x43\x2a\x83\xeb\x05\x4e\x37\xa0\xf7\x03\x5f\x95\x2e\x13\x02\xef\x3e\xbf\x34\x40\xda\x6b\xba\x86\x34\x5d\xad\x43\x55\xc3\x99\xe3\xcd\xef\x30\xdf\x1f\xac\xed\xc1\xab\xe8\xc2\x5b\x2c\x60\xcd\x35\x78\x99\x49\x21\xb6\x5b\x41\xd7\x3e\xda\x97\x18\xb0\x8a\x01\xb1\xf8\x9b\x37\xea\xe2\x0f\x70\x7c\xea\x00\x75\x83\x98\xf9\x80\x99\x70\x67\x8b\x3a\xe6\x7f\x6b\xc7\x8c\x2d\x6a\x79\xcd\xfd\xb0\x56\xd3\x76\xcc\x80\x5a\x2d\x37\xbb\x4f\xcd\xec\x72\xcd\xb2\x7e\x13\xdd\x34\x59\x1c\x36\xf1\x0c\x9f\xd7\x69\xf0\x3a\xe7\x18\x2d\xcd\x07\xd1\xe6\xe9\x35\xda\x2c\x0f\xa8\x11\x44\xb5\x36\x1a\xd5\x6e\x71\x44\x1d\x36\x56\x40\x10\x9d\xc8\x31\xd1\x38\xf5\xf8\xf8\x4b\x7e\xaf\x41\x6f\x31\xe0\xae\x83\xdc\xa0\xda\x6e\x87\x3f\xbc\xf8\x61\xf8\xf8\xb8\xdd\x2a\xf9\x52\x02\x5b\x08\xde\xb8\x3f\x63\x80\xa4\x1a\x47\x81\xf2\xaf\xf0\x17\xf7\x81\xdc\x33\x87\xcd\xef\x8e\xb7\x62\xc5\xb8\xb9\x82\xc8\xe5\x50\xf9\xe4\xf1\x85\x42\xca\x29\xc8\xa9\x10\xf2\x44\xef\x80\xd4\x31\xd3\xd6\x99\x3b\xb0\x72\x80\x5e\xcc\xf5\x43\x81\x7a\x4c\x46\xa7\x42\x04\xc0\x0f\x6e\xe4\x40\x3c\x3e\x26\x4b\x96\xc1\x7c\x45\xe7\x77\x33\xef\x73\xb4\xcf\x90\xdc\xcb\xb0\xb5\x62\xb6\x4d\xdd\x38\x9e\xca\x3b\x94\xe6\x13\xff\x7c\xe6\x78\xb0\xe4\x31\x24\x1f\x64\x28\x23\x1c\x32\x7c\xe7\x05\x9a\x63\x9a\x21\x8a\x85\x6f\xcf\xff\x75\x83\xdb\x54\xaa\x45\x4e\x51\x49\xa2\x95\xb2\x7e\x08\xce\xc7\x0d\x6c\x73\xcf\xc4\x8a\xe0\x1d\x78\x13\xba\x06\xd0\xc5\x7f\xd4\x96\xd2\x76\x0b\x99\xce\x92\x92\x6f\xd9\x8f\xe4\xdb\xb9\x17\x50\xf2\xe6\x98\xa8\x50\x74\x76\x75\x6b\x7e\x80\x1c\x12\x04\x10\x01\x48\x28\x5f\x9b\x17\xe6\x6f\xf4\x01\x85\x24\x12\xc9\x80\xee\xec\x61\x70\x61\x28\x91\x8f\x84\x1d\xf9\x34\xfa\x92\x4a\x64\x89\xeb\xdd\x07\x96\xce\x20\x4f\x62\x3a\xa8\x97\xb0\x2b\x7b\xc5\x7b\x70\xb8\xab\x91\xb2\x57\xc8\xa9\x4f\x83\x39\x44\x18\x95\x00\xab\xbc\x27\x79\x7a\x1c\x31\xb8\xe5\x54\x19\x32\xc3\x56\x3d\x4d\xed\x82\x9c\x1c\x56\x8a\x9b\x07\x5e\x14\x42\x3e\xec\x5b\x06\x48\xf5\x31\xb5\xa1\xd5\xb2\x4c\x60\x1e\x15\x85\x51\x4f\x4b\xa5\x19\x8a\x20\xe5\x74\xd9\x3f\x4d\x5c\x50\xa1\x1c\xfa\x0f\x5e\xc4\x9a\xe1\xcd\x4b\x02\x99\x12\xf3\xa9\x46\x42\xa5\x7c\x6a\x27\x1b\x38\x47\x68\x25\x56\x69\x1c\x8b\xef\xf0\xc5\x0e\x67\x93\xcd\x70\xc4\x33\x7b\x63\xa1\xcb\x32\x59\xd8\xc9\x77\xb9\x4d\xb1\x7d\xb8\xc0\x90\x1e\x9e\x09\x0e\x55\x17\x2e\x43\x6d\x55\xf9\x34\xda\x1a\x16\xc9\xde\x70\x5d\x70\x89\x10\x45\x05\x17\x91\xa0\xbd\x1c\xc0\x36\xd8\x64\x2f\x66\x0f\x90\x09\xff\xd1\x12\xa0\x6c\xc6\xef\x12\x70\x1a\xbd\x93\x44\xb2\x00\x85\x4d\xcc\x11\x0b\x2e\xac\x35\x4d\x60\x4a\xf3\xd3\x50\x65\xb3\x40\xee\x28\x16\x00\x2b\x87\x52\x35\x96\x90\x88\xfd\x7f\x9a\x1b\xdd\x48\x6e\x10\x9f\x7c\xf2\xf2\x15\x72\x1b\x28\x76\x36\xdd\x48\x46\xdf\x97\x02\xdd\x9e\xdc\xb4\x6e\x92\xdb\xdf\xd3\xcc\xb4\x56\xe5\xcc\x0a\x58\x12\xd9\xea\x74\x63\x31\xa7\x5d\x17\x98\x4e\x30\xcc\xff\xeb\xce\xb8\xff\x4b\x9d\xdc\x09\xfe\x18\x24\xee\x79\xa5\x9e\xb4\xe3\x38\xf5\x84\x95\x16\xb2\x13\x2a\x45\xce\xd7\x0b\x34\x95\xea\x2f\x56\x71\x72\x43\xa2\x2b\xed\xb4\x57\x56\x60\xad\xb9\x39\xbe\xb8\x3d\x37\xed\x05\x31\xde\xdd\x18\xc4\x90\x53\xa0\x1b\x9d\xf3\xab\x7e\xe8\x9c\x7e\x3c\x9d\x7c\xe8\x81\xce\xed\xcd\x78\xd4\x03\x99\xe9\xe5\xf4\x14\xc5\xe9\x84\x5b\x19\xb4\xe9\x03\xbc\x98\xeb\xd9\xdd\xd1\x4b\x51\x69\x06\x5f\xaa\xed\xd3\xe1\x97\xe6\xf7\x44\x00\xd6\x88\x5b\x25\x38\x4c\x16\x01\xa5\x2d\xbb\x84\x7b\x43\x18\x76\xdd\x0b\xc3\x26\x52\xc9\xfe\x40\x4c\x3b\x61\x5f\x09\x56\xd4\x33\xff\x35\xf6\x64\x24\xbf\xeb\x57\xa4\x2a\x85\xd6\xe7\x5e\xe8\x0a\x6a\xd7\x36\x27\x79\x1b\xc9\x66\x32\x10\xed\xd1\x0f\xa3\xd1\x1e\xdd\x64\x48\xe9\x2d\x35\xca\x62\x42\x77\x8c\x61\xfb\x25\x48\x6c\x91\x00\xcc\xa4\x2c\x3d\x62\x0b\x33\x0b\x2d\x2c\x93\x19\x65\x61\xa5\xf9\xc4\x8e\xc8\xc6\x10\xc2\xa2\xac\xa4\xe5\x84\x06\x42\x23\xea\x08\x6b\xe2\xb6\xee\x72\x19\x36\x9f\x90\xd0\xa7\x1d\x87\x2c\xf1\x4e\x33\x97\xf5\x99\x7e\x64\x56\x46\x13\x98\x79\xc1\xc2\xaa\x9c\xac\xc5\x75\x05\x91\x75\x36\xc7\xc6\xdb\xc9\xf4\x86\x00\xd0\x11\x4e\xe7\x5e\xfa\x13\x37\xd0\xac\x5f\x6d\x1c\xe5\x36\x98\x67\x47\xc3\xf4\x93\x13\xfc\x62\xb3\xe7\xa2\x67\xb7\x70\x30\x2a\x5f\x5c\x3a\x81\x38\x42\xde\xfe\x3e\x1d\xdf\x90\xb5\x67\x87\x8e\x47\x7e\x7a\xdf\xc1\x80\x6f\x73\x22\x7e\xf7\xd3\xfb\x83\xcb\xd8\xda\x8e\x6d\x85\xec\x02\xad\xac\xdf\xec\x8d\x2d\x68\x10\x78\x41\x47\x68\x55\x34\x76\x62\xab\x6a\x76\x00\x70\xd5\x84\xbf\xa2\x6b\xe5\xf0\x3e\x27\x78\xad\x40\x56\x95\x7e\xd6\xcf\xbc\x1e\x10\xaa\x0a\x3a\x9f\x88\x7d\x39\x28\x36\x61\xde\x07\xf4\x55\xc0\x5d\x1f\xec\xbb\x81\x5a\x1a\x85\xfa\x40\x35\xdf\x9a\xdf\x51\xd1\x11\xd6\x34\x91\x9d\xb8\xa6\xdb\x1d\x00\xd8\x22\xca\x5f\x91\xad\x7a\x88\xbf\x42\xdb\x33\x81\x36\x95\x57\x7d\x59\x80\xab\x96\xe1\x59\xc0\x5c\x06\x96\xba\xe3\xdc\x9a\xae\xf7\xf8\xa8\x09\xbd\x12\x94\x3b\x1f\x9f\x67\xf1\x0d\xde\x9a\xbf\x31\xd7\x4e\x00\x0e\x2b\x51\x14\xb4\xdd\xc1\xf3\x22\xb4\xc5\x5d\x1a\x43\x01\x76\x78\xd7\x66\x9b\x0a\x3b\xec\xb5\xcf\xae\x3b\xb6\xdd\xa2\xc2\x6e\xbd\xee\xb2\x03\xc1\x2f\xf0\xed\x2f\xe5\xfe\x68\xee\x36\xed\x6b\xf6\x91\x9e\x60\x03\x29\x72\xec\xee\x53\x24\xae\x09\x02\x42\xa9\x42\x14\xf9\xf8\xa5\xb2\xb4\x2a\x47\x48\xbf\x90\x15\x74\xcc\x5d\x92\x99\x70\x07\xb2\x16\x25\xfa\xe4\x6c\x45\xf5\x1e\xd1\xde\xfc\xdb\xcb\xcb\x0f\x26\x40\xee\x82\x2d\x81\x8b\xf9\xd6\xf3\x1c\xf9\xe6\x0c\xc9\xa9\xf9\x83\x44\xd4\xde\x9b\xa2\x15\xd5\x37\x16\xdf\x18\xb2\x12\xa2\x8e\xbe\xfa\xb1\x05\x96\x85\xc4\x65\x11\x52\xb0\xb3\xab\xdb\xa8\x3e\xee\x24\xaa\x8e\xc0\xcb\x6c\xa5\x8e\xae\xb6\x52\xe4\x8c\x3d\x54\x01\x99\xe3\x7a\x10\x55\xe2\x61\x24\x32\x1c\x2d\xbc\x60\x9d\xfe\x1e\xbf\xf2\x02\xf6\x17\x56\x05\x39\x03\xf9\x6a\xe6\x05\x52\x55\x4f\x08\x6f\x2d\x1f\x15\x64\xf8\x67\x48\x83\x07\xac\x0b\x59\x9f\xce\xb1\xd0\x30\xa3\xa5\x5c\x7a\xcb\x62\x95\x68\x47\x53\x8a\x30\xe0\xe1\x6c\xcd\xe2\x5f\x32\xe9\xbb\x54\x9f\xd4\xb8\x23\xd3\xc1\x32\xf0\x42\x7f\x80\x05\x7a\xd4\xce\x2e\xec\x53\x2d\x71\x6c\x64\x43\x12\x5f\x0d\xf8\x3a\x2a\xaa\x94\xf7\xf9\x82\xe4\x72\xe7\x00\x21\x69\x95\x3d\x63\x4f\x50\xce\x20\x59\xe9\x5f\x1d\x55\xf9\x42\x44\xae\xc2\x13\x62\x69\x7e\x95\xb6\xc9\x17\x4d\xa6\x85\xc4\x72\x96\xb1\xac\x83\xc5\x49\x85\x77\x28\x62\x76\xb4\x6d\xc6\xf1\xa5\xdd\x5c\x3e\x55\x59\xdb\x44\xc2\xa2\x1c\x88\x2c\x70\x57\xac\xf5\x2c\x54\xf0\x97\xfe\xf2\xa2\xd1\x18\xd7\xc1\x0e\x88\x1a\x50\xbe\x32\xb2\x9f\xdf\xae\xc6\xd7\x93\xcb\x91\xa9\x5f\x82\x66\x25\xc5\x64\xc8\xf3\x45\x41\xb4\xf8\x32\x3f\x11\xdb\xfa\x49\xf5\xa4\xdb\x51\xa1\xa0\xcb\xea\x32\xf5\x6f\x29\xa9\xe4\x45\x5b\x34\xb5\x17\x4f\x00\xa6\xf6\xe2\xa0\x58\x0a\xe4\xab\xa0\x34\xd7\x49\x7f\x7c\x15\xe6\x14\x33\xcc\xc4\x35\x77\xc0\xac\xbd\xd8\x8d\xb2\x15\x3a\x76\x02\xd9\xff\x16\x44\xad\xb0\xdd\x9e\x80\x5a\xed\x0e\xb1\x2c\xbb\xf1\x34\xc1\xd2\xd1\xbb\x1e\xa1\xb4\x81\x6c\x59\x09\x10\x45\x47\xef\x9e\x35\x88\xda\x8b\x5d\x18\x5a\xf8\x45\x10\xe6\xc6\x3c\xed\x83\x0e\xcb\x1b\x2d\x9e\xaa\x63\x37\x4c\x3b\xc5\xe4\xe2\x72\x34\xc6\xca\x12\xed\x1d\x19\x03\x96\xeb\x62\x35\xac\xc9\x10\x31\xf5\xdc\x8f\xea\xda\xc8\x37\x7a\x27\x3f\x46\x1c\x4a\xc0\x88\x7c\x41\xc2\xe4\xd7\x0f\x5d\x62\x56\xc5\x4c\xac\x81\xb0\x12\x6a\x3b\x87\xae\x0e\x14\x77\x38\x64\xae\x38\x20\x5f\x0c\xde\x4a\x9e\x64\xa8\x3a\x09\x94\xfe\xe6\x59\x94\xa7\xb7\x48\xcd\x9e\x22\x52\xb3\xc3\x46\x6a\xd6\x2e\x52\xb3\xd6\x91\x9a\x35\x88\xd4\x15\x3a\x7e\x8d\xd4\x71\xef\xea\x48\x5d\x61\xbb\x3d\x23\x75\xb5\x3b\xc4\xb2\xb4\x89\xd4\x93\x3e\x23\x75\x03\xd9\x8a\x91\x7a\xf2\xbc\x23\x35\x3b\x5c\xa4\x66\x45\x78\xbd\x3a\x3d\xfb\x6d\x3c\x3d\x44\x24\x64\x18\x09\x23\xf2\x4d\x63\x75\x89\x84\xe3\xeb\xeb\xcb\xeb\x83\x09\xa8\xa9\x77\x90\x6f\x72\xb0\x5c\x42\x0a\x38\x39\x6c\x2e\x51\x81\x15\x7b\xe4\x12\xb5\xce\xd5\x21\x76\xe7\xbf\x3b\x35\xcf\x26\xea\x9c\xa9\x93\x40\x99\x4f\xfc\x9d\xe4\x99\xf4\x91\xdd\xb0\xa7\xca\x6e\xf0\x37\xac\x87\x4f\x6f\xf0\x53\xc0\x21\xf3\x1b\xf5\xa9\xa1\x3c\xc1\x39\xa7\x6b\x2f\x78\x68\x98\xc7\x00\xa1\xdd\x89\x4c\x95\x36\x5f\x37\x76\xf7\x4b\x6e\xaa\xec\xb9\x67\x76\x53\xe3\x0c\xb1\x34\xbb\xd3\x9b\x3c\x55\x7e\x6f\xf9\x7d\x4a\x29\xe9\xe5\xc5\xac\xd5\x69\xa4\x33\xaa\x4a\x8d\x28\x41\xa2\xcf\x38\x0f\x02\x2d\x0e\xb7\xed\x5b\xe5\x45\xfb\x6d\xfb\xc6\xdf\x1a\xfb\xdd\xf6\xf5\xf9\x13\xa0\xad\xcf\x0f\x0a\xb6\x40\xbe\x0a\x6b\xaf\x02\x6f\x4e\x39\xa7\xbc\x21\xdc\xfa\x7c\x37\xda\x56\x68\xd3\x0e\x6c\xd5\x81\x35\x55\x40\x9b\x2e\xe7\x7f\x5e\x98\xdb\x75\xc6\xf9\xbc\xfc\x10\x87\x2f\x0b\xf8\x15\x43\xba\x27\x92\x56\xfb\x63\x1a\x1c\x77\xc1\xfd\x8d\x9c\xee\xf0\xd7\xf5\xc4\x88\xce\x03\x6a\x49\xb4\xcd\xaf\x6c\x53\xcb\xda\x46\x52\x7e\x98\x9c\x4f\xa6\xa6\xcf\x5d\xf3\x03\x4c\x8c\x6a\x09\x07\x8d\x85\x4b\xd6\xdd\xfd\xc9\x76\xee\x05\x0d\xd6\xdb\x89\x18\xbe\x13\xf2\xe4\xb8\x84\xac\xd8\xbb\xa2\x4d\x5f\x78\x5f\xe1\x44\xfb\xc1\xbd\xcf\x0f\x81\xf6\x9b\xe5\x13\xa0\xfd\x66\x79\x50\xb4\x07\xf2\x55\x68\xff\x51\x1d\x68\x45\x96\x8e\x37\xc3\x73\x59\x85\x25\xc2\xa6\xc8\xbf\x59\xee\x46\xfe\x0a\xcd\xbe\x22\x7f\x13\xe4\xdf\x2c\x9f\x23\xf2\x57\x0c\xe9\x9e\xc8\x5f\xed\x9b\xb1\x2c\x55\xc8\xff\x54\x28\x55\xa1\xf0\x7e\x28\xb5\x59\xf6\x8d\x52\xbc\x65\xe9\xa3\x0f\x09\x5e\x52\xf7\x18\x05\x84\x4c\xed\x23\x36\x31\xaf\x26\x23\xac\xc6\x73\x97\x49\x05\xa4\x1f\xd5\x76\xfb\xcc\xae\x3b\xd8\x45\x63\x5e\x9a\x54\xb3\x92\xc3\xb8\xc7\x6d\xae\x87\x7e\x9d\xfa\xaf\x49\x29\xb9\x22\xc5\x69\x90\x14\x93\xfb\x3c\x84\x7b\xa9\x00\xc9\xe9\xd9\xb0\xae\x3c\xd1\x2a\x60\xb0\xfe\x17\x0f\x2d\xbb\x5d\xb0\x79\xc3\x82\xcf\xb8\xcb\x0d\xfb\xab\x6d\x97\x6b\xca\xc1\x75\xeb\x4a\x3d\xb1\x10\x34\xfa\x91\xee\x49\xd4\x6d\xaa\x8f\xba\xac\xff\x55\x77\xbb\x01\x48\xaa\xf9\x7f\x92\x03\x80\x45\xfc\x3b\x07\xa0\x53\x95\x69\x94\x0a\x94\x16\x99\xbe\x7a\x92\x53\x1c\x7c\x4e\x0c\x50\xac\xe3\x61\x07\x48\xe5\xb6\x1f\x2a\x37\xe3\xeb\xee\x67\x41\x48\xad\xae\x27\x3d\x50\xb9\x98\x9c\x8d\x7b\x20\xf3\x71\x72\x3d\xed\x81\xcc\xb5\xdc\xca\xef\x4a\x65\x3a\x39\x1f\x63\x66\xa3\x26\x56\x0f\x36\x3a\x95\xf4\xba\x1e\x74\xa1\x23\x44\xf7\x1a\xe2\x28\x67\xd1\x09\x31\x84\x97\x6b\xf5\x24\x15\x15\x65\x42\xa5\x93\x8d\xd4\x35\xa4\x1b\x44\x77\x57\x0f\xa2\xf0\x91\x39\x9b\x3c\xd5\x1e\x12\x9a\xe8\x78\xcd\x13\xcd\x25\xfe\x3d\xc2\x49\x7c\x54\x5b\x9c\x22\xca\x5c\x08\x4f\xde\x83\x94\x26\x66\xa4\x0e\x5a\x4b\x1d\x61\x67\x10\x30\xd1\x9c\xae\x3c\x07\x72\x48\x3c\x57\xee\x2a\xb9\x55\x39\xca\x15\x05\x50\xb7\x11\xb1\x92\xdd\x3c\xf5\xec\xa3\xe5\x84\x99\x33\xd8\x4a\x0e\x74\x6b\x1b\xbb\x53\x25\xfa\xd0\x9b\x87\x7e\x62\x92\x94\x6e\x32\x07\x2f\x3b\xbc\xcb\xd4\x29\xc7\x89\xbc\xfc\x2e\x65\x9f\x21\xd2\x6a\x21\xc8\x66\x29\x3f\x9e\x18\x2a\x03\xb8\xa3\x0f\x98\xdd\xab\x47\x51\x1c\x00\x81\x3c\x07\x19\xa0\x48\x3f\xc7\x61\x43\x1f\x4e\xaa\x96\x11\x63\xec\x51\x09\xe2\x0d\xa4\x88\x53\x19\x9d\xab\xac\xad\xf9\x2a\xc9\x55\x34\x97\x73\x78\x88\xfd\x72\xbf\xd7\x80\xa7\xe6\x2d\x80\x65\x92\xae\xe8\xd3\x7a\x75\xce\x12\x86\xac\xf0\xa3\x8d\x16\xb9\x44\x42\x5f\x86\xb2\x7f\x24\xe4\x25\xe1\xef\xab\xd3\x95\x48\xb8\x0b\x2b\x17\x5b\x4b\x5b\x5d\x05\xde\x86\xd9\x34\xa8\x6b\xd9\x50\x5a\x48\xed\x84\x8e\xbc\xaf\x13\x71\x71\x8d\x47\x0d\x92\x28\x54\x2b\x7a\x43\x56\x7a\x68\x16\xcc\xc1\x75\xb3\x58\x29\xae\x3f\xff\x7f\xc2\x16\x5f\xe1\x9b\x2a\xce\xa5\x01\xbf\x81\xcb\x70\xc8\x7a\xe7\xab\x26\xf9\x2f\x5b\x10\xfa\x27\x29\x3a\xac\x4d\xf0\xac\xc6\x0a\xb8\xce\x4d\x10\x92\xa1\x58\xd7\x3e\xc6\xdc\x7d\x93\x98\x68\xa5\x10\x27\x31\x6a\xcd\x9d\x3d\x20\xb4\xcb\xcf\x66\x58\xfe\x87\x2d\xc9\x2b\x4c\xd1\x2a\x5f\xfa\xda\x43\x2b\x1b\x48\x0f\xab\x7c\x6b\xb3\x80\xce\x85\xfc\x9e\xd7\xd3\xaf\x64\x62\x27\xe8\x10\xe2\x52\x61\x19\x0d\xbe\xca\x7c\xe6\x53\x07\xc6\x80\x3b\x9d\x3a\x6c\xe9\x9e\xa9\xe7\xca\xce\x85\x33\x4a\x93\x7e\xf0\x4c\x37\x8d\xc1\x52\x6f\xb8\xe5\x10\x1e\x1b\x5a\x01\x15\xe9\xd6\x0a\xc5\x53\x27\x2f\x5a\x5a\xa3\x55\xac\xce\x7f\x02\x00\x00\xff\xff\x94\x2e\xb3\xb5\xf9\x67\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 26617, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
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
	"index.html": indexHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
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

