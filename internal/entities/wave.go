package entities

import "github.com/Giovanny472/gwav/model"

type wavedescrip struct {

	// заголовок
	headwav model.Header

	//  данные
	datawwav model.Data
}

var wavedescrp *wavedescrip

func NewWave(h model.Header, d model.Data) model.Wave {

	if wavedescrp == nil {
		wavedescrp = &wavedescrip{headwav: h, datawwav: d}
	}
	return wavedescrp
}

func (wd *wavedescrip) Head() model.Header {
	return wd.headwav
}

func (wd *wavedescrip) Data() model.Data {
	return wd.datawwav
}
