package reader

import (
	"io"
	"os"

	"github.com/Giovanny472/gwav/model"
)

type readwav struct {

	// путь файла-wav
	pathfile string

	// содержание файла
	filebytes []byte
}

var readw *readwav

func NewReader(path string) (model.Reader, error) {

	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

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
	defer file.Close()

	rw.filebytes, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &rw.filebytes, nil
}
