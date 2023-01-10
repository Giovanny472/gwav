package gwav

//**************************************//
//    			RIFF 					//
//**************************************//

type headRiff struct {

	// RIFF header
	// секция для слова RIFF
	// 4 байта
	riff [4]uint8

	// RIFF Chunk Data Size
	// Размер данных секции RIFF
	// 4 байта
	chunkSizeRIFF uint32

	// WAVE header
	// секция для слова WAVE типа RIFF
	// 4 байта
	wave [4]uint8
}

var hr *headRiff

func NewHeadRiff(rf [4]uint8, sr uint32, wv [4]uint8) Riff {
	if hr == nil {
		hr = &headRiff{riff: rf, chunkSizeRIFF: sr, wave: wv}
		return hr
	}
	return hr
}

func (hrf *headRiff) ChunkRIFF() [4]uint8 {
	return hr.riff
}

func (hrf *headRiff) ChunkSizeRIFF() uint32 {
	return hrf.chunkSizeRIFF
}

func (hrf *headRiff) WAVE() [4]uint8 {
	return hrf.wave
}

//**************************************//
//    			FMT 					//
//**************************************//

// Структура, которая описывает формат звуковых данных:
type headFmt struct {

	// секция для слова FMT
	// 4 байта
	fmt [4]uint8

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

func NewHeadFmt(fm [4]uint8, sfm uint32, af uint16, nc uint16, sr uint32, br uint32, ba uint16, bs uint16) Fmt {
	if hfmt == nil {
		hfmt = &headFmt{fmt: fm, chunkSizeFMT: sfm, audioFormat: af, numChannels: nc,
			sampleRate: sr, byteRate: br, blockAlign: ba, bitsPerSample: bs}
		return hfmt
	}
	return hfmt
}

func (hf *headFmt) ChunckFmt() [4]uint8 {
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
