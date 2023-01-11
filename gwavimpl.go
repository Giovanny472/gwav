package gwav

import (
	"github.com/Giovanny472/gwav/internal"
	"github.com/Giovanny472/gwav/model"
)

type gwavаfacade struct {

	// читатель файла-wav
	reader model.Reader

	// parser
	parser model.Parser
}

var gwf *gwavаfacade

func NewReaderWav(pathFile string) (GWav, error) {

	rd, err := internal.NewReader(pathFile)
	if err != nil {
		return nil, err
	}

	// создание
	if gwf != nil {
		gwf = &gwavаfacade{reader: rd,
			parser: internal.NewParser()}
	}

	return gwf, nil
}

func (gw *gwavаfacade) Read() (*[]byte, error) {
	filebytes, err := gw.reader.Read()
	if err != nil {
		return nil, err
	}

	return filebytes, nil
}
