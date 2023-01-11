package internal

import "github.com/Giovanny472/gwav/model"

type readwav struct {
}

var readw *readwav

func NewReader(path string) (model.Reader, error) {

	if readw == nil {
		readw = &readwav{}
	}
	return readw, nil
}

func (r *readwav) Read() (*[]byte, error) {

	return nil, nil
}
