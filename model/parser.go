package model

type WavConst string

const (
	ConstRIFF WavConst = "RIFF"
	ConstWAVE WavConst = "WAVE"
	ConstFMT  WavConst = "fmt"
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
	IdxEndChunkSzFmt   indexDataWav = 19

	IdxStartAudioformat indexDataWav = 20
	IdxEndAudioformat   indexDataWav = 21

	IdxStartNumChannels indexDataWav = 22
	IdxEndNumChannels   indexDataWav = 23

	IdxStartSampleRate indexDataWav = 24
	IdxEndSampleRate   indexDataWav = 27

	IdxStartByteRate indexDataWav = 28
	IdxEndByteRate   indexDataWav = 31

	IdxStartBlockAlign indexDataWav = 32
	IdxEndBlockAlign   indexDataWav = 33

	IdxStartBitsPerSample indexDataWav = 34
	IdxEndBitsPerSample   indexDataWav = 35

	// константы для данных
	IdxStartWordData indexDataWav = 36
	IdxEndWordData   indexDataWav = 39

	IdxStartChunkSzData indexDataWav = 40
	IdxEndChunkSzData   indexDataWav = 43

	IdxStartChunkData indexDataWav = 44
)

type Parser interface {
	// данные файла-wav
	Parse(dw *[]byte) (Wave, error)
}
