/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package search

import (
	"testing"
)

func TestBool(t *testing.T) {
	var list []bool
	index, found := Bool(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, false)
	index, found = Bool(list, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Bool(list, true)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	list = append(list, true)
	index, found = Bool(list, true)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	list = list[1:]
	index, found = Bool(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Bool(list, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestBoolDesc(t *testing.T) {
	var list []bool
	index, found := BoolDesc(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, true)
	index, found = BoolDesc(list, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = BoolDesc(list, false)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	list = append(list, false)
	index, found = BoolDesc(list, false)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	list = list[1:]
	index, found = BoolDesc(list, true)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = BoolDesc(list, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestByteA(t *testing.T) {
	var list []byte
	index, found := Byte(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40)
	index, found = Byte(list, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteB(t *testing.T) {
	var list []byte
	index, found := Byte(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40, 50)
	index, found = Byte(list, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteDescA(t *testing.T) {
	var list []byte
	index, found := ByteDesc(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10)
	index, found = ByteDesc(list, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteDescB(t *testing.T) {
	var list []byte
	index, found := ByteDesc(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10, 0)
	index, found = ByteDesc(list, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128(t *testing.T) {
	var list []complex128
	index, found := Complex128(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 10+5i, 20, 20+5i)
	index, found = Complex128(list, 5+5i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, 10+5i)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, 10+10i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, 20+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, 20+5i)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128Desc(t *testing.T) {
	var list []complex128
	index, found := Complex128Desc(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 20+5i, 20, 10+5i, 10)
	index, found = Complex128Desc(list, 20+10i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 20+5i)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 15+5i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 10+5i)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 10+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestFloat32(t *testing.T) {
	var list []float32
	index, found := Float32(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40)
	index, found = Float32(list, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestBoolRngL(t *testing.T) {
	list := []bool{false}
	from := boolRngL(list, false, 0, 0)
	if from != 0 {
		t.Error(from)
	}
	list = append(list, false, false, false, false, true, true)
	from = boolRngL(list, false, 0, 3)
	if from != 0 {
		t.Error(from)
	}
	from = boolRngL(list, true, 0, 5)
	if from != 5 {
		t.Error(from)
	}
}

func TestBoolRngR(t *testing.T) {
	list := []bool{false}
	from := boolRngR(list, false, 0, 0)
	if from != 1 {
		t.Error(from)
	}
	list = append(list, false, false, false, false, true, true, true)
	from = boolRngR(list, false, 2, 6)
	if from != 5 {
		t.Error(from)
	}
	from = boolRngR(list, true, 6, 6)
	if from != 7 {
		t.Error(from)
	}
	list = append(list, true)
	from = boolRngR(list, true, 6, 7)
	if from != 8 {
		t.Error(from)
	}
}

func TestBoolRng(t *testing.T) {
	list := []bool{false, false, false, false, false, true, true, true}
	from, to, found := BoolRng(list, false)
	if from != 0 || to != 5 || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRng(list, true)
	if from != 5 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	list[5], list[6], list[7] = false, false, false
	from, to, found = BoolRng(list, false)
	if from != 0 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRng(list, true)
	if from != len(list) || to != len(list)+1 || found {
		t.Error(from, to, found)
	}
}

func TestBoolRngDesc(t *testing.T) {
	list := []bool{true, true, true, true, true, false, false, false}
	from, to, found := BoolRngDesc(list, true)
	if from != 0 || to != 5 || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRngDesc(list, false)
	if from != 5 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	list[5], list[6], list[7] = true, true, true
	from, to, found = BoolRngDesc(list, true)
	if from != 0 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRngDesc(list, false)
	if from != len(list) || to != len(list)+1 || found {
		t.Error(from, to, found)
	}
}

func TestBoolIdx(t *testing.T) {
	list := []bool{false, true, true, false}
	var indices []int
	index, found := BoolIdx(list, indices, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 3)
	index, found = BoolIdx(list, indices, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = BoolIdx(list, indices, true)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 1)
	index, found = BoolIdx(list, indices, true)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	indices = indices[1:]
	index, found = BoolIdx(list, indices, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = BoolIdx(list, indices, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestBoolIdxDesc(t *testing.T) {
	list := []bool{false, true, true, false}
	var indices []int
	index, found := BoolIdxDesc(list, indices, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 1)
	index, found = BoolIdxDesc(list, indices, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = BoolIdxDesc(list, indices, false)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 3)
	index, found = BoolIdxDesc(list, indices, false)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	indices = indices[1:]
	index, found = BoolIdxDesc(list, indices, true)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = BoolIdxDesc(list, indices, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestByteIdxA(t *testing.T) {
	list := []byte{10, 20, 30, 40}
	var indices []int
	index, found := ByteIdx(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 0, 1, 2, 3)
	index, found = ByteIdx(list, indices, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteIdxB(t *testing.T) {
	list := []byte{10, 20, 30, 40, 50}
	var indices []int
	index, found := ByteIdx(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = ByteIdx(list, indices, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdx(list, indices, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteIdxDescA(t *testing.T) {
	list := []byte{10, 20, 30, 40, 50}
	var indices []int
	index, found := ByteIdxDesc(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 3, 2, 1, 0)
	index, found = ByteIdxDesc(list, indices, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteIdxDescB(t *testing.T) {
	list := []byte{10, 20, 30, 40, 0}
	var indices []int
	index, found := ByteIdxDesc(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 3, 2, 1, 0, 4)
	index, found = ByteIdxDesc(list, indices, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteIdxDesc(list, indices, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128Idx(t *testing.T) {
	list := []complex128{10, 10 + 5i, 20, 20 + 5i}
	var indices []int
	index, found := Complex128Idx(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 0, 1, 2, 3)
	index, found = Complex128Idx(list, indices, 5+5i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Idx(list, indices, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Idx(list, indices, 10+5i)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Idx(list, indices, 10+10i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Idx(list, indices, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Idx(list, indices, 20+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Idx(list, indices, 20+5i)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128IdxDesc(t *testing.T) {
	list := []complex128{10, 10 + 5i, 20, 20 + 5i}
	var indices []int
	index, found := Complex128IdxDesc(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	indices = append(indices, 3, 2, 1, 0)
	index, found = Complex128IdxDesc(list, indices, 20+10i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128IdxDesc(list, indices, 20+5i)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128IdxDesc(list, indices, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128IdxDesc(list, indices, 15+5i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128IdxDesc(list, indices, 10+5i)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128IdxDesc(list, indices, 10+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128IdxDesc(list, indices, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteRngL(t *testing.T) {
	list := []byte{50}
	from := byteRngL(list, 50, 0, 0)
	if from != 0 {
		t.Error(from)
	}
	list = append(list, 50, 60, 60, 60, 70, 70, 70, 80)
	from = byteRngL(list, 60, 0, 3)
	if from != 2 {
		t.Error(from)
	}
	from = byteRngL(list, 70, 2, 6)
	if from != 5 {
		t.Error(from)
	}
}

func TestByteRngR(t *testing.T) {
	list := []byte{50}
	from := byteRngR(list, 50, 0, 0)
	if from != 1 {
		t.Error(from)
	}
	list = append(list, 50, 60, 60, 60, 70, 70, 70, 80)
	from = byteRngR(list, 60, 2, 8)
	if from != 5 {
		t.Error(from)
	}
	from = byteRngR(list, 70, 5, 8)
	if from != 8 {
		t.Error(from)
	}
}

func TestByteRng(t *testing.T) {
	list := []byte{50, 50, 54, 60, 60, 60, 60, 70, 70, 70, 80}
	from, to, found := ByteRng(list, 50)
	if from != 0 || to != 2 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 54)
	if from != 2 || to != 3 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 60)
	if from != 3 || to != 7 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 70)
	if from != 7 || to != 10 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 80)
	if from != 10 || to != 11 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 100)
	if from != 11 || to != 12 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 55)
	if from != 3 || to != 4 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 65)
	if from != 7 || to != 8 || found {
		t.Error(from, to, found)
	}
}

func TestByteRngDesc(t *testing.T) {
	list := []byte{80, 80, 79, 70, 70, 70, 70, 60, 60, 60, 50}
	from, to, found := ByteRngDesc(list, 80)
	if from != 0 || to != 2 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 79)
	if from != 2 || to != 3 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 70)
	if from != 3 || to != 7 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 60)
	if from != 7 || to != 10 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 50)
	if from != 10 || to != 11 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 40)
	if from != 11 || to != 12 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 75)
	if from != 3 || to != 4 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 65)
	if from != 7 || to != 8 || found {
		t.Error(from, to, found)
	}
}

func TestByteIdxRngA(t *testing.T) {
	list := []byte{50, 50, 54, 60, 60, 60, 60, 70, 70, 70, 80}
	var indices []int
	from, to, found := ByteIdxRng(list, indices, 50)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	indices = append(indices, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	from, to, found = ByteIdxRng(list, indices, 40)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 50)
	if from != 0 || to != 2 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 54)
	if from != 2 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 56)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 60)
	if from != 3 || to != 7 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 65)
	if from != 7 || to != 8 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 80)
	if from != 10 || to != 11 || found != true {
		t.Error(from, to, found)
	}
}

func TestByteIdxRngB(t *testing.T) {
	list := []byte{50, 50, 54, 60, 60, 60, 60, 70, 70, 70, 80, 90}
	var indices []int
	from, to, found := ByteIdxRng(list, indices, 50)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	indices = append(indices, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	from, to, found = ByteIdxRng(list, indices, 40)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 50)
	if from != 0 || to != 2 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 54)
	if from != 2 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 56)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 60)
	if from != 3 || to != 7 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 65)
	if from != 7 || to != 8 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRng(list, indices, 90)
	if from != 11 || to != 12 || found != true {
		t.Error(from, to, found)
	}
}

func TestByteIdxRngDescA(t *testing.T) {
	list := []byte{50, 50, 54, 60, 60, 60, 60, 70, 70, 70, 80}
	var indices []int
	from, to, found := ByteIdxRngDesc(list, indices, 80)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	indices = append(indices, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0)
	from, to, found = ByteIdxRngDesc(list, indices, 90)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 80)
	if from != 0 || to != 1 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 70)
	if from != 1 || to != 4 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 65)
	if from != 4 || to != 5 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 60)
	if from != 4 || to != 8 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 54)
	if from != 8 || to != 9 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 50)
	if from != 9 || to != 11 || found != true {
		t.Error(from, to, found)
	}
}

func TestByteIdxRngDescB(t *testing.T) {
	list := []byte{50, 50, 54, 60, 60, 60, 60, 70, 70, 70, 80, 90}
	var indices []int
	from, to, found := ByteIdxRngDesc(list, indices, 80)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	indices = append(indices, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0)
	from, to, found = ByteIdxRngDesc(list, indices, 95)
	if from != 0 || to != 1 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 80)
	if from != 1 || to != 2 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 70)
	if from != 2 || to != 5 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 65)
	if from != 5 || to != 6 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 60)
	if from != 5 || to != 9 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 54)
	if from != 9 || to != 10 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = ByteIdxRngDesc(list, indices, 50)
	if from != 10 || to != 12 || found != true {
		t.Error(from, to, found)
	}
}

func TestStringOff(t *testing.T) {
	var list []string
	from, to, found := StringOff(list, "asdf", 0)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	from, to, found = StringOff(list, "asdf", 10)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	from, to, found = StringOff(list, "a", 0)
	if from != 1 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringOff(list, "b", 0)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	from, to, found = StringOff(list, "a", 1)
	if from != 1 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringOff(list, "b", 1)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringOff(list, "", 1)
	if from != 0 || to != len(list) || found != true {
		t.Error(from, to, found)
	}
}

func TestStringIdxOffDesc(t *testing.T) {
	list := []string{"", "aa", "ab", "dc", "dd"}
	var indices []int
	from, to, found := StringIdxOffDesc(list, indices, "asdf", 0)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	from, to, found = StringIdxOffDesc(list, indices, "asdf", 10)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	indices = append(indices, 4, 3, 2, 1, 0)
	from, to, found = StringIdxOffDesc(list, indices, "a", 0)
	if from != 2 || to != 4 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringIdxOffDesc(list, indices, "b", 0)
	if from != 2 || to != 3 || found != false {
		t.Error(from, to, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	from, to, found = StringIdxOffDesc(list, indices, "a", 1)
	if from != 2 || to != 4 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringIdxOffDesc(list, indices, "b", 1)
	if from != 2 || to != 3 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringIdxOffDesc(list, indices, "", 1)
	if from != 0 || to != len(indices) || found != true {
		t.Error(from, to, found)
	}
}
