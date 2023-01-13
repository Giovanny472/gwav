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

	// файл wav
	wav model.Wave
}

var gwf *gwavаfacade

func NewReaderWav(pathFile string) (GWav, error) {

	rd, err := internal.NewReader(pathFile)
	if err != nil {
		return nil, err
	}

	// создание WAVE с помощью parser и builder
	if gwf == nil {
		gwf = &gwavаfacade{reader: rd,

			// создание объекта parser и передается builder
			// для создания экземпляра wave
			parser: internal.NewParser(internal.NewBuildWav())}
	}

	return gwf, nil
}

func (gw *gwavаfacade) Read() (*[]byte, error) {

	// читаем файл-wav
	filebytes, err := gw.reader.Read()
	if err != nil {
		return nil, err
	}

	return filebytes, nil
}

func (gw *gwavаfacade) Parse(data *[]byte) (model.Wave, error) {
	var err error
	gw.wav, err = gw.parser.Parse(data)
	return gw.wav, err
}
