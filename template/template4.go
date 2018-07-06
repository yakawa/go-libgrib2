package template

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/yakawa/go-libgrib2/common"
)

type Template4_0 struct {
	ParameterCategory               common.Grib2CodeValue
	ParameterNumber                 common.Grib2CodeValue
	TypeOfGeneratingProcess         common.Grib2CodeValue
	BackgroundProcess               common.Grib2NumericalValue
	GeneratingProcessIdentifier     common.Grib2NumericalValue
	HoursAfterDataCutoff            common.Grib2NumericalValue
	MinutesAfterDataCutoff          common.Grib2NumericalValue
	IndicatorOfUnitOfTimeRange      common.Grib2CodeValue
	ForecastTime                    common.Grib2NumericalValue
	TypeOfFirstFixedSurface         common.Grib2CodeValue
	ScaleFactorOfFirstFixedSurface  common.Grib2NumericalValue
	ScaledValueOfFirstFixedSurface  common.Grib2NumericalValue
	TypeOfSecondFixedSurface        common.Grib2CodeValue
	ScaleFactorOfSecondFixedSurface common.Grib2NumericalValue
	ScaledValueOfSecondFixedSurface common.Grib2NumericalValue
}

func ReadTemplate4_0(data []byte) Template4_0 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)

	return Template4_0{
		ParameterCategory:               common.ReadCodeValue(o10),
		ParameterNumber:                 common.ReadCodeValue(o11),
		TypeOfGeneratingProcess:         common.ReadCodeValue(o12),
		BackgroundProcess:               common.ReadUnsignedNumericalValue(o13),
		GeneratingProcessIdentifier:     common.ReadUnsignedNumericalValue(o14),
		HoursAfterDataCutoff:            common.ReadUnsignedNumericalValue(o15),
		MinutesAfterDataCutoff:          common.ReadUnsignedNumericalValue(o17),
		IndicatorOfUnitOfTimeRange:      common.ReadCodeValue(o18),
		ForecastTime:                    common.ReadSignedNumericalValue(o19),
		TypeOfFirstFixedSurface:         common.ReadCodeValue(o23),
		ScaleFactorOfFirstFixedSurface:  common.ReadSignedNumericalValue(o24),
		ScaledValueOfFirstFixedSurface:  common.ReadUnsignedNumericalValue(o25),
		TypeOfSecondFixedSurface:        common.ReadCodeValue(o29),
		ScaleFactorOfSecondFixedSurface: common.ReadSignedNumericalValue(o30),
		ScaledValueOfSecondFixedSurface: common.ReadUnsignedNumericalValue(o31),
	}

}

type Template4_1 struct {
	ParameterCategory               common.Grib2CodeValue
	ParameterNumber                 common.Grib2CodeValue
	TypeOfGeneratingProcess         common.Grib2CodeValue
	BackgroundProcess               common.Grib2NumericalValue
	GeneratingProcessIdentifier     common.Grib2NumericalValue
	HoursAfterDataCutoff            common.Grib2NumericalValue
	MinutesAfterDataCutoff          common.Grib2NumericalValue
	IndicatorOfUnitOfTimeRange      common.Grib2CodeValue
	ForecastTime                    common.Grib2NumericalValue
	TypeOfFirstFixedSurface         common.Grib2CodeValue
	ScaleFactorOfFirstFixedSurface  common.Grib2NumericalValue
	ScaledValueOfFirstFixedSurface  common.Grib2NumericalValue
	TypeOfSecondFixedSurface        common.Grib2CodeValue
	ScaleFactorOfSecondFixedSurface common.Grib2NumericalValue
	ScaledValueOfSecondFixedSurface common.Grib2NumericalValue
	TypeOfEnsembleForecast          common.Grib2CodeValue
	PerturbationNumber              common.Grib2NumericalValue
	NumberOfForecastsInEnsemble     common.Grib2NumericalValue
}

