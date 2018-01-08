package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import (
	"fmt"
	"path/filepath"
	"strings"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {
		if strings.EqualFold(f.Name, className) {
			fmt.Printf("is same : ")
			rc, err := f.Open()
			defer rc.Close()
			if err != nil {
				return nil, nil, err
			}
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			fmt.Println(data)
			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found:" + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
