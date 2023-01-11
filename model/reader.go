package model

// интерфейс для чтения файл-wav
type Reader interface {

	// для чтения файла-wav
	Read() (*[]byte, error)
}
