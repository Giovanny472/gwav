package internal

import "github.com/Giovanny472/gwav/model"

type buildwav struct {
}

var buildw *buildwav

func NewBuildWav() model.BuilderWav {
	if buildw == nil {
		buildw = &buildwav{}
	}

	return buildw
}

// создание riff
func (bw *buildwav) BuildHeadRiff() model.Riff {
	return nil
}

// создание fmt
func (bw *buildwav) BuildHeadFmt() model.Fmt {
	return nil
}

// созданние данных
func (bw *buildwav) BuildData() model.Data {
	return nil
}
