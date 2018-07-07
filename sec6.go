package libgrib2

/*
 Bitmap Section
*/
import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

// Section6 : Bitmap Section
type Section6 struct {
	BitmapIndicator common.Grib2CodeValue // Bit-map indicator (see Code Table 6.0 and Note 1)
	Bitmap          interface{}
}

func readSection6(data []byte) Section6 {
	var o6 uint8

	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o6)

	return Section6{
		BitmapIndicator: common.ReadCodeValue(o6),
	}
}
