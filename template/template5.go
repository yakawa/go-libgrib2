package template

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

type Template5_0 struct {
	Oct12 float32
	Oct16 common.Grib2NumericalValue
	Oct18 common.Grib2NumericalValue
	Oct20 common.Grib2NumericalValue
	Oct21 common.Grib2FlagValue
}

func ReadTemplate5_0(data []byte) Template5_0 {
	var o12 float32
	var o16 uint16
	var o18 uint16
	var o20 uint8
	var o21 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o21)

	return Template5_0{
		Oct12: o12,
		Oct16: common.ReadNumericalValue(o16),
		Oct18: common.ReadNumericalValue(o18),
		Oct20: common.ReadNumericalValue(o20),
		Oct21: common.ReadFlagValue(o21, 5, 1),
	}
}
