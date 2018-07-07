package template

/*
 Template for Section4
*/
import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/yakawa/go-libgrib2/common"
)

// Template4_0 : Analysis or forecast at a horizontal level or in a horizontal layer at a point in time
type Template4_0 struct {
	ParameterCategory               common.Grib2CodeValue      // Parameter category (see Code Table 4.1).
	ParameterNumber                 common.Grib2CodeValue      // Parameter number (see Code Table 4.2).
	TypeOfGeneratingProcess         common.Grib2CodeValue      // Type of generating process (see Code Table 4.3)
	BackgroundProcess               common.Grib2NumericalValue // Background generating process identifier (defined by originating centre)
	GeneratingProcessIdentifier     common.Grib2NumericalValue // Analysis or forecast generating processes identifier (defined by originating centre)
	HoursAfterDataCutoff            common.Grib2NumericalValue // Hours of observational data cutoff after reference time (see Note 1)
	MinutesAfterDataCutoff          common.Grib2NumericalValue // Minutes of observational data cutoff after reference time
	IndicatorOfUnitOfTimeRange      common.Grib2CodeValue      // Indicator of unit of time range (see Code Table 4.4)
	ForecastTime                    common.Grib2NumericalValue // Forecast time in units defined by octet 18
	TypeOfFirstFixedSurface         common.Grib2CodeValue      // Type of first fixed surface (see Code Table 4.5)
	ScaleFactorOfFirstFixedSurface  common.Grib2NumericalValue // Scale factor of first fixed surface
	ScaledValueOfFirstFixedSurface  common.Grib2NumericalValue // Scaled value of first fixed surface
	TypeOfSecondFixedSurface        common.Grib2CodeValue      // Type of second fixed surface (see Code Table 4.5)
	ScaleFactorOfSecondFixedSurface common.Grib2NumericalValue // Scale factor of second fixed surface
	ScaledValueOfSecondFixedSurface common.Grib2NumericalValue // Scaled value of second fixed surface
}

// ReadTemplate4_0 : Read Template 4.0
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

// Template4_1 : Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer at a point in time
type Template4_1 struct {
	ParameterCategory               common.Grib2CodeValue      // Parameter category (see Code Table 4.1)
	ParameterNumber                 common.Grib2CodeValue      // Parameter number (see Code Table 4.2)
	TypeOfGeneratingProcess         common.Grib2CodeValue      // Type of generating process (see Code Table 4.3)
	BackgroundProcess               common.Grib2NumericalValue // Background generating process identifier (defined by originating Centre)
	GeneratingProcessIdentifier     common.Grib2NumericalValue // Forecast generating process identifier (defined by originating Centre)
	HoursAfterDataCutoff            common.Grib2NumericalValue // Hours after reference time of data cutoff (see Note 1)
	MinutesAfterDataCutoff          common.Grib2NumericalValue // Minutes after reference time of data cutoff
	IndicatorOfUnitOfTimeRange      common.Grib2CodeValue      // Indicator of unit of time range (see Code Table 4.4)
	ForecastTime                    common.Grib2NumericalValue // Forecast time in units defined by octet 18
	TypeOfFirstFixedSurface         common.Grib2CodeValue      // Type of first fixed surface (see Code Table 4.5)
	ScaleFactorOfFirstFixedSurface  common.Grib2NumericalValue // Scale factor of first fixed surface
	ScaledValueOfFirstFixedSurface  common.Grib2NumericalValue // Scaled value of first fixed surface
	TypeOfSecondFixedSurface        common.Grib2CodeValue      // Type of second fixed surface (see Code Table 4.5)
	ScaleFactorOfSecondFixedSurface common.Grib2NumericalValue // Scale factor of second fixed surface
	ScaledValueOfSecondFixedSurface common.Grib2NumericalValue // Scaled value of second fixed surface
	TypeOfEnsembleForecast          common.Grib2CodeValue      // Type of ensemble forecast (see Code Table 4.6)
	PerturbationNumber              common.Grib2NumericalValue // Perturbation number
	NumberOfForecastsInEnsemble     common.Grib2NumericalValue // Number of forecasts in ensemble
}

// ReadTemplate4_1 : Read Template 4.1
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

