package internal

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/Giovanny472/gwav/model"
)

type parsewav struct {

	// RIFF
	сhunkRIFF     string
	сhunkSizeRIFF uint32
	сhunkWave     string

	//  fmt
	сhunkFmt     string
	сhunkSizeFmt uint32
	audioformat  uint16

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

	// cлово RIFF
	pw.сhunkRIFF = string((*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff])
	if pw.сhunkRIFF != string(model.ConstRIFF) {
		return nil, errors.New("слово RIFF не найдено. Значение: " + pw.сhunkRIFF)
	}

	// size RIFF
	szRiff := (*dw)[model.IdxStartChunkSzRiff:model.IdxEndChunkSzRiff]
	pw.сhunkSizeRIFF = binary.LittleEndian.Uint32(szRiff)

	// слово WAVE
	pw.сhunkWave = string((*dw)[model.IdxStartWordWave:model.IdxEndWordWave])
	if pw.сhunkWave != string(model.ConstWAVE) {
		return nil, errors.New("слово WAVE не найдено. Значение: " + pw.сhunkWave)
	}

	// cлово fmt
	pw.сhunkFmt = string((*dw)[model.IdxStartWordFmt:model.IdxEndWordFmt])
	if pw.сhunkRIFF != string(model.ConstRIFF) {
		return nil, errors.New("слово FMT не найдено. Значение:" + pw.сhunkFmt)
	}

	// size fmt
	szFmt := (*dw)[model.IdxStartChunkSzFmt:model.IdxEndChunkSzFmt]
	pw.сhunkSizeFmt = binary.LittleEndian.Uint32(szFmt)

	// audioformat
	audioFormat := (*dw)[model.IdxStartAudioformat:model.IdxEndAudioformat]
	pw.audioformat = binary.LittleEndian.Uint16(audioFormat)

	fmt.Println("strRiff -->", pw.сhunkRIFF)
	fmt.Println("strsize -->", pw.сhunkSizeRIFF)
	fmt.Println("strWAVE -->", pw.сhunkWave)
	fmt.Println("strFmt -->", pw.сhunkFmt)
	fmt.Println("сhunkSizeFmt -->", pw.сhunkSizeFmt)
	fmt.Println("audioformat -->", pw.audioformat)

	//fmt.Println(pw.chunkRIFF)

	return nil, nil
}
