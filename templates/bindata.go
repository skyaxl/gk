// Code generated by go-bindata.
// sources:
// tmpl/file.tmpl
// tmpl/partials/constants.tmpl
// tmpl/partials/func.tmpl
// tmpl/partials/func_parameters.tmpl
// tmpl/partials/func_results.tmpl
// tmpl/partials/func_return.tmpl
// tmpl/partials/imports.tmpl
// tmpl/partials/interface.tmpl
// tmpl/partials/interface_func.tmpl
// tmpl/partials/interface_stub.tmpl
// tmpl/partials/struct.tmpl
// tmpl/partials/struct_function.tmpl
// tmpl/partials/vars.tmpl
// tmpl/service.tmpl
// DO NOT EDIT!

package template

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _tmplFileTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x43\xe8\x41\x41\x8a\x67\xc1\x93\xa7\x3d\x28\x82\xe0\x55\x86\xec\xb4\x16\xdb\x69\x4d\x66\x73\x09\xf9\xef\x12\x9b\x2d\x69\x68\xd9\x5b\x99\xf7\xde\xf7\x5e\x33\xa3\xfe\xc1\x8e\xc0\xfb\xe6\x7d\xf9\x0c\xc1\xfb\xbe\x85\x4e\xe0\x6e\x20\x86\xe6\x34\xce\x93\x11\x7b\x0f\x8f\x21\x54\xde\x0b\x8d\xf3\x80\x42\xa0\xfa\x45\x50\xab\x25\x26\x89\xcf\x25\xe0\x65\x62\x2b\xc8\x7b\x08\x7d\x95\x54\x66\x3b\xc0\x7c\xa2\xd9\x21\x38\x34\x31\x1c\xc5\x83\xdc\x89\x85\x4c\x8b\x9a\xd6\xb4\x41\xee\x08\xea\xfe\xa1\x76\xf0\xf4\x9c\x3b\x62\x36\xfb\xc1\xeb\x5d\x41\xed\x32\xfa\x4e\xc9\x87\x98\x8b\x96\xc3\x86\x24\x6f\xf1\xf6\xff\x78\x9b\xfd\x4a\xf2\x3d\x9d\x0b\x76\x06\x4f\x7a\x8a\xd1\x2f\xd4\x2e\x15\x36\x6f\x38\x12\x28\x55\xbc\x59\x7b\x61\xbd\xd6\x0e\x96\x0a\x79\xd9\xf5\x15\x5d\xd2\x4f\xbc\x19\x58\x6d\x87\x56\xd5\x5f\x00\x00\x00\xff\xff\x09\xcf\x6b\x86\x40\x02\x00\x00"

func tmplFileTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplFileTmpl,
		"tmpl/file.tmpl",
	)
}

func tmplFileTmpl() (*asset, error) {
	bytes, err := tmplFileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/file.tmpl", size: 576, mode: os.FileMode(420), modTime: time.Unix(1491089311, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsConstantsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x48\x2f\x51\xd0\xc8\x49\xcd\x53\x50\xd1\x54\x30\xa8\xad\x4d\xce\xcf\x2b\x2e\x51\xd0\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xe5\xaa\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x20\xec\xcc\x34\x05\x95\x32\x3d\x8f\xc4\xe2\xb0\xc4\x9c\xd2\xd4\xda\x5a\x5b\x88\x3c\x94\x57\x5d\x9d\x9a\x97\x02\xa7\xb8\x34\xa1\x0c\x40\x00\x00\x00\xff\xff\x13\xc6\xee\xf4\x7e\x00\x00\x00"

func tmplPartialsConstantsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsConstantsTmpl,
		"tmpl/partials/constants.tmpl",
	)
}

