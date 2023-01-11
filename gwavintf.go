package gwav

import "github.com/Giovanny472/gwav/model"

type GWav interface {

	// читать файл-wav
	Read() (*[]byte, error)

	// Парсе файла-wav
	Parse(data *[]byte) (model.Wave, error)
}
