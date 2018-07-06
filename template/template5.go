package template

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

type Template5_0 struct {
	ReferenceValue            float32
	BinaryScaleFactor         common.Grib2NumericalValue
	DecimalScaleFactor        common.Grib2NumericalValue
	BitsPerValue              common.Grib2NumericalValue
	TypeOfOriginalFieldValues common.Grib2CodeValue
}

func ReadTemplate5_0(data []byte) Template5_0 {
	var o12 float32
	var o16 uint16
	var o18 uint16
	var o20 uint8
	var o21 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o21)

	return Template5_0{
		ReferenceValue:            o12,
		BinaryScaleFactor:         common.ReadSignedNumericalValue(o16),
		DecimalScaleFactor:        common.ReadSignedNumericalValue(o18),
		BitsPerValue:              common.ReadUnsignedNumericalValue(o20),
		TypeOfOriginalFieldValues: common.ReadCodeValue(o21),
	}
}

type Template5_2 struct {
	ReferenceValue                    float32
	BinaryScaleFactor                 common.Grib2NumericalValue
	DecimalScaleFactor                common.Grib2NumericalValue
	BitsPerValue                      common.Grib2NumericalValue
	TypeOfOriginalFieldValues         common.Grib2CodeValue
	GroupSplittingMethodUsed          common.Grib2CodeValue
	MissingValueManagementUsed        common.Grib2CodeValue
	PrimaryMissingValueSubstitute     common.Grib2NumericalValue
	SecondaryMissingValueSubstitute   common.Grib2NumericalValue
	NumberOfGroupsOfDataValues        common.Grib2NumericalValue
	ReferenceForGroupWidth            common.Grib2NumericalValue
	NumberOfBitsUsedForTheGroupWidths common.Grib2NumericalValue
	ReferenceForGroupLength           common.Grib2NumericalValue
	LengthIncrementForTheGroupLengths common.Grib2NumericalValue
	TrueLengthOfLastGroup             common.Grib2NumericalValue
	NumberOfBitsForScaledGroupLengths common.Grib2NumericalValue
}

func ReadTemplate5_2(data []byte) Template5_2 {
	var o12 float32
	var o16, o18 uint16
	var o20, o21, o22, o23 uint8
	var o24, o28, o32 uint32
	var o36, o37 uint8
	var o38 uint32
	var o42 uint8
	var o43 uint32
	var o47 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o21)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o28)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o32)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o38)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o47)

	return Template5_2{
		ReferenceValue:                    o12,
		BinaryScaleFactor:                 common.ReadSignedNumericalValue(o16),
		DecimalScaleFactor:                common.ReadSignedNumericalValue(o18),
		BitsPerValue:                      common.ReadUnsignedNumericalValue(o20),
		TypeOfOriginalFieldValues:         common.ReadCodeValue(o21),
		GroupSplittingMethodUsed:          common.ReadCodeValue(o22),
		MissingValueManagementUsed:        common.ReadCodeValue(o23),
		PrimaryMissingValueSubstitute:     common.ReadUnsignedNumericalValue(o24),
		SecondaryMissingValueSubstitute:   common.ReadUnsignedNumericalValue(o28),
		NumberOfGroupsOfDataValues:        common.ReadUnsignedNumericalValue(o32),
		ReferenceForGroupWidth:            common.ReadUnsignedNumericalValue(o36),
		NumberOfBitsUsedForTheGroupWidths: common.ReadUnsignedNumericalValue(o37),
		ReferenceForGroupLength:           common.ReadUnsignedNumericalValue(o38),
		LengthIncrementForTheGroupLengths: common.ReadUnsignedNumericalValue(o42),
		TrueLengthOfLastGroup:             common.ReadUnsignedNumericalValue(o43),
		NumberOfBitsForScaledGroupLengths: common.ReadUnsignedNumericalValue(o47),
	}
}

type Template5_3 struct {
	ReferenceValue                    float32
	BinaryScaleFactor                 common.Grib2NumericalValue
	DecimalScaleFactor                common.Grib2NumericalValue
	BitsPerValue                      common.Grib2NumericalValue
	TypeOfOriginalFieldValues         common.Grib2CodeValue
	GroupSplittingMethodUsed          common.Grib2CodeValue
	MissingValueManagementUsed        common.Grib2CodeValue
	PrimaryMissingValueSubstitute     common.Grib2NumericalValue
	SecondaryMissingValueSubstitute   common.Grib2NumericalValue
	NumberOfGroupsOfDataValues        common.Grib2NumericalValue
	ReferenceForGroupWidth            common.Grib2NumericalValue
	NumberOfBitsUsedForTheGroupWidths common.Grib2NumericalValue
	ReferenceForGroupLength           common.Grib2NumericalValue
	LengthIncrementForTheGroupLengths common.Grib2NumericalValue
	TrueLengthOfLastGroup             common.Grib2NumericalValue
	NumberOfBitsForScaledGroupLengths common.Grib2NumericalValue
	OrderOfSpatialDifferencing        common.Grib2CodeValue
	NumberOfOctetsExtraDescriptors    common.Grib2NumericalValue
}

func ReadTemplate5_3(data []byte) Template5_3 {
	var o12 float32
	var o16, o18 uint16
	var o20, o21, o22, o23 uint8
	var o24, o28, o32 uint32
	var o36, o37 uint8
	var o38 uint32
	var o42 uint8
	var o43 uint32
	var o47, o48, o49 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o21)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o28)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o32)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o38)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o47)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o48)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o49)

	return Template5_3{
		ReferenceValue:                    o12,
		BinaryScaleFactor:                 common.ReadSignedNumericalValue(o16),
		DecimalScaleFactor:                common.ReadSignedNumericalValue(o18),
		BitsPerValue:                      common.ReadUnsignedNumericalValue(o20),
		TypeOfOriginalFieldValues:         common.ReadCodeValue(o21),
		GroupSplittingMethodUsed:          common.ReadCodeValue(o22),
		MissingValueManagementUsed:        common.ReadCodeValue(o23),
		PrimaryMissingValueSubstitute:     common.ReadUnsignedNumericalValue(o24),
		SecondaryMissingValueSubstitute:   common.ReadUnsignedNumericalValue(o28),
		NumberOfGroupsOfDataValues:        common.ReadUnsignedNumericalValue(o32),
		ReferenceForGroupWidth:            common.ReadUnsignedNumericalValue(o36),
		NumberOfBitsUsedForTheGroupWidths: common.ReadUnsignedNumericalValue(o37),
		ReferenceForGroupLength:           common.ReadUnsignedNumericalValue(o38),
		LengthIncrementForTheGroupLengths: common.ReadUnsignedNumericalValue(o42),
		TrueLengthOfLastGroup:             common.ReadUnsignedNumericalValue(o43),
		NumberOfBitsForScaledGroupLengths: common.ReadUnsignedNumericalValue(o47),
		OrderOfSpatialDifferencing:        common.ReadCodeValue(o48),
		NumberOfOctetsExtraDescriptors:    common.ReadUnsignedNumericalValue(o49),
	}
}