// Template4_8 : Average, accumulation, and/or extreme values or other statistically processed values at a horizontal level or in a horizontal layer in a continuous or non- continuous time interval
type Template4_8 struct {
	ParameterCategory                   common.Grib2CodeValue      // Parameter category (see Code Table 4.1)
	ParameterNumber                     common.Grib2CodeValue      // Parameter number (see Code Table 4.2)
	TypeOfGeneratingProcess             common.Grib2CodeValue      // Type of generating process (see Code Table 4.3)
	BackgroundProcess                   common.Grib2NumericalValue // Background generating process identifier (defined by originating Centre)
	GeneratingProcessIdentifier         common.Grib2NumericalValue // Analysis or Forecast generating process identifier (defined by originating Centre)
	HoursAfterDataCutoff                common.Grib2NumericalValue // Hours after reference time of data cut-off (see Note 1)
	MinutesAfterDataCutoff              common.Grib2NumericalValue // Minutes after reference time of data cut-off
	IndicatorOfUnitOfTimeRange          common.Grib2CodeValue      // Indicator of unit of time range (see Code Table 4.4)
	ForecastTime                        common.Grib2NumericalValue // Forecast time in units defined by octet 18 (see Note 2)
	TypeOfFirstFixedSurface             common.Grib2CodeValue      // Type of first fixed surface (see Code Table 4.5)
	ScaleFactorOfFirstFixedSurface      common.Grib2NumericalValue // Scale factor of first fixed surface
	ScaledValueOfFirstFixedSurface      common.Grib2NumericalValue // Scaled value of first fixed surface
	TypeOfSecondFixedSurface            common.Grib2CodeValue      // Type of second fixed surface (see Code Table 4.5)
	ScaleFactorOfSecondFixedSurface     common.Grib2NumericalValue // Scale factor of second fixed surface
	ScaledValueOfSecondFixedSurface     common.Grib2NumericalValue // Scaled value of second fixed surface
	EndOfOverallTimeInterval            time.Time                  // Time of end of overall time interval
	NumberOfTimeRange                   common.Grib2NumericalValue // n - Number of time range specifications describing the time intervals used to calculate the statistically processed field
	NumberOfMissingInStatisticalProcess common.Grib2NumericalValue // Total number of data values missing in statistical process.
	TypeOfStatisticalProcessing         common.Grib2CodeValue      // Statistical process used to calculate the processed field from the field at each time increment during the time range (see Code Table 4.10)
	TypeOfTimeIncrement                 common.Grib2CodeValue      // Type of time increment between successive fields used in the statistical processing (see Code Table 4.11)
	IndicatorOfUnitForTimeRange         common.Grib2CodeValue      // Indicator of unit of time for time range over which statistical processing is done (see Code Table 4.4)
	LengthOfTimeRange                   common.Grib2NumericalValue // Length of the time range over which statistical processing is done, in units defined by the previous octet
	IndicatorOfUnitForTimeIncrement     common.Grib2CodeValue      // Indicator of unit of time for the increment between the successive fields used (see Code Table 4.4)
	TimeIncrement                       common.Grib2NumericalValue // Time increment between successive fields, in units defined by the previous octet (see Notes 3 and 4)
}

// ReadTemplate4_8 : Read Template 4.8
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

// Template4_11 : Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer in a continuous or non-continuous time interval
type Template4_11 struct {
	ParameterCategory                   common.Grib2CodeValue      // Parameter category (see Code table 4.1)
	ParameterNumber                     common.Grib2CodeValue      // Parameter number (see Code table 4.2)
	TypeOfGeneratingProcess             common.Grib2CodeValue      // Type of generating process (see Code table 4.3)
	BackgroundProcess                   common.Grib2NumericalValue // Background generating process identifier (defined by originating centre)
	GeneratingProcessIdentifier         common.Grib2NumericalValue // Forecast generating process identifier (defined by originating centre)
	HoursAfterDataCutoff                common.Grib2NumericalValue // Hours after reference time of data cut-off (see Note 1)
	MinutesAfterDataCutoff              common.Grib2NumericalValue // Minutes after reference time of data cut-off
	IndicatorOfUnitOfTimeRange          common.Grib2CodeValue      // Indicator of unit of time range (see Code table 4.4)
	ForecastTime                        common.Grib2NumericalValue // Forecast time in units defined by octet 18 (see Note 2)
	TypeOfFirstFixedSurface             common.Grib2CodeValue      // Type of first fixed surface (see Code table 4.5)
	ScaleFactorOfFirstFixedSurface      common.Grib2NumericalValue // Scale factor of first fixed surface
	ScaledValueOfFirstFixedSurface      common.Grib2NumericalValue // Scaled value of first fixed surface
	TypeOfSecondFixedSurface            common.Grib2CodeValue      // Type of second fixed surface (see Code table 4.5)
	ScaleFactorOfSecondFixedSurface     common.Grib2NumericalValue // Scale factor of second fixed surface
	ScaledValueOfSecondFixedSurface     common.Grib2NumericalValue // Scaled value of second fixed surface
	TypeOfEnsembleForecast              common.Grib2CodeValue      // Type of ensemble forecast (see Code table 4.6)
	PerturbationNumber                  common.Grib2NumericalValue // Perturbation number
	NumberOfForecastsInEnsemble         common.Grib2NumericalValue // Number of forecasts in ensemble
	EndOfOverallTimeInterval            time.Time                  // end of overall time interval
	NumberOfTimeRange                   common.Grib2NumericalValue // n - number of time range specifications describing the time intervals used to calculate the statistically- processed field
	NumberOfMissingInStatisticalProcess common.Grib2NumericalValue // Total number of data values missing in statistical process
	TypeOfStatisticalProcessing         common.Grib2CodeValue      // Statistical process used to calculate the processed field from the field at each time increment during the time range (see Code table 4.10)
	TypeOfTimeIncrement                 common.Grib2CodeValue      // Type of time increment between successive fields used in the statistical processing (see Code table 4.11)
	IndicatorOfUnitForTimeRange         common.Grib2CodeValue      // Indicator of unit of time for time range over which statistical processing is done (see Code table 4.4)
	LengthOfTimeRange                   common.Grib2NumericalValue // Length of the time range over which statistical processing is done, in units defined by the previous octet
	IndicatorOfUnitForTimeIncrement     common.Grib2CodeValue      // Indicator of unit of time for the increment between the successive fields used (see Code table 4.4)
	TimeIncrement                       common.Grib2NumericalValue // Time increment between successive fields, in units defined by the previous octet (see Note 3)
}

// ReadTemplate4_11 : Read Template 4.11
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
