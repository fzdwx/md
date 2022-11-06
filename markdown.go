package main

import (
	"errors"
	"github.com/fzdwx/md/utils"
	"io"
	"os"
	"path"
	"path/filepath"
)

func filePathToMd(file string) (*markdown, error) {
	if file == "" {
		return defaultMd(), nil
	}

	var err error
	m := &markdown{fileName: defaultName("")}
	m.filePath, err = filepath.Abs(file)
	if err == nil {
		m.fileName = defaultName(path.Base(file))
		err = m.loadBody()
		m.hasFile = err == nil
	}

	return m, err
}

func defaultMd() *markdown {
	return &markdown{
		body:     "",
		filePath: "",
		fileName: defaultName(""),
	}
}

func defaultName(filename string) string {
	if filename == "" {
		return "UNKNOWN"
	}
	return filename
}

type markdown struct {
	body     string
	filePath string
	fileName string
	hasFile  bool
}

func (m *markdown) mustLoadBody() {
	_ = m.loadBody()
}

// loadBody from local fs load body
func (m *markdown) loadBody() error {
	if m.filePath == "" {
		return errors.New("filePath not found")
	}

	r, err := os.Open(m.fileName)
	defer r.Close()
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	m.body = utils.CleanCr(string(bytes))
	return nil
}
