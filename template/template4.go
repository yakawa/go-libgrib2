package template

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/yakawa/go-libgrib2/common"
)

type Template4_0 struct {
	Oct10 common.Grib2FlagValue
	Oct11 common.Grib2FlagValue
	Oct12 common.Grib2FlagValue
	Oct13 common.Grib2FlagValue
	Oct14 common.Grib2FlagValue
	Oct15 common.Grib2NumericalValue
	Oct17 common.Grib2NumericalValue
	Oct18 common.Grib2FlagValue
	Oct19 common.Grib2NumericalValue
	Oct23 common.Grib2FlagValue
	Oct24 common.Grib2NumericalValue
	Oct25 common.Grib2NumericalValue
	Oct29 common.Grib2FlagValue
	Oct30 common.Grib2NumericalValue
	Oct31 common.Grib2NumericalValue
}

func ReadTemplate4_0(data []byte) Template4_0 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)

	return Template4_0{
		Oct10: common.ReadFlagValue(o10, 4, 1),
		Oct11: common.ReadFlagValue(o11, 4, 2),
		Oct12: common.ReadFlagValue(o12, 4, 3),
		Oct13: common.ReadFlagValue(o13, 99, 0),
		Oct14: common.ReadFlagValue(o14, 99, 0),
		Oct15: common.ReadNumericalValue(o15),
		Oct17: common.ReadNumericalValue(o17),
		Oct18: common.ReadFlagValue(o18, 4, 4),
		Oct19: common.ReadNumericalValue(o19),
		Oct23: common.ReadFlagValue(o23, 4, 5),
		Oct24: common.ReadNumericalValue(o24),
		Oct25: common.ReadNumericalValue(o25),
		Oct29: common.ReadFlagValue(o29, 4, 5),
		Oct30: common.ReadNumericalValue(o30),
		Oct31: common.ReadNumericalValue(o31),
	}

}

type Template4_1 struct {
	Oct10 common.Grib2FlagValue
	Oct11 common.Grib2FlagValue
	Oct12 common.Grib2FlagValue
	Oct13 common.Grib2FlagValue
	Oct14 common.Grib2FlagValue
	Oct15 common.Grib2NumericalValue
	Oct17 common.Grib2NumericalValue
	Oct18 common.Grib2FlagValue
	Oct19 common.Grib2NumericalValue
	Oct23 common.Grib2FlagValue
	Oct24 common.Grib2NumericalValue
	Oct25 common.Grib2NumericalValue
	Oct29 common.Grib2FlagValue
	Oct30 common.Grib2NumericalValue
	Oct31 common.Grib2NumericalValue
	Oct35 common.Grib2FlagValue
	Oct36 common.Grib2NumericalValue
	Oct37 common.Grib2NumericalValue
}

func ReadTemplate4_1(data []byte) Template4_1 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32
	var o35, o36, o37 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &o35)
	_ = binary.Read(bytes.NewBuffer(data[26:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[27:]), binary.BigEndian, &o37)

	return Template4_1{
		Oct10: common.ReadFlagValue(o10, 4, 1),
		Oct11: common.ReadFlagValue(o11, 4, 2),
		Oct12: common.ReadFlagValue(o12, 4, 3),
		Oct13: common.ReadFlagValue(o13, 99, 0),
		Oct14: common.ReadFlagValue(o14, 99, 0),
		Oct15: common.ReadNumericalValue(o15),
		Oct17: common.ReadNumericalValue(o17),
		Oct18: common.ReadFlagValue(o18, 4, 4),
		Oct19: common.ReadNumericalValue(o19),
		Oct23: common.ReadFlagValue(o23, 4, 5),
		Oct24: common.ReadNumericalValue(o24),
		Oct25: common.ReadNumericalValue(o25),
		Oct29: common.ReadFlagValue(o29, 4, 5),
		Oct30: common.ReadNumericalValue(o30),
		Oct31: common.ReadNumericalValue(o31),
		Oct35: common.ReadFlagValue(o35, 4, 6),
		Oct36: common.ReadNumericalValue(o36),
		Oct37: common.ReadNumericalValue(o37),
	}

}

