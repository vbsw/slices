/*
 *          Copyright 2020, Vitali Baumtrok.
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
func Bool(list []bool, value ...bool) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Byte returns true, if list contains value.
func Byte(list []byte, value ...byte) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Complex64 returns true, if list contains value.
func Complex64(list []complex64, value ...complex64) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Complex128 returns true, if list contains value.
func Complex128(list []complex128, value ...complex128) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Error returns true, if list contains value.
func Error(list []error, value ...error) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Float32 returns true, if list contains value.
func Float32(list []float32, value ...float32) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Float64 returns true, if list contains value.
func Float64(list []float64, value ...float64) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Int returns true, if list contains value.
func Int(list []int, value ...int) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Int8 returns true, if list contains value.
func Int8(list []int8, value ...int8) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Int16 returns true, if list contains value.
func Int16(list []int16, value ...int16) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Int32 returns true, if list contains value.
func Int32(list []int32, value ...int32) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Int64 returns true, if list contains value.
func Int64(list []int64, value ...int64) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Interface returns true, if list contains value.
func Interface(list []interface{}, value ...interface{}) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Pointer returns true, if list contains value.
func Pointer(list []unsafe.Pointer, value ...unsafe.Pointer) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// Rune returns true, if list contains value.
func Rune(list []rune, value ...rune) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// String returns true, if list contains value.
func String(list []string, value ...string) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// UInt returns true, if list contains value.
func UInt(list []uint, value ...uint) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// UInt8 returns true, if list contains value.
func UInt8(list []uint8, value ...uint8) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// UInt16 returns true, if list contains value.
func UInt16(list []uint16, value ...uint16) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// UInt32 returns true, if list contains value.
func UInt32(list []uint32, value ...uint32) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}

// UInt64 returns true, if list contains value.
func UInt64(list []uint64, value ...uint64) bool {
	found := -1
	for i, val := range value {
		for _, listValue := range list {
			if listValue == val {
				found++
				break
			}
		}
		if found < i {
			return false
		}
	}
	return true
}
