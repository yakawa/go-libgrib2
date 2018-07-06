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

type Template5_2 struct {
	Oct12 float32
	Oct16 common.Grib2NumericalValue
	Oct18 common.Grib2NumericalValue
	Oct20 common.Grib2NumericalValue
	Oct21 common.Grib2FlagValue
	Oct22 common.Grib2FlagValue
	Oct23 common.Grib2FlagValue
	Oct24 common.Grib2NumericalValue
	Oct28 common.Grib2NumericalValue
	Oct32 common.Grib2NumericalValue
	Oct36 common.Grib2NumericalValue
	Oct37 common.Grib2NumericalValue
	Oct38 common.Grib2NumericalValue
	Oct42 common.Grib2NumericalValue
	Oct43 common.Grib2NumericalValue
	Oct47 common.Grib2NumericalValue
}

func ReadTemplate5_2(data []byte) Template5_2 {
	var o12 float32
	var o16, o18 uint16
	var o20, o21, o22, o23 uint8
	var o24, o28, o32 uint32
	var o36, o37 uint8
	var o38 uint32
	var o42 uint8
	var o43 uint32
	var o47 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o21)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o28)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o32)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o38)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o47)

	return Template5_2{
		Oct12: o12,
		Oct16: common.ReadNumericalValue(o16),
		Oct18: common.ReadNumericalValue(o18),
		Oct20: common.ReadNumericalValue(o20),
		Oct21: common.ReadFlagValue(o21, 5, 1),
		Oct22: common.ReadFlagValue(o22, 5, 4),
		Oct23: common.ReadFlagValue(o23, 5, 5),
		Oct24: common.ReadNumericalValue(o24),
		Oct28: common.ReadNumericalValue(o28),
		Oct32: common.ReadNumericalValue(o32),
		Oct36: common.ReadNumericalValue(o36),
		Oct37: common.ReadNumericalValue(o37),
		Oct38: common.ReadNumericalValue(o38),
		Oct42: common.ReadNumericalValue(o42),
		Oct43: common.ReadNumericalValue(o43),
		Oct47: common.ReadNumericalValue(o47),
	}
}

type Template5_3 struct {
	Oct12 float32
	Oct16 common.Grib2NumericalValue
	Oct18 common.Grib2NumericalValue
	Oct20 common.Grib2NumericalValue
	Oct21 common.Grib2FlagValue
	Oct22 common.Grib2FlagValue
	Oct23 common.Grib2FlagValue
	Oct24 common.Grib2NumericalValue
	Oct28 common.Grib2NumericalValue
	Oct32 common.Grib2NumericalValue
	Oct36 common.Grib2NumericalValue
	Oct37 common.Grib2NumericalValue
	Oct38 common.Grib2NumericalValue
	Oct42 common.Grib2NumericalValue
	Oct43 common.Grib2NumericalValue
	Oct47 common.Grib2NumericalValue
	Oct48 common.Grib2NumericalValue
	Oct49 common.Grib2NumericalValue
}

func ReadTemplate5_3(data []byte) Template5_3 {
	var o12 float32
	var o16, o18 uint16
	var o20, o21, o22, o23 uint8
	var o24, o28, o32 uint32
	var o36, o37 uint8
	var o38 uint32
	var o42 uint8
	var o43 uint32
	var o47, o48, o49 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o21)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o28)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o32)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o38)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o47)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o48)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o49)

	return Template5_3{
		Oct12: o12,
		Oct16: common.ReadNumericalValue(o16),
		Oct18: common.ReadNumericalValue(o18),
		Oct20: common.ReadNumericalValue(o20),
		Oct21: common.ReadFlagValue(o21, 5, 1),
		Oct22: common.ReadFlagValue(o22, 5, 4),
		Oct23: common.ReadFlagValue(o23, 5, 5),
		Oct24: common.ReadNumericalValue(o24),
		Oct28: common.ReadNumericalValue(o28),
		Oct32: common.ReadNumericalValue(o32),
		Oct36: common.ReadNumericalValue(o36),
		Oct37: common.ReadNumericalValue(o37),
		Oct38: common.ReadNumericalValue(o38),
		Oct42: common.ReadNumericalValue(o42),
		Oct43: common.ReadNumericalValue(o43),
		Oct47: common.ReadNumericalValue(o47),
		Oct48: common.ReadNumericalValue(o48),
		Oct49: common.ReadNumericalValue(o49),
	}
}
