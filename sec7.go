package libgrib2

/*
 Data Section
*/
import (
	"log"
	"math"

	"github.com/yakawa/go-libgrib2/common"
	"github.com/yakawa/go-libgrib2/template"
)

// Section7 : Data Section
type Section7 struct {
	Data []interface{}
}

func readSection7(data []byte, sec5 Section5, sec6 Section6) Section7 {
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
		val = decodeComplexAndSpatialDifferencing(data[5:], size, sec5.Template.(template.Template5_3))
	default:
		log.Fatalf("Cannot Decode Template 5.%d", sec5.DataRepresentationTemplateNumber.Val)
	}

	return Section7{
		Data: val,
	}
}

func decodeComplexAndSpatialDifferencing(data []byte, size int, tmp template.Template5_3) []interface{} {
	var vals []interface{}
	var pos = 0

	sd := tmp.OrderOfSpatialDifferencing.Val
	if sd != 1 && sd != 2 {
		log.Fatalf("Not Supported, Order of Spatial Differencing: %d", sd)
		return vals
	}

	gsm := tmp.GroupSplittingMethodUsed.Val
	if gsm != 1 {
		log.Fatalf("Not Support for Group Split Method: %d", gsm)
		return vals
	}

	var g1 uint64
	var gMin int64
	var h1, h2 uint64
	var hMin int64
	ne := tmp.NumberOfOctetsExtraDescriptors.Val

	if sd == 1 {
		g1 = readNUint(data[pos+0:], ne)
		gMin = readNInt(data[pos+ne*1:], ne)
		pos += 2 * ne
	} else if sd == 2 {
		h1 = readNUint(data[pos+0:], ne)
		h2 = readNUint(data[pos+ne*1:], ne)
		hMin = readNInt(data[pos+ne*2:], ne)
		pos += 3 * ne
	} else {
	}

	ng := tmp.NumberOfGroupsOfDataValues.Val
	nBits := tmp.BitsPerValue.Val

	var X1 []uint64
	for i := 0; i < ng; i++ {
		var v uint64

		startByte := (i * nBits) / 8
		startBit := (i * nBits) % 8
		endByte := (i*nBits+nBits)/8 + 1

		v, _ = readVal(data[pos+startByte:pos+endByte], startBit, nBits)
		X1 = append(X1, v)
	}
	pos += (ng*nBits + 7) / 8

	nBitsGw := tmp.NumberOfBitsUsedForTheGroupWidths.Val
	refGw := tmp.ReferenceForGroupWidth.Val
	var gw []uint64
	for i := 0; i < ng; i++ {
		var v uint64

		startByte := (i * nBitsGw) / 8
		startBit := (i * nBitsGw) % 8
		endByte := (i*nBitsGw+nBitsGw)/8 + 1

		v, _ = readVal(data[pos+startByte:pos+endByte], startBit, nBitsGw)
		gw = append(gw, v+uint64(refGw))
	}
	pos += (ng*nBitsGw + 7) / 8

	nBitsGl := tmp.NumberOfBitsForScaledGroupLengths.Val

	var gl []uint64
	refGl := tmp.ReferenceForGroupLength.Val
	lenInc := tmp.LengthIncrementForTheGroupLengths.Val
	for i := 0; i < ng; i++ {
		var v uint64

		startByte := (i * nBitsGl) / 8
		startBit := (i * nBitsGl) % 8
		endByte := (i*nBitsGl+nBitsGl)/8 + 1

		v, _ = readVal(data[pos+startByte:pos+endByte], startBit, nBitsGl)
		gl = append(gl, v*uint64(lenInc)+uint64(refGl))
	}
	gl[ng-1] = uint64(tmp.TrueLengthOfLastGroup.Val)
	pos += (ng*nBitsGl + 7) / 8

	var dataLocation []uint64
	dataLocation = append(dataLocation, 0)
	d := 0
	for i := 0; i < ng; i++ {
		for j := 0; uint64(j) < gl[i]; j++ {
			dataLocation = append(dataLocation, dataLocation[d]+gw[i])
			d++
		}
	}
	if d != size {
		log.Fatalf("size mismatch expected %d but got %d\n", size, d)
	}

	var X2 []int64
	d = 0
	for i := 0; i < ng; i++ {
		for j := 0; j < int(gl[i]); j++ {
			startByte := (dataLocation[d]) / 8
			startBit := (dataLocation[d]) % 8
			endByte := (dataLocation[d]+gw[i])/8 + 1

			v, _ := readVal(data[pos+int(startByte):pos+int(endByte)], int(startBit), int(gw[i]))
			X2 = append(X2, int64(v))
			d++
		}
	}

	if tmp.MissingValueManagementUsed.Val == 0 {
		d = 0
		for i := 0; i < ng; i++ {
			for j := 0; j < int(gl[i]); j++ {
				X2[d] += int64(X1[i])
				d++
			}
		}
	} else {
		log.Fatal("Not impliment\n")
	}

	if tmp.OrderOfSpatialDifferencing.Val == 1 {
		X2[0] = int64(g1)
		for i := 1; i < size; i++ {
			X2[i] = X2[i] + gMin + X2[i-1]
		}
	} else if tmp.OrderOfSpatialDifferencing.Val == 2 {
		X2[0] = int64(h1)
		X2[1] = int64(h2)
		for i := 2; i < size; i++ {
			X2[i] = X2[i] + hMin + int64(2*X2[i-1]) - int64(X2[i-2])
		}
	}

	R := tmp.ReferenceValue
	E := tmp.BinaryScaleFactor.Val
	D := tmp.DecimalScaleFactor.Val

	for i := 0; i < size; i++ {
		var v float64
		v = (float64(R) + float64(X2[i])*math.Pow(float64(2.0), float64(E))) / (math.Pow10(D))
		vals = append(vals, common.Grib2FloatValue{Val: float32(v), Missing: false})
	}

	return vals
}

func readNUint(data []byte, n int) uint64 {
	var v uint64
	v = 0
	for i := 0; i < n; i++ {
		v = v*256 + uint64(data[i])
	}
	return v
}

func readNInt(data []byte, n int) int64 {
	var v uint64
	v = 0
	sign := 1
	for i := 0; i < n; i++ {
		if i == 0 {
			if (data[0] & 0x80) != 0 {
				sign = -1
				data[0] = data[0] & 0x7F
			}
		}
		v = v*256 + uint64(data[i])
	}
	return int64(sign) * int64(v)
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
	if nBits == 0 {
		return val, false
	}
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
