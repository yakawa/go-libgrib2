package libgrib2

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

type Section5 struct {
	Oct6           common.Grib2NumericalValue
	TemplateNumber common.Grib2FlagValue
	Template       interface{}
}

func ReadSection5(data []byte) Section5 {
	var o6 uint32
	var tn uint16

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o6)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &tn)

	var tmp interface{}
	switch tn {
	case uint16(0):
		tmp = template.ReadTemplate5_0(data[11:])
	case uint16(0xFFFF):
		tmp = nil
	default:
		log.Fatalf("Not Impliment Template 5.%d", tn)
	}

	return Section5{
		Oct6:           common.ReadNumericalValue(o6),
		TemplateNumber: common.ReadFlagValue(tn, 5, 0),
		Template:       tmp,
	}
}
