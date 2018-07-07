package libgrib2

/*
 Grid Definition Section (GDS)
*/

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

// Section3 : Grid Definition Section (GDS)
type Section3 struct {
	SourceOfGridDefinition          common.Grib2CodeValue      // Source of grid definition (see Code Table 3.0 and Note 1)
	NumberOfDataPoints              common.Grib2NumericalValue // Number of data points
	NumberOfOctetsForNumberOfPoints common.Grib2NumericalValue // Number of octets for optional list of numbers (see Note 2)
	InterpretationOfNumberOfPoints  common.Grib2CodeValue      // Interpretation of list of numbers (see Code Table 3.11)
	GridDefinitionTemplateNumber    common.Grib2CodeValue      // Grid Definition Template Number (= N) (see Code Table 3.1)
	Template                        interface{}
}

func readSection3(data []byte) Section3 {
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
		log.Printf("Not Impliment Template 3.%d", tn)
	}

	return Section3{
		SourceOfGridDefinition:          common.ReadCodeValue(o6),
		NumberOfDataPoints:              common.ReadUnsignedNumericalValue(o7),
		NumberOfOctetsForNumberOfPoints: common.ReadUnsignedNumericalValue(o11),
		InterpretationOfNumberOfPoints:  common.ReadCodeValue(o12),
		GridDefinitionTemplateNumber:    common.ReadCodeValue(tn),
		Template:                        tmp,
	}
}
