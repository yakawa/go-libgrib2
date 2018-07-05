package libgrib2

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

// Section0
type Section0 struct {
	Oct7 common.Grib2FlagValue
	Oct8 common.Grib2NumericalValue
	Oct9 common.Grib2NumericalValue
}

func ReadSection0(data []byte) Section0 {
	var o7 uint8
	var o8 uint8
	var o9 uint64

	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o7)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o8)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o9)

	return Section0{
		Oct7: common.ReadFlagValue(o7, 0, 0),
		Oct8: common.ReadNumericalValue(o8),
		Oct9: common.ReadNumericalValue(o9),
	}
}