func ReadTemplate4_1(data []byte) Template4_1 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32
	var o35, o36, o37 uint8

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &o35)
	_ = binary.Read(bytes.NewBuffer(data[26:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[27:]), binary.BigEndian, &o37)

	return Template4_1{
		ParameterCategory:               common.ReadCodeValue(o10),
		ParameterNumber:                 common.ReadCodeValue(o11),
		TypeOfGeneratingProcess:         common.ReadCodeValue(o12),
		BackgroundProcess:               common.ReadUnsignedNumericalValue(o13),
		GeneratingProcessIdentifier:     common.ReadUnsignedNumericalValue(o14),
		HoursAfterDataCutoff:            common.ReadUnsignedNumericalValue(o15),
		MinutesAfterDataCutoff:          common.ReadUnsignedNumericalValue(o17),
		IndicatorOfUnitOfTimeRange:      common.ReadCodeValue(o18),
		ForecastTime:                    common.ReadSignedNumericalValue(o19),
		TypeOfFirstFixedSurface:         common.ReadCodeValue(o23),
		ScaleFactorOfFirstFixedSurface:  common.ReadSignedNumericalValue(o24),
		ScaledValueOfFirstFixedSurface:  common.ReadUnsignedNumericalValue(o25),
		TypeOfSecondFixedSurface:        common.ReadCodeValue(o29),
		ScaleFactorOfSecondFixedSurface: common.ReadSignedNumericalValue(o30),
		ScaledValueOfSecondFixedSurface: common.ReadUnsignedNumericalValue(o31),
		TypeOfEnsembleForecast:          common.ReadCodeValue(o35),
		PerturbationNumber:              common.ReadUnsignedNumericalValue(o36),
		NumberOfForecastsInEnsemble:     common.ReadUnsignedNumericalValue(o37),
	}

}

type Template4_8 struct {
	ParameterCategory                   common.Grib2CodeValue
	ParameterNumber                     common.Grib2CodeValue
	TypeOfGeneratingProcess             common.Grib2CodeValue
	BackgroundProcess                   common.Grib2NumericalValue
	GeneratingProcessIdentifier         common.Grib2NumericalValue
	HoursAfterDataCutoff                common.Grib2NumericalValue
	MinutesAfterDataCutoff              common.Grib2NumericalValue
	IndicatorOfUnitOfTimeRange          common.Grib2CodeValue
	ForecastTime                        common.Grib2NumericalValue
	TypeOfFirstFixedSurface             common.Grib2CodeValue
	ScaleFactorOfFirstFixedSurface      common.Grib2NumericalValue
	ScaledValueOfFirstFixedSurface      common.Grib2NumericalValue
	TypeOfSecondFixedSurface            common.Grib2CodeValue
	ScaleFactorOfSecondFixedSurface     common.Grib2NumericalValue
	ScaledValueOfSecondFixedSurface     common.Grib2NumericalValue
	EndOfOverallTimeInterval            time.Time
	NumberOfTimeRange                   common.Grib2NumericalValue
	NumberOfMissingInStatisticalProcess common.Grib2NumericalValue
	TypeOfStatisticalProcessing         common.Grib2CodeValue
	TypeOfTimeIncrement                 common.Grib2CodeValue
	IndicatorOfUnitForTimeRange         common.Grib2CodeValue
	LengthOfTimeRange                   common.Grib2NumericalValue
	IndicatorOfUnitForTimeIncrement     common.Grib2CodeValue
	TimeIncrement                       common.Grib2NumericalValue
}

