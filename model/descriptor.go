package model

// интерфейс для чтения файла wav
type Wave interface {
	Header
	Data
}

type Header interface {

	// заголовок
	Riff

	// fmt - описывает формат звуковых данных:
	Fmt
}

type Riff interface {

	// RIFF header

	// секция для слова RIFF
	// 4 байта (0-3 позиция). Big-endian form
	ChunkRIFF() [4]uint8

	// RIFF Chunk Data Size
	// Размер данных секции RIFF
	// 4 байта (4-7 позиция). Little-endian form
	ChunkSizeRIFF() uint32

	// WAVE header
	// секция для слова WAVE типа RIFF
	// 4 байта (8-11 позиция). Big-endian form
	WAVE() [4]uint8
}

type Fmt interface {

	// FMT header
	// секция для слова FMT
	// "fmt" описывает формат звуковых данных:
	// 4 байта (12-15 позиция). Big-endian form
	ChunckFmt() [4]uint8

	// RIFF Chunk Data Size
	// Размер данных секции FMT
	// 4 байта (16-19 позиция). Little-endian form
	ChunkSizeFMT() uint32

	// Compression Code
	// Код типа сжатия аудиоданных
	// 0 (0x0000)	Неизвестный формат
	// 1 (0x0001)	PCM / несжатые данные
	// 2 (0x0002)	Microsoft ADPCM
	// 6 (0x0006)	ITU G.711 a-law
	// 7 (0x0007)	ITU G.711 µ-law
	// 17 (0x0011)	IMA ADPCM
	// 20 (0x0016)	ITU G.723 ADPCM (Yamaha)
	// 49 (0x0031)	GSM 6.10
	// 64 (0x0040)	ITU G.721 ADPCM
	// 80 (0x0050)	MPEG
	// 65,535 (0xFFFF)	Экспериментальный формат
	// 2 байта (20-21 позиция). Little-endian form
	AudioFormat() uint16

	// Number of channels
	// Количество каналов. Моно = 1, Стерео = 2
	// 2 байта (22-23 позиция). Little-endian form
	NumChannels() uint16

	// Sample rate
	// Частота дискретизации. 8000 Гц, 44100 Гц и т.д.
	// 4 байта (24-27 позиция). Little-endian form
	SampleRate() uint32

	// byteRate
	// Количество байт в секунду
	// sampleRate * numChannels * bitsPerSample/8
	// Количество байт в секунду = Частота дискретизации × Размер блока
	// 4 байта (28-31 позиция). Little-endian form
	ByteRate() uint32

	// Block align
	// Размер блока. Количество байтов для одного sample, включая все каналы.
	// numChannels * bitsPerSample/8
	// Размер блока = Количество каналов * Количество значащих бит на выборку / 8
	// 2 байта (32-33 позиция). Little-endian form
	BlockAlign() uint16

	// Significant bits per sample
	// Количество значащих бит на выборку. 8 бит, 16 бит и т.д.
	// 2 байта (34-35 позиция). Little-endian form
	BitsPerSample() uint16
}

type Data interface {

	// секция для слова data
	// 4 байта (36-39 позиция). Big-endian form
	ChunkData() [4]uint8

	// NumSamples * NumChannels * BitsPerSample/8
	// количество байт в данных
	// 4 байта (40-43 позиция). Little-endian form
	ChunkSizeData() uint32

	// данные могут быть: is uint8, int16, or float32.
	//  (44 - ...n позиция). Little-endian form
	ListData() []any
}
