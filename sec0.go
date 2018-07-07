package libgrib2

/*
 Indicator Section
*/

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

// Section0 : Indicator Section.
type Section0 struct {
	Discipline    common.Grib2CodeValue      // GRIB Master Table Number (see Code Table 0.0)
	EditionNumber common.Grib2NumericalValue // GRIB Edition Number
	TotalLength   common.Grib2NumericalValue // Total length of GRIB message in octets (including Section 0)
}

func readSection0(data []byte) Section0 {
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
