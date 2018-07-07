package common

// Grib2NumericalValue : Numerical Value with Missing Flag
type Grib2NumericalValue struct {
	Val     int
	Missing bool
}

// Grib2CodeValue : Code Value with Missing
type Grib2CodeValue struct {
	Val     uint
	Missing bool
}

// Grib2FlagValue : Flag Value with Missing
type Grib2FlagValue struct {
	Val     uint
	Missing bool
}

// Grib2FloatValue : Float Value with Missing
type Grib2FloatValue struct {
	Val     float32
	Missing bool
}

// Grib2IntegerValue : Integer Value with Missing
type Grib2IntegerValue struct {
	Val     int
	Missing bool
}
