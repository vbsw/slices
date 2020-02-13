/*
 *          Copyright 2018, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package remove provides the "remove" function for slices of basic types.
package remove

import (
	"unsafe"
)

// Bool removes an element from list.
func Bool(list []bool, index int) []bool {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Byte removes an element from list.
func Byte(list []byte, index int) []byte {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex64 removes an element from list.
func Complex64(list []complex64, index int) []complex64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex128 removes an element from list.
func Complex128(list []complex128, index int) []complex128 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Error removes an element from list.
func Error(list []error, index int) []error {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float32 removes an element from list.
func Float32(list []float32, index int) []float32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float64 removes an element from list.
func Float64(list []float64, index int) []float64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int removes an element from list.
func Int(list []int, index int) []int {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int8 removes an element from list.
func Int8(list []int8, index int) []int8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int16 removes an element from list.
func Int16(list []int16, index int) []int16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int32 removes an element from list.
func Int32(list []int32, index int) []int32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int64 removes an element from list.
func Int64(list []int64, index int) []int64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Interface removes an element from list.
func Interface(list []interface{}, index int) []interface{} {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Pointer removes an element from list.
func Pointer(list []unsafe.Pointer, index int) []unsafe.Pointer {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Rune removes an element from list.
func Rune(list []rune, index int) []rune {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// String removes an element from list.
func String(list []string, index int) []string {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt removes an element from list.
func UInt(list []uint, index int) []uint {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt8 removes an element from list.
func UInt8(list []uint8, index int) []uint8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt16 removes an element from list.
func UInt16(list []uint16, index int) []uint16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt32 removes an element from list.
func UInt32(list []uint32, index int) []uint32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt64 removes an element from list.
func UInt64(list []uint64, index int) []uint64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}
