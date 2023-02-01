package gwav

import (
	"github.com/Giovanny472/gwav/internal/builder"
	"github.com/Giovanny472/gwav/internal/parser"
	"github.com/Giovanny472/gwav/internal/reader"
	"github.com/Giovanny472/gwav/model"
)

type gwavаfacade struct {

	// читатель файла-wav
	reader model.Reader

	// parser
	parse model.Parser

	// файл wav
	wav model.Wave
}

var gwf *gwavаfacade

func NewReaderWav(pathFile string) (GWav, error) {

	rd, err := reader.NewReader(pathFile)
	if err != nil {
		return nil, err
	}

	// создание WAVE с помощью parser и builder
	if gwf == nil {
		gwf = &gwavаfacade{reader: rd,

			// создание объекта parser и передается builder
			// для создания экземпляра wave
			parse: parser.NewParser(builder.NewBuildWav())}
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
	gw.wav, err = gw.parse.Parse(data)
	return gw.wav, err
}
