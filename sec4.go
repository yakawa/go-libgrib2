package libgrib2

/*
 Product Definition Section (PDS)
*/

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

// Section4 : Product Definition Section (PDS)
type Section4 struct {
	NV                              common.Grib2NumericalValue // Number of coordinate values after template or number of information according to 3D vertical coordinate GRIB2 message
	ProductDefinitionTemplateNumber common.Grib2CodeValue      // Product Definition Template Number (see Code Table 4.0)
	Template                        interface{}
}

func readSection4(data []byte) Section4 {
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
		NV: common.ReadUnsignedNumericalValue(o6),
		ProductDefinitionTemplateNumber: common.ReadCodeValue(tn),
		Template:                        tmp,
	}
}
