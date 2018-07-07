package libgrib2

/*
 Identification Section
*/

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/yakawa/go-libgrib2/common"
)

// Section1 : Identification Section
type Section1 struct {
	Centre                          common.Grib2CodeValue      // Identification of originating/generating centre (see Common Code Table C-11)
	SubCentre                       common.Grib2NumericalValue // Identification of originating/generating sub-centre (allocated by originating/generating Centre)
	TablesVersion                   common.Grib2CodeValue      // GRIB Master Tables Version Number (see Code Table 1.0 and Note 1)
	LocalTablesVersion              common.Grib2CodeValue      // Version number of GRIB Local Tables used to augment Master Tables (see Code Table 1.1 and Note 2)
	SignificanceOfReferenceTime     common.Grib2CodeValue      // Significance of Reference Time (see Code Table 1.2)
	ReferenceTime                   time.Time                  // Reference time of data
	ProductionStatusOfProcessedData common.Grib2CodeValue      // Production status of processed data in this GRIB message (see Code Table 1.3)
	TypeOfProcessedData             common.Grib2CodeValue      // Type of processed data in this GRIB message (see Code Table 1.4)
}

func readSection1(data []byte) Section1 {
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
