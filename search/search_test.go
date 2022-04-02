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
	list[0] = true
	list = list[:1]
	index, found = Bool(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Bool(list, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestBoolRev(t *testing.T) {
	var list []bool
	index, found := BoolRev(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, true)
	index, found = BoolRev(list, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = BoolRev(list, false)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	list = append(list, false)
	index, found = BoolRev(list, false)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	list[0] = false
	list = list[:1]
	index, found = BoolRev(list, true)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = BoolRev(list, false)
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

func TestByteRevA(t *testing.T) {
	var list []byte
	index, found := ByteRev(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10)
	index, found = ByteRev(list, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteRevB(t *testing.T) {
	var list []byte
	index, found := ByteRev(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10, 0)
	index, found = ByteRev(list, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteRev(list, 10)
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

func TestComplex128Rev(t *testing.T) {
	var list []complex128
	index, found := Complex128Rev(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 20+5i, 20, 10+5i, 10)
	index, found = Complex128Rev(list, 20+10i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Rev(list, 20+5i)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Rev(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Rev(list, 15+5i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Rev(list, 10+5i)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Rev(list, 10+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Rev(list, 10)
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
