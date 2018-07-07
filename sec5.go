package libgrib2

/*
 Data Representation Section
*/
import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

// Section5 : Data Representation Section (PDS)
type Section5 struct {
	NumberOfValues                   common.Grib2NumericalValue // Number of data points where one or more values are specified in Section 7 when a bit map is present, total number of data points when a bit map is absent.
	DataRepresentationTemplateNumber common.Grib2CodeValue      // Data Representation Template Number (see Code Table 5.0)
	Template                         interface{}
}

func readSection5(data []byte) Section5 {
	var o6 uint32
	var tn uint16

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o6)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &tn)

	var tmp interface{}
	switch tn {
	case uint16(0):
		tmp = template.ReadTemplate5_0(data[11:])
	case uint16(2):
		tmp = template.ReadTemplate5_2(data[11:])
	case uint16(3):
		tmp = template.ReadTemplate5_3(data[11:])
	case uint16(0xFFFF):
		tmp = nil
	default:
		log.Fatalf("Not Impliment Template 5.%d", tn)
	}

	return Section5{
		NumberOfValues:                   common.ReadUnsignedNumericalValue(o6),
		DataRepresentationTemplateNumber: common.ReadCodeValue(tn),
		Template: tmp,
	}
}
