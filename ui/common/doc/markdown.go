package doc

import (
	"errors"
	"github.com/fzdwx/md/utils"
	"github.com/gookit/goutil/fsutil"
	"io"
	"os"
	"path"
	"path/filepath"
)

const (
	UNKNOWN = "UNKNOWN"
)

func FilePathToMd(file string) (*Markdown, error) {
	if file == "" {
		return DefaultMd(), nil
	}

	var err error
	m := &Markdown{FileName: DefaultName("")}
	m.filePath, err = filepath.Abs(file)
	if err == nil {
		m.FileName = DefaultName(path.Base(file))
		err = m.LoadBody()
		m.hasFile = err == nil
	}

	return m, err
}

func DefaultMd() *Markdown {
	return &Markdown{
		Body:     "",
		filePath: "",
		FileName: DefaultName(""),
	}
}

func DefaultName(filename string) string {
	if filename == "" {
		return UNKNOWN
	}
	return filename
}

type Markdown struct {
	Body     string
	filePath string
	FileName string
	hasFile  bool
}

func (m *Markdown) MustLoadBody() {
	_ = m.LoadBody()
}

// LoadBody from local fs load Body
func (m *Markdown) LoadBody() error {
	if m.filePath == "" {
		return errors.New("filePath not found")
	}

	r, err := os.Open(m.FileName)
	defer r.Close()
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	m.Body = utils.CleanCr(string(bytes))
	return nil
}

func (m *Markdown) NoName() bool {
	return m.FileName == "" || m.FileName == UNKNOWN
}

func (m *Markdown) Save() error {
	if m.filePath == "" {
		m.filePath = m.FileName
	}

	return fsutil.WriteFile(m.filePath, m.Body, 0666, fsutil.FsCWTFlags)
}
