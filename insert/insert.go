/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package insert provides the "insert" function for slices of basic types.
package insert

import (
	"unsafe"
)

// Bool inserts value in list.
func Bool(list []bool, index int, value ...bool) []bool {
	var extendedList []bool
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]bool, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Byte inserts value in list.
func Byte(list []byte, index int, value ...byte) []byte {
	var extendedList []byte
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]byte, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Complex64 inserts value in list.
func Complex64(list []complex64, index int, value ...complex64) []complex64 {
	var extendedList []complex64
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]complex64, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Complex128 inserts value in list.
func Complex128(list []complex128, index int, value ...complex128) []complex128 {
	var extendedList []complex128
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]complex128, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Error inserts value in list.
func Error(list []error, index int, value ...error) []error {
	var extendedList []error
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]error, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Float32 inserts value in list.
func Float32(list []float32, index int, value ...float32) []float32 {
	var extendedList []float32
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]float32, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Float64 inserts value in list.
func Float64(list []float64, index int, value ...float64) []float64 {
	var extendedList []float64
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]float64, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int inserts value in list.
func Int(list []int, index int, value ...int) []int {
	var extendedList []int
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]int, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int8 inserts value in list.
func Int8(list []int8, index int, value ...int8) []int8 {
	var extendedList []int8
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]int8, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int16 inserts value in list.
func Int16(list []int16, index int, value ...int16) []int16 {
	var extendedList []int16
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]int16, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int32 inserts value in list.
func Int32(list []int32, index int, value ...int32) []int32 {
	var extendedList []int32
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]int32, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int64 inserts value in list.
func Int64(list []int64, index int, value ...int64) []int64 {
	var extendedList []int64
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]int64, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Interface inserts value in list.
func Interface(list []interface{}, index int, value ...interface{}) []interface{} {
	var extendedList []interface{}
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]interface{}, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Pointer inserts value in list.
func Pointer(list []unsafe.Pointer, index int, value ...unsafe.Pointer) []unsafe.Pointer {
	var extendedList []unsafe.Pointer
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]unsafe.Pointer, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Rune inserts value in list.
func Rune(list []rune, index int, value ...rune) []rune {
	var extendedList []rune
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]rune, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// String inserts value in list.
func String(list []string, index int, value ...string) []string {
	var extendedList []string
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]string, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt inserts value in list.
func UInt(list []uint, index int, value ...uint) []uint {
	var extendedList []uint
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]uint, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt8 inserts value in list.
func UInt8(list []uint8, index int, value ...uint8) []uint8 {
	var extendedList []uint8
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]uint8, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt16 inserts value in list.
func UInt16(list []uint16, index int, value ...uint16) []uint16 {
	var extendedList []uint16
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]uint16, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt32 inserts value in list.
func UInt32(list []uint32, index int, value ...uint32) []uint32 {
	var extendedList []uint32
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]uint32, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt64 inserts value in list.
func UInt64(list []uint64, index int, value ...uint64) []uint64 {
	var extendedList []uint64
	if cap(list)-len(list) >= len(value) {
		extendedList = list[:len(list)+len(value)]
	} else {
		extendedList = make([]uint64, len(list)+len(value))
		copy(extendedList[:index], list[:index])
	}
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}
