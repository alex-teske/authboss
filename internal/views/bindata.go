package views

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
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _layout_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x95\xc1\x72\xd3\x30\x10\x86\xef\x79\x0a\x8d\x0e\xdc\x2c\xd1\x26\xc0\x4c\xeb\xe4\x06\x2f\xc0\x13\xc8\xd6\x3a\x56\x90\x25\xa3\x5d\xb7\xc9\x78\xfa\xee\xc8\x96\x1b\x5c\xd7\x40\x87\x03\x33\xe4\x60\x67\x37\xda\xff\xdf\xfd\x34\x52\xf2\x9a\x1a\x7b\xd8\xe4\x35\x28\x7d\xd8\xb0\xf8\xc9\xad\x71\xdf\x58\x1d\xa0\xda\x73\x29\x1d\x90\x76\x4a\x14\xde\x13\x52\x50\x6d\xa9\x9d\x28\x7d\x23\x2b\xef\x28\x53\x8f\x80\xbe\x01\xb9\x13\xef\xc5\x56\x96\x88\x2f\xd2\x22\x26\x38\x0b\x60\xf7\x1c\xe9\x62\x01\x6b\x00\xe2\x4c\xae\xd9\x34\xea\x3c\x28\xbf\xb2\xb9\x26\xe4\x56\x6c\xc5\xcd\xe8\x71\xcd\x89\xc6\xb8\x3f\x98\x60\x19\x4c\x4b\x0c\x43\xb9\xe7\x35\x51\x8b\x77\x52\xaa\x93\x3a\x8b\xa3\xf7\x47\x0b\xaa\x35\x38\xfa\x0c\x39\x69\x4d\x81\xf2\xf4\xbd\x83\x70\x91\xb7\xe2\x26\x8e\x94\x82\xd1\xe7\x84\xfc\x90\xcb\xa4\xb7\x22\xfe\xd6\x11\x6e\xe5\x69\x39\x41\x54\x66\x74\x69\x61\xcf\x09\xce\x24\x4f\xea\x41\x25\xe5\xb9\x61\x2e\xd3\x0e\xe5\x85\xd7\x97\xf8\xd2\xe6\x81\x95\x56\x21\xee\x79\x19\x99\x2b\xe3\x20\x64\x95\xed\x8c\xe6\xa9\xbb\xbe\x37\x15\x13\x5f\xe2\x92\xfa\x6b\x57\x96\x80\xf8\xf4\x94\xda\x9e\x95\x06\xff\x38\x2d\x5f\xfe\x52\x7a\x9b\x9d\x31\xf3\x55\x85\x40\xd9\x96\x0d\x71\xa3\xb3\x8f\xb3\xe5\xcb\x12\x65\x21\x10\x1b\x9f\x19\x26\xcb\x29\xd2\x06\x1b\x83\xa8\x0a\x0b\x9c\x8d\xdb\xb4\xe7\x8d\x0a\x47\xe3\x32\xf2\xed\x1d\xfb\xf4\xa1\x3d\xdf\x2f\x94\x47\xf5\xa2\x23\xf2\x6e\xa2\x93\x02\x7e\xed\xd0\x7a\x8c\x72\x5a\x91\x7a\x36\x98\x7a\x88\xdc\xb0\x55\xee\xf0\x8e\x4c\x03\x78\x1f\x21\x0e\x51\x2e\x93\xc0\x6b\x9b\xbe\x6f\x83\x71\xb4\x4a\xeb\xda\x8a\x8c\x93\xce\x50\xfd\x0c\x67\x5f\xfb\x1e\x9c\x9e\x0a\x67\xfc\x3f\x87\xe0\xc3\xbf\xa4\xaf\x95\x3b\x42\xf8\x1f\xe1\xcf\x51\xfd\x35\xfa\x25\xe3\xb5\xa9\x6f\x17\x53\x2f\xe9\x47\xda\x13\xfd\xdd\x33\xfd\xdd\x6f\xe8\xc7\x19\xc1\xb2\xf1\x99\x69\xa8\x54\x67\x69\x0d\xe9\xb2\x22\x1b\x0e\xb5\x71\xc7\xe1\xa4\xbf\x18\xf2\xd7\x15\xc3\xf9\x5f\x91\x4e\x0c\x08\x9a\xd6\x2a\x02\xc6\x55\x47\x75\xe1\x87\xbb\x51\x2c\x70\xae\x20\x7d\x13\xe5\xeb\x6b\xba\x81\x64\xfa\xe7\xd8\x6c\x7e\x04\x00\x00\xff\xff\x06\xe2\x8e\xc9\x43\x06\x00\x00")

func layout_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_layout_tpl,
		"layout.tpl",
	)
}

