package libgrib2

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

type Section4 struct {
	Oct6           common.Grib2NumericalValue
	TemplateNumber common.Grib2FlagValue
	Template       interface{}
}

func ReadSection4(data []byte) Section4 {
	var o6 uint16
	var tn uint16

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o6)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &tn)

	var tmp interface{}
	switch tn {
	case uint16(0):
		tmp = template.ReadTemplate4_0(data[9:])
	case uint16(1):
		tmp = template.ReadTemplate4_1(data[9:])
	case uint16(8):
		tmp = template.ReadTemplate4_8(data[9:])
	case uint16(11):
		tmp = template.ReadTemplate4_11(data[9:])
	case uint16(0xFFFF):
		tmp = nil
	default:
		log.Printf("Not Impliment Template 4.%d", tn)
	}

	return Section4{
		Oct6:           common.ReadNumericalValue(o6),
		TemplateNumber: common.ReadFlagValue(tn, 4, 0),
		Template:       tmp,
	}
}
