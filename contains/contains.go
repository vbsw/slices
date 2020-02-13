/*
 *          Copyright 2018, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package contains provides the "contains" function for slices of basic types.
package contains

import (
	"unsafe"
)

// Bool returns true, if list contains value.
func Bool(list []bool, value bool) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Byte returns true, if list contains value.
func Byte(list []byte, value byte) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Complex64 returns true, if list contains value.
func Complex64(list []complex64, value complex64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Complex128 returns true, if list contains value.
func Complex128(list []complex128, value complex128) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Error returns true, if list contains value.
func Error(list []error, value error) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Float32 returns true, if list contains value.
func Float32(list []float32, value float32) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Float64 returns true, if list contains value.
func Float64(list []float64, value float64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int returns true, if list contains value.
func Int(list []int, value int) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int8 returns true, if list contains value.
func Int8(list []int8, value int8) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int16 returns true, if list contains value.
func Int16(list []int16, value int16) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int32 returns true, if list contains value.
func Int32(list []int32, value int32) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int64 returns true, if list contains value.
func Int64(list []int64, value int64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Interface returns true, if list contains value.
func Interface(list []interface{}, value interface{}) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Pointer returns true, if list contains value.
func Pointer(list []unsafe.Pointer, value unsafe.Pointer) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Rune returns true, if list contains value.
func Rune(list []rune, value rune) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// String returns true, if list contains value.
func String(list []string, value string) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt returns true, if list contains value.
func UInt(list []uint, value uint) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt8 returns true, if list contains value.
func UInt8(list []uint8, value uint8) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt16 returns true, if list contains value.
func UInt16(list []uint16, value uint16) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt32 returns true, if list contains value.
func UInt32(list []uint32, value uint32) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt64 returns true, if list contains value.
func UInt64(list []uint64, value uint64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}
