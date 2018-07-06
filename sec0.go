package libgrib2

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

// Section0
type Section0 struct {
	Discipline    common.Grib2CodeValue
	EditionNumber common.Grib2NumericalValue
	TotalLength   common.Grib2NumericalValue
}

func ReadSection0(data []byte) Section0 {
	var o7 uint8
	var o8 uint8
	var o9 uint64

	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o7)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o8)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o9)

	return Section0{
		Discipline:    common.ReadCodeValue(o7),
		EditionNumber: common.ReadUnsignedNumericalValue(o8),
		TotalLength:   common.ReadUnsignedNumericalValue(o9),
	}
}