func layout_tpl() (*asset, error) {
	bytes, err := layout_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "layout.tpl", size: 1603, mode: os.FileMode(438), modTime: time.Unix(1423332529, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _layoutemail_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x49\xcd\xb5\xab\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x2c\x2d\xc9\x48\xca\x2f\x2e\x56\x52\xd0\xab\xad\xb5\xd1\x07\xca\x02\x02\x00\x00\xff\xff\x3a\xdb\x96\xd1\x22\x00\x00\x00")

func layoutemail_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_layoutemail_tpl,
		"layoutEmail.tpl",
	)
}

func layoutemail_tpl() (*asset, error) {
	bytes, err := layoutemail_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "layoutEmail.tpl", size: 34, mode: os.FileMode(438), modTime: time.Unix(1423334931, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _login_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x94\x4b\x8e\xa3\x30\x10\x86\xf7\x23\xcd\x1d\x2c\xef\x09\x17\x00\xa4\x59\xcc\x6e\x46\x13\x4d\xba\x0f\x60\x4c\x11\xac\xd8\x2e\xcb\xd8\x79\x08\x71\xf7\x36\xcf\x06\xf2\xe8\x6d\x47\xb2\x54\xaa\xd4\xff\xa7\xfe\x2f\x86\xa4\x44\xab\x08\xe3\x4e\xa0\x4e\x69\x2c\xf1\x28\x34\x25\x0a\x5c\x85\x45\x4a\xf7\xff\x0e\x6f\x34\xfb\xf9\x83\x84\x4f\x52\x88\x33\xe1\x92\xd5\x75\x4a\x3b\x51\x74\xb4\xe8\x4d\xd3\x88\x92\xec\x7e\x5b\x8b\xb6\x6d\x49\xc5\xea\x08\xba\xba\x69\x40\x17\x6d\x3b\x69\xb7\x7a\xa1\x8d\x77\x83\xc1\x72\xa4\x1f\xab\x0d\xd3\x0f\xe6\x22\x56\x14\xa8\x69\x96\x88\x79\x09\x46\x4a\x16\xf9\x1a\x6c\xe8\xc6\x22\x9c\x4e\xba\xb5\xeb\x2d\x88\xbb\x19\x48\xa9\x83\xab\xa3\xab\x0c\x1c\xb5\xb3\x28\x29\xd1\x4c\x85\x81\xce\xac\xab\x28\x31\x92\x71\xa8\x50\x16\x60\x53\xfa\x3e\xb7\xcf\x4c\xfa\x30\xd7\x34\xbb\xa9\xb7\x09\x19\x87\x94\x13\xb0\x65\xfd\x3d\xe1\x49\xe4\xa7\xaf\xe1\x8d\xf4\x4c\x10\x5e\xd0\x16\x2f\x09\x7e\x0e\xad\x08\xee\xa7\xf6\x13\x56\x77\xcb\x57\x20\x4d\x94\x0f\xfb\x05\xda\x23\xa3\xd5\x96\x4b\xfd\x40\xf2\x50\xe1\xe5\x3f\x28\x50\x39\x84\xe1\x7b\xf2\xbc\x02\x7e\xca\xf1\xba\xda\x42\xb2\x1c\xe4\xab\x5b\x33\xab\xc6\x88\x56\xcd\xf7\xc0\x59\x0f\x34\x23\xd3\x6f\x92\xbf\xb0\x8c\xb7\x74\x5e\x2f\xdb\xff\xbf\xe3\x17\xb9\x77\x0e\xe7\xdc\xb9\xd3\x24\x9c\xc8\x58\xa1\x98\xbd\xf5\xf5\x80\x61\xdc\xa6\xf6\xb9\x12\x8e\x66\x7f\xba\x27\x35\x89\x07\xf5\x03\x0a\x1c\xcf\x0b\x08\x6c\xeb\x2f\x85\x3e\x3d\x35\x27\x95\x85\x32\xbc\x0d\xec\xe0\x42\xb3\xd1\x8e\xfc\xe2\x1c\xbd\x76\x49\xcc\xb6\x51\x92\xb8\xbb\x0d\xd9\x47\x00\x00\x00\xff\xff\x94\xfa\xa3\xd7\x4f\x04\x00\x00")

func login_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_login_tpl,
		"login.tpl",
	)
}

