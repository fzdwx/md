package main

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
)

func filePathToMd(file string) (*markdown, error) {
	var err error
	m := &markdown{}
	m.filePath, err = filepath.Abs(file)
	if err == nil {
		m.fileName = path.Base(file)
		err = m.load()
	}

	return m, err
}

type markdown struct {
	body     string
	filePath string
	fileName string
}

func (m *markdown) mustLoad() {
	_ = m.load()
}

func (m *markdown) load() error {
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

	m.body = string(bytes)
	return nil
}
