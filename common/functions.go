package common

import (
	"log"
)

func ReadNumericalValue(v interface{}) Grib2NumericalValue {
	switch v.(type) {
	case int:
		return Grib2NumericalValue{Val: v.(int), Missing: false}
	case uint8:
		if v.(uint8)&0xFF == 0xFF {
			return Grib2NumericalValue{Val: -1, Missing: true}
		} else if v.(uint8)&0x80 == 0 {
			return Grib2NumericalValue{Val: int(v.(uint8)), Missing: false}
		} else {
			return Grib2NumericalValue{Val: -1 * int(v.(uint8)&0x7f), Missing: false}
		}
	case uint16:
		if v.(uint16)&0xFFFF == 0xFFFF {
			return Grib2NumericalValue{Val: -1, Missing: true}
		} else if v.(uint16)&0x8000 == 0 {
			return Grib2NumericalValue{Val: int(v.(uint16)), Missing: false}
		} else {
			return Grib2NumericalValue{Val: -1 * int(v.(uint16)&0x7fff), Missing: false}
		}
	case uint32:
		if v.(uint32)&0xFFFFFFFF == 0xFFFFFFFF {
			return Grib2NumericalValue{Val: -1, Missing: true}
		} else if v.(uint32)&0x80000000 == 0 {
			return Grib2NumericalValue{Val: int(v.(uint32)), Missing: false}
		} else {
			return Grib2NumericalValue{-1 * int(v.(uint32)&0x7fffffff), false}
		}
	case uint64:
		if v.(uint64)&0xFFFFFFFFFFFFFFFF == 0xFFFFFFFFFFFFFFFF {
			return Grib2NumericalValue{Val: -1, Missing: true}
		} else if v.(uint64)&0x8000000000000000 == 0 {
			return Grib2NumericalValue{Val: int(v.(uint64)), Missing: false}
		} else {
			return Grib2NumericalValue{Val: -1 * int(v.(uint64)&0x7FFFFFFFFFFFFFFF), Missing: false}
		}
	default:
		log.Fatalf("Unknown Type: %s", v)
		return Grib2NumericalValue{}
	}
}

func ReadFlagValue(v interface{}, sec uint, number uint) Grib2FlagValue {
	switch v.(type) {
	case uint8:
		if v.(uint8)&0xFF == 0xFF {
			return Grib2FlagValue{Val: uint(v.(uint8)), Missing: true}
		}
		return Grib2FlagValue{Val: uint(v.(uint8)), Missing: false}
	case uint16:
		if v.(uint16)&0xFFFF == 0xFFFF {
			return Grib2FlagValue{Val: uint(v.(uint16)), Missing: true}
		}
		return Grib2FlagValue{Val: uint(v.(uint16)), Missing: false}
	case uint32:
		if v.(uint32)&0xFFFFFFFF == 0xFFFFFFFF {
			return Grib2FlagValue{Val: uint(v.(uint32)), Missing: true}
		}
		return Grib2FlagValue{Val: uint(v.(uint32)), Missing: false}
	case uint64:
		if v.(uint64)&0xFFFFFFFFFFFFFFFF == 0xFFFFFFFFFFFFFFFF {
			return Grib2FlagValue{Val: uint(v.(uint64)), Missing: true}
		}
		return Grib2FlagValue{Val: uint(v.(uint64)), Missing: false}
	default:
		log.Fatal("Unknown Type")
		return Grib2FlagValue{}
	}
}
