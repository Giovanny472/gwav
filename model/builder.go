package model

type BuilderWav interface {
	// создание riff
	BuildHeadRiff(rf string, sr uint32, wv string)

	// создание fmt
	BuildHeadFmt(fm string, sfm uint32, af uint16, nc uint16,
		sr uint32, br uint32, ba uint16, bs uint16)

	// созданние данных
	BuildData(cd string, cs uint32, ld map[WavChannels][]uint32)

	// cоздание wav-файл
	WavFile() Wave
}
