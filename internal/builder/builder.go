package builder

import (
	"github.com/Giovanny472/gwav/internal/entities"
	"github.com/Giovanny472/gwav/model"
)

type buildwav struct {
	br model.Riff
	bf model.Fmt
	bd model.Data
}

var buildw *buildwav

func NewBuildWav() model.BuilderWav {
	if buildw == nil {
		buildw = &buildwav{}
	}
	return buildw
}

// создание riff
func (bw *buildwav) BuildHeadRiff(rf string, sr uint32, wv string) {
	bw.br = entities.NewHeadRiff(rf, sr, wv)
}

// создание fmt
func (bw *buildwav) BuildHeadFmt(fm string, sfm uint32, af uint16, nc uint16,
	sr uint32, br uint32, ba uint16, bs uint16) {

	bw.bf = entities.NewHeadFmt(fm, sfm, af, nc, sr, br, ba, bs)
}

// созданние данных
func (bw *buildwav) BuildData(cd string, cs uint32, ld map[model.WavChannels][]uint32) {
	bw.bd = entities.NewDataWav(cd, cs, ld)
}

func (bw *buildwav) WavFile() model.Wave {

	return entities.NewWave(entities.NewHeadWav(bw.br, bw.bf), bw.bd)
}
