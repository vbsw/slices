/*
 *          Copyright 2018, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package clone provides the "clone" function for slices of basic types.
package clone

import (
	"unsafe"
)

// Bool makes a copy of the given list. The copy is returned.
func Bool(list []bool) []bool {
	listCopy := make([]bool, len(list))
	copy(listCopy, list)
	return listCopy
}

// Byte makes a copy of the given list. The copy is returned.
func Byte(list []byte) []byte {
	listCopy := make([]byte, len(list))
	copy(listCopy, list)
	return listCopy
}

// Complex64 makes a copy of the given list. The copy is returned.
func Complex64(list []complex64) []complex64 {
	listCopy := make([]complex64, len(list))
	copy(listCopy, list)
	return listCopy
}

// Complex128 makes a copy of the given list. The copy is returned.
func Complex128(list []complex128) []complex128 {
	listCopy := make([]complex128, len(list))
	copy(listCopy, list)
	return listCopy
}

// Error makes a copy of the given list. The copy is returned.
func Error(list []error) []error {
	listCopy := make([]error, len(list))
	copy(listCopy, list)
	return listCopy
}

// Float32 makes a copy of the given list. The copy is returned.
func Float32(list []float32) []float32 {
	listCopy := make([]float32, len(list))
	copy(listCopy, list)
	return listCopy
}

// Float64 makes a copy of the given list. The copy is returned.
func Float64(list []float64) []float64 {
	listCopy := make([]float64, len(list))
	copy(listCopy, list)
	return listCopy
}

// Int makes a copy of the given list. The copy is returned.
func Int(list []int) []int {
	listCopy := make([]int, len(list))
	copy(listCopy, list)
	return listCopy
}

// Int8 makes a copy of the given list. The copy is returned.
func Int8(list []int8) []int8 {
	listCopy := make([]int8, len(list))
	copy(listCopy, list)
	return listCopy
}

// Int16 makes a copy of the given list. The copy is returned.
func Int16(list []int16) []int16 {
	listCopy := make([]int16, len(list))
	copy(listCopy, list)
	return listCopy
}

// Int32 makes a copy of the given list. The copy is returned.
func Int32(list []int32) []int32 {
	listCopy := make([]int32, len(list))
	copy(listCopy, list)
	return listCopy
}

// Int64 makes a copy of the given list. The copy is returned.
func Int64(list []int64) []int64 {
	listCopy := make([]int64, len(list))
	copy(listCopy, list)
	return listCopy
}

// Interface makes a copy of the given list. The copy is returned.
func Interface(list []interface{}) []interface{} {
	listCopy := make([]interface{}, len(list))
	copy(listCopy, list)
	return listCopy
}

// Pointer makes a copy of the given list. The copy is returned.
func Pointer(list []unsafe.Pointer) []unsafe.Pointer {
	listCopy := make([]unsafe.Pointer, len(list))
	copy(listCopy, list)
	return listCopy
}

// Rune makes a copy of the given list. The copy is returned.
func Rune(list []rune) []rune {
	listCopy := make([]rune, len(list))
	copy(listCopy, list)
	return listCopy
}

// String makes a copy of the given list. The copy is returned.
func String(list []string) []string {
	listCopy := make([]string, len(list))
	copy(listCopy, list)
	return listCopy
}

// UInt makes a copy of the given list. The copy is returned.
func UInt(list []uint) []uint {
	listCopy := make([]uint, len(list))
	copy(listCopy, list)
	return listCopy
}

// UInt8 makes a copy of the given list. The copy is returned.
func UInt8(list []uint8) []uint8 {
	listCopy := make([]uint8, len(list))
	copy(listCopy, list)
	return listCopy
}

// UInt16 makes a copy of the given list. The copy is returned.
func UInt16(list []uint16) []uint16 {
	listCopy := make([]uint16, len(list))
	copy(listCopy, list)
	return listCopy
}

// UInt32 makes a copy of the given list. The copy is returned.
func UInt32(list []uint32) []uint32 {
	listCopy := make([]uint32, len(list))
	copy(listCopy, list)
	return listCopy
}

// UInt64 makes a copy of the given list. The copy is returned.
func UInt64(list []uint64) []uint64 {
	listCopy := make([]uint64, len(list))
	copy(listCopy, list)
	return listCopy
}