func ReadTemplate4_8(data []byte) Template4_8 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32
	var year uint16
	var month, day, hour, minutes, seconds uint8
	var o42 uint8
	var o43 uint32
	var o47, o48, o49 uint8
	var o50 uint32
	var o54 uint8
	var o55 uint32

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &year)
	_ = binary.Read(bytes.NewBuffer(data[27:]), binary.BigEndian, &month)
	_ = binary.Read(bytes.NewBuffer(data[28:]), binary.BigEndian, &day)
	_ = binary.Read(bytes.NewBuffer(data[29:]), binary.BigEndian, &hour)
	_ = binary.Read(bytes.NewBuffer(data[30:]), binary.BigEndian, &seconds)
	_ = binary.Read(bytes.NewBuffer(data[31:]), binary.BigEndian, &minutes)
	_ = binary.Read(bytes.NewBuffer(data[32:]), binary.BigEndian, &o42)
	_ = binary.Read(bytes.NewBuffer(data[33:]), binary.BigEndian, &o43)
	_ = binary.Read(bytes.NewBuffer(data[37:]), binary.BigEndian, &o47)
	_ = binary.Read(bytes.NewBuffer(data[38:]), binary.BigEndian, &o48)
	_ = binary.Read(bytes.NewBuffer(data[39:]), binary.BigEndian, &o49)
	_ = binary.Read(bytes.NewBuffer(data[40:]), binary.BigEndian, &o50)
	_ = binary.Read(bytes.NewBuffer(data[44:]), binary.BigEndian, &o54)
	_ = binary.Read(bytes.NewBuffer(data[45:]), binary.BigEndian, &o55)

	loc, _ := time.LoadLocation("UTC")

	return Template4_8{
		ParameterCategory:                   common.ReadCodeValue(o10),
		ParameterNumber:                     common.ReadCodeValue(o11),
		TypeOfGeneratingProcess:             common.ReadCodeValue(o12),
		BackgroundProcess:                   common.ReadUnsignedNumericalValue(o13),
		GeneratingProcessIdentifier:         common.ReadUnsignedNumericalValue(o14),
		HoursAfterDataCutoff:                common.ReadUnsignedNumericalValue(o15),
		MinutesAfterDataCutoff:              common.ReadUnsignedNumericalValue(o17),
		IndicatorOfUnitOfTimeRange:          common.ReadCodeValue(o18),
		ForecastTime:                        common.ReadSignedNumericalValue(o19),
		TypeOfFirstFixedSurface:             common.ReadCodeValue(o23),
		ScaleFactorOfFirstFixedSurface:      common.ReadSignedNumericalValue(o24),
		ScaledValueOfFirstFixedSurface:      common.ReadUnsignedNumericalValue(o25),
		TypeOfSecondFixedSurface:            common.ReadCodeValue(o29),
		ScaleFactorOfSecondFixedSurface:     common.ReadSignedNumericalValue(o30),
		ScaledValueOfSecondFixedSurface:     common.ReadUnsignedNumericalValue(o31),
		EndOfOverallTimeInterval:            time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minutes), int(seconds), 0, loc),
		NumberOfTimeRange:                   common.ReadUnsignedNumericalValue(o42),
		NumberOfMissingInStatisticalProcess: common.ReadUnsignedNumericalValue(o43),
		TypeOfStatisticalProcessing:         common.ReadCodeValue(o47),
		TypeOfTimeIncrement:                 common.ReadCodeValue(o48),
		IndicatorOfUnitForTimeRange:         common.ReadCodeValue(o49),
		LengthOfTimeRange:                   common.ReadUnsignedNumericalValue(o50),
		IndicatorOfUnitForTimeIncrement:     common.ReadCodeValue(o54),
		TimeIncrement:                       common.ReadUnsignedNumericalValue(o55),
	}

}