type Template4_8 struct {
	Oct10 common.Grib2FlagValue
	Oct11 common.Grib2FlagValue
	Oct12 common.Grib2FlagValue
	Oct13 common.Grib2FlagValue
	Oct14 common.Grib2FlagValue
	Oct15 common.Grib2NumericalValue
	Oct17 common.Grib2NumericalValue
	Oct18 common.Grib2FlagValue
	Oct19 common.Grib2NumericalValue
	Oct23 common.Grib2FlagValue
	Oct24 common.Grib2NumericalValue
	Oct25 common.Grib2NumericalValue
	Oct29 common.Grib2FlagValue
	Oct30 common.Grib2NumericalValue
	Oct31 common.Grib2NumericalValue
	Oct35 time.Time
	Oct42 common.Grib2NumericalValue
	Oct43 common.Grib2NumericalValue
	Oct47 common.Grib2FlagValue
	Oct48 common.Grib2FlagValue
	Oct49 common.Grib2FlagValue
	Oct50 common.Grib2NumericalValue
	Oct54 common.Grib2FlagValue
	Oct55 common.Grib2NumericalValue
}

func ReadTemplate4_8(data []byte) Template4_8 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32
	var year uint16
	var month, day, hour, minutes, seconds uint8
	var o42 uint8
	var o43 uint32
	var o47, o48, o49 uint8
	var o50 uint32
	var o54 uint8
	var o55 uint32

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &year)
	_ = binary.Read(bytes.NewBuffer(data[27:]), binary.BigEndian, &month)
	_ = binary.Read(bytes.NewBuffer(data[28:]), binary.BigEndian, &day)
	_ = binary.Read(bytes.NewBuffer(data[29:]), binary.BigEndian, &hour)
	_ = binary.Read(bytes.NewBuffer(data[30:]), binary.BigEndian, &seconds)
	_ = binary.Read(bytes.NewBuffer(data[31:]), binary.BigEndian, &minutes)
	_ = binary.Read(bytes.NewBuffer(data[32:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[33:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[37:]), binary.BigEndian, &o47)
	_ = binary.Read(bytes.NewBuffer(data[38:]), binary.BigEndian, &o48)
	_ = binary.Read(bytes.NewBuffer(data[39:]), binary.BigEndian, &o49)
	_ = binary.Read(bytes.NewBuffer(data[40:]), binary.BigEndian, &o50)
	_ = binary.Read(bytes.NewBuffer(data[44:]), binary.BigEndian, &o54)
	_ = binary.Read(bytes.NewBuffer(data[45:]), binary.BigEndian, &o55)

	loc, _ := time.LoadLocation("UTC")

	return Template4_8{
		Oct10: common.ReadFlagValue(o10, 4, 1),
		Oct11: common.ReadFlagValue(o11, 4, 2),
		Oct12: common.ReadFlagValue(o12, 4, 3),
		Oct13: common.ReadFlagValue(o13, 99, 0),
		Oct14: common.ReadFlagValue(o14, 99, 0),
		Oct15: common.ReadNumericalValue(o15),
		Oct17: common.ReadNumericalValue(o17),
		Oct18: common.ReadFlagValue(o18, 4, 4),
		Oct19: common.ReadNumericalValue(o19),
		Oct23: common.ReadFlagValue(o23, 4, 5),
		Oct24: common.ReadNumericalValue(o24),
		Oct25: common.ReadNumericalValue(o25),
		Oct29: common.ReadFlagValue(o29, 4, 5),
		Oct30: common.ReadNumericalValue(o30),
		Oct31: common.ReadNumericalValue(o31),
		Oct35: time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minutes), int(seconds), 0, loc),
		Oct42: common.ReadNumericalValue(o42),
		Oct43: common.ReadNumericalValue(o43),
		Oct47: common.ReadFlagValue(o47, 4, 10),
		Oct48: common.ReadFlagValue(o48, 4, 11),
		Oct49: common.ReadFlagValue(o49, 4, 4),
		Oct50: common.ReadNumericalValue(o50),
		Oct54: common.ReadFlagValue(o54, 4, 4),
		Oct55: common.ReadNumericalValue(o55),
	}

}

