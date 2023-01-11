package internal

import "github.com/Giovanny472/gwav/model"

type parsewav struct {
}

var parsew *parsewav

func NewParser() model.Parser {

	if parsew == nil {
		parsew = &parsewav{}
	}
	return parsew
}

func (pw *parsewav) Parse(dw *[]byte) model.Wave {

	return nil
}