func tmplPartialsConstantsTmpl() (*asset, error) {
	bytes, err := tmplPartialsConstantsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/constants.tmpl", size: 126, mode: os.FileMode(420), modTime: time.Unix(1490947670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFuncTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xb1\x6a\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\xea\xae\xf1\xe5\x1b\xd2\xa5\xcd\x0f\x1c\x8e\x3d\x76\x04\xf2\xca\x91\x56\x85\x59\xf6\xdf\x83\x1c\x08\x5c\xf7\x76\x67\x18\xde\xda\x64\x86\x99\x72\x3f\xd2\xa4\x44\x88\xa2\x2c\xeb\x34\xf3\xd9\xa3\x80\xd1\x1d\x66\x71\x85\x64\xc5\x8d\x3f\x18\xdf\xf3\x72\x22\x84\xbb\xbb\xd9\x75\x74\x60\xaa\x74\xb7\xe1\xf1\xc0\xc7\x7e\x24\xee\x14\xc5\x99\x5b\xc1\x57\xab\x51\x58\x2b\x52\xde\xe2\x8c\x6f\x16\x0e\xd7\xe0\xa6\xb8\x25\x0a\xc6\x4f\xd6\x96\xb4\xde\xf1\xe6\x5e\xa8\xad\xc8\x8b\x51\xf7\x78\xfe\xfd\xc3\x7f\xd9\x7d\x30\xa3\x2c\xee\xe8\x7e\x17\xfd\x06\x00\x00\xff\xff\xef\x37\xdc\xd3\xcc\x00\x00\x00"

func tmplPartialsFuncTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFuncTmpl,
		"tmpl/partials/func.tmpl",
	)
}

func tmplPartialsFuncTmpl() (*asset, error) {
	bytes, err := tmplPartialsFuncTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func.tmpl", size: 204, mode: os.FileMode(420), modTime: time.Unix(1491089883, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFunc_parametersTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xad\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x52\x41\xc2\x99\x69\x0a\x79\xf9\x25\x0a\x1a\x39\x89\xc5\x25\x0a\x2a\x99\x0a\x2a\x9a\xb5\xb5\x3a\xd5\xd5\xa9\x79\x29\x20\x59\x30\x05\x08\x00\x00\xff\xff\x15\x5b\x78\xe8\x50\x00\x00\x00"

func tmplPartialsFunc_parametersTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFunc_parametersTmpl,
		"tmpl/partials/func_parameters.tmpl",
	)
}

func tmplPartialsFunc_parametersTmpl() (*asset, error) {
	bytes, err := tmplPartialsFunc_parametersTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func_parameters.tmpl", size: 80, mode: os.FileMode(420), modTime: time.Unix(1490904745, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFunc_resultsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x48\x2f\x51\xd0\xc8\x49\xcd\x53\x50\xd1\x54\x30\xa8\xad\xd5\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xad\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x52\x41\xc2\x99\x69\x0a\x79\xf9\x20\xfd\x89\xc5\x25\x0a\x2a\x99\x0a\x2a\x9a\xb5\xb5\x3a\xd5\xd5\xa9\x79\x29\x20\x59\x30\xa5\x09\xa5\x01\x01\x00\x00\xff\xff\xcc\x6e\x26\x22\x6c\x00\x00\x00"

func tmplPartialsFunc_resultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFunc_resultsTmpl,
		"tmpl/partials/func_results.tmpl",
	)
}

func tmplPartialsFunc_resultsTmpl() (*asset, error) {
	bytes, err := tmplPartialsFunc_resultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func_results.tmpl", size: 108, mode: os.FileMode(420), modTime: time.Unix(1490905608, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsFunc_returnTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xad\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\x05\x31\x33\xd3\x14\xf2\xf2\x4b\x14\x34\x72\x12\x8b\x4b\x14\x54\x32\x15\x54\x34\x6b\x6b\x75\xaa\xab\x53\xf3\x52\x40\xb2\x60\x0a\x10\x00\x00\xff\xff\xbc\x66\xba\xd5\x44\x00\x00\x00"

func tmplPartialsFunc_returnTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsFunc_returnTmpl,
		"tmpl/partials/func_return.tmpl",
	)
}

