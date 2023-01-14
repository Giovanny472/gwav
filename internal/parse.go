package internal

import (
	"errors"
	"fmt"

	"github.com/Giovanny472/gwav/model"
)

type parsewav struct {

	// RIFF
	сhunkRIFF     string
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

	pw.сhunkRIFF = string((*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff])
	if pw.сhunkRIFF != string(model.ConstRIFF) {
		return nil, errors.New("слово RIFF не найдено. значение: " + pw.сhunkRIFF)
	}

	b := (*dw)[model.IdxStartChunkSzRiff:model.IdxEndChunkSzRiff]
	c := (*dw)[model.IdxStartWordWave:model.IdxEndWordWave]

	//pw.сhunkRIFF[0] = (*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff]

	//pw.сhunkRIFF[0]
	fmt.Println("strRiff -->", pw.сhunkRIFF)
	fmt.Printf("val b-->  %T , val: %v  \n", b, b)
	fmt.Printf("val c-->  %T , val: %v  \n", c, c)

	//fmt.Println(pw.chunkRIFF)

	return nil, nil
}
