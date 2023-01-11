package internal

import (
	"io"
	"os"

	"github.com/Giovanny472/gwav/model"
)

type readwav struct {

	// путь файла-wav
	pathfile string
}

var readw *readwav

func NewReader(path string) (model.Reader, error) {

	if readw == nil {
		readw = &readwav{pathfile: path}
	}
	return readw, nil
}

func (rw *readwav) Read() (*[]byte, error) {

	file, err := os.Open(rw.pathfile)
	if err != nil {
		return nil, err
	}

	filebytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &filebytes, nil
}
