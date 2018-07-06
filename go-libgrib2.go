package libgrib2

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Grib2 struct {
	sec1 Section1
	sec2 Section2
	sec3 Section3
	sec4 Section4
	sec5 Section5
	sec6 Section6
	sec7 Section7
}

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

		sec0 := ReadSection0(data[pos : pos+16])
		pos += 16
		var sec1 Section1
		var sec2 Section2
		var sec3 Section3
		var sec4 Section4
		var sec5 Section5
		var sec6 Section6

		for pos < sec0.TotalLength.Val {

			secN, size, _ := readSectionHeader(data[pos : pos+5])
			switch secN {
			case 1:
				sec1 = ReadSection1(data[pos : pos+size])
			case 2:
				ReadSection2(data[pos : pos+size])
			case 3:
				sec3 = ReadSection3(data[pos : pos+size])
			case 4:
				sec4 = ReadSection4(data[pos : pos+size])
			case 5:
				sec5 = ReadSection5(data[pos : pos+size])
			case 6:
				sec6 = ReadSection6(data[pos : pos+size])
			case 7:
				sec7 := ReadSection7(data[pos:pos+size], sec5, sec6)
				grib2 := Grib2{
					sec1: sec1,
					sec2: sec2,
					sec3: sec3,
					sec4: sec4,
					sec5: sec5,
					sec6: sec6,
					sec7: sec7,
				}
				Grib2s = append(Grib2s, grib2)
			default:
				if string(data[pos:pos+4]) == "7777" {
					size = 4
				}
			}
			pos += size
		}

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
