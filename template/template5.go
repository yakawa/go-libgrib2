package template

/*
 Template for Section5
*/
import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

// Template5_0 : Grid point data - simple packing Note: For most templates, details of the packing process are described in regulation 92.9.4
type Template5_0 struct {
	ReferenceValue            float32                    // Reference value (R) (IEEE 32-bit floating-point value)
	BinaryScaleFactor         common.Grib2NumericalValue // Binary scale factor (E)
	DecimalScaleFactor        common.Grib2NumericalValue // Decimal scale factor (D)
	BitsPerValue              common.Grib2NumericalValue // Number of bits used for each packed value for simple packing, or for each group reference value for complex packing or spatial differencing
	TypeOfOriginalFieldValues common.Grib2CodeValue      // Type of original field values (see Code Table 5.1)
}

// ReadTemplate5_0 : Read Template 5.0
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

// Template5_2 : Grid point data - complex packing Note: For most templates, details of the packing process are described in regulation 92.9.4
type Template5_2 struct {
	ReferenceValue                    float32                    // Reference value (R) (IEEE 32-bit floating-point value)
	BinaryScaleFactor                 common.Grib2NumericalValue // Binary scale factor (E)
	DecimalScaleFactor                common.Grib2NumericalValue // Decimal scale factor (D)
	BitsPerValue                      common.Grib2NumericalValue // Number of bits used for each packed value for simple packing, or for each group reference value for complex packing or spatial differencing
	TypeOfOriginalFieldValues         common.Grib2CodeValue      // Type of original field values (see Code Table 5.1)
	GroupSplittingMethodUsed          common.Grib2CodeValue      // Group splitting method used (see Code Table 5.4)
	MissingValueManagementUsed        common.Grib2CodeValue      // Missing value management used (see Code Table 5.5)
	PrimaryMissingValueSubstitute     common.Grib2NumericalValue // Primary missing value substitute
	SecondaryMissingValueSubstitute   common.Grib2NumericalValue // Secondary missing value substitute
	NumberOfGroupsOfDataValues        common.Grib2NumericalValue // NG - Number of groups of data values into which field is split
	ReferenceForGroupWidth            common.Grib2NumericalValue // Reference for group widths (see Note 12)
	NumberOfBitsUsedForTheGroupWidths common.Grib2NumericalValue // Number of bits used for the group widths (after the reference value in octet 36 has been removed)
	ReferenceForGroupLength           common.Grib2NumericalValue // Reference for group lengths (see Note 13)
	LengthIncrementForTheGroupLengths common.Grib2NumericalValue // Length increment for the group lengths (see Note 14)
	TrueLengthOfLastGroup             common.Grib2NumericalValue // True length of last group
	NumberOfBitsForScaledGroupLengths common.Grib2NumericalValue // Number of bits used for the scaled group lengths (after subtraction of the reference value given in octets 38-41 and division by the length increment given in octet 42)
}

// ReadTemplate5_2 : Read Template 5.2
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
	_ = binary.Read(bytes.NewBuffer(data[10:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[11:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[12:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[16:]), binary.BigEndian, &o28)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o32)
	_ = binary.Read(bytes.NewBuffer(data[24:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[26:]), binary.BigEndian, &o38)
	_ = binary.Read(bytes.NewBuffer(data[30:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[31:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[35:]), binary.BigEndian, &o47)

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

// Template5_3 : Grid point data - complex packing and spatial differencing Note: For most templates, details of the packing process are described in regulation 92.9.4
type Template5_3 struct {
	ReferenceValue                    float32                    // Reference value (R) (IEEE 32-bit floating-point value)
	BinaryScaleFactor                 common.Grib2NumericalValue // Binary scale factor (E)
	DecimalScaleFactor                common.Grib2NumericalValue // Decimal scale factor (D)
	BitsPerValue                      common.Grib2NumericalValue // Number of bits used for each packed value for simple packing, or for each group reference value for complex packing or spatial differencing
	TypeOfOriginalFieldValues         common.Grib2CodeValue      // Type of original field values (see Code Table 5.1)
	GroupSplittingMethodUsed          common.Grib2CodeValue      // Group splitting method used (see Code Table 5.4)
	MissingValueManagementUsed        common.Grib2CodeValue      // Missing value management used (see Code Table 5.5)
	PrimaryMissingValueSubstitute     common.Grib2NumericalValue // Primary missing value substitute
	SecondaryMissingValueSubstitute   common.Grib2NumericalValue // Secondary missing value substitute
	NumberOfGroupsOfDataValues        common.Grib2NumericalValue // NG - Number of groups of data values into which field is split
	ReferenceForGroupWidth            common.Grib2NumericalValue // Reference for group widths (see Note 12)
	NumberOfBitsUsedForTheGroupWidths common.Grib2NumericalValue // Number of bits used for the group widths (after the reference value in octet 36 has been removed)
	ReferenceForGroupLength           common.Grib2NumericalValue // Reference for group lengths (see Note 13)
	LengthIncrementForTheGroupLengths common.Grib2NumericalValue // Length increment for the group lengths (see Note 14)
	TrueLengthOfLastGroup             common.Grib2NumericalValue // True length of last group
	NumberOfBitsForScaledGroupLengths common.Grib2NumericalValue // Number of bits used for the scaled group lengths (after subtraction of the reference value given in octets 38-41 and division by the length increment given in octet 42)
	OrderOfSpatialDifferencing        common.Grib2CodeValue      // Order of spatial differencing (see Code Table 5.6)
	NumberOfOctetsExtraDescriptors    common.Grib2NumericalValue // Number of octets required in the Data Section to specify extra descriptors needed for spatial differencing (octets 6-ww in Data Template 7.3)
}

// ReadTemplate5_3 : Read Template 5.3
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
	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o20)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o21)
	_ = binary.Read(bytes.NewBuffer(data[10:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[11:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[12:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[16:]), binary.BigEndian, &o28)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o32)
	_ = binary.Read(bytes.NewBuffer(data[24:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[26:]), binary.BigEndian, &o38)
	_ = binary.Read(bytes.NewBuffer(data[30:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[31:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[35:]), binary.BigEndian, &o47)
	_ = binary.Read(bytes.NewBuffer(data[36:]), binary.BigEndian, &o48)
	_ = binary.Read(bytes.NewBuffer(data[37:]), binary.BigEndian, &o49)

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
