package libgrib2

import (
	"fmt"
	"log"
	"math"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

type Section7 struct {
	Data []interface{}
}

func ReadSection7(data []byte, sec5 Section5, sec6 Section6) Section7 {
	var val []interface{}
	size := sec5.NumberOfValues.Val
	switch sec5.DataRepresentationTemplateNumber.Val {
	case 0:
		R := sec5.Template.(template.Template5_0).ReferenceValue
		E := sec5.Template.(template.Template5_0).BinaryScaleFactor.Val
		D := sec5.Template.(template.Template5_0).DecimalScaleFactor.Val
		nBits := sec5.Template.(template.Template5_0).BitsPerValue.Val
		val = decodeSimplePack(data[5:], size, R, E, D, nBits)
	case 3:
		/* R := sec5.Template.(template.Template5_0).Oct12
		E := sec5.Template.(template.Template5_0).Oct16.Val
		D := sec5.Template.(template.Template5_0).Oct18.Val
		nBits := sec5.Template.(template.Template5_0).Oct20.Val */
		fmt.Printf("Sec5: %+v\n", sec5)
		fmt.Printf("Sec6: %+v\n", sec6)
	default:
		log.Fatalf("Cannot Decode Template 5.%d", sec5.DataRepresentationTemplateNumber.Val)
	}

	return Section7{
		Data: val,
	}
}

func decodeSimplePack(data []byte, size int, R float32, E int, D int, nBits int) []interface{} {
	var vals []interface{}

	for bits := 0; bits < size*8; bits += nBits {
		startByte := bits / 8
		startBit := bits % 8
		endByte := (bits+nBits)/8 + 1

		X, missing := readVal(data[startByte:endByte], startBit, nBits)

		Y := (float64(R) + float64(X)*math.Pow(float64(2.0), float64(E))) / (math.Pow10(D))
		vals = append(vals, common.Grib2FloatValue{Val: float32(Y), Missing: missing})
	}

	return vals
}

func readVal(data []byte, sb int, nBits int) (uint64, bool) {
	var val uint64
	val = 0
	readingCount := 0
	missing := uint64(1)
	for n := 0; n < nBits; n++ {
		missing = missing * 2
	}
	missing = missing - 1
	for readingByte := 0; readingByte < len(data); readingByte++ {
		targetByte := data[readingByte]
		for readingBit := 0; readingBit < 8; readingBit++ {
			if sb <= 0 {
				readingCount++
				val = val * 2
				if targetByte&0x80 == 0x80 {
					val++
				}
			}
			targetByte = targetByte * 2
			sb--
			if readingCount == nBits {
				if val == missing {
					return val, true
				}
				return val, false
			}
		}
	}
	if val == missing {
		return val, true
	}
	return val, false

}
