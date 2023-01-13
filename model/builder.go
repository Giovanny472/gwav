package model

type BuilderWav interface {
	// создание riff
	BuildHeadRiff() Riff

	// создание fmt
	BuildHeadFmt() Fmt

	// созданние данных
	BuildData() Data
}
