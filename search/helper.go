/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package search

import (
	"unsafe"
)

func boolRngL(list []bool, element bool, left, right int) int {
	if element {
		for left <= right {
			middle := (left + right) / 2
			value := list[middle]
			if value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return left
}

func boolRngR(list []bool, element bool, left, right int) int {
	if !element {
		for left <= right {
			middle := (left + right) / 2
			value := list[middle]
			if value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return left
	}
	return right + 1
}

func boolRngRevL(list []bool, element bool, left, right int) int {
	if !element {
		for left <= right {
			middle := (left + right) / 2
			value := list[middle]
			if !value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return left
}

func boolRngRevR(list []bool, element bool, left, right int) int {
	if element {
		for left <= right {
			middle := (left + right) / 2
			value := list[middle]
			if !value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return left
	}
	return right + 1
}

func boolIdxRngL(list []bool, indices []int, element bool, left, right int) int {
	if element {
		for left <= right {
			middle := (left + right) / 2
			valueIndex := indices[middle]
			value := list[valueIndex]
			if value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return left
}

func boolIdxRngR(list []bool, indices []int, element bool, left, right int) int {
	if !element {
		for left <= right {
			middle := (left + right) / 2
			valueIndex := indices[middle]
			value := list[valueIndex]
			if value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return left
	}
	return right + 1
}

func boolIdxRngRevL(list []bool, indices []int, element bool, left, right int) int {
	if !element {
		for left <= right {
			middle := (left + right) / 2
			valueIndex := indices[middle]
			value := list[valueIndex]
			if !value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return left
}

func boolIdxRngRevR(list []bool, indices []int, element bool, left, right int) int {
	if element {
		for left <= right {
			middle := (left + right) / 2
			valueIndex := indices[middle]
			value := list[valueIndex]
			if !value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return left
	}
	return right + 1
}

func byteRngL(list []byte, element byte, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func byteRngR(list []byte, element byte, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func byteIdxRngL(list []byte, indices []int, element byte, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func byteIdxRngR(list []byte, indices []int, element byte, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func complex128RngL(list []complex128, element complex128, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func complex128RngR(list []complex128, element complex128, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func complex128IdxRngL(list []complex128, indices []int, element complex128, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func complex128IdxRngR(list []complex128, indices []int, element complex128, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func complex64RngL(list []complex64, element complex64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func complex64RngR(list []complex64, element complex64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func complex64IdxRngL(list []complex64, indices []int, element complex64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func complex64IdxRngR(list []complex64, indices []int, element complex64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func float32RngL(list []float32, element float32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func float32RngR(list []float32, element float32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func float32IdxRngL(list []float32, indices []int, element float32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func float32IdxRngR(list []float32, indices []int, element float32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func float64RngL(list []float64, element float64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func float64RngR(list []float64, element float64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func float64IdxRngL(list []float64, indices []int, element float64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func float64IdxRngR(list []float64, indices []int, element float64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func intRngL(list []int, element int, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func intRngR(list []int, element int, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func intIdxRngL(list []int, indices []int, element int, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func intIdxRngR(list []int, indices []int, element int, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int16RngL(list []int16, element int16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int16RngR(list []int16, element int16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int16IdxRngL(list []int16, indices []int, element int16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int16IdxRngR(list []int16, indices []int, element int16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int32RngL(list []int32, element int32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int32RngR(list []int32, element int32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int32IdxRngL(list []int32, indices []int, element int32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int32IdxRngR(list []int32, indices []int, element int32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int64RngL(list []int64, element int64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int64RngR(list []int64, element int64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int64IdxRngL(list []int64, indices []int, element int64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int64IdxRngR(list []int64, indices []int, element int64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int8RngL(list []int8, element int8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int8RngR(list []int8, element int8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int8IdxRngL(list []int8, indices []int, element int8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int8IdxRngR(list []int8, indices []int, element int8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func pointerRngL(list []unsafe.Pointer, element unsafe.Pointer, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func pointerRngR(list []unsafe.Pointer, element unsafe.Pointer, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func pointerIdxRngL(list []unsafe.Pointer, indices []int, element unsafe.Pointer, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func pointerIdxRngR(list []unsafe.Pointer, indices []int, element unsafe.Pointer, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func runeRngL(list []rune, element rune, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func runeRngR(list []rune, element rune, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func runeIdxRngL(list []rune, indices []int, element rune, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func runeIdxRngR(list []rune, indices []int, element rune, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func stringRngL(list []string, element string, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func stringRngR(list []string, element string, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func stringIdxRngL(list []string, indices []int, element string, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func stringIdxRngR(list []string, indices []int, element string, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uintRngL(list []uint, element uint, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uintRngR(list []uint, element uint, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uintIdxRngL(list []uint, indices []int, element uint, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uintIdxRngR(list []uint, indices []int, element uint, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint16RngL(list []uint16, element uint16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint16RngR(list []uint16, element uint16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint16IdxRngL(list []uint16, indices []int, element uint16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint16IdxRngR(list []uint16, indices []int, element uint16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint32RngL(list []uint32, element uint32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint32RngR(list []uint32, element uint32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint32IdxRngL(list []uint32, indices []int, element uint32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint32IdxRngR(list []uint32, indices []int, element uint32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint64RngL(list []uint64, element uint64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint64RngR(list []uint64, element uint64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint64IdxRngL(list []uint64, indices []int, element uint64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint64IdxRngR(list []uint64, indices []int, element uint64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint8RngL(list []uint8, element uint8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint8RngR(list []uint8, element uint8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint8IdxRngL(list []uint8, indices []int, element uint8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint8IdxRngR(list []uint8, indices []int, element uint8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uintptrRngL(list []uintptr, element uintptr, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uintptrRngR(list []uintptr, element uintptr, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uintptrIdxRngL(list []uintptr, indices []int, element uintptr, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uintptrIdxRngR(list []uintptr, indices []int, element uintptr, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		valueIndex := indices[middle]
		value := list[valueIndex]
		if element == value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
