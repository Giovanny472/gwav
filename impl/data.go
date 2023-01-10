package impl

import "github.com/Giovanny472/gwav/model"

//**************************************//
//    			DATA 					//
//**************************************//

type datawav struct {

	// секция для слова data
	// 4 байта (36-39 позиция). Big-endian form
	chunkData [4]uint8

	// NumSamples * NumChannels * BitsPerSample/8
	// количество байт в данных
	// 4 байта (40-43 позиция). Little-endian form
	chunkSizeData uint32

	// данные могут быть: is uint8, int16, or float32.
	//  (44 - ...n позиция). Little-endian form
	listData []any
}

var dwav *datawav

func NewDataWav(cd [4]uint8, cs uint32, ld []any) Data {
	if dwav == nil {
		dwav = &datawav{chunkData: cd, chunkSizeData: cs, listData: ld}
		return dwav
	}
	return dwav
}

func (dw *datawav) ChunkData() [4]uint8 {
	return dw.chunkData
}

func (dw *datawav) ChunkSizeData() uint32 {
	return dw.chunkSizeData
}

func (dw *datawav) ListData() []any {
	return dw.listData
}
