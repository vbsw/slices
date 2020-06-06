/*
 *       Copyright 2018, 2020, Vitali Baumtrok.
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
func Bool(list []bool, index int, value bool) []bool {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// BoolD2 inserts value in list.
func BoolD2(list [][]bool, index int, value []bool) [][]bool {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// BoolD3 inserts value in list.
func BoolD3(list [][][]bool, index int, value [][]bool) [][][]bool {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// BoolD4 inserts value in list.
func BoolD4(list [][][][]bool, index int, value [][][]bool) [][][][]bool {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// BoolD5 inserts value in list.
func BoolD5(list [][][][][]bool, index int, value [][][][]bool) [][][][][]bool {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Byte inserts value in list.
func Byte(list []byte, index int, value byte) []byte {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ByteD2 inserts value in list.
func ByteD2(list [][]byte, index int, value []byte) [][]byte {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ByteD3 inserts value in list.
func ByteD3(list [][][]byte, index int, value [][]byte) [][][]byte {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ByteD4 inserts value in list.
func ByteD4(list [][][][]byte, index int, value [][][]byte) [][][][]byte {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ByteD5 inserts value in list.
func ByteD5(list [][][][][]byte, index int, value [][][][]byte) [][][][][]byte {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex64 inserts value in list.
func Complex64(list []complex64, index int, value complex64) []complex64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex64D2 inserts value in list.
func Complex64D2(list [][]complex64, index int, value []complex64) [][]complex64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex64D3 inserts value in list.
func Complex64D3(list [][][]complex64, index int, value [][]complex64) [][][]complex64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex64D4 inserts value in list.
func Complex64D4(list [][][][]complex64, index int, value [][][]complex64) [][][][]complex64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex64D5 inserts value in list.
func Complex64D5(list [][][][][]complex64, index int, value [][][][]complex64) [][][][][]complex64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex128 inserts value in list.
func Complex128(list []complex128, index int, value complex128) []complex128 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex128D2 inserts value in list.
func Complex128D2(list [][]complex128, index int, value []complex128) [][]complex128 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex128D3 inserts value in list.
func Complex128D3(list [][][]complex128, index int, value [][]complex128) [][][]complex128 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex128D4 inserts value in list.
func Complex128D4(list [][][][]complex128, index int, value [][][]complex128) [][][][]complex128 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Complex128D5 inserts value in list.
func Complex128D5(list [][][][][]complex128, index int, value [][][][]complex128) [][][][][]complex128 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Error inserts value in list.
func Error(list []error, index int, value error) []error {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ErrorD2 inserts value in list.
func ErrorD2(list [][]error, index int, value []error) [][]error {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ErrorD3 inserts value in list.
func ErrorD3(list [][][]error, index int, value [][]error) [][][]error {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ErrorD4 inserts value in list.
func ErrorD4(list [][][][]error, index int, value [][][]error) [][][][]error {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// ErrorD5 inserts value in list.
func ErrorD5(list [][][][][]error, index int, value [][][][]error) [][][][][]error {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float32 inserts value in list.
func Float32(list []float32, index int, value float32) []float32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float32D2 inserts value in list.
func Float32D2(list [][]float32, index int, value []float32) [][]float32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float32D3 inserts value in list.
func Float32D3(list [][][]float32, index int, value [][]float32) [][][]float32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float32D4 inserts value in list.
func Float32D4(list [][][][]float32, index int, value [][][]float32) [][][][]float32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float32D5 inserts value in list.
func Float32D5(list [][][][][]float32, index int, value [][][][]float32) [][][][][]float32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float64 inserts value in list.
func Float64(list []float64, index int, value float64) []float64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float64D2 inserts value in list.
func Float64D2(list [][]float64, index int, value []float64) [][]float64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float64D3 inserts value in list.
func Float64D3(list [][][]float64, index int, value [][]float64) [][][]float64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float64D4 inserts value in list.
func Float64D4(list [][][][]float64, index int, value [][][]float64) [][][][]float64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Float64D5 inserts value in list.
func Float64D5(list [][][][][]float64, index int, value [][][][]float64) [][][][][]float64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int inserts value in list.
func Int(list []int, index int, value int) []int {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// IntD2 inserts value in list.
func IntD2(list [][]int, index int, value []int) [][]int {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// IntD3 inserts value in list.
func IntD3(list [][][]int, index int, value [][]int) [][][]int {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// IntD4 inserts value in list.
func IntD4(list [][][][]int, index int, value [][][]int) [][][][]int {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// IntD5 inserts value in list.
func IntD5(list [][][][][]int, index int, value [][][][]int) [][][][][]int {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int8 inserts value in list.
func Int8(list []int8, index int, value int8) []int8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int8D2 inserts value in list.
func Int8D2(list [][]int8, index int, value []int8) [][]int8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int8D3 inserts value in list.
func Int8D3(list [][][]int8, index int, value [][]int8) [][][]int8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int8D4 inserts value in list.
func Int8D4(list [][][][]int8, index int, value [][][]int8) [][][][]int8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int8D5 inserts value in list.
func Int8D5(list [][][][][]int8, index int, value [][][][]int8) [][][][][]int8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int16 inserts value in list.
func Int16(list []int16, index int, value int16) []int16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int16D2 inserts value in list.
func Int16D2(list [][]int16, index int, value []int16) [][]int16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int16D3 inserts value in list.
func Int16D3(list [][][]int16, index int, value [][]int16) [][][]int16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int16D4 inserts value in list.
func Int16D4(list [][][][]int16, index int, value [][][]int16) [][][][]int16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int16D5 inserts value in list.
func Int16D5(list [][][][][]int16, index int, value [][][][]int16) [][][][][]int16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int32 inserts value in list.
func Int32(list []int32, index int, value int32) []int32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int32D2 inserts value in list.
func Int32D2(list [][]int32, index int, value []int32) [][]int32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int32D3 inserts value in list.
func Int32D3(list [][][]int32, index int, value [][]int32) [][][]int32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int32D4 inserts value in list.
func Int32D4(list [][][][]int32, index int, value [][][]int32) [][][][]int32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int32D5 inserts value in list.
func Int32D5(list [][][][][]int32, index int, value [][][][]int32) [][][][][]int32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int64 inserts value in list.
func Int64(list []int64, index int, value int64) []int64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int64D2 inserts value in list.
func Int64D2(list [][]int64, index int, value []int64) [][]int64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int64D3 inserts value in list.
func Int64D3(list [][][]int64, index int, value [][]int64) [][][]int64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int64D4 inserts value in list.
func Int64D4(list [][][][]int64, index int, value [][][]int64) [][][][]int64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Int64D5 inserts value in list.
func Int64D5(list [][][][][]int64, index int, value [][][][]int64) [][][][][]int64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Interface inserts value in list.
func Interface(list []interface{}, index int, value interface{}) []interface{} {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// InterfaceD2 inserts value in list.
func InterfaceD2(list [][]interface{}, index int, value []interface{}) [][]interface{} {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// InterfaceD3 inserts value in list.
func InterfaceD3(list [][][]interface{}, index int, value [][]interface{}) [][][]interface{} {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// InterfaceD4 inserts value in list.
func InterfaceD4(list [][][][]interface{}, index int, value [][][]interface{}) [][][][]interface{} {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// InterfaceD5 inserts value in list.
func InterfaceD5(list [][][][][]interface{}, index int, value [][][][]interface{}) [][][][][]interface{} {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Pointer inserts value in list.
func Pointer(list []unsafe.Pointer, index int, value unsafe.Pointer) []unsafe.Pointer {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// PointerD2 inserts value in list.
func PointerD2(list [][]unsafe.Pointer, index int, value []unsafe.Pointer) [][]unsafe.Pointer {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// PointerD3 inserts value in list.
func PointerD3(list [][][]unsafe.Pointer, index int, value [][]unsafe.Pointer) [][][]unsafe.Pointer {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// PointerD4 inserts value in list.
func PointerD4(list [][][][]unsafe.Pointer, index int, value [][][]unsafe.Pointer) [][][][]unsafe.Pointer {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// PointerD5 inserts value in list.
func PointerD5(list [][][][][]unsafe.Pointer, index int, value [][][][]unsafe.Pointer) [][][][][]unsafe.Pointer {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// Rune inserts value in list.
func Rune(list []rune, index int, value rune) []rune {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// RuneD2 inserts value in list.
func RuneD2(list [][]rune, index int, value []rune) [][]rune {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// RuneD3 inserts value in list.
func RuneD3(list [][][]rune, index int, value [][]rune) [][][]rune {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// RuneD4 inserts value in list.
func RuneD4(list [][][][]rune, index int, value [][][]rune) [][][][]rune {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// RuneD5 inserts value in list.
func RuneD5(list [][][][][]rune, index int, value [][][][]rune) [][][][][]rune {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// String inserts value in list.
func String(list []string, index int, value string) []string {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// StringD2 inserts value in list.
func StringD2(list [][]string, index int, value []string) [][]string {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// StringD3 inserts value in list.
func StringD3(list [][][]string, index int, value [][]string) [][][]string {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// StringD4 inserts value in list.
func StringD4(list [][][][]string, index int, value [][][]string) [][][][]string {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// StringD5 inserts value in list.
func StringD5(list [][][][][]string, index int, value [][][][]string) [][][][][]string {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt inserts value in list.
func UInt(list []uint, index int, value uint) []uint {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UIntD2 inserts value in list.
func UIntD2(list [][]uint, index int, value []uint) [][]uint {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UIntD3 inserts value in list.
func UIntD3(list [][][]uint, index int, value [][]uint) [][][]uint {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UIntD4 inserts value in list.
func UIntD4(list [][][][]uint, index int, value [][][]uint) [][][][]uint {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UIntD5 inserts value in list.
func UIntD5(list [][][][][]uint, index int, value [][][][]uint) [][][][][]uint {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt8 inserts value in list.
func UInt8(list []uint8, index int, value uint8) []uint8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt8D2 inserts value in list.
func UInt8D2(list [][]uint8, index int, value []uint8) [][]uint8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt8D3 inserts value in list.
func UInt8D3(list [][][]uint8, index int, value [][]uint8) [][][]uint8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt8D4 inserts value in list.
func UInt8D4(list [][][][]uint8, index int, value [][][]uint8) [][][][]uint8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt8D5 inserts value in list.
func UInt8D5(list [][][][][]uint8, index int, value [][][][]uint8) [][][][][]uint8 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt16 inserts value in list.
func UInt16(list []uint16, index int, value uint16) []uint16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt16D2 inserts value in list.
func UInt16D2(list [][]uint16, index int, value []uint16) [][]uint16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt16D3 inserts value in list.
func UInt16D3(list [][][]uint16, index int, value [][]uint16) [][][]uint16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt16D4 inserts value in list.
func UInt16D4(list [][][][]uint16, index int, value [][][]uint16) [][][][]uint16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt16D5 inserts value in list.
func UInt16D5(list [][][][][]uint16, index int, value [][][][]uint16) [][][][][]uint16 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt32 inserts value in list.
func UInt32(list []uint32, index int, value uint32) []uint32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt32D2 inserts value in list.
func UInt32D2(list [][]uint32, index int, value []uint32) [][]uint32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt32D3 inserts value in list.
func UInt32D3(list [][][]uint32, index int, value [][]uint32) [][][]uint32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt32D4 inserts value in list.
func UInt32D4(list [][][][]uint32, index int, value [][][]uint32) [][][][]uint32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt32D5 inserts value in list.
func UInt32D5(list [][][][][]uint32, index int, value [][][][]uint32) [][][][][]uint32 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt64 inserts value in list.
func UInt64(list []uint64, index int, value uint64) []uint64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt64D2 inserts value in list.
func UInt64D2(list [][]uint64, index int, value []uint64) [][]uint64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt64D3 inserts value in list.
func UInt64D3(list [][][]uint64, index int, value [][]uint64) [][][]uint64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt64D4 inserts value in list.
func UInt64D4(list [][][][]uint64, index int, value [][][]uint64) [][][][]uint64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}

// UInt64D5 inserts value in list.
func UInt64D5(list [][][][][]uint64, index int, value [][][][]uint64) [][][][][]uint64 {
	extendedList := append(list, value)
	copy(extendedList[index+1:], list[index:])
	extendedList[index] = value
	return extendedList
}
