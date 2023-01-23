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
	dataaudio     map[model.WavChannels][]uint32

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

	// читаем заголовок
	pw.readHeader(dw)

	// читаем данные
	pw.readData(dw)

	// отображение данных
	fmt.Println("Riff: ", pw.сhunkRIFF)
	fmt.Println("Riffsize: ", pw.сhunkSizeRIFF)
	fmt.Println("WAVE: ", pw.сhunkWave)

	fmt.Println("Fmt: ", pw.сhunkFmt)
	fmt.Println("сhunkSizeFmt: ", pw.сhunkSizeFmt)
	fmt.Println("audioformat: ", pw.audioformat)
	fmt.Println("numChannels: ", pw.numChannels)
	fmt.Println("sampleRate: ", pw.sampleRate)
	fmt.Println("byteRate: ", pw.byteRate)
	fmt.Println("blockAling: ", pw.blockAling)
	fmt.Println("bitperSample: ", pw.bitperSample)

	fmt.Println("data :", pw.сhunkData)
	fmt.Println("data size: ", pw.сhunkSizeData)
	fmt.Println("dataaudio: ", pw.dataaudio)

	return nil, nil
}

func (pw *parsewav) readHeader(dw *[]byte) error {

	//*******************************
	//    З А Г О Л О В О К
	//*******************************

	// cлово RIFF
	pw.сhunkRIFF = string((*dw)[model.IdxStartWordRiff:model.IdxEndWordRiff])
	if pw.сhunkRIFF != string(model.ConstRIFF) {
		return errors.New("слово RIFF не найдено. Значение: " + pw.сhunkRIFF)
	}

	// size RIFF
	szRiff := (*dw)[model.IdxStartChunkSzRiff:model.IdxEndChunkSzRiff]
	pw.сhunkSizeRIFF = binary.LittleEndian.Uint32(szRiff)

	// слово WAVE
	pw.сhunkWave = string((*dw)[model.IdxStartWordWave:model.IdxEndWordWave])
	if pw.сhunkWave != string(model.ConstWAVE) {
		return errors.New("слово WAVE не найдено. Значение: " + pw.сhunkWave)
	}

	// cлово fmt
	pw.сhunkFmt = string((*dw)[model.IdxStartWordFmt:model.IdxEndWordFmt])
	if pw.сhunkRIFF != string(model.ConstRIFF) {
		return errors.New("слово FMT не найдено. Значение:" + pw.сhunkFmt)
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

	return nil
}

func (pw *parsewav) readData(dw *[]byte) error {

	//*******************************
	//    Д А Н Н Ы Е
	//*******************************

	// chunk data
	pw.сhunkData = string((*dw)[model.IdxStartWordData:model.IdxEndWordData])
	if pw.сhunkData != string(model.ConstData) {
		return errors.New("слово data не найдено. Значение: " + pw.сhunkData)
	}

	// size RIFF
	szData := (*dw)[model.IdxStartChunkSzData:model.IdxEndChunkSzData]
	pw.сhunkSizeData = binary.LittleEndian.Uint32(szData)

	// data audio
	datawav := (*dw)[model.IdxStartChunkData:]

	// ТРЕБУЮТСЯ:
	// [1] bitsample   = 16, 32, 64  // количество бит для одного sample
	// [2] blockAling  = 2, 4        // Количество байтов для одного sample
	// [3] numchannels = 1, 2        // Количество каналов. Моно = 1, Стерео = 2

	// настройка slice для хранения sample
	pw.setSizeSliceData()

	// настройка на основе  размер slice для хранения
	for idx := 0; idx < len(datawav); idx += int(pw.blockAling) {

		// получаем один SAMPLE с помощью BLOCKALIGN
		sample := datawav[idx : idx+int(pw.blockAling)]

		// назначение sample
		pw.setSample(sample)

	}

	return nil

}

func (pw *parsewav) setSample(sp []byte) {

	// значение

	// назначение sample на основе numchannel
	switch pw.numChannels {

	// МОНО
	case uint16(model.ConstMono):
		{
			// littleendian
			val := binary.LittleEndian.Uint32(sp)

			// назначение
			pw.dataaudio[model.ConstMonoCh] = append(pw.dataaudio[model.ConstMonoCh], val)
		}

	case uint16(model.ConstStereo):
		{

		}
	}
}

func (pw *parsewav) setSizeSliceData() {

	// найдем размер slice
	sz := int(pw.сhunkSizeData) / int(pw.blockAling)

	// определим тип данных для slice
	//tdslice := pw.blockAling * pw.bitperSample

	// настройка slice для хранения sample
	switch pw.numChannels {

	// МОНО
	case uint16(model.ConstMono):
		pw.dataaudio[model.ConstMonoCh] = make([]uint32, sz)
		pw.dataaudio[model.ConstStereoRightCh] = nil
		pw.dataaudio[model.ConstStereoLeftCh] = nil

	// STEREO
	case uint16(model.ConstStereo):

		pw.dataaudio[model.ConstMonoCh] = nil
		pw.dataaudio[model.ConstStereoRightCh] = make([]uint32, sz)
		pw.dataaudio[model.ConstStereoLeftCh] = make([]uint32, sz)
	}

}