type Template4_11 struct {
	ParameterCategory                   common.Grib2CodeValue
	ParameterNumber                     common.Grib2CodeValue
	TypeOfGeneratingProcess             common.Grib2CodeValue
	BackgroundProcess                   common.Grib2NumericalValue
	GeneratingProcessIdentifier         common.Grib2NumericalValue
	HoursAfterDataCutoff                common.Grib2NumericalValue
	MinutesAfterDataCutoff              common.Grib2NumericalValue
	IndicatorOfUnitOfTimeRange          common.Grib2CodeValue
	ForecastTime                        common.Grib2NumericalValue
	TypeOfFirstFixedSurface             common.Grib2CodeValue
	ScaleFactorOfFirstFixedSurface      common.Grib2NumericalValue
	ScaledValueOfFirstFixedSurface      common.Grib2NumericalValue
	TypeOfSecondFixedSurface            common.Grib2CodeValue
	ScaleFactorOfSecondFixedSurface     common.Grib2NumericalValue
	ScaledValueOfSecondFixedSurface     common.Grib2NumericalValue
	TypeOfEnsembleForecast              common.Grib2CodeValue
	PerturbationNumber                  common.Grib2NumericalValue
	NumberOfForecastsInEnsemble         common.Grib2NumericalValue
	EndOfOverallTimeInterval            time.Time
	NumberOfTimeRange                   common.Grib2NumericalValue
	NumberOfMissingInStatisticalProcess common.Grib2NumericalValue
	TypeOfStatisticalProcessing         common.Grib2CodeValue
	TypeOfTimeIncrement                 common.Grib2CodeValue
	IndicatorOfUnitForTimeRange         common.Grib2CodeValue
	LengthOfTimeRange                   common.Grib2NumericalValue
	IndicatorOfUnitForTimeIncrement     common.Grib2CodeValue
	TimeIncrement                       common.Grib2NumericalValue
}

