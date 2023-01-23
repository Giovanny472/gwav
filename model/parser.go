package model

type WavTypeBits uint

const (
	Const16Bits = 16
	Const32Bits = 32
	Const64Bits = 64
)

type WavTypeAduioFormat uint

const (
	ConstAudioFormatPCM       = 1
	ConstAudioFormatIEEEFLOAT = 3
	ConstAudioFormatALAW      = 6
)

type WavTypeChannel uint

const (
	ConstMono   WavTypeChannel = 1
	ConstStereo WavTypeChannel = 2
)

type WavChannels uint

const (
	ConstMonoCh        WavChannels = 0
	ConstStereoRightCh WavChannels = 1
	ConstStereoLeftCh  WavChannels = 2
)

type WavConst string

const (
	ConstRIFF WavConst = "RIFF"
	ConstWAVE WavConst = "WAVE"
	ConstFMT  WavConst = "fmt"
	ConstData WavConst = "data"
)

type indexDataWav int

const (

	// константы для RIFF header
	IdxStartWordRiff indexDataWav = 0
	IdxEndWordRiff   indexDataWav = 4

	IdxStartChunkSzRiff indexDataWav = 4
	IdxEndChunkSzRiff   indexDataWav = 8

	IdxStartWordWave indexDataWav = 8
	IdxEndWordWave   indexDataWav = 12

	// константы для FMT header
	IdxStartWordFmt indexDataWav = 12
	IdxEndWordFmt   indexDataWav = 16

	IdxStartChunkSzFmt indexDataWav = 16
	IdxEndChunkSzFmt   indexDataWav = 20

	IdxStartAudioformat indexDataWav = 20
	IdxEndAudioformat   indexDataWav = 22

	IdxStartNumChannels indexDataWav = 22
	IdxEndNumChannels   indexDataWav = 24

	IdxStartSampleRate indexDataWav = 24
	IdxEndSampleRate   indexDataWav = 28

	IdxStartByteRate indexDataWav = 28
	IdxEndByteRate   indexDataWav = 32

	IdxStartBlockAlign indexDataWav = 32
	IdxEndBlockAlign   indexDataWav = 34

	IdxStartBitsPerSample indexDataWav = 34
	IdxEndBitsPerSample   indexDataWav = 36

	// константы для данных
	IdxStartWordData indexDataWav = 36
	IdxEndWordData   indexDataWav = 40

	IdxStartChunkSzData indexDataWav = 40
	IdxEndChunkSzData   indexDataWav = 44

	IdxStartChunkData indexDataWav = 44
)

type Parser interface {
	// данные файла-wav
	Parse(dw *[]byte) (Wave, error)
}