func login_tpl() (*asset, error) {
	bytes, err := login_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "login.tpl", size: 1103, mode: os.FileMode(438), modTime: time.Unix(1423332525, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_complete_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x51\x41\xcf\x9b\x30\x0c\xbd\xf7\x57\x58\x51\xaf\x94\xfb\x04\x5c\xa6\x1d\xa7\x55\x6a\xff\x40\x20\xa6\x44\x0d\x49\x66\x42\xb7\x2a\xca\x7f\x5f\x20\xd0\x75\xa8\x95\x26\x4d\x3b\x7c\x48\x11\xb1\x63\x3f\xfb\xbd\x57\xb4\x86\x7a\xe0\x8d\x93\x46\x97\x2c\x27\x6c\xcc\x0d\x29\x6f\x4c\x6f\x15\x3a\x64\xd0\xa3\xeb\x8c\x28\xd9\xf1\xdb\xe9\xcc\xaa\x1d\xc4\xaf\x90\xda\x8e\x0e\xdc\xdd\x62\xc9\x3a\x29\x04\x6a\x06\x9a\xf7\x31\x72\xe6\x3a\x05\x37\xae\xc6\x18\x79\x7f\x38\x4f\x89\x10\x18\xe4\xa9\xd7\xfb\xbd\xe5\xc3\xf0\xc3\x90\xf8\x42\x34\xc0\xa7\x12\x0e\xf1\xf2\x95\xdb\xc3\x9a\x0f\x21\x4d\x11\xf2\x06\x8d\x8a\xc9\x92\x4d\x4b\x66\x17\x32\xa3\xf5\x5e\xb6\xf0\x07\x44\x08\xd0\xf1\x21\x43\x22\x43\xde\xa3\x8e\xfd\xcb\x9e\x5b\x94\x79\xed\x04\xf3\x54\x31\x57\x0d\x96\xeb\x17\x65\x19\x17\xc2\x68\x56\x15\xf2\xb1\x09\x87\x96\x67\xca\x34\xd7\x98\xcd\x65\x3c\x53\xeb\x06\x2d\xe9\xf3\xbc\x7b\x63\xb4\x23\xa3\xd8\x22\x9a\xc3\x9f\x6e\x95\x6c\xe5\xc2\xc0\x2a\xde\x60\x67\x94\x40\x8a\x7a\x3f\xd2\x84\xdf\x47\x49\x28\x56\x09\xe7\x11\x79\xe4\xf5\x3b\xf4\x9e\xb8\xbe\x20\xec\xa3\x0a\x93\xa4\x1b\x81\xde\x73\xed\x50\xd9\xac\x4e\x74\xbc\xb7\x24\xb5\x9b\x41\x42\xd8\x12\x5b\xa4\xdd\x3d\x4d\x5f\x0d\x8d\xe4\x5a\x49\xfd\xf1\x8d\xaf\x9b\xe7\xbf\xb0\xf7\x05\xe0\x47\x77\x79\x43\x69\x63\xf6\xe7\xf4\x0a\xff\x66\xfa\x4b\xd9\xfe\x93\xf7\xf3\xb5\x1e\x9d\x33\x0f\xbc\xda\x69\x88\x27\x8b\x48\x3d\xa7\xfb\x7c\x4f\xf0\x8b\x1e\xc3\x58\xf7\xd2\xb1\xea\x34\xff\x8b\x3c\xb5\x57\xbb\x22\x9f\xd4\xab\x7e\x05\x00\x00\xff\xff\x36\xd8\xe2\xc3\x8b\x04\x00\x00")

func recover_complete_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_recover_complete_tpl,
		"recover-complete.tpl",
	)
}

func recover_complete_tpl() (*asset, error) {
	bytes, err := recover_complete_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover-complete.tpl", size: 1163, mode: os.FileMode(438), modTime: time.Unix(1423332541, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_html_email = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x29\x2e\x29\xca\xcf\x4b\xb7\xab\xae\xd6\xf3\xc9\xcc\xcb\xae\xad\xb5\xd1\x87\x8a\x00\x02\x00\x00\xff\xff\xe1\x46\x1b\xff\x1a\x00\x00\x00")

func recover_html_email_bytes() ([]byte, error) {
	return bindata_read(
		_recover_html_email,
		"recover-html.email",
	)
}

func recover_html_email() (*asset, error) {
	bytes, err := recover_html_email_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover-html.email", size: 26, mode: os.FileMode(438), modTime: time.Unix(1422773459, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_text_email = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\xd6\xf3\xc9\xcc\xcb\xae\xad\x05\x04\x00\x00\xff\xff\x41\xf7\xa1\x3d\x09\x00\x00\x00")

func recover_text_email_bytes() ([]byte, error) {
	return bindata_read(
		_recover_text_email,
		"recover-text.email",
	)
}