func tmplPartialsFunc_returnTmpl() (*asset, error) {
	bytes, err := tmplPartialsFunc_returnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/func_return.tmpl", size: 68, mode: os.FileMode(420), modTime: time.Unix(1490905876, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsImportsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x50\xa9\xad\xcd\xcc\x2d\xc8\x2f\x2a\xd1\xe0\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x05\xc9\x81\xd5\xe4\xe5\x97\x28\x68\xa4\x16\x2a\xa8\x94\xe9\xf9\x25\xe6\xa6\x2a\x28\x29\x69\x82\x64\xa0\xdc\xda\x5a\x85\xea\xea\xd4\xbc\x14\xa8\x50\x48\x65\x41\x6a\x6d\x2d\x17\x54\x88\x4b\x13\xca\x00\x04\x00\x00\xff\xff\x34\xce\x15\xc7\x6d\x00\x00\x00"

func tmplPartialsImportsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsImportsTmpl,
		"tmpl/partials/imports.tmpl",
	)
}

func tmplPartialsImportsTmpl() (*asset, error) {
	bytes, err := tmplPartialsImportsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/imports.tmpl", size: 109, mode: os.FileMode(420), modTime: time.Unix(1491090111, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsInterfaceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcb\xc1\x09\xc2\x40\x14\x04\xd0\xfb\x56\x31\x84\x3d\x4a\x0a\x10\x2c\x41\x5b\x90\x25\x99\x68\xc0\x7c\x43\xfc\x06\x64\x98\xde\xbd\xa5\x80\x97\xbf\x95\x90\xfa\x5b\x5b\x68\x63\x8e\xe4\x36\xb5\x81\x50\x91\xb6\x16\x0f\xa2\xce\xa7\xba\xe3\x7c\x41\x7f\x65\x3e\xdf\xe3\xc7\x96\x92\xcb\xfa\x6a\x49\x74\x07\xb9\x4f\xdf\x18\x3a\xd4\xdd\x2e\x12\x63\xb4\x8b\xff\x01\x00\x00\xff\xff\xe7\xc0\x6a\x14\x60\x00\x00\x00"

func tmplPartialsInterfaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsInterfaceTmpl,
		"tmpl/partials/interface.tmpl",
	)
}

func tmplPartialsInterfaceTmpl() (*asset, error) {
	bytes, err := tmplPartialsInterfaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/interface.tmpl", size: 96, mode: os.FileMode(420), modTime: time.Unix(1490947770, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsInterface_funcTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xd5\xa8\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x2b\xcd\x4b\x8e\x2f\x48\x2c\x4a\xcc\x4d\x2d\x49\x2d\x2a\x56\x52\xd0\x0b\x80\x73\x6a\x6b\x35\x31\x94\x16\xa5\x16\x97\xe6\x94\x80\xd4\x05\x41\x58\xb5\xb5\x80\x00\x00\x00\xff\xff\xc1\x2e\xd4\xb5\x59\x00\x00\x00"

func tmplPartialsInterface_funcTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsInterface_funcTmpl,
		"tmpl/partials/interface_func.tmpl",
	)
}

func tmplPartialsInterface_funcTmpl() (*asset, error) {
	bytes, err := tmplPartialsInterface_funcTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/interface_func.tmpl", size: 89, mode: os.FileMode(420), modTime: time.Unix(1490906014, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsInterface_stubTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x2a\x2e\x29\x2a\x4d\x2e\x51\x52\xd0\xab\xad\xe5\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\xd0\xf3\xd4\xf3\x4d\x2d\xc9\xc8\x4f\x29\xae\xad\x45\xd5\x54\x9a\x14\x9f\x56\x9a\x97\x5c\x92\x99\x9f\xa7\xa4\xa0\x52\x56\x5b\xcb\xc5\x55\x5d\x9d\x9a\x97\x52\x5b\x0b\x08\x00\x00\xff\xff\x39\x10\x5c\x96\x5d\x00\x00\x00"

func tmplPartialsInterface_stubTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsInterface_stubTmpl,
		"tmpl/partials/interface_stub.tmpl",
	)
}

