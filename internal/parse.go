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
	numChannels  uint16
	sampleRate   uint32
	byteRate     uint32
	blockAling   uint16
	bitperSample uint16

	// data
	сhunkData     string
	сhunkSizeData uint32
	dataaudio     []byte

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

	//*******************************
	//    З А Г О Л О В О К
	//*******************************

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

	// numChannels
	numch := (*dw)[model.IdxStartNumChannels:model.IdxEndNumChannels]
	pw.numChannels = binary.LittleEndian.Uint16(numch)

	// samplerate
	samplerate := (*dw)[model.IdxStartSampleRate:model.IdxEndSampleRate]
	pw.sampleRate = binary.LittleEndian.Uint32(samplerate)

	// byterate
	byterate := (*dw)[model.IdxStartByteRate:model.IdxEndByteRate]
	pw.byteRate = binary.LittleEndian.Uint32(byterate)

	// blockalign
	blockalign := (*dw)[model.IdxStartBlockAlign:model.IdxEndBlockAlign]
	pw.blockAling = binary.LittleEndian.Uint16(blockalign)

	// bitspersample
	bitssample := (*dw)[model.IdxStartBitsPerSample:model.IdxEndBitsPerSample]
	pw.bitperSample = binary.LittleEndian.Uint16(bitssample)

	//*******************************
	//    Д А Н Н Ы Е
	//*******************************

	// chunk data
	pw.сhunkData = string((*dw)[model.IdxStartWordData:model.IdxEndWordData])
	if pw.сhunkData != string(model.ConstData) {
		return nil, errors.New("слово data не найдено. Значение: " + pw.сhunkData)
	}

	// size RIFF
	szData := (*dw)[model.IdxStartChunkSzData:model.IdxEndChunkSzData]
	pw.сhunkSizeData = binary.LittleEndian.Uint32(szData)

	// data audio
	dataAudio := (*dw)[model.IdxStartChunkData:]
	pw.dataaudio = dataAudio
	//pw.dataaudio = binary.LittleEndian.Uint32(dataAudio)

	//*******************************
	//    ОТОБРАЖЕНИЕ
	//*******************************

	fmt.Println("strRiff -->", pw.сhunkRIFF)
	fmt.Println("strsize -->", pw.сhunkSizeRIFF)
	fmt.Println("strWAVE -->", pw.сhunkWave)

	fmt.Println("strFmt -->", pw.сhunkFmt)
	fmt.Println("сhunkSizeFmt -->", pw.сhunkSizeFmt)
	fmt.Println("audioformat -->", pw.audioformat)
	fmt.Println("numChannels -->", pw.numChannels)
	fmt.Println("samplerate -->", pw.sampleRate)
	fmt.Println("byteRate -->", pw.byteRate)
	fmt.Println("blockAling -->", pw.blockAling)
	fmt.Println("bitperSample -->", pw.bitperSample)

	fmt.Println("strdata -->", pw.сhunkData)
	fmt.Println("sizeData -->", pw.сhunkSizeData)
	fmt.Println("Data_audio -->", pw.dataaudio)
	fmt.Println("Data_lenght -->", len(pw.dataaudio))

	return nil, nil
}
