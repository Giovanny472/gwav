package entities

import "github.com/Giovanny472/gwav/model"

//**************************************//
//    			RIFF 					//
//**************************************//
type headwav struct {

	// заголовок riff
	riff model.Riff

	// заголовок fmt
	fmt model.Fmt
}

var headw *headwav

func NewHeadWav(r model.Riff, f model.Fmt) model.Header {

	if headw == nil {
		headw = &headwav{riff: r, fmt: f}
	}
	return headw
}

func (hw *headwav) Riff() model.Riff {
	return hw.riff
}

func (hw *headwav) Fmt() model.Fmt {
	return hw.fmt
}

//**************************************//
//    			RIFF 					//
//**************************************//

type headRiff struct {

	// RIFF header
	// секция для слова RIFF
	// 4 байта
	riff string

	// RIFF Chunk Data Size
	// Размер данных секции RIFF
	// 4 байта
	chunkSizeRIFF uint32

	// WAVE header
	// секция для слова WAVE типа RIFF
	// 4 байта
	wave string
}

var hr *headRiff

func NewHeadRiff(rf string, sr uint32, wv string) model.Riff {
	if hr == nil {
		hr = &headRiff{riff: rf, chunkSizeRIFF: sr, wave: wv}
	}
	return hr
}

func (hrf *headRiff) ChunkRIFF() string {
	return hr.riff
}

func (hrf *headRiff) ChunkSizeRIFF() uint32 {
	return hrf.chunkSizeRIFF
}

func (hrf *headRiff) ChunkWave() string {
	return hrf.wave
}

//**************************************//
//    			FMT 					//
//**************************************//

// Структура, которая описывает формат звуковых данных:
type headFmt struct {

	// секция для слова FMT
	// 4 байта
	fmt string

	// RIFF Chunk Data Size
	// Размер данных секции FMT
	// 4 байта
	chunkSizeFMT uint32

	// Compression Code
	// 2 байта
	audioFormat uint16

	// Number of channels
	// 2 байта
	numChannels uint16

	// Sample rate
	// Частота дискретизации. 8000 Гц, 44100 Гц и т.д.
	// 4 байта
	sampleRate uint32

	// byteRate
	// Количество байт в секунду
	// sampleRate * numChannels * bitsPerSample/8
	// Количество байт в секунду = Частота дискретизации × Размер блока
	// 4 байта
	byteRate uint32

	// Block align
	// Размер блока. Количество байтов для одного sample, включая все каналы.
	// numChannels * bitsPerSample/8
	// Размер блока = Количество каналов * Количество значащих бит на выборку / 8
	// 2 байта
	blockAlign uint16

	// Significant bits per sample
	// Количество значащих бит на выборку. 8 бит, 16 бит и т.д.
	bitsPerSample uint16
}

var hfmt *headFmt

func NewHeadFmt(fm string, sfm uint32, af uint16, nc uint16,
	sr uint32, br uint32, ba uint16, bs uint16) model.Fmt {

	if hfmt == nil {
		hfmt = &headFmt{fmt: fm, chunkSizeFMT: sfm, audioFormat: af, numChannels: nc,
			sampleRate: sr, byteRate: br, blockAlign: ba, bitsPerSample: bs}
	}
	return hfmt
}

func (hf *headFmt) ChunckFmt() string {
	return hf.fmt
}

func (hf *headFmt) ChunkSizeFMT() uint32 {
	return hf.chunkSizeFMT
}

func (hf *headFmt) AudioFormat() uint16 {
	return hf.audioFormat
}

func (hf *headFmt) NumChannels() uint16 {
	return hf.numChannels
}

func (hf *headFmt) SampleRate() uint32 {
	return hf.sampleRate
}

func (hf *headFmt) ByteRate() uint32 {
	return hf.byteRate
}

func (hf *headFmt) BlockAlign() uint16 {
	return hf.blockAlign
}

func (hf *headFmt) BitsPerSample() uint16 {
	return hf.bitsPerSample
}
