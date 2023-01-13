package model

type indexDataWav int

const (

	// константы для RIFF header
	idxStartWordRiff indexDataWav = 0
	idxEndWordRiff   indexDataWav = 3

	idxStartChunkSzRiff indexDataWav = 4
	idxEndChunkSzRiff   indexDataWav = 7

	idxStartWordWave indexDataWav = 8
	idxEndWordWave   indexDataWav = 11

	// константы для FMT header
	idxStartWordFmt indexDataWav = 12
	idxEndWordFmt   indexDataWav = 15

	idxStartChunkSzFmt indexDataWav = 16
	idxEndChunkSzFmt   indexDataWav = 19

	idxStartAudioformat indexDataWav = 20
	idxEndAudioformat   indexDataWav = 21

	idxStartNumChannels indexDataWav = 22
	idxEndNumChannels   indexDataWav = 23

	idxStartSampleRate indexDataWav = 24
	idxEndSampleRate   indexDataWav = 27

	idxStartByteRate indexDataWav = 28
	idxEndByteRate   indexDataWav = 31

	idxStartBlockAlign indexDataWav = 32
	idxEndBlockAlign   indexDataWav = 33

	idxStartBitsPerSample indexDataWav = 34
	idxEndBitsPerSample   indexDataWav = 35

	// константы для данных
	idxStartWordData indexDataWav = 36
	idxEndWordData   indexDataWav = 39

	idxStartChunkSzData indexDataWav = 40
	idxEndChunkSzData   indexDataWav = 43

	idxStartChunkData indexDataWav = 44
)

type Parser interface {
	// данные файла-wav
	Parse(dw *[]byte) (Wave, error)
}
