/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package count provides the "count" function for slices of basic types.
package count

import (
	"unsafe"
)

// Bool returns count of value in list.
func Bool(list []bool, value ...bool) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Byte returns count of value in list.
func Byte(list []byte, value ...byte) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Complex64 returns count of value in list.
func Complex64(list []complex64, value ...complex64) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Complex128 returns count of value in list.
func Complex128(list []complex128, value ...complex128) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Error returns count of value in list.
func Error(list []error, value ...error) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Float32 returns count of value in list.
func Float32(list []float32, value ...float32) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Float64 returns count of value in list.
func Float64(list []float64, value ...float64) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Int returns count of value in list.
func Int(list []int, value ...int) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Int8 returns count of value in list.
func Int8(list []int8, value ...int8) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Int16 returns count of value in list.
func Int16(list []int16, value ...int16) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Int32 returns count of value in list.
func Int32(list []int32, value ...int32) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Int64 returns count of value in list.
func Int64(list []int64, value ...int64) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Interface returns count of value in list.
func Interface(list []interface{}, value ...interface{}) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Pointer returns count of value in list.
func Pointer(list []unsafe.Pointer, value ...unsafe.Pointer) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// Rune returns count of value in list.
func Rune(list []rune, value ...rune) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// String returns count of value in list.
func String(list []string, value ...string) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// UInt returns count of value in list.
func UInt(list []uint, value ...uint) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// UInt8 returns count of value in list.
func UInt8(list []uint8, value ...uint8) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// UInt16 returns count of value in list.
func UInt16(list []uint16, value ...uint16) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// UInt32 returns count of value in list.
func UInt32(list []uint32, value ...uint32) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}

// UInt64 returns count of value in list.
func UInt64(list []uint64, value ...uint64) int {
	valueCount := 0
	for _, val := range value {
		for _, listValue := range list {
			if listValue == val {
				valueCount++
			}
		}
	}
	return valueCount
}
