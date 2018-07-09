/*
Package libgrib2 : GRIB2 Reader Library for Go-lang.
*/
package libgrib2

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// Grib2 : All message for GRIB2.
type Grib2 struct {
	Sec1 Section1
	Sec2 Section2
	Sec3 Section3
	Sec4 Section4
	Sec5 Section5
	Sec6 Section6
	Sec7 Section7
}

// Read is function for reading grib2 file
func Read(data []byte) ([]Grib2, error) {
	var Grib2s []Grib2

	if data == nil {
		return Grib2s, errors.New("Data is missing")
	}

	dataSize := len(data)

	if dataSize < 4 {
		return Grib2s, errors.New("File size is too small")
	}

	for pos := 0; pos < dataSize; pos++ {
		if string(data[pos:pos+4]) != "GRIB" {
			continue
		}
		gPos := 0
		sec0 := readSection0(data[pos : pos+16])

		gPos += 16
		var sec1 Section1
		var sec2 Section2
		var sec3 Section3
		var sec4 Section4
		var sec5 Section5
		var sec6 Section6

		for gPos < sec0.TotalLength.Val {

			secN, size, _ := readSectionHeader(data[pos+gPos : pos+gPos+5])
			switch secN {
			case 1:
				sec1 = readSection1(data[pos+gPos : pos+gPos+size])
			case 2:
				sec2 = readSection2(data[pos+gPos : pos+gPos+size])
			case 3:
				sec3 = readSection3(data[pos+gPos : pos+gPos+size])
			case 4:
				sec4 = readSection4(data[pos+gPos : pos+gPos+size])
			case 5:
				sec5 = readSection5(data[pos+gPos : pos+gPos+size])
			case 6:
				sec6 = readSection6(data[pos+gPos : pos+gPos+size])
			case 7:
				sec7 := readSection7(data[pos+gPos:pos+gPos+size], sec5, sec6)
				grib2 := Grib2{
					Sec1: sec1,
					Sec2: sec2,
					Sec3: sec3,
					Sec4: sec4,
					Sec5: sec5,
					Sec6: sec6,
					Sec7: sec7,
				}
				Grib2s = append(Grib2s, grib2)
			default:
				if string(data[pos:pos+4]) == "7777" {
					size = 4
				}
			}
			gPos += size
		}
		pos += sec0.TotalLength.Val

	}
	return Grib2s, nil
}

func readSectionHeader(data []byte) (int, int, error) {
	var size uint32
	var sec uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &size)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &sec)

	return int(sec), int(size), nil
}
