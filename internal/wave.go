package internal

import "github.com/Giovanny472/gwav/model"

type wavedescrip struct {

	// заголовок
	headwav model.Header

	//  данные
	datawwav model.Data
}

var wavedescrp *wavedescrip

func NewWave() model.Wave {

	if wavedescrp == nil {
		wavedescrp = &wavedescrip{}
	}
	return wavedescrp
}

func (wd *wavedescrip) Head() model.Header {
	return wd.headwav
}

func (wd *wavedescrip) Data() model.Data {
	return wd.datawwav
}