func recover_text_email() (*asset, error) {
	bytes, err := recover_text_email_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover-text.email", size: 9, mode: os.FileMode(438), modTime: time.Unix(1422773459, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x94\xc1\x6e\xa3\x30\x10\x86\xef\x79\x8a\x91\x95\x2b\xe1\xbe\x02\x2e\xd1\x1e\x57\xbb\xda\x36\x0f\x30\x80\x09\x56\x8c\x6d\x0d\x26\x6a\x64\xf9\xdd\x6b\x20\xa1\x89\x45\xa4\xf6\xd0\x43\x91\x90\x8c\x19\x7f\xe3\xff\xff\x2d\x67\x8d\xa6\x0e\xb0\xb2\x42\xab\x9c\xa5\xc4\x2b\x7d\xe6\xc4\xa0\xe3\xb6\xd5\x75\xce\xfe\xfd\x7d\x79\x65\xc5\x06\xc2\xe3\xdc\x76\xe8\x39\x29\xec\xf8\x6f\xa2\x1e\x7e\xe5\xb0\x0b\x83\x3f\x68\x76\xb7\x79\xef\xa7\xca\xac\x16\x67\xa8\x24\xf6\x7d\xce\x46\x7e\x72\x24\x3d\x18\xe7\x44\x03\x0f\x08\xef\xa1\xc5\x3e\xe1\x44\x9a\x9c\xe3\xaa\xf6\xfe\xda\x2b\xa6\x08\x65\x06\x3b\x63\xee\x2a\xa6\xaa\xde\xa0\x5a\x29\x4b\xb0\xae\xb5\x62\x45\x26\x96\x9d\x20\x34\x98\x8c\xfd\xc3\x6c\x2a\xc2\x3b\x2e\x8d\x68\x13\xe1\x61\xef\x95\x56\x96\xb4\x64\x60\x2f\x86\xe7\xcc\xf2\x37\xcb\x60\x54\x90\xb3\x9b\x16\x06\x46\x62\xc5\x5b\x2d\x6b\x4e\x39\x3b\x2c\xd3\x67\x94\x43\xa8\x73\x6e\x77\x58\x1c\x62\x90\xde\x69\x4c\x83\xc8\x8f\x4f\xe7\x08\xd5\x91\xc3\x36\x58\x32\xfa\x1b\xb9\xf5\x5c\x78\xcb\xa5\x49\x4a\xa9\xab\x13\x2b\x9c\x33\x24\x94\x9d\x20\xde\xc7\x2a\xaf\x3e\x6f\xee\xba\xdf\xd2\x0d\x4a\x1b\x41\xdd\xe1\x49\xc8\xd1\xef\x4f\x64\xbd\x02\xfc\xe9\x91\x47\x92\xa2\xe4\xf7\xf3\x5f\x58\x3b\x01\xfb\xd8\xbe\x2f\x1d\x84\x55\x2b\xbf\xe9\x3c\x8c\xc3\xf9\xbb\x1c\xac\xd5\x0b\xb4\xb4\x0a\xc2\x9b\x04\x5c\x87\x74\x99\xc6\x73\x8f\xab\x51\xfd\x50\x76\xc2\xb2\xe2\xff\x7c\x8b\x64\xe9\xbc\x7e\x26\x66\x18\x73\xa4\x50\xa7\xa7\x10\x68\x89\x37\xe1\x46\x92\xfa\x28\x42\xa6\x7b\x54\x15\x97\x59\x8a\xc5\x26\x4b\xc7\x90\x8a\xf7\x00\x00\x00\xff\xff\xc2\xb1\x79\x5b\xba\x04\x00\x00")

func recover_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_recover_tpl,
		"recover.tpl",
	)
}

func recover_tpl() (*asset, error) {
	bytes, err := recover_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover.tpl", size: 1210, mode: os.FileMode(438), modTime: time.Unix(1423332554, 0)}
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
	"layout.tpl": layout_tpl,
	"layoutEmail.tpl": layoutemail_tpl,
	"login.tpl": login_tpl,
	"recover-complete.tpl": recover_complete_tpl,
	"recover-html.email": recover_html_email,
	"recover-text.email": recover_text_email,
	"recover.tpl": recover_tpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"layout.tpl": &_bintree_t{layout_tpl, map[string]*_bintree_t{
	}},
	"layoutEmail.tpl": &_bintree_t{layoutemail_tpl, map[string]*_bintree_t{
	}},
	"login.tpl": &_bintree_t{login_tpl, map[string]*_bintree_t{
	}},
	"recover-complete.tpl": &_bintree_t{recover_complete_tpl, map[string]*_bintree_t{
	}},
	"recover-html.email": &_bintree_t{recover_html_email, map[string]*_bintree_t{
	}},
	"recover-text.email": &_bintree_t{recover_text_email, map[string]*_bintree_t{
	}},
	"recover.tpl": &_bintree_t{recover_tpl, map[string]*_bintree_t{
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

