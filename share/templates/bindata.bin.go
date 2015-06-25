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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\xff\x73\xdb\xb6\x92\xff\x3d\x7f\x05\x8f\x7d\xef\xe6\xe5\xcd\x13\x75\xc9\xb5\x77\x99\xd4\xf6\x8c\x6d\xc9\x89\xa6\xb1\xad\xb3\xe5\xdc\xf5\xa7\x0e\x25\x42\x12\x62\x8a\x64\x09\x52\x8e\xab\xf1\xff\x7e\xbb\xf8\xc2\xaf\x20\x4d\x8a\xb4\xe3\x76\xd2\x99\x8e\x45\x12\xd8\x5d\x2c\x76\x3f\xbb\x0b\x02\xcc\xc1\xbf\x8d\x2e\x4f\x67\xbf\x4e\xc7\xc6\x3a\xda\xb8\xaf\x8e\x0e\xf8\x1f\xc3\x80\x1f\xc4\x76\xe0\x07\xfe\xdc\x90\xc8\x36\x16\x6b\x3b\x64\x24\x3a\x34\xe3\x68\x39\x78\x67\x66\x1f\xad\xa3\x28\x18\x90\xdf\x63\xba\x3d\x34\xff\x6f\x70\x73\x3c\x38\xf5\x37\x81\x1d\xd1\xb9\x4b\x4c\x63\xe1\x7b\x11\xf1\xa0\xdf\x64\x7c\x48\x9c\x15\xc9\xf5\xf4\xec\x0d\x39\x34\xb7\x94\xdc\x05\x7e\x18\x65\x1a\xdf\x51\x27\x5a\x1f\x3a\x64\x4b\x17\x64\xc0\x2f\xfe\x65\x50\x8f\x46\xd4\x76\x07\x6c\x61\xbb\xe4\xf0\x8d\x22\x14\xd1\xc8\x25\xfc\x37\x5c\xed\x76\xd6\xc8\x8e\x6c\xeb\xa3\xcf\x22\x24\xfe\xf0\x60\xc0\x2f\x20\x79\x30\x4c\xdb\x1d\x1d\xb8\xd4\xbb\x35\x42\xe2\x1e\x9a\x14\x58\x9a\x46\x74\x1f\x80\x1c\x74\x63\xaf\xc8\x30\xf0\x56\xa6\xb1\x0e\xc9\xf2\xd0\x1c\x2e\xed\x2d\x36\xb0\xf0\x5e\xa9\x2b\x8b\xee\x5d\xc2\xd6\x84\x44\x8a\x40\x44\xbe\x46\xc3\x05\x63\x49\x7f\xf8\x3d\xa4\x9e\x43\xbe\x5a\x78\x57\x52\x60\x8b\x90\x06\x51\xb6\xcb\x17\x7b\x6b\x8b\xbb\x66\x51\xcf\x06\x0b\x17\x40\xe8\x0b\x1b\x86\xa8\xe1\x90\xc0\xaf\xb7\xd6\x1b\xeb\xcd\x3b\x75\xc3\xda\x50\xcf\xfa\x02\x3c\x1d\x18\xf8\x60\x63\x53\x4f\xb4\xdf\xed\xe8\xd2\xb0\x66\xc7\x1f\x3e\x8c\x47\x73\xea\x3d\x3c\x40\x3b\x29\x8c\xe8\xb1\xdb\x11\x97\x81\x86\x80\xc3\x70\x43\xdd\x5b\xf9\x90\x3f\xf0\x9c\x87\x07\x53\x29\xf5\x60\x28\x84\x93\xf2\x0f\xa5\x69\x1c\x1d\xcc\x7d\xe7\x5e\xde\xf4\xec\xad\xb1\x70\x6d\xc6\x0e\x4d\xf8\x39\xb7\x43\x43\xfc\x19\x38\x64\x69\xc7\x6e\xa4\x2e\x97\xf4\x2b\x71\x06\x91\x1f\x98\x46\xe8\xc3\x34\x62\x6b\xba\x02\x63\x81\x69\x48\xf8\x39\x34\x21\x86\x16\x01\x43\x22\xd0\xd3\x8d\xa9\xa3\xda\x14\x5a\x39\xa1\x1f\x38\xfe\x9d\xa7\xb8\xa0\x84\x24\x4c\x1b\x73\x61\xe3\x28\xf2\xbd\xbc\x90\x20\xc8\x6a\xe5\x12\xb0\x3b\xd7\xb5\x03\x46\x1c\x35\x93\xa2\xb1\xd4\xa9\x68\x84\xb2\x88\x56\xea\xb6\x1d\xae\x70\xa2\x7e\x90\xb4\x92\xc7\x19\xb6\x7c\xbe\x03\x3b\x61\xcb\xc2\x81\xef\xb9\xf7\xf9\x26\xd0\x68\x26\xe4\x48\x95\x01\x3a\x87\x6e\x35\x94\xd0\x2e\x07\xc0\xb6\x44\xea\xc5\xf4\x1c\x0a\x25\xe6\x26\xc1\x36\xa8\x93\x4e\xd7\x3c\xb4\x3d\x50\x79\x61\x12\xd5\xa4\x48\xb5\xca\x46\xb9\x99\x50\x4d\x4d\xc3\x0e\xa9\x3d\x20\x5f\x81\xbb\x43\x80\x72\x14\xc6\x24\xf1\xbe\xe2\x44\xa0\xc1\x20\xfb\xb5\x04\x87\xd2\x30\xca\xe8\x71\x30\x84\x4e\xc5\x71\xd9\xb9\x21\xc5\x6e\x69\x00\x1b\xe2\xc5\xca\xbe\xc5\x6f\x2e\xa6\x6b\xcf\x89\xeb\x12\x67\x7e\x5f\x54\x41\x81\x83\x4b\x4b\x24\xa5\x45\x4b\xa2\x41\x48\x18\x80\x5a\xce\x6d\x92\xee\x61\xec\x79\xd4\x5b\x1d\x0c\x5d\x5a\xa6\xdb\xa0\x3f\x4c\x52\x69\x46\x00\xa5\x8a\xed\x32\xfa\x1a\x4d\xae\x67\x57\x93\x13\x54\x97\x5d\x36\x8d\x92\x18\xc3\xd8\xcd\x69\x30\xa7\xe3\x74\x9a\x8a\x6e\x95\x42\x82\xb8\x61\xd4\xf8\x5d\x76\x5a\x10\x9a\x64\x53\xf8\xa3\xd1\x35\x32\xa3\x81\x59\x84\x06\xcd\x98\x93\x11\x4f\xa6\x38\x58\x9d\x82\x0b\xaa\x2b\x6b\x57\x0d\x4f\xb5\x70\x6d\x1d\x67\xa3\xe0\x0d\xad\x1c\xa0\x3c\x53\x09\x57\xe0\x56\x7a\x9a\x19\xd7\xa7\x63\x9d\xcd\x6b\xa6\xa9\xac\xe6\x7d\xac\x5f\x27\x4d\x47\xf3\xe7\x24\xe2\x20\xa2\x1b\x52\x9a\xa0\x84\x7c\x43\x32\x12\xae\x04\x35\x0d\x4e\x69\xbd\x22\xa7\xcf\x1b\xde\x55\xe7\x18\x52\xa7\x5a\x09\xf3\x0e\xa2\x6f\x59\x76\xa3\x2a\x8b\x57\x3f\x43\xba\x5a\x47\xd5\x50\x53\x67\xf6\x19\xda\x2e\x65\xd1\x80\x7a\x90\x06\x91\x4a\xaf\x52\xa4\x2b\x74\x2a\xe0\xf9\x07\x4c\x02\xf4\xda\xc3\x47\xed\x34\xf6\x18\xaf\x80\x55\xb1\x0a\x58\xbf\x9c\xb6\xab\x2a\x4e\x5b\x7b\x05\x60\x1f\x3d\x85\x29\x94\x3d\xc6\xb0\x17\x11\xdd\x96\x43\x5c\x21\x02\x6f\x49\xc8\xd0\xfa\x2b\x62\xf0\x3e\x61\x17\xab\x01\xf6\x7e\x38\xbc\xbb\xbb\xb3\x20\x8e\x86\xf0\xbf\xb5\xf0\x37\x43\x91\x87\x43\xd6\xea\x12\x9b\x11\x36\x74\xed\x88\x30\x5d\x54\xb9\xbc\x9e\x8d\x2f\x66\xba\x50\xd2\x0d\x6c\xd4\x58\x9f\x04\x71\x24\xf1\x3e\x20\xa7\x1f\x35\x72\x6a\x0a\x86\x3e\x8f\xaf\xae\x27\x97\x17\xcf\x83\x43\x85\x70\x9e\xb9\x84\x0b\xc0\x0a\x59\x31\xa8\x88\xc4\x6b\x01\x7d\xb6\x6f\x04\xb6\xe3\x40\x2a\x83\xf5\x02\x23\x5b\x18\xf7\x3d\x5b\x6b\xcb\x84\xd0\xbf\x2b\x96\x06\x48\x7b\x43\x36\x90\xa6\x8b\x3a\x54\x34\x9c\xbb\xfe\xe2\x16\xf3\xfd\xc1\xc6\x19\xbc\x55\x3f\xfc\xe5\x12\x6a\xae\xc1\x9b\x5c\x0a\xb1\xdb\x45\x64\x13\xa0\x7e\x0d\x13\xaa\x18\x10\x8b\xbd\x7f\x2f\x7e\xfc\x06\x86\x4f\x5c\xa0\x6e\x1a\x56\x31\x60\xe6\xe4\xca\xb1\xfb\xcf\x7c\x86\x92\x69\xc6\xc9\xbd\x29\x61\x33\x1a\x71\x31\xeb\x11\xa6\x6d\xf0\x1e\xdc\x48\x41\x3f\xc6\x3c\x82\xbc\x1c\x59\x99\xc6\xd2\x0f\xc1\x8a\xe3\xcd\xe6\x9e\x2e\x41\x9d\x4b\xba\x5a\xac\xc9\xe2\x76\xee\x7f\xe5\xcf\x06\x0c\x4b\xd7\x1f\xd4\xb3\x9a\xe4\x0f\x69\x0a\x27\x10\xa4\x79\x3d\x49\x7e\x37\xcc\xa5\xed\x62\x36\xf6\x8f\x2f\x0c\xca\x29\x61\x5f\xa7\x2e\x85\xc9\xb1\x3e\x52\x87\x08\xba\x93\xb3\xd7\x50\x7a\x0b\x0c\x52\xe5\xa4\x42\x48\xba\xac\x49\x25\x25\xa5\x99\x3d\x9f\x9c\x59\x33\x2c\xd9\xf5\x59\x25\xf5\x82\x38\x12\x50\xa6\x1f\x6b\x51\x6d\xe9\x13\x51\xe1\xa5\xd7\x1c\x30\xd6\xd4\x71\x88\x97\x20\x19\xbf\xd2\xd8\x3e\xea\xfe\x95\x4e\x92\xbe\xd9\x71\xf3\xa8\xcc\x1d\x93\xe9\x53\x7c\x93\xc9\x11\x04\x1f\x9f\x9b\x44\x50\xc1\xbb\x50\xf2\x67\x79\x82\xd1\x6c\x14\x9b\xb5\x1f\xd2\x3f\xd0\x47\xdd\x01\xde\xd6\x23\x58\xc6\xae\xb1\xd1\x60\x15\xfa\x71\x30\x40\x57\x26\x4e\x05\x4e\xe5\x5c\x06\xec\x8e\x77\x31\x92\x5f\x03\xb6\x29\x44\x22\x51\x5a\x56\x04\xf5\x92\xef\x00\x21\x4e\x4c\xad\x45\xc0\x98\xa9\x57\x11\xa7\xb3\x93\x5a\x98\xba\xea\xf6\x1f\xb9\x12\x75\xf6\xd1\x54\xa6\x06\xde\x35\xe6\xe1\x16\xd1\x4c\x78\x97\x43\x19\x5e\x38\x4d\xfc\x6b\xef\x91\x15\xfc\x52\xc8\x80\x39\xe2\xe4\x8c\x57\x40\xb5\x03\xd6\x16\x12\xb5\x8f\xf6\xb2\x9e\xdd\xee\x8e\x46\xeb\xbc\xb2\xae\xc8\x12\x42\xed\x1a\xa5\xd4\xe2\x78\x28\x9e\x73\xfc\x96\x0a\xac\x96\x69\x88\x92\x68\x5c\xa3\x94\x71\x63\xb0\xd1\x1a\x65\x36\xc5\xa0\xcb\x01\x03\x81\x17\xeb\x64\x1d\x69\x8e\x3e\x9d\x94\x96\x67\xb3\xe3\x93\x6b\x8b\x9e\x4d\x8f\x4f\x7f\x19\xcf\xae\xad\x1b\xea\x45\x3a\xd7\x14\x74\xed\x74\xda\x03\x7b\x71\x4b\xa2\x6a\xa7\x98\x8a\xe7\xfa\x1c\xa0\x2a\x0b\xd8\x4f\xf6\xf1\xd5\xd5\xe5\x55\x0b\xd1\x49\x18\xfa\x61\xb5\xe4\x63\xfe\xb8\x17\xc1\x55\x62\x5c\x2f\xff\xc9\xaf\xb3\x71\x0b\xf1\xe7\xf7\x90\x84\x55\x4a\x7f\x82\x4f\xdb\x0b\x5f\xce\xbe\x2a\xab\x70\xcd\x2a\x55\x7d\xf8\x50\xa6\x92\x51\x0f\xa8\x82\x83\x10\xd4\x75\x59\xf0\xd1\x9a\xe3\x6b\x1d\x3e\x89\x98\xcd\x1f\x03\x3a\x55\x04\x18\xa9\x76\x0c\xec\x35\x26\x2f\x03\x7f\x47\x07\x91\x8b\x97\xb9\x11\xcb\x74\x50\xbf\x06\x52\x9d\xef\x25\xfd\x7f\x93\xe9\x24\x4f\xfb\x4a\x8b\x9e\x35\x53\x94\xea\x5e\xda\x7a\x0b\xd5\x67\xbc\xe9\xa9\x34\x2f\x59\x34\x53\xfc\x23\xde\x9d\xd3\xbb\x18\xed\xde\x6a\x17\xdd\x7b\xd0\xba\x70\xd1\x16\x4a\x4f\x21\xe0\xa9\x74\x2e\x38\x34\x53\x79\x3d\x20\xe5\x34\xce\x47\xba\xb7\xc2\x79\xef\xbd\xf5\xad\x5b\x2e\xcf\xdf\xaa\x5c\xee\x5d\x04\x71\x6d\xad\xd6\xae\x36\x03\x6a\x2f\xb9\x36\x73\x6a\x6a\x33\xe7\x89\x6b\xb3\x51\x4d\x6d\xe6\x34\xab\xcd\x46\x6d\x6a\xb3\xf2\x58\xff\xd2\xb5\x99\xd3\xa5\x36\x1b\x7d\xaf\xcd\xfe\xe4\xb5\xd9\xa8\xb6\x36\xd3\xf9\xd7\xde\x23\xab\xac\xcd\x46\x7f\x8a\xda\x6c\xf4\x02\x6b\x33\xa7\xbe\xbe\x19\x89\x70\xec\x9c\x4d\x2e\x2e\x47\xcd\x0b\x04\x07\x06\xe6\x3b\x35\x15\xc2\x84\x3f\xee\x52\xdf\x38\x8d\xea\x9b\x44\xfe\x56\xf5\x8d\xf3\x67\xab\x6f\x12\x7d\x67\xb4\x53\x91\xee\xe9\x66\xb4\x2a\xdf\x1b\xed\x97\xef\x15\x59\xd4\x26\x7c\x4d\x2d\x2c\x93\xf1\xa9\xd1\xee\x99\xf2\xa9\xee\x9d\x73\x6c\xa7\x90\x63\x37\x50\xfa\xe3\x39\x76\x57\x9d\x37\xc8\xb1\x1b\x3a\x45\x4e\xe3\x5d\x72\x6c\xe7\xf9\x73\xec\xfc\x3b\x18\x75\xd1\x64\xa3\x55\x93\xb7\x2c\x01\xab\x4b\xdc\xdf\x15\x5e\xb2\xb4\x4c\xe4\x03\xa6\xcd\xe3\x9f\x6e\x44\xed\xc7\xf0\xe4\xd5\xc2\x76\x55\x5d\x2d\xa8\x67\x4f\x55\x2d\x7c\xfe\x50\x5d\x2d\xe8\xde\x75\x1f\x7d\x16\xaf\xb9\x8d\x95\xeb\xcf\x71\xb7\x66\x64\x47\xb1\x26\x3a\x94\xca\x84\xf2\x20\xff\xd2\x65\x42\x32\x6f\xfb\x94\x09\x7c\x52\xf6\x2b\x13\xc4\xd6\x0d\x5e\x22\x18\x73\x3f\xe4\x46\xe1\x43\xc2\xbe\x11\xb7\x30\x7d\xd4\xed\x14\xd1\x98\x7a\x0f\x29\xe0\xe7\x0f\x7d\xa4\x80\x2f\xbf\xa8\xc9\xbb\x7c\xde\x7d\x9f\xbd\xd4\xd9\x27\xf3\xaf\xca\xb2\xfb\xc8\xd6\xb6\x6d\x7d\xa0\xa5\xf5\x27\x61\x7b\xbb\xda\x33\x62\x6f\x57\x2f\x20\x56\xd7\x6e\x1b\x57\x1d\xb6\x76\x68\xa0\xba\x8c\x43\x43\xa6\x37\x0f\x0f\x3f\x17\xf7\x6f\xcb\x6d\xdb\xb8\x93\x9b\x6f\xfa\xdf\xed\x86\xff\x7c\xf5\xcf\x21\x7a\x9a\x18\x70\x59\x03\x0b\x97\x06\x26\x0c\x5e\x18\x24\x90\x9e\x8c\x8e\xa3\x28\x84\xc8\x00\x17\x7c\xf2\x1e\x1e\x3a\x42\x27\x88\xf1\x85\x7d\x3d\x75\x7d\x06\xfc\x39\x1f\x64\x28\xdc\x0a\x18\x9e\xf9\xa1\xe4\x98\x65\x88\x62\xe1\xd3\xf3\xff\xbd\xc6\xdd\xf8\xa2\xc5\x2b\x51\x1b\xcf\x40\x47\x69\x1d\x7c\x24\x2d\xa4\xc9\x60\x83\x18\xf2\x02\x66\x62\x1b\x0e\x5d\x78\x05\x61\x5f\x80\x91\x40\x24\xbe\x73\x7e\xb7\x83\x48\xb7\x22\xc6\xdf\xe8\xbf\x8c\xbf\x2d\xfc\x90\x18\xef\x0f\x95\xb9\x4e\x6f\xac\x4f\x94\xa1\x00\x51\xb8\xdb\xdd\x92\x7b\x83\xaf\x70\x02\xa9\xf9\xfd\xe0\xc2\x14\xed\xad\x0b\x21\xed\x41\xe4\x28\x17\x48\x01\xd8\xf0\xfc\xbb\xd0\x96\x7b\xe4\x40\xfc\xa4\xc7\xc1\x30\x72\x2a\x7b\x25\xa7\x0c\xd0\xea\x53\x55\xa9\xee\x37\x8c\x84\xa7\xe2\xde\xab\xc4\xf2\xd3\x47\x39\xfb\x3e\xea\xc2\xe8\xfa\x9e\x49\x3e\x46\x91\x11\x3c\xea\x8f\xcf\xc4\x71\x49\xc5\x80\xf0\x51\x05\xa3\x61\x14\x66\x2c\x22\xff\xa7\x89\x7d\x08\x3c\xc0\xc9\xc5\x1f\x89\xd0\x78\xf1\xc6\x60\x11\x78\x1b\x49\xb1\x06\xc6\x95\x39\x4d\x03\x9c\x95\x77\x47\xeb\xac\xdf\x27\x57\xf8\xe0\x11\x73\xe0\xcd\x70\xc2\x72\xfb\xf3\x63\x8f\xe6\xc2\xf6\xd1\xdf\x0b\x68\xb5\x0f\x17\x98\xad\xa7\x67\x82\x53\xd5\x85\xcb\x50\x6a\x95\xdf\x55\xc7\x53\xa2\xf4\x7c\x4a\x1d\xba\x2b\x77\x17\xe8\x1e\xa5\xe8\xc8\x27\xb0\x0d\x70\xa8\xca\xae\x25\x7a\x38\x94\xdd\xa6\xc8\x31\x3a\xe3\x44\x4a\xe8\xa1\x8a\x4e\x81\x20\x0e\x0d\xf9\x11\x06\xd1\xdb\x1a\xd1\xf0\x82\x1f\x58\x28\xb8\x51\x01\x42\x6a\x94\xc0\x91\xf4\x1f\xf8\xc7\x78\xf3\x16\xb9\x0d\x04\x3b\x87\x6c\x73\x8c\xc8\x16\x19\xbd\xd6\xa2\xd0\x5e\xdc\xfe\x2b\xcb\x4c\x3b\x2a\x3d\xb3\x12\x42\x80\xdb\xf3\x5e\xc7\x5b\x9b\xba\xed\xba\x80\x27\xc1\x0c\xff\xbb\x37\x67\xc1\xcf\x07\x2c\x0e\x52\x24\xa9\x1e\x43\xcc\x48\x40\xc2\x05\x24\x25\x30\x92\xf4\xc2\x48\x28\x4e\xc5\x0d\x2e\xbd\xa2\xd9\x08\xea\x94\x54\x33\x3f\xb2\xb3\x03\xe9\x04\x5a\xf9\x55\x87\x8e\xc8\xa5\xd5\x4b\xb4\xc6\xcd\xb3\x78\x3e\xce\x34\xd4\x2f\x69\xd3\x9f\xe0\x27\xb3\xa6\x76\x68\x6f\x98\x35\xbe\xb8\x39\xb7\x9c\xa5\x61\x9e\x5d\x9b\x86\x69\x56\x54\x00\xad\xa9\x9d\x4f\xfb\xa4\x76\xfc\xf9\x78\xf2\xa9\x37\x6a\x37\xd7\xe3\x51\x6f\xc4\x66\x97\xb3\x63\x14\xad\x13\xe8\xe5\xa0\xaa\x0f\xe4\x93\x8b\x88\x1d\xa1\x4f\x50\xd1\x60\x9f\x5c\xe2\x7c\x36\xf0\x93\xfc\x9e\x09\xfd\x2a\xb8\xed\x05\x7f\x13\xa8\x9c\x49\xcb\x2e\x71\xef\xf8\x87\x24\x3b\x03\xa0\x78\x2d\xd2\x1f\x02\xe6\x16\xba\xbb\x27\x6f\xc9\x9b\xd7\xc2\x69\xd3\xa3\x11\x3f\xb7\x5c\x91\x06\x95\x5a\x9f\xfb\xb1\x17\x11\xa7\xb6\xb9\x76\x5d\xe6\x88\x47\xba\x3d\xfa\x61\xb8\xdb\xa3\x1b\x8f\x47\xbd\xa5\x5d\x79\xc8\xe8\x0e\x41\x74\xbf\xe4\x8b\x2e\x53\xfc\x99\x54\xa4\x5e\x34\x97\x7a\x49\x17\xa5\x4b\x2b\x87\x3a\xed\x61\x86\xaa\xb4\x27\x4f\xb1\xb9\xbf\x43\x87\x11\x71\x23\x7b\xe2\xb5\xee\x72\x19\x47\x6d\xfa\xb4\xe3\x90\x27\xde\xc9\x69\x69\x9f\x69\x4b\xae\xe0\x9a\x80\xd3\x85\x4b\xbb\xd2\x4f\xcb\xe5\x8a\x21\x5f\x32\x9d\x4c\x66\xd7\x06\xc0\x9d\xc1\x70\x39\x36\x7b\x26\x7d\xe2\xd5\x17\x31\x07\x85\xb3\x33\xf3\x83\x61\xf6\xce\x11\x1e\x46\xdb\xb3\x96\x7a\x5c\x38\x98\x95\x6f\x2e\x5d\x84\x10\x62\xf0\x17\x6d\xc6\xc6\x77\x62\xd7\x37\x7e\xfc\xd0\x41\x81\x27\x05\x11\xff\xfe\xe3\x87\x27\x97\xb1\xb5\x1e\xdb\x0a\xd9\x05\x55\x69\xbf\x79\x5d\xb2\x01\xb3\x13\xaa\x0a\x1a\x1a\x58\x95\x9b\x43\xfb\xc7\x55\x49\xf8\x3b\xb0\x3e\x3e\xb3\x2f\x09\x59\x2b\x40\x55\x25\xa6\x75\x4e\xd7\x03\x38\x55\xa1\xe6\x33\xb1\xd7\xe3\x61\x13\xe6\x7d\xa0\x5e\x05\xd2\xf5\xc1\xbe\x1b\x9e\x65\x01\xa8\x0f\x40\x53\x1b\xf9\x3b\x21\x9a\x24\xa2\x81\x34\x75\xce\xa0\x7f\x4c\x53\x94\xbf\x83\x5a\x83\xd9\xfd\x8e\x6a\x2f\x04\xd5\x44\x36\xf5\x6d\xb1\xad\x5a\x86\x17\x81\x70\x39\x44\xea\x0e\x71\x1b\xb2\xd9\xe3\xf5\x25\xf4\x4a\x01\xee\x7c\x7c\x5e\x82\x36\x3c\x4b\x2f\x40\xed\x96\xe2\x07\x99\xb0\x87\xf5\x0b\x45\xa1\x0a\xa8\x96\x79\xd2\x10\x05\xb0\xc3\x59\x9b\x65\x2b\xec\xd0\xeb\xa2\xbd\x24\xd8\x65\xc9\x0a\x49\xf4\xba\x64\x0f\x04\xbf\xc1\x7b\xc6\x8c\x77\xe0\x94\xb4\x69\x5f\xb3\xae\xf4\x0c\x0b\x4a\xca\xee\xbb\x7b\x50\x72\x82\x08\x08\x69\xf7\xf6\x15\xb6\xfb\x54\x6d\xe1\x93\x43\xb1\x0d\xdc\xd7\x72\x32\xe5\xaf\xa9\x71\x8b\x82\x6e\x79\xff\xe4\xf2\xf2\x93\xc5\xd6\xfe\x9d\xdc\x05\x17\xc4\x46\x69\xe3\x5e\xb2\xb7\xbb\xf4\xc4\xe4\xfb\x22\xc4\x66\xa5\x8f\x21\x59\x36\xe7\x92\xbe\x36\x3f\x9d\xde\xa8\x6d\x7a\x99\x7d\x72\xe9\xde\x9d\xdc\x51\x2d\xb5\x79\x6d\xcf\x91\xc1\x10\x0a\x1b\x78\xcc\x54\x10\xb1\x5f\x2d\xb3\xcd\xa3\x70\xb4\x45\xb3\x6f\x0d\x05\xc1\xbf\xc7\x0b\xfc\xb2\x4a\x95\x24\xff\x13\x93\xf0\x3e\xe1\xc3\x2b\xf5\xec\xce\x3f\x21\xc8\x80\xc5\xf3\x0d\x4d\xbe\xe9\x28\xaf\x32\x7d\x9a\x6e\x37\x6b\xb4\xc9\x4c\x7c\x16\x86\x5f\x17\x37\x8d\x36\x37\x9b\x35\x6e\x0c\x44\xb5\x26\xf6\x21\x4c\x24\xbb\x07\xad\x95\x85\x48\x82\x19\x55\x49\xa1\xd4\x0e\x33\x5b\x2f\x6b\x51\x54\xdc\x05\x93\x1e\x8d\x81\x2b\x2e\x99\x3a\x1c\xd3\x51\x48\xf1\xf9\x21\xad\x98\xc9\xc9\x98\x84\x3f\xe2\x0f\x5c\x95\x0f\xac\x35\x3b\xa0\xd8\x68\xbe\xeb\xc0\x29\xdd\xd0\x58\x1e\xcd\x74\x7c\x35\xb9\x1c\x59\xb2\x09\x1f\x50\x69\x53\x58\x66\xff\x5d\x6e\x57\x98\xfc\x59\xf4\xd2\xd6\x96\x53\xed\x8d\x8f\x6c\x9d\xc8\xee\x8c\xd3\x88\xc5\x7f\xb4\x85\x5e\xfc\xb0\xce\xb7\x83\x5e\xcc\x84\x9e\x1e\x7a\x81\x4b\xaa\xe3\x73\xb2\xf1\xc3\xfb\x86\xe8\x0b\x3d\xf7\x47\x5f\x3e\xb8\xef\xe8\xdb\x37\xfa\x72\xb5\xf6\x89\xbe\x59\xf3\x48\x84\x7a\x04\x7d\x1b\xca\xca\xee\xec\xa0\x5f\x61\x91\x22\xef\x43\x97\x23\x89\xed\x7b\x0f\x90\x18\x48\xed\xc5\xa3\x34\x17\xff\x39\x51\x9a\x5b\xd8\x7e\x28\x9d\xd4\x11\xfd\xa2\x74\xc0\xbe\x21\x48\x07\xec\x19\x30\x3a\xc8\x6c\x2a\x9d\x86\xfe\x82\x40\x19\xcd\x1a\xa2\x74\xc0\xf6\x07\x69\x1c\x5c\x3b\x8c\x6e\x7e\xaa\xe3\x25\x42\x75\x3f\x2e\x19\xb0\xda\xcf\x3d\xbc\xa0\x68\x81\xf3\xdb\x27\xfe\x66\xec\x34\x0b\xa5\x6d\x32\xf5\xe9\x35\x87\x08\xf8\xeb\xf9\xd1\x88\x2c\x42\x62\x73\x20\xcf\x65\xeb\x55\xa2\xba\xe0\x18\x75\x82\x7e\x9a\x9c\x4f\x66\x56\xc0\xbc\xb2\x98\x83\xfd\x24\x4c\x8b\x8a\x46\x02\x42\x76\x55\xab\xc9\x1a\x01\x55\x29\x91\xb2\x0f\xdc\x98\xa9\xb3\x05\x6d\xa3\xd4\x53\x46\x08\xb4\xaa\xfd\x02\x44\xc0\xfa\x8e\x0f\xac\xe5\x0a\x64\x00\xe8\x9a\x2e\x3f\x2a\x5d\x17\x97\x20\x03\xf5\x5a\x25\xa0\xb8\x00\x89\x9d\xac\xe9\x64\x84\xeb\x60\xde\xaa\xe6\x24\x85\x0c\x34\xa0\xcd\xa4\x4f\xb3\x95\xc6\xa4\xc7\x4d\xa1\x87\x7c\x9c\xf9\xaf\xcd\xcb\x9c\x80\xc5\x0c\x10\x5a\x37\x0a\x79\x8d\x5b\xfb\x1b\xbe\xd9\x49\x47\x15\x52\xc8\xd4\xa3\xfb\x96\xdd\x2e\xe8\xa2\xe1\xba\x6b\xd2\xe5\x9a\xfe\xd1\xb6\xcb\x15\x61\x60\xa2\x5e\xcd\x6b\x24\x5c\x77\x55\xfb\xe2\x8e\x54\xb7\x99\xfc\x7a\x76\xfd\x66\xcb\x56\x13\xf0\x23\x9f\x00\x7c\x89\x56\x37\x01\xc5\x57\x6b\x9d\x16\x72\x95\x7f\x69\xd7\x71\xdf\x3e\xe3\xae\x6b\x84\x08\x18\x69\x2f\x5b\x91\x91\xd6\x4d\x9f\xb4\xae\xc7\x57\x7d\xed\xe0\xe6\xe3\xbc\x9a\xf4\x46\xeb\x62\x72\x3a\xee\x8d\xd8\xe7\xc9\xd5\xac\x37\x62\x57\xe3\xeb\xde\x68\xcd\x26\xe7\x63\x4c\x3b\x85\x23\xf6\xa6\xbb\x63\x4e\xb5\xeb\x86\x75\x19\x52\xba\x2f\xee\xab\x84\x52\xd6\x2e\x10\x82\xc4\x0d\x1e\x64\x0d\x4b\x9e\x3d\x94\x09\x60\xe6\x37\xa4\x80\x86\x6c\x2a\x6e\x98\x32\xe2\xe4\xfe\x85\x94\x4c\x7b\x48\x32\xd5\x47\xbe\x8f\xe4\x79\xe7\xe4\x25\xe2\x51\x72\x92\x32\x49\xe1\x79\x7e\x8a\xdf\x2c\x80\x34\x33\x61\x24\xce\x41\x66\x0e\x7d\x9a\x06\xe8\x67\x41\xd6\xbe\x0b\x39\xfe\x21\x04\xd3\x69\x7a\xc9\x53\x9d\x80\x40\x18\x70\x10\xbe\x44\xc2\x90\xb9\xf9\xd9\x76\xe3\xe4\xee\x2b\xae\x6f\xcd\x89\xcb\xb6\xb1\x3e\xf3\x4a\x0d\x7a\xcb\x53\x7a\x56\x92\xb2\x24\xc7\x31\x39\xf7\x87\x07\x75\x8e\xab\x05\x87\xed\x8a\xef\x01\xe1\xf0\x19\x1a\x90\x09\xe0\x51\x5e\x71\x2b\x89\x0a\xbe\x8b\x54\x0f\xcd\x9f\x92\x00\x22\xbf\x7c\x2e\xbe\x46\xc0\x3f\x20\x5a\x09\xe7\x0d\x24\x48\xf2\x19\x99\xb0\x6c\xec\xc5\x3a\x4d\x58\x24\x97\x73\xb8\x89\xfd\x4a\x89\x8b\xfc\xf4\xbf\xcc\x5e\xe2\x98\xc7\x1d\x24\x61\xdd\xf0\xcc\x42\x97\x55\x34\x8d\x67\xff\x9d\x92\x2f\x12\x7e\x5d\x9d\xb3\xe0\xfb\x4a\x6c\x76\x61\x17\x02\xac\xb6\x15\x14\xbf\x5b\xea\x90\xb0\xae\x65\x33\x69\xdf\xa5\xd2\xe2\x97\x21\x48\x56\x5c\xf9\xf3\x1a\xef\xd7\x8a\xde\x8c\xd5\x4f\xff\x91\xf2\x5a\x52\x97\x04\x76\xb4\xd6\xb0\x93\x73\x87\x2d\x7e\xc3\x26\x8f\x46\xfd\x06\xd6\x22\xbf\x21\xd5\x20\xff\xad\x3c\xef\x9e\x35\x5c\xe7\x75\x05\x18\x17\x1c\xc4\xc8\x91\xae\x6b\x9f\x60\xe9\xbe\x29\x8d\x3a\x12\x9f\xa4\x34\xe2\x80\x7d\xfe\x3b\x30\x5d\xde\x53\xd3\xe2\x9b\xe4\xf4\x11\xe6\x6e\x95\x0f\x03\x69\xaa\x95\x0d\xb8\xd9\x55\x3e\x75\x68\x48\x16\x11\x5f\x8c\xef\xe9\xb5\x74\x62\x0d\x1d\x42\x57\x21\x66\x95\x83\x4d\x7d\xa0\xea\x1c\xa6\xb2\x4b\x28\x2d\x43\x14\xff\xcc\x92\xa8\xc5\xa5\xbd\x67\x3e\x06\x25\xee\x27\xdf\x7e\x9a\xf2\x30\xf5\x1a\xbf\x30\xb3\xc5\x68\xc1\x3b\x8b\x9b\x89\x6d\x9a\xed\x0a\xd2\x4c\xba\x82\xa6\xba\xce\xbd\xd8\x10\x67\x5d\x80\xc7\xb1\x4b\x57\x9e\x3c\x45\x2e\x18\xc8\x8f\x9f\xc1\x33\x5c\x37\xc8\x7d\x75\x40\x45\x37\x59\x5b\x67\x3e\x35\x50\x3a\xa2\x6e\x9d\xda\x21\x89\xb2\xad\xb9\xf0\xd9\x03\xe9\xb6\xb4\x85\x75\x32\x9c\xff\x0f\x00\x00\xff\xff\x6f\x39\x54\x8c\x94\x71\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 29076, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
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