func tmplPartialsInterface_stubTmpl() (*asset, error) {
	bytes, err := tmplPartialsInterface_stubTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/interface_stub.tmpl", size: 93, mode: os.FileMode(420), modTime: time.Unix(1490908658, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsStructTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\x55\x28\x2e\x29\x2a\x4d\x2e\x51\xa8\x56\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\xd0\x0b\x4b\x2c\x2a\xae\xad\xe5\xaa\xae\x56\x29\x83\xa9\x06\xb3\x43\x2a\x0b\x52\x6b\x6b\xab\xab\x53\xf3\x52\x6a\x6b\x15\x6a\x01\x01\x00\x00\xff\xff\x65\x47\x74\x51\x51\x00\x00\x00"

func tmplPartialsStructTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsStructTmpl,
		"tmpl/partials/struct.tmpl",
	)
}

func tmplPartialsStructTmpl() (*asset, error) {
	bytes, err := tmplPartialsStructTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/struct.tmpl", size: 81, mode: os.FileMode(420), modTime: time.Unix(1491062995, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsStruct_functionTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x31\x4e\xc4\x40\x0c\x45\xfb\x9c\xe2\x6b\xaa\xa4\x99\xe5\x0c\x74\x34\x14\x40\xbf\x0a\xd9\x9f\x25\xd2\xc4\x09\x33\x9e\x22\xb2\x7c\x77\x34\x41\x02\x6d\x65\xfb\xc9\xf6\x7f\x73\x95\x09\xbd\x59\x7c\xd7\x5c\x27\x8d\xaf\xe3\x4a\x77\xfc\x83\x8f\x63\xa7\xfb\x00\x33\xe5\xba\xa7\x51\x89\xb0\x88\x32\xcf\xe3\xc4\x6b\x3b\x0f\x88\xee\x66\xcb\x0c\xd9\x14\x3d\xbf\x11\x9f\xb7\xdb\x81\x10\x86\xc6\xcf\xa1\x35\x4c\x85\xee\xd6\x5d\x2e\x78\x59\xf7\xc4\x95\xa2\x38\xb6\x9a\xf1\x59\xcb\x22\x2c\x05\x69\xbb\x2f\x13\xbe\x98\xd9\x9d\x0f\xef\x8a\x3e\x51\x10\xdf\x58\x6a\xd2\x32\xe0\xc9\x3d\x53\x6b\x96\x07\xa1\xa6\x71\xfd\xe5\xe1\x6f\xd9\xbd\x33\xa3\xdc\xfc\x0c\x6f\xf5\x27\x00\x00\xff\xff\x1d\x73\x76\x5d\xed\x00\x00\x00"

func tmplPartialsStruct_functionTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsStruct_functionTmpl,
		"tmpl/partials/struct_function.tmpl",
	)
}

func tmplPartialsStruct_functionTmpl() (*asset, error) {
	bytes, err := tmplPartialsStruct_functionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/struct_function.tmpl", size: 237, mode: os.FileMode(420), modTime: time.Unix(1491089803, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplPartialsVarsTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xce\x4c\x53\x48\x2f\x51\xd0\xc8\x49\xcd\x53\x50\xd1\x54\x30\xa8\xad\x2d\x4b\x2c\x52\xd0\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x29\x53\xb0\xb2\x55\x50\xa9\xad\xe5\xaa\xae\x56\x29\xd3\xf3\x4b\xcc\x4d\xad\xad\x55\x00\xb3\x43\x2a\x0b\x20\xec\xcc\x34\x05\x95\x32\x3d\x8f\xc4\xe2\xb0\xc4\x9c\xd2\xd4\xda\x5a\x5b\x88\x3c\x94\x57\x5d\x9d\x9a\x97\x02\xa7\xb8\x34\xa1\x0c\x40\x00\x00\x00\xff\xff\xd8\x54\xdd\xc2\x7c\x00\x00\x00"

func tmplPartialsVarsTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplPartialsVarsTmpl,
		"tmpl/partials/vars.tmpl",
	)
}

