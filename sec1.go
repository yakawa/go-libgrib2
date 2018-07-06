package libgrib2

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/yakawa/go-libgrib2/common"
)

type Section1 struct {
	Centre                          common.Grib2CodeValue
	SubCentre                       common.Grib2NumericalValue
	TablesVersion                   common.Grib2CodeValue
	LocalTablesVersion              common.Grib2CodeValue
	SignificanceOfReferenceTime     common.Grib2CodeValue
	ReferenceTime                   time.Time
	ProductionStatusOfProcessedData common.Grib2CodeValue
	TypeOfProcessedData             common.Grib2CodeValue
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
		Centre:                          common.ReadCodeValue(o6),
		SubCentre:                       common.ReadUnsignedNumericalValue(o8),
		TablesVersion:                   common.ReadCodeValue(o10),
		LocalTablesVersion:              common.ReadCodeValue(o11),
		SignificanceOfReferenceTime:     common.ReadCodeValue(o12),
		ReferenceTime:                   time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minutes), int(seconds), 0, loc),
		ProductionStatusOfProcessedData: common.ReadCodeValue(o20),
		TypeOfProcessedData:             common.ReadCodeValue(o21),
	}
}
