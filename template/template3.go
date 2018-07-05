package template

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

type Template3_0 struct {
	Oct15 common.Grib2FlagValue
	Oct16 common.Grib2NumericalValue
	Oct17 common.Grib2NumericalValue
	Oct21 common.Grib2NumericalValue
	Oct22 common.Grib2NumericalValue
	Oct26 common.Grib2NumericalValue
	Oct27 common.Grib2NumericalValue
	Oct31 common.Grib2NumericalValue
	Oct35 common.Grib2NumericalValue
	Oct39 common.Grib2NumericalValue
	Oct43 common.Grib2NumericalValue
	Oct47 common.Grib2NumericalValue
	Oct51 common.Grib2NumericalValue
	Oct55 common.Grib2FlagValue
	Oct56 common.Grib2NumericalValue
	Oct60 common.Grib2NumericalValue
	Oct64 common.Grib2NumericalValue
	Oct68 common.Grib2NumericalValue
	Oct72 common.Grib2FlagValue
}

func ReadTemplate3_0(data []byte) Template3_0 {
	var o15, o16 uint8
	var o17 uint32
	var o21 uint8
	var o22 uint32
	var o26 uint8
	var o27 uint32
	var o31, o35, o39, o43, o47, o51 uint32
	var o55 uint8
	var o56, o60, o64, o68 uint32
	var o72 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o21)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[11:]), binary.BigEndian, &o26)
	_ = binary.Read(bytes.NewBuffer(data[12:]), binary.BigEndian, &o27)
	_ = binary.Read(bytes.NewBuffer(data[16:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o35)
	_ = binary.Read(bytes.NewBuffer(data[24:]), binary.BigEndian, &o39)
	_ = binary.Read(bytes.NewBuffer(data[28:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[32:]), binary.BigEndian, &o47)
	_ = binary.Read(bytes.NewBuffer(data[36:]), binary.BigEndian, &o51)
	_ = binary.Read(bytes.NewBuffer(data[40:]), binary.BigEndian, &o55)
	_ = binary.Read(bytes.NewBuffer(data[41:]), binary.BigEndian, &o56)
	_ = binary.Read(bytes.NewBuffer(data[45:]), binary.BigEndian, &o60)
	_ = binary.Read(bytes.NewBuffer(data[49:]), binary.BigEndian, &o64)
	_ = binary.Read(bytes.NewBuffer(data[53:]), binary.BigEndian, &o68)
	_ = binary.Read(bytes.NewBuffer(data[57:]), binary.BigEndian, &o72)

	return Template3_0{
		Oct15: common.ReadFlagValue(o15, 3, 2),
		Oct16: common.ReadNumericalValue(o16),
		Oct17: common.ReadNumericalValue(o17),
		Oct21: common.ReadNumericalValue(o21),
		Oct22: common.ReadNumericalValue(o22),
		Oct26: common.ReadNumericalValue(o26),
		Oct27: common.ReadNumericalValue(o27),
		Oct31: common.ReadNumericalValue(o31),
		Oct35: common.ReadNumericalValue(o35),
		Oct39: common.ReadNumericalValue(o39),
		Oct43: common.ReadNumericalValue(o43),
		Oct47: common.ReadNumericalValue(o47),
		Oct51: common.ReadNumericalValue(o51),
		Oct55: common.ReadFlagValue(o55, 3, 3),
		Oct56: common.ReadNumericalValue(o56),
		Oct60: common.ReadNumericalValue(o60),
		Oct64: common.ReadNumericalValue(o64),
		Oct68: common.ReadNumericalValue(o68),
		Oct72: common.ReadFlagValue(o72, 3, 4),
	}

}
