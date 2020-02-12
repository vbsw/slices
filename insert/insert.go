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
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Byte inserts value in list.
func Byte(list []byte, index int, value ...byte) []byte {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Complex64 inserts value in list.
func Complex64(list []complex64, index int, value ...complex64) []complex64 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Complex128 inserts value in list.
func Complex128(list []complex128, index int, value ...complex128) []complex128 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Error inserts value in list.
func Error(list []error, index int, value ...error) []error {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Float32 inserts value in list.
func Float32(list []float32, index int, value ...float32) []float32 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Float64 inserts value in list.
func Float64(list []float64, index int, value ...float64) []float64 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int inserts value in list.
func Int(list []int, index int, value ...int) []int {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int8 inserts value in list.
func Int8(list []int8, index int, value ...int8) []int8 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int16 inserts value in list.
func Int16(list []int16, index int, value ...int16) []int16 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int32 inserts value in list.
func Int32(list []int32, index int, value ...int32) []int32 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Int64 inserts value in list.
func Int64(list []int64, index int, value ...int64) []int64 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Interface inserts value in list.
func Interface(list []interface{}, index int, value ...interface{}) []interface{} {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Pointer inserts value in list.
func Pointer(list []unsafe.Pointer, index int, value ...unsafe.Pointer) []unsafe.Pointer {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// Rune inserts value in list.
func Rune(list []rune, index int, value ...rune) []rune {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// String inserts value in list.
func String(list []string, index int, value ...string) []string {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt inserts value in list.
func UInt(list []uint, index int, value ...uint) []uint {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt8 inserts value in list.
func UInt8(list []uint8, index int, value ...uint8) []uint8 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt16 inserts value in list.
func UInt16(list []uint16, index int, value ...uint16) []uint16 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt32 inserts value in list.
func UInt32(list []uint32, index int, value ...uint32) []uint32 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}

// UInt64 inserts value in list.
func UInt64(list []uint64, index int, value ...uint64) []uint64 {
	extendedList := append(list, value...)
	copy(extendedList[index+len(value):], list[index:])
	copy(extendedList[index:], value)
	return extendedList
}
