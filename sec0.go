package libgrib2

import (
	"bytes"
	"encoding/binary"
)

// Section0
type Section0 struct {
	Discipline uint8
	Edition    uint8
	Length     uint64
}

func ReadSection0(data []byte) Section0 {
	var discipline uint8
	var edition uint8
	var length uint64

	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &discipline)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &edition)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &length)

	return Section0{
		Discipline: discipline,
		Edition:    edition,
		Length:     length,
	}
}
