package libgrib2

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

type Section3 struct {
	Oct6           common.Grib2FlagValue
	Oct7           common.Grib2NumericalValue
	Oct11          common.Grib2NumericalValue
	Oct12          common.Grib2FlagValue
	TemplateNumber common.Grib2FlagValue
	Template       interface{}
}

func ReadSection3(data []byte) Section3 {
	var o6 uint8
	var o7 uint32
	var o11, o12 uint8
	var tn uint16

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o6)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o7)
	_ = binary.Read(bytes.NewBuffer(data[10:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[11:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[12:]), binary.BigEndian, &tn)

	var tmp interface{}
	switch tn {
	case uint16(0):
		tmp = template.ReadTemplate3_0(data[14:])
	case uint16(0xFFFF):
		tmp = nil
	default:
		log.Fatalf("Not Impliment Template 3.%d", tn)
	}

	return Section3{
		Oct6:           common.ReadFlagValue(o6, 3, 0),
		Oct7:           common.ReadNumericalValue(o7),
		Oct11:          common.ReadNumericalValue(o11),
		Oct12:          common.ReadFlagValue(o12, 3, 11),
		TemplateNumber: common.ReadFlagValue(tn, 3, 1),
		Template:       tmp,
	}
}
