package libgrib2

import (
	"bytes"
	"encoding/binary"
	"time"
)

type Section1 struct {
	Center        uint16
	SubCenter     uint16
	MasterVersion uint8
	LocalVersion  uint8
	RefTimeMean   uint8
	Status        uint8
	TypeOfData    uint8
	RefTime       time.Time
}

func ReadSection1(data []byte) Section1 {
	var center, subCenter uint16
	var masterVersion, localVersion uint8
	var year uint16
	var month, day, hour, minutes, seconds uint8
	var meaning, status, tod uint8

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &center)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &subCenter)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &masterVersion)
	_ = binary.Read(bytes.NewBuffer(data[10:]), binary.BigEndian, &localVersion)
	_ = binary.Read(bytes.NewBuffer(data[11:]), binary.BigEndian, &meaning)
	_ = binary.Read(bytes.NewBuffer(data[12:]), binary.BigEndian, &year)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &month)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &day)
	_ = binary.Read(bytes.NewBuffer(data[16:]), binary.BigEndian, &hour)
	_ = binary.Read(bytes.NewBuffer(data[17:]), binary.BigEndian, &minutes)
	_ = binary.Read(bytes.NewBuffer(data[18:]), binary.BigEndian, &seconds)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &status)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &tod)
	loc, _ := time.LoadLocation("UTC")

	return Section1{
		Center:        center,
		SubCenter:     subCenter,
		MasterVersion: masterVersion,
		LocalVersion:  localVersion,
		RefTimeMean:   meaning,
		RefTime:       time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minutes), int(seconds), 0, loc),
		Status:        status,
		TypeOfData:    tod,
	}
}
