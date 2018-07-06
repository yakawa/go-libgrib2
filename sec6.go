package libgrib2

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

type Section6 struct {
	BitmapIndicator common.Grib2CodeValue
	Bitmap          interface{}
}

func ReadSection6(data []byte) Section6 {
	var o6 uint8

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o6)

	return Section6{
		BitmapIndicator: common.ReadCodeValue(o6),
	}
}