func ReadTemplate4_11(data []byte) Template4_11 {
	var o10, o11, o12, o13, o14 uint8
	var o15 uint16
	var o17 uint8
	var o18 uint8
	var o19 uint32
	var o23, o24 uint8
	var o25 uint32
	var o29, o30 uint8
	var o31 uint32
	var o35, o36, o37 uint8
	var year uint16
	var month, day, hour, minutes, seconds uint8
	var o45 uint8
	var o46 uint32
	var o50, o51, o52 uint8
	var o53 uint32
	var o57 uint8
	var o58 uint32

	_ = binary.Read(bytes.NewBuffer(data[0:]), binary.BigEndian, &o10)
	_ = binary.Read(bytes.NewBuffer(data[1:]), binary.BigEndian, &o11)
	_ = binary.Read(bytes.NewBuffer(data[2:]), binary.BigEndian, &o12)
	_ = binary.Read(bytes.NewBuffer(data[3:]), binary.BigEndian, &o13)
	_ = binary.Read(bytes.NewBuffer(data[4:]), binary.BigEndian, &o14)
	_ = binary.Read(bytes.NewBuffer(data[5:]), binary.BigEndian, &o15)
	_ = binary.Read(bytes.NewBuffer(data[7:]), binary.BigEndian, &o17)
	_ = binary.Read(bytes.NewBuffer(data[8:]), binary.BigEndian, &o18)
	_ = binary.Read(bytes.NewBuffer(data[9:]), binary.BigEndian, &o19)
	_ = binary.Read(bytes.NewBuffer(data[13:]), binary.BigEndian, &o23)
	_ = binary.Read(bytes.NewBuffer(data[14:]), binary.BigEndian, &o24)
	_ = binary.Read(bytes.NewBuffer(data[15:]), binary.BigEndian, &o25)
	_ = binary.Read(bytes.NewBuffer(data[19:]), binary.BigEndian, &o29)
	_ = binary.Read(bytes.NewBuffer(data[20:]), binary.BigEndian, &o30)
	_ = binary.Read(bytes.NewBuffer(data[21:]), binary.BigEndian, &o31)
	_ = binary.Read(bytes.NewBuffer(data[25:]), binary.BigEndian, &o35)
	_ = binary.Read(bytes.NewBuffer(data[26:]), binary.BigEndian, &o36)
	_ = binary.Read(bytes.NewBuffer(data[27:]), binary.BigEndian, &o37)
	_ = binary.Read(bytes.NewBuffer(data[28:]), binary.BigEndian, &year)
	_ = binary.Read(bytes.NewBuffer(data[30:]), binary.BigEndian, &month)
	_ = binary.Read(bytes.NewBuffer(data[31:]), binary.BigEndian, &day)
	_ = binary.Read(bytes.NewBuffer(data[32:]), binary.BigEndian, &hour)
	_ = binary.Read(bytes.NewBuffer(data[33:]), binary.BigEndian, &seconds)
	_ = binary.Read(bytes.NewBuffer(data[34:]), binary.BigEndian, &minutes)
	_ = binary.Read(bytes.NewBuffer(data[35:]), binary.BigEndian, &o45)
	_ = binary.Read(bytes.NewBuffer(data[36:]), binary.BigEndian, &o46)
	_ = binary.Read(bytes.NewBuffer(data[40:]), binary.BigEndian, &o50)
	_ = binary.Read(bytes.NewBuffer(data[41:]), binary.BigEndian, &o51)
	_ = binary.Read(bytes.NewBuffer(data[42:]), binary.BigEndian, &o52)
	_ = binary.Read(bytes.NewBuffer(data[43:]), binary.BigEndian, &o53)
	_ = binary.Read(bytes.NewBuffer(data[47:]), binary.BigEndian, &o57)
	_ = binary.Read(bytes.NewBuffer(data[48:]), binary.BigEndian, &o58)

	loc, _ := time.LoadLocation("UTC")

	return Template4_11{
		ParameterCategory:                   common.ReadCodeValue(o10),
		ParameterNumber:                     common.ReadCodeValue(o11),
		TypeOfGeneratingProcess:             common.ReadCodeValue(o12),
		BackgroundProcess:                   common.ReadUnsignedNumericalValue(o13),
		GeneratingProcessIdentifier:         common.ReadUnsignedNumericalValue(o14),
		HoursAfterDataCutoff:                common.ReadUnsignedNumericalValue(o15),
		MinutesAfterDataCutoff:              common.ReadUnsignedNumericalValue(o17),
		IndicatorOfUnitOfTimeRange:          common.ReadCodeValue(o18),
		ForecastTime:                        common.ReadSignedNumericalValue(o19),
		TypeOfFirstFixedSurface:             common.ReadCodeValue(o23),
		ScaleFactorOfFirstFixedSurface:      common.ReadSignedNumericalValue(o24),
		ScaledValueOfFirstFixedSurface:      common.ReadUnsignedNumericalValue(o25),
		TypeOfSecondFixedSurface:            common.ReadCodeValue(o29),
		ScaleFactorOfSecondFixedSurface:     common.ReadSignedNumericalValue(o30),
		ScaledValueOfSecondFixedSurface:     common.ReadUnsignedNumericalValue(o31),
		TypeOfEnsembleForecast:              common.ReadCodeValue(o35),
		PerturbationNumber:                  common.ReadUnsignedNumericalValue(o36),
		NumberOfForecastsInEnsemble:         common.ReadUnsignedNumericalValue(o37),
		EndOfOverallTimeInterval:            time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minutes), int(seconds), 0, loc),
		NumberOfTimeRange:                   common.ReadUnsignedNumericalValue(o45),
		NumberOfMissingInStatisticalProcess: common.ReadUnsignedNumericalValue(o46),
		TypeOfStatisticalProcessing:         common.ReadCodeValue(o50),
		TypeOfTimeIncrement:                 common.ReadCodeValue(o51),
		IndicatorOfUnitForTimeRange:         common.ReadCodeValue(o52),
		LengthOfTimeRange:                   common.ReadUnsignedNumericalValue(o53),
		IndicatorOfUnitForTimeIncrement:     common.ReadCodeValue(o57),
		TimeIncrement:                       common.ReadUnsignedNumericalValue(o58),
	}

}
