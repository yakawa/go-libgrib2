package libgrib2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func Read(data []byte) error {
	if data == nil {
		return errors.New("Data is missing")
	}

	dataSize := len(data)

	if dataSize < 4 {
		return errors.New("File size is too small")
	}

	for pos := 0; pos < dataSize; pos++ {
		if string(data[pos:pos+4]) != "GRIB" {
			continue
		}

		sec0 := ReadSection0(data[pos : pos+16])
		fmt.Printf("Sec0: %+v\n", sec0)
		pos += 16
		for i := 0; uint64(i) < sec0.Length; i++ {
			secN, size, _ := readSectionHeader(data[pos : pos+5])
			switch secN {
			case 1:
				sec1 := ReadSection1(data[pos : pos+size])
				fmt.Printf("Sec1: %+v\n", sec1)
			case 2:
				ReadSection2(data[pos : pos+size])
			case 3:
				ReadSection3(data[pos : pos+size])
			case 4:
				ReadSection4(data[pos : pos+size])
			case 5:
				ReadSection5(data[pos : pos+size])
			case 6:
				ReadSection6(data[pos : pos+size])
			case 7:
				ReadSection7(data[pos : pos+size])
			default:
				panic("")
			}
			pos += size
		}

	}
	return nil
}

func readSectionHeader(data []byte) (int, int, error) {
	var size uint32
	var sec uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &size)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &sec)

	return int(sec), int(size), nil
}
