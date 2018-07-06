package libgrib2

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

type Section3 struct {
	SourceOfGridDefinition          common.Grib2CodeValue
	NumberOfDataPoints              common.Grib2NumericalValue
	NumberOfOctetsForNumberOfPoints common.Grib2NumericalValue
	InterpretationOfNumberOfPoints  common.Grib2CodeValue
	GridDefinitionTemplateNumber    common.Grib2CodeValue
	Template                        interface{}
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
