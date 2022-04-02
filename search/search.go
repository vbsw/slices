/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package search provides the "binary search" function for slices of basic types.
package search

import (
	"unsafe"
)

// Bool searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and sorted from false to true.
func Bool(list []bool, element bool) (int, bool) {
	if len(list) > 0 {
		if !list[0] {
			if !element {
				return 0, true
			}
			return 1, len(list) >= 2
		}
		return 0, element
	}
	return 0, false
}

// BoolRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and sorted from true to false.
func BoolRev(list []bool, element bool) (int, bool) {
	if len(list) > 0 {
		if list[0] {
			if element {
				return 0, true
			}
			return 1, len(list) >= 2
		}
		return 0, !element
	}
	return 0, false
}

// Byte searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Byte(list []byte, element byte) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// ByteRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func ByteRev(list []byte, element byte) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Complex128(list []complex128, element complex128) (int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal || (elementReal == valueReal && elementImag > valueImag) {
			left = middle + 1
		} else if elementReal < valueReal || (elementReal == valueReal && elementImag < valueImag) {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Complex128Rev(list []complex128, element complex128) (int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal || (elementReal == valueReal && elementImag < valueImag) {
			left = middle + 1
		} else if elementReal > valueReal || (elementReal == valueReal && elementImag > valueImag) {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Complex64(list []complex64, element complex64) (int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal || (elementReal == valueReal && elementImag > valueImag) {
			left = middle + 1
		} else if elementReal < valueReal || (elementReal == valueReal && elementImag < valueImag) {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Complex64Rev(list []complex64, element complex64) (int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal || (elementReal == valueReal && elementImag < valueImag) {
			left = middle + 1
		} else if elementReal > valueReal || (elementReal == valueReal && elementImag > valueImag) {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float32 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Float32(list []float32, element float32) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float32Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Float32Rev(list []float32, element float32) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float64 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Float64(list []float64, element float64) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float64Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Float64Rev(list []float64, element float64) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Int(list []int, element int) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// IntRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func IntRev(list []int, element int) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int16 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Int16(list []int16, element int16) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int16Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Int16Rev(list []int16, element int16) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int32 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Int32(list []int32, element int32) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int32Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Int32Rev(list []int32, element int32) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int64 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Int64(list []int64, element int64) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int64Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Int64Rev(list []int64, element int64) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int8 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Int8(list []int8, element int8) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int8Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func Int8Rev(list []int8, element int8) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Pointer searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Pointer(list []unsafe.Pointer, element unsafe.Pointer) (int, bool) {
	left := 0
	right := len(list) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		valueUIntPtr := uintptr(list[middle])
		if elementUIntPtr > valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr < valueUIntPtr {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// PointerRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func PointerRev(list []unsafe.Pointer, element unsafe.Pointer) (int, bool) {
	left := 0
	right := len(list) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		valueUIntPtr := uintptr(list[middle])
		if elementUIntPtr < valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr > valueUIntPtr {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Rune searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func Rune(list []rune, element rune) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// RuneRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func RuneRev(list []rune, element rune) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// String searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func String(list []string, element string) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// StringRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func StringRev(list []string, element string) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func UInt(list []uint, element uint) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UIntRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func UIntRev(list []uint, element uint) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt16 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func UInt16(list []uint16, element uint16) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt16Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func UInt16Rev(list []uint16, element uint16) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt32 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func UInt32(list []uint32, element uint32) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt32Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func UInt32Rev(list []uint32, element uint32) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt64 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func UInt64(list []uint64, element uint64) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt64Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func UInt64Rev(list []uint64, element uint64) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt8 searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func UInt8(list []uint8, element uint8) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt8Rev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func UInt8Rev(list []uint8, element uint8) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UIntPtr searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in ascending order.
func UIntPtr(list []uintptr, element uintptr) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UIntPtrRev searches an element in list and returns its index. Second value returned
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Precondition is, elements in list are unique and in descending order.
func UIntPtrRev(list []uintptr, element uintptr) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}
