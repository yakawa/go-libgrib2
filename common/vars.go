package common

type Grib2NumericalValue struct {
	Val     int
	Missing bool
}

type Grib2FlagValue struct {
	Val     uint
	Missing bool
}

type Grib2FloatValue struct {
	Val     float32
	Missing bool
}

type Grib2IntegerValue struct {
	Val     int
	Missing bool
}
