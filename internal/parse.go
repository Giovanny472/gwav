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

	a := (*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff]
	b := (*dw)[model.IdxStartChunkSzRiff:model.IdxEndChunkSzRiff]
	c := (*dw)[model.IdxStartWordWave:model.IdxEndWordWave]

	//pw.сhunkRIFF[0] = (*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff]

	//pw.сhunkRIFF[0]
	fmt.Printf("val a-->  %T", a)
	fmt.Printf("val b-->  %T", b)
	fmt.Printf("val c-->  %T", c)
	//fmt.Println(pw.chunkRIFF)

	return nil, nil
}
