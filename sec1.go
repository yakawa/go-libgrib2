package libgrib2

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/yakawa/go-libgrib2/common"
)

type Section1 struct {
	Oct6  uint16
	Oct8  uint16
	Oct10 common.Grib2FlagValue
	Oct11 common.Grib2FlagValue
	Oct12 common.Grib2FlagValue
	Oct13 time.Time
	Oct20 common.Grib2FlagValue
	Oct21 common.Grib2FlagValue
}

func ReadSection1(data []byte) Section1 {
	var o6, o8 uint16
	var o10, o11 uint8
	var year uint16
	var month, day, hour, minutes, seconds uint8
	var o12, o20, o21 uint8

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o6)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o8)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[10:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[11:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[12:]), binary.BigEndian, &year)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &month)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &day)
	_ = binary.Read(bytes.NewBuffer(data[16:]), binary.BigEndian, &hour)
	_ = binary.Read(bytes.NewBuffer(data[17:]), binary.BigEndian, &minutes)
	_ = binary.Read(bytes.NewBuffer(data[18:]), binary.BigEndian, &seconds)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o21)
	loc, _ := time.LoadLocation("UTC")

	return Section1{
		Oct6:  o6,
		Oct8:  o8,
		Oct10: common.ReadFlagValue(o10, 1, 0),
		Oct11: common.ReadFlagValue(o11, 1, 1),
		Oct12: common.ReadFlagValue(o12, 1, 2),
		Oct13: time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minutes), int(seconds), 0, loc),
		Oct20: common.ReadFlagValue(o20, 1, 3),
		Oct21: common.ReadFlagValue(o21, 1, 4),
	}
}
