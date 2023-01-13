package internal

import (
	"fmt"

	"github.com/Giovanny472/gwav/model"
)

type parsewav struct {

	// RIFF
	сhunkRIFF     [4]uint8
	сhunkSizeRIFF uint32
	сhunkWave     [4]uint8

	// для создания объекта wav
	buildwav model.BuilderWav
}

var parsew *parsewav

func NewParser(bw model.BuilderWav) model.Parser {
	if parsew == nil {
		parsew = &parsewav{buildwav: bw}
	}
	return parsew
}

func (pw *parsewav) Parse(dw *[]byte) (model.Wave, error) {

	p := (*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff]
	//pw.сhunkRIFF[0] = (*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff]

	//pw.сhunkRIFF[0]
	fmt.Printf("val -->  %T", p)
	//fmt.Println(pw.chunkRIFF)

	return nil, nil
}
