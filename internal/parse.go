package internal

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Giovanny472/gwav/model"
)

type parsewav struct {

	// RIFF
	сhunkRIFF     string
	сhunkSizeRIFF int
	сhunkWave     string
	сhunkFmt      string

	// для создания объекта wav
	buildwav model.BuilderWav

	// принимаеи ошибки
	err error
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
		return nil, errors.New("слово RIFF не найдено. Значение: " + pw.сhunkRIFF)
	}

	strSzRiff := string((*dw)[model.IdxStartChunkSzRiff:model.IdxEndChunkSzRiff])
	pw.сhunkSizeRIFF, pw.err = strconv.Atoi(strSzRiff)
	if pw.err != nil {
		return nil, errors.New("ошибка при получании sizeRiff. Значение: " + pw.сhunkWave)
	}

	pw.сhunkWave = string((*dw)[model.IdxStartWordWave:model.IdxEndWordWave])
	if pw.сhunkWave != string(model.ConstWAVE) {
		return nil, errors.New("слово WAVE не найдено. Значение: " + pw.сhunkWave)
	}

	pw.сhunkFmt = string((*dw)[model.IdxStartWordFmt:model.IdxEndWordFmt])
	if pw.сhunkRIFF != string(model.ConstRIFF) {
		return nil, errors.New("слово FMT не найдено. Значение:" + pw.сhunkFmt)
	}

	//pw.сhunkRIFF[0]
	fmt.Println("strRiff -->", pw.сhunkRIFF)
	fmt.Println("strsize -->", pw.сhunkSizeRIFF)
	fmt.Println("strWAVE -->", pw.сhunkWave)
	fmt.Println("strFmt -->", pw.сhunkFmt)

	//fmt.Println(pw.chunkRIFF)

	return nil, nil
}
