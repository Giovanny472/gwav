package model

type Parser interface {
	// данные файла-wav
	Parse(dw *[]byte) (Wave, error)
}
