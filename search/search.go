/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package search provides the "binary search" function for slices of basic types.
package search

import (
	"strings"
	"unsafe"
)

// Bool searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and sorted from false to true.
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

// BoolDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and sorted from true to false.
func BoolDesc(list []bool, element bool) (int, bool) {
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

// BoolIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and sorted from false to true by indices.
func BoolIdx(list []bool, indices []int, element bool) (int, bool) {
	if len(indices) > 0 {
		valueIndex := indices[0]
		if !list[valueIndex] {
			if !element {
				return 0, true
			}
			return 1, len(indices) >= 2
		}
		return 0, element
	}
	return 0, false
}

// BoolIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and sorted from true to false by indices.
func BoolIdxDesc(list []bool, indices []int, element bool) (int, bool) {
	if len(indices) > 0 {
		valueIndex := indices[0]
		if list[valueIndex] {
			if element {
				return 0, true
			}
			return 1, len(indices) >= 2
		}
		return 0, !element
	}
	return 0, false
}

// BoolIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be sorted from false to true by indices.
func BoolIdxRng(list []bool, indices []int, element bool) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[0]
		value := list[valueIndex]
		if element != value {
			if element {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			from := boolIdxRngL(list, indices, element, left, middle-1)
			to := boolIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// BoolIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be sorted from true to false by indices.
func BoolIdxRngDesc(list []bool, indices []int, element bool) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[0]
		value := list[valueIndex]
		if element != value {
			if !element {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			from := boolIdxRngDescL(list, indices, element, left, middle-1)
			to := boolIdxRngDescR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// BoolRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be sorted from false to true.
func BoolRng(list []bool, element bool) (int, int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element != value {
			if element {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			from := boolRngL(list, element, left, middle-1)
			to := boolRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// BoolRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be sorted from true to false.
func BoolRngDesc(list []bool, element bool) (int, int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element != value {
			if !element {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			from := boolRngDescL(list, element, left, middle-1)
			to := boolRngDescR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Byte searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// ByteDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func ByteDesc(list []byte, element byte) (int, bool) {
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

// ByteIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func ByteIdx(list []byte, indices []int, element byte) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// ByteIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func ByteIdxDesc(list []byte, indices []int, element byte) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// ByteIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func ByteIdxRng(list []byte, indices []int, element byte) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := byteIdxRngL(list, indices, element, left, middle-1)
			to := byteIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// ByteIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func ByteIdxRngDesc(list []byte, indices []int, element byte) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := byteIdxRngL(list, indices, element, left, middle-1)
			to := byteIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// ByteRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func ByteRng(list []byte, element byte) (int, int, bool) {
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
			from := byteRngL(list, element, left, middle-1)
			to := byteRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// ByteRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func ByteRngDesc(list []byte, element byte) (int, int, bool) {
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
			from := byteRngL(list, element, left, middle-1)
			to := byteRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex128 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Complex128Desc(list []complex128, element complex128) (int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Complex128Idx(list []complex128, indices []int, element complex128) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Complex128IdxDesc(list []complex128, indices []int, element complex128) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Complex128IdxRng(list []complex128, indices []int, element complex128) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			from := complex128IdxRngL(list, indices, element, left, middle-1)
			to := complex128IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex128IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Complex128IdxRngDesc(list []complex128, indices []int, element complex128) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			from := complex128IdxRngL(list, indices, element, left, middle-1)
			to := complex128IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex128Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Complex128Rng(list []complex128, element complex128) (int, int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			from := complex128RngL(list, element, left, middle-1)
			to := complex128RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex128RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Complex128RngDesc(list []complex128, element complex128) (int, int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			from := complex128RngL(list, element, left, middle-1)
			to := complex128RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex64 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Complex64Desc(list []complex64, element complex64) (int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Complex64Idx(list []complex64, indices []int, element complex64) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Complex64IdxDesc(list []complex64, indices []int, element complex64) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Complex64IdxRng(list []complex64, indices []int, element complex64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			from := complex64IdxRngL(list, indices, element, left, middle-1)
			to := complex64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex64IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Complex64IdxRngDesc(list []complex64, indices []int, element complex64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			from := complex64IdxRngL(list, indices, element, left, middle-1)
			to := complex64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex64Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Complex64Rng(list []complex64, element complex64) (int, int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			from := complex64RngL(list, element, left, middle-1)
			to := complex64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex64RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Complex64RngDesc(list []complex64, element complex64) (int, int, bool) {
	left := 0
	right := len(list) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			from := complex64RngL(list, element, left, middle-1)
			to := complex64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float32 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// Float32Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Float32Desc(list []float32, element float32) (int, bool) {
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

// Float32Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Float32Idx(list []float32, indices []int, element float32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Float32IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Float32IdxDesc(list []float32, indices []int, element float32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Float32IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Float32IdxRng(list []float32, indices []int, element float32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := float32IdxRngL(list, indices, element, left, middle-1)
			to := float32IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float32IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Float32IdxRngDesc(list []float32, indices []int, element float32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := float32IdxRngL(list, indices, element, left, middle-1)
			to := float32IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float32Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Float32Rng(list []float32, element float32) (int, int, bool) {
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
			from := float32RngL(list, element, left, middle-1)
			to := float32RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float32RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Float32RngDesc(list []float32, element float32) (int, int, bool) {
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
			from := float32RngL(list, element, left, middle-1)
			to := float32RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float64 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// Float64Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Float64Desc(list []float64, element float64) (int, bool) {
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

// Float64Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Float64Idx(list []float64, indices []int, element float64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Float64IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Float64IdxDesc(list []float64, indices []int, element float64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Float64IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Float64IdxRng(list []float64, indices []int, element float64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := float64IdxRngL(list, indices, element, left, middle-1)
			to := float64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float64IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Float64IdxRngDesc(list []float64, indices []int, element float64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := float64IdxRngL(list, indices, element, left, middle-1)
			to := float64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float64Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Float64Rng(list []float64, element float64) (int, int, bool) {
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
			from := float64RngL(list, element, left, middle-1)
			to := float64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float64RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Float64RngDesc(list []float64, element float64) (int, int, bool) {
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
			from := float64RngL(list, element, left, middle-1)
			to := float64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// IntDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func IntDesc(list []int, element int) (int, bool) {
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

// IntIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func IntIdx(list []int, indices []int, element int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// IntIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func IntIdxDesc(list []int, indices []int, element int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// IntIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func IntIdxRng(list []int, indices []int, element int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := intIdxRngL(list, indices, element, left, middle-1)
			to := intIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// IntIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func IntIdxRngDesc(list []int, indices []int, element int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := intIdxRngL(list, indices, element, left, middle-1)
			to := intIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// IntRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func IntRng(list []int, element int) (int, int, bool) {
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
			from := intRngL(list, element, left, middle-1)
			to := intRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// IntRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func IntRngDesc(list []int, element int) (int, int, bool) {
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
			from := intRngL(list, element, left, middle-1)
			to := intRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int16 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// Int16Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int16Desc(list []int16, element int16) (int, bool) {
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

// Int16Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Int16Idx(list []int16, indices []int, element int16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int16IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Int16IdxDesc(list []int16, indices []int, element int16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int16IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Int16IdxRng(list []int16, indices []int, element int16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := int16IdxRngL(list, indices, element, left, middle-1)
			to := int16IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int16IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Int16IdxRngDesc(list []int16, indices []int, element int16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := int16IdxRngL(list, indices, element, left, middle-1)
			to := int16IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int16Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int16Rng(list []int16, element int16) (int, int, bool) {
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
			from := int16RngL(list, element, left, middle-1)
			to := int16RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int16RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int16RngDesc(list []int16, element int16) (int, int, bool) {
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
			from := int16RngL(list, element, left, middle-1)
			to := int16RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int32 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// Int32Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int32Desc(list []int32, element int32) (int, bool) {
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

// Int32Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Int32Idx(list []int32, indices []int, element int32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int32IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Int32IdxDesc(list []int32, indices []int, element int32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int32IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Int32IdxRng(list []int32, indices []int, element int32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := int32IdxRngL(list, indices, element, left, middle-1)
			to := int32IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int32IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Int32IdxRngDesc(list []int32, indices []int, element int32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := int32IdxRngL(list, indices, element, left, middle-1)
			to := int32IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int32Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int32Rng(list []int32, element int32) (int, int, bool) {
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
			from := int32RngL(list, element, left, middle-1)
			to := int32RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int32RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int32RngDesc(list []int32, element int32) (int, int, bool) {
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
			from := int32RngL(list, element, left, middle-1)
			to := int32RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int64 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// Int64Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int64Desc(list []int64, element int64) (int, bool) {
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

// Int64Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Int64Idx(list []int64, indices []int, element int64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int64IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Int64IdxDesc(list []int64, indices []int, element int64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int64IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Int64IdxRng(list []int64, indices []int, element int64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := int64IdxRngL(list, indices, element, left, middle-1)
			to := int64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int64IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Int64IdxRngDesc(list []int64, indices []int, element int64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := int64IdxRngL(list, indices, element, left, middle-1)
			to := int64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int64Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int64Rng(list []int64, element int64) (int, int, bool) {
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
			from := int64RngL(list, element, left, middle-1)
			to := int64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int64RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int64RngDesc(list []int64, element int64) (int, int, bool) {
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
			from := int64RngL(list, element, left, middle-1)
			to := int64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int8 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// Int8Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int8Desc(list []int8, element int8) (int, bool) {
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

// Int8Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func Int8Idx(list []int8, indices []int, element int8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int8IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func Int8IdxDesc(list []int8, indices []int, element int8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// Int8IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func Int8IdxRng(list []int8, indices []int, element int8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := int8IdxRngL(list, indices, element, left, middle-1)
			to := int8IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int8IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func Int8IdxRngDesc(list []int8, indices []int, element int8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := int8IdxRngL(list, indices, element, left, middle-1)
			to := int8IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int8Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int8Rng(list []int8, element int8) (int, int, bool) {
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
			from := int8RngL(list, element, left, middle-1)
			to := int8RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int8RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int8RngDesc(list []int8, element int8) (int, int, bool) {
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
			from := int8RngL(list, element, left, middle-1)
			to := int8RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Pointer searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// PointerDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func PointerDesc(list []unsafe.Pointer, element unsafe.Pointer) (int, bool) {
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

// PointerIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func PointerIdx(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		valueUIntPtr := uintptr(list[valueIndex])
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

// PointerIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func PointerIdxDesc(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		valueUIntPtr := uintptr(list[valueIndex])
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

// PointerIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func PointerIdxRng(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		valueUIntPtr := uintptr(list[valueIndex])
		if elementUIntPtr > valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr < valueUIntPtr {
			right = middle - 1
		} else {
			from := pointerIdxRngL(list, indices, element, left, middle-1)
			to := pointerIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// PointerIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func PointerIdxRngDesc(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		valueUIntPtr := uintptr(list[valueIndex])
		if elementUIntPtr < valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr > valueUIntPtr {
			right = middle - 1
		} else {
			from := pointerIdxRngL(list, indices, element, left, middle-1)
			to := pointerIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// PointerRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func PointerRng(list []unsafe.Pointer, element unsafe.Pointer) (int, int, bool) {
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
			from := pointerRngL(list, element, left, middle-1)
			to := pointerRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// PointerRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func PointerRngDesc(list []unsafe.Pointer, element unsafe.Pointer) (int, int, bool) {
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
			from := pointerRngL(list, element, left, middle-1)
			to := pointerRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Rune searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// RuneDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func RuneDesc(list []rune, element rune) (int, bool) {
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

// RuneIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func RuneIdx(list []rune, indices []int, element rune) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// RuneIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func RuneIdxDesc(list []rune, indices []int, element rune) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// RuneIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func RuneIdxRng(list []rune, indices []int, element rune) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := runeIdxRngL(list, indices, element, left, middle-1)
			to := runeIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// RuneIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func RuneIdxRngDesc(list []rune, indices []int, element rune) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := runeIdxRngL(list, indices, element, left, middle-1)
			to := runeIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// RuneRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func RuneRng(list []rune, element rune) (int, int, bool) {
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
			from := runeRngL(list, element, left, middle-1)
			to := runeRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// RuneRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func RuneRngDesc(list []rune, element rune) (int, int, bool) {
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
			from := runeRngL(list, element, left, middle-1)
			to := runeRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// String searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// StringDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func StringDesc(list []string, element string) (int, bool) {
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

// StringIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func StringIdx(list []string, indices []int, element string) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// StringIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func StringIdxDesc(list []string, indices []int, element string) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// StringIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func StringIdxRng(list []string, indices []int, element string) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := stringIdxRngL(list, indices, element, left, middle-1)
			to := stringIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// StringIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func StringIdxRngDesc(list []string, indices []int, element string) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := stringIdxRngL(list, indices, element, left, middle-1)
			to := stringIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// StringRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func StringRng(list []string, element string) (int, int, bool) {
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
			from := stringRngL(list, element, left, middle-1)
			to := stringRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// StringRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func StringRngDesc(list []string, element string) (int, int, bool) {
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
			from := stringRngL(list, element, left, middle-1)
			to := stringRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// StringIdxOff searches elements in list starting at offset matching element, i.e.
// list[indices[i]][offset:offset+len(element)] == element. Returns start and end index
// in indices, i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in ascending order by indices
// regarding offset.
func StringIdxOff(list []string, indices []int, element string, offset int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	if len(element) > 0 {
		for left <= right {
			middle := (left + right) / 2
			valueIndex := indices[middle]
			value := list[valueIndex]
			if len(value) >= offset {
				value = value[offset:]
				if element > value {
					left = middle + 1
				} else if strings.HasPrefix(value, element) {
					from := stringIdxOffL(list, indices, element, left, middle-1, offset)
					to := stringIdxOffR(list, indices, element, middle+1, right, offset)
					return from, to, true
				} else {
					right = middle - 1
				}
			} else {
				left = middle + 1
			}
		}
	} else if len(list) > 0 {
		return 0, len(list), true
	}
	return left, left + 1, false
}

// StringIdxOffDesc searches elements in list starting at offset matching element, i.e.
// list[indices[i]][offset:offset+len(element)] == element. Returns start and end index
// in indices, i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in descending order by indices
// regarding offset.
func StringIdxOffDesc(list []string, indices []int, element string, offset int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	if len(element) > 0 {
		for left <= right {
			middle := (left + right) / 2
			valueIndex := indices[middle]
			value := list[valueIndex]
			if len(value) >= offset {
				value = value[offset:]
				if element > value {
					right = middle - 1
				} else if strings.HasPrefix(value, element) {
					from := stringIdxOffL(list, indices, element, left, middle-1, offset)
					to := stringIdxOffR(list, indices, element, middle+1, right, offset)
					return from, to, true
				} else {
					left = middle + 1
				}
			} else {
				right = middle - 1
			}
		}
	} else if len(list) > 0 {
		return 0, len(list), true
	}
	return left, left + 1, false
}

// StringOff searches elements in list starting at offset matching element, i.e.
// list[i][offset:offset+len(element)] == element. Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in ascending order
// regarding offset.
func StringOff(list []string, element string, offset int) (int, int, bool) {
	left := 0
	right := len(list) - 1
	if len(element) > 0 {
		for left <= right {
			middle := (left + right) / 2
			value := list[middle]
			if len(value) >= offset {
				value = value[offset:]
				if element > value {
					left = middle + 1
				} else if strings.HasPrefix(value, element) {
					from := stringOffL(list, element, left, middle-1, offset)
					to := stringOffR(list, element, middle+1, right, offset)
					return from, to, true
				} else {
					right = middle - 1
				}
			} else {
				left = middle + 1
			}
		}
	} else if len(list) > 0 {
		return 0, len(list), true
	}
	return left, left + 1, false
}

// StringOff searches elements in list starting at offset matching element, i.e.
// list[i][offset:offset+len(element)] == element. Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in descending order
// regarding offset.
func StringOffDesc(list []string, element string, offset int) (int, int, bool) {
	left := 0
	right := len(list) - 1
	if len(element) > 0 {
		for left <= right {
			middle := (left + right) / 2
			value := list[middle]
			if len(value) >= offset {
				value = value[offset:]
				if element > value {
					right = middle - 1
				} else if strings.HasPrefix(value, element) {
					from := stringOffL(list, element, left, middle-1, offset)
					to := stringOffR(list, element, middle+1, right, offset)
					return from, to, true
				} else {
					left = middle + 1
				}
			} else {
				right = middle - 1
			}
		}
	} else if len(list) > 0 {
		return 0, len(list), true
	}
	return left, left + 1, false
}

// UInt searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// UIntDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UIntDesc(list []uint, element uint) (int, bool) {
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

// UIntIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func UIntIdx(list []uint, indices []int, element uint) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UIntIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func UIntIdxDesc(list []uint, indices []int, element uint) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UIntIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func UIntIdxRng(list []uint, indices []int, element uint) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := uintIdxRngL(list, indices, element, left, middle-1)
			to := uintIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func UIntIdxRngDesc(list []uint, indices []int, element uint) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := uintIdxRngL(list, indices, element, left, middle-1)
			to := uintIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UIntRng(list []uint, element uint) (int, int, bool) {
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
			from := uintRngL(list, element, left, middle-1)
			to := uintRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UIntRngDesc(list []uint, element uint) (int, int, bool) {
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
			from := uintRngL(list, element, left, middle-1)
			to := uintRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt16 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// UInt16Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt16Desc(list []uint16, element uint16) (int, bool) {
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

// UInt16Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func UInt16Idx(list []uint16, indices []int, element uint16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt16IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func UInt16IdxDesc(list []uint16, indices []int, element uint16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt16IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func UInt16IdxRng(list []uint16, indices []int, element uint16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := uint16IdxRngL(list, indices, element, left, middle-1)
			to := uint16IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt16IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func UInt16IdxRngDesc(list []uint16, indices []int, element uint16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := uint16IdxRngL(list, indices, element, left, middle-1)
			to := uint16IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt16Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt16Rng(list []uint16, element uint16) (int, int, bool) {
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
			from := uint16RngL(list, element, left, middle-1)
			to := uint16RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt16RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt16RngDesc(list []uint16, element uint16) (int, int, bool) {
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
			from := uint16RngL(list, element, left, middle-1)
			to := uint16RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt32 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// UInt32Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt32Desc(list []uint32, element uint32) (int, bool) {
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

// UInt32Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func UInt32Idx(list []uint32, indices []int, element uint32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt32IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func UInt32IdxDesc(list []uint32, indices []int, element uint32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt32IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func UInt32IdxRng(list []uint32, indices []int, element uint32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := uint32IdxRngL(list, indices, element, left, middle-1)
			to := uint32IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt32IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func UInt32IdxRngDesc(list []uint32, indices []int, element uint32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := uint32IdxRngL(list, indices, element, left, middle-1)
			to := uint32IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt32Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt32Rng(list []uint32, element uint32) (int, int, bool) {
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
			from := uint32RngL(list, element, left, middle-1)
			to := uint32RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt32RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt32RngDesc(list []uint32, element uint32) (int, int, bool) {
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
			from := uint32RngL(list, element, left, middle-1)
			to := uint32RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt64 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// UInt64Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt64Desc(list []uint64, element uint64) (int, bool) {
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

// UInt64Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func UInt64Idx(list []uint64, indices []int, element uint64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt64IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func UInt64IdxDesc(list []uint64, indices []int, element uint64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt64IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func UInt64IdxRng(list []uint64, indices []int, element uint64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := uint64IdxRngL(list, indices, element, left, middle-1)
			to := uint64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt64IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func UInt64IdxRngDesc(list []uint64, indices []int, element uint64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := uint64IdxRngL(list, indices, element, left, middle-1)
			to := uint64IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt64Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt64Rng(list []uint64, element uint64) (int, int, bool) {
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
			from := uint64RngL(list, element, left, middle-1)
			to := uint64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt64RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt64RngDesc(list []uint64, element uint64) (int, int, bool) {
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
			from := uint64RngL(list, element, left, middle-1)
			to := uint64RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt8 searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// UInt8Desc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt8Desc(list []uint8, element uint8) (int, bool) {
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

// UInt8Idx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func UInt8Idx(list []uint8, indices []int, element uint8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt8IdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func UInt8IdxDesc(list []uint8, indices []int, element uint8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UInt8IdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func UInt8IdxRng(list []uint8, indices []int, element uint8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := uint8IdxRngL(list, indices, element, left, middle-1)
			to := uint8IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt8IdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func UInt8IdxRngDesc(list []uint8, indices []int, element uint8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := uint8IdxRngL(list, indices, element, left, middle-1)
			to := uint8IdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt8Rng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt8Rng(list []uint8, element uint8) (int, int, bool) {
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
			from := uint8RngL(list, element, left, middle-1)
			to := uint8RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt8RngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt8RngDesc(list []uint8, element uint8) (int, int, bool) {
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
			from := uint8RngL(list, element, left, middle-1)
			to := uint8RngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntPtr searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
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

// UIntPtrDesc searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UIntPtrDesc(list []uintptr, element uintptr) (int, bool) {
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

// UIntPtrIdx searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in ascending order by indices.
func UIntPtrIdx(list []uintptr, indices []int, element uintptr) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UIntPtrIdxDesc searches an element in list and returns its index in indices. Second
// return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Indices must reference elements in list,
// elements must be unique and in descending order by indices.
func UIntPtrIdxDesc(list []uintptr, indices []int, element uintptr) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
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

// UIntPtrIdxRng searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order by indices.
func UIntPtrIdxRng(list []uintptr, indices []int, element uintptr) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element > value {
			left = middle + 1
		} else if element < value {
			right = middle - 1
		} else {
			from := uintptrIdxRngL(list, indices, element, left, middle-1)
			to := uintptrIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntPtrIdxRngDesc searches an element in list and returns its start and end index in indices.
// Start index is inclusive, end index is exclusive. Third return value is true, if element
// is in list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order by indices.
func UIntPtrIdxRngDesc(list []uintptr, indices []int, element uintptr) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element < value {
			left = middle + 1
		} else if element > value {
			right = middle - 1
		} else {
			from := uintptrIdxRngL(list, indices, element, left, middle-1)
			to := uintptrIdxRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntPtrRng searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UIntPtrRng(list []uintptr, element uintptr) (int, int, bool) {
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
			from := uintptrRngL(list, element, left, middle-1)
			to := uintptrRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntPtrRngDesc searches an element in list and returns its start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UIntPtrRngDesc(list []uintptr, element uintptr) (int, int, bool) {
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
			from := uintptrRngL(list, element, left, middle-1)
			to := uintptrRngR(list, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}
