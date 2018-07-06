package template

import (
	"bytes"
	"encoding/binary"

	"github.com/yakawa/go-libgrib2/common"
)

type Template3_0 struct {
	ShapeOfTheEarth                        common.Grib2CodeValue
	ScaleFactorOfRadiusOfSphericalEarth    common.Grib2NumericalValue
	ScaledValuedOfRadiusOfSphericalEarth   common.Grib2NumericalValue
	ScaleFactorOfEarthMajorAxis            common.Grib2NumericalValue
	ScaledValueOfEarthMajorAxis            common.Grib2NumericalValue
	ScaleFactorOfEarthMinorAxis            common.Grib2NumericalValue
	ScaledValueOfEarthMinorAxis            common.Grib2NumericalValue
	Ni                                     common.Grib2NumericalValue
	Nj                                     common.Grib2NumericalValue
	BasicAngleOfTheInitialProductionDomain common.Grib2NumericalValue
	SubdivisionsOfBasicAngle               common.Grib2NumericalValue
	LatitudeOfFirstGridPoint               common.Grib2NumericalValue
	LongitudeOfFirstGridPoint              common.Grib2NumericalValue
	ResolutionAndComponentFlags            common.Grib2FlagValue
	LatitudeOfLastGridPoint                common.Grib2NumericalValue
	LongitudeOfLastGridPoint               common.Grib2NumericalValue
	IDirectionIncrement                    common.Grib2NumericalValue
	JDirectionIncrement                    common.Grib2NumericalValue
	ScanningMode                           common.Grib2FlagValue
}

func ReadTemplate3_0(data []byte) Template3_0 {
	var o15, o16 uint8
	var o17 uint32
	var o21 uint8
	var o22 uint32
	var o26 uint8
	var o27 uint32
	var o31, o35, o39, o43, o47, o51 uint32
	var o55 uint8
	var o56, o60, o64, o68 uint32
	var o72 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o16)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[6:]), binary.BigEndian, &o21)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o22)
	_ = binary.Read(bytes.NewBuffer(data[11:]), binary.BigEndian, &o26)
	_ = binary.Read(bytes.NewBuffer(data[12:]), binary.BigEndian, &o27)
	_ = binary.Read(bytes.NewBuffer(data[16:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o35)
	_ = binary.Read(bytes.NewBuffer(data[24:]), binary.BigEndian, &o39)
	_ = binary.Read(bytes.NewBuffer(data[28:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[32:]), binary.BigEndian, &o47)
	_ = binary.Read(bytes.NewBuffer(data[36:]), binary.BigEndian, &o51)
	_ = binary.Read(bytes.NewBuffer(data[40:]), binary.BigEndian, &o55)
	_ = binary.Read(bytes.NewBuffer(data[41:]), binary.BigEndian, &o56)
	_ = binary.Read(bytes.NewBuffer(data[45:]), binary.BigEndian, &o60)
	_ = binary.Read(bytes.NewBuffer(data[49:]), binary.BigEndian, &o64)
	_ = binary.Read(bytes.NewBuffer(data[53:]), binary.BigEndian, &o68)
	_ = binary.Read(bytes.NewBuffer(data[57:]), binary.BigEndian, &o72)

	return Template3_0{
		ShapeOfTheEarth:                      common.ReadCodeValue(o15),
		ScaleFactorOfRadiusOfSphericalEarth:  common.ReadUnsignedNumericalValue(o16),
		ScaledValuedOfRadiusOfSphericalEarth: common.ReadUnsignedNumericalValue(o17),
		ScaleFactorOfEarthMajorAxis:          common.ReadUnsignedNumericalValue(o21),
		ScaledValueOfEarthMajorAxis:          common.ReadUnsignedNumericalValue(o22),
		ScaleFactorOfEarthMinorAxis:          common.ReadUnsignedNumericalValue(o26),
		ScaledValueOfEarthMinorAxis:          common.ReadUnsignedNumericalValue(o27),
		Ni: common.ReadUnsignedNumericalValue(o31),
		Nj: common.ReadUnsignedNumericalValue(o35),
		BasicAngleOfTheInitialProductionDomain: common.ReadUnsignedNumericalValue(o39),
		SubdivisionsOfBasicAngle:               common.ReadUnsignedNumericalValue(o43),
		LatitudeOfFirstGridPoint:               common.ReadSignedNumericalValue(o47),
		LongitudeOfFirstGridPoint:              common.ReadSignedNumericalValue(o51),
		ResolutionAndComponentFlags:            common.ReadFlagValue(o55),
		LatitudeOfLastGridPoint:                common.ReadSignedNumericalValue(o56),
		LongitudeOfLastGridPoint:               common.ReadSignedNumericalValue(o60),
		IDirectionIncrement:                    common.ReadUnsignedNumericalValue(o64),
		JDirectionIncrement:                    common.ReadUnsignedNumericalValue(o68),
		ScanningMode:                           common.ReadFlagValue(o72),
	}

}