func tmplPartialsVarsTmpl() (*asset, error) {
	bytes, err := tmplPartialsVarsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/partials/vars.tmpl", size: 124, mode: os.FileMode(420), modTime: time.Unix(1490946107, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplServiceTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\x2d\x2a\xcb\x4c\x4e\xe5\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\xcc\x2d\xc8\x2f\x2a\x29\x56\x52\xd0\xf3\x84\xb0\x6a\x6b\x51\xe5\xf3\x4a\x52\x8b\xd2\x12\x93\x53\x95\x14\xf4\x82\xf5\x3c\x6b\x6b\xb9\xb0\x4b\xc7\x17\x97\x94\x26\x81\xd4\xd4\xd6\x02\x02\x00\x00\xff\xff\x1d\x16\x0b\x7c\x6f\x00\x00\x00"

func tmplServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplServiceTmpl,
		"tmpl/service.tmpl",
	)
}

func tmplServiceTmpl() (*asset, error) {
	bytes, err := tmplServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/service.tmpl", size: 111, mode: os.FileMode(420), modTime: time.Unix(1490908189, 0)}
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
	"tmpl/file.tmpl": tmplFileTmpl,
	"tmpl/partials/constants.tmpl": tmplPartialsConstantsTmpl,
	"tmpl/partials/func.tmpl": tmplPartialsFuncTmpl,
	"tmpl/partials/func_parameters.tmpl": tmplPartialsFunc_parametersTmpl,
	"tmpl/partials/func_results.tmpl": tmplPartialsFunc_resultsTmpl,
	"tmpl/partials/func_return.tmpl": tmplPartialsFunc_returnTmpl,
	"tmpl/partials/imports.tmpl": tmplPartialsImportsTmpl,
	"tmpl/partials/interface.tmpl": tmplPartialsInterfaceTmpl,
	"tmpl/partials/interface_func.tmpl": tmplPartialsInterface_funcTmpl,
	"tmpl/partials/interface_stub.tmpl": tmplPartialsInterface_stubTmpl,
	"tmpl/partials/struct.tmpl": tmplPartialsStructTmpl,
	"tmpl/partials/struct_function.tmpl": tmplPartialsStruct_functionTmpl,
	"tmpl/partials/vars.tmpl": tmplPartialsVarsTmpl,
	"tmpl/service.tmpl": tmplServiceTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"file.tmpl": &bintree{tmplFileTmpl, map[string]*bintree{}},
		"partials": &bintree{nil, map[string]*bintree{
			"constants.tmpl": &bintree{tmplPartialsConstantsTmpl, map[string]*bintree{}},
			"func.tmpl": &bintree{tmplPartialsFuncTmpl, map[string]*bintree{}},
			"func_parameters.tmpl": &bintree{tmplPartialsFunc_parametersTmpl, map[string]*bintree{}},
			"func_results.tmpl": &bintree{tmplPartialsFunc_resultsTmpl, map[string]*bintree{}},
			"func_return.tmpl": &bintree{tmplPartialsFunc_returnTmpl, map[string]*bintree{}},
			"imports.tmpl": &bintree{tmplPartialsImportsTmpl, map[string]*bintree{}},
			"interface.tmpl": &bintree{tmplPartialsInterfaceTmpl, map[string]*bintree{}},
			"interface_func.tmpl": &bintree{tmplPartialsInterface_funcTmpl, map[string]*bintree{}},
			"interface_stub.tmpl": &bintree{tmplPartialsInterface_stubTmpl, map[string]*bintree{}},
			"struct.tmpl": &bintree{tmplPartialsStructTmpl, map[string]*bintree{}},
			"struct_function.tmpl": &bintree{tmplPartialsStruct_functionTmpl, map[string]*bintree{}},
			"vars.tmpl": &bintree{tmplPartialsVarsTmpl, map[string]*bintree{}},
		}},
		"service.tmpl": &bintree{tmplServiceTmpl, map[string]*bintree{}},
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