type Template4_11 struct {
	Oct10 common.Grib2FlagValue
	Oct11 common.Grib2FlagValue
	Oct12 common.Grib2FlagValue
	Oct13 common.Grib2FlagValue
	Oct14 common.Grib2FlagValue
	Oct15 common.Grib2NumericalValue
	Oct17 common.Grib2NumericalValue
	Oct18 common.Grib2FlagValue
	Oct19 common.Grib2NumericalValue
	Oct23 common.Grib2FlagValue
	Oct24 common.Grib2NumericalValue
	Oct25 common.Grib2NumericalValue
	Oct29 common.Grib2FlagValue
	Oct30 common.Grib2NumericalValue
	Oct31 common.Grib2NumericalValue
	Oct35 common.Grib2FlagValue
	Oct36 common.Grib2NumericalValue
	Oct37 common.Grib2NumericalValue
	Oct38 time.Time
	Oct45 common.Grib2NumericalValue
	Oct46 common.Grib2NumericalValue
	Oct50 common.Grib2FlagValue
	Oct51 common.Grib2FlagValue
	Oct52 common.Grib2FlagValue
	Oct53 common.Grib2NumericalValue
	Oct57 common.Grib2FlagValue
	Oct58 common.Grib2NumericalValue
}

func ReadTemplate4_11(data []byte) Template4_11 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32
	var o35, o36, o37 uint8
	var year uint16
	var month, day, hour, minutes, seconds uint8
	var o45 uint8
	var o46 uint32
	var o50, o51, o52 uint8
	var o53 uint32
	var o57 uint8
	var o58 uint32

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &o35)
	_ = binary.Read(bytes.NewBuffer(data[26:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[27:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[28:]), binary.BigEndian, &year)
	_ = binary.Read(bytes.NewBuffer(data[30:]), binary.BigEndian, &month)
	_ = binary.Read(bytes.NewBuffer(data[31:]), binary.BigEndian, &day)
	_ = binary.Read(bytes.NewBuffer(data[32:]), binary.BigEndian, &hour)
	_ = binary.Read(bytes.NewBuffer(data[33:]), binary.BigEndian, &seconds)
	_ = binary.Read(bytes.NewBuffer(data[34:]), binary.BigEndian, &minutes)
	_ = binary.Read(bytes.NewBuffer(data[35:]), binary.BigEndian, &o45)
	_ = binary.Read(bytes.NewBuffer(data[36:]), binary.BigEndian, &o46)
	_ = binary.Read(bytes.NewBuffer(data[40:]), binary.BigEndian, &o50)
	_ = binary.Read(bytes.NewBuffer(data[41:]), binary.BigEndian, &o51)
	_ = binary.Read(bytes.NewBuffer(data[42:]), binary.BigEndian, &o52)
	_ = binary.Read(bytes.NewBuffer(data[43:]), binary.BigEndian, &o53)
	_ = binary.Read(bytes.NewBuffer(data[47:]), binary.BigEndian, &o57)
	_ = binary.Read(bytes.NewBuffer(data[48:]), binary.BigEndian, &o58)

	loc, _ := time.LoadLocation("UTC")

	return Template4_11{
		Oct10: common.ReadFlagValue(o10, 4, 1),
		Oct11: common.ReadFlagValue(o11, 4, 2),
		Oct12: common.ReadFlagValue(o12, 4, 3),
		Oct13: common.ReadFlagValue(o13, 99, 0),
		Oct14: common.ReadFlagValue(o14, 99, 0),
		Oct15: common.ReadNumericalValue(o15),
		Oct17: common.ReadNumericalValue(o17),
		Oct18: common.ReadFlagValue(o18, 4, 4),
		Oct19: common.ReadNumericalValue(o19),
		Oct23: common.ReadFlagValue(o23, 4, 5),
		Oct24: common.ReadNumericalValue(o24),
		Oct25: common.ReadNumericalValue(o25),
		Oct29: common.ReadFlagValue(o29, 4, 5),
		Oct30: common.ReadNumericalValue(o30),
		Oct31: common.ReadNumericalValue(o31),
		Oct35: common.ReadFlagValue(o35, 4, 6),
		Oct36: common.ReadNumericalValue(o36),
		Oct37: common.ReadNumericalValue(o37),
		Oct38: time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minutes), int(seconds), 0, loc),
		Oct45: common.ReadNumericalValue(o45),
		Oct46: common.ReadNumericalValue(o46),
		Oct50: common.ReadFlagValue(o50, 4, 10),
		Oct51: common.ReadFlagValue(o51, 4, 11),
		Oct52: common.ReadFlagValue(o52, 4, 4),
		Oct53: common.ReadNumericalValue(o53),
		Oct57: common.ReadFlagValue(o57, 4, 4),
		Oct58: common.ReadNumericalValue(o58),
	}

}
