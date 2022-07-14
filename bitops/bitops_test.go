package bitops

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPerimeterMask(t *testing.T) {
	Print(PerimeterMask(), 127)
	fmt.Printf("\n\n\n")
	t.Errorf("not implemented")
}
func TestRankMasks(t *testing.T) {
	var i int8
	for i = 0; i < 8; i++ {
		fmt.Printf("\n\n\ni=%v\n", i)
		Print(rankMasks[i], 127)
	}
	t.Errorf("not implemented")
}

func TestFileMasks(t *testing.T) {
	var i int8
	for i = 0; i < 8; i++ {
		fmt.Printf("\n\n\ni=%v\n", i)
		Print(fileMasks[i], 127)
	}
	t.Errorf("not implemented")
}
func TestMask(t *testing.T) {
	v := Mask(1)
	if v != 2 {
		t.Errorf("expected Mask(1) to be 2")
	}
}

func TestFindSetBits(t *testing.T) {
	var v uint64
	v = SetBit(v, 2)
	v = SetBit(v, 4)
	v = SetBit(v, 6)
	result := FindSetBits(v)
	expected := []int8{2, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to be %v\n", result, expected)
	}
}

func TestFindSetBitsNone(t *testing.T) {
	var v uint64
	result := FindSetBits(v)
	var expected []int8
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to be %v\n", result, expected)
	}
}

func TestFindSetBitsExtrema(t *testing.T) {
	var v uint64
	v = SetBit(v, 0)
	v = SetBit(v, 63)
	result := FindSetBits(v)
	expected := []int8{0, 63}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to be %v\n", result, expected)
	}
}

func TestPrint(t *testing.T) {
	var val uint64 = 5
	Print(val, 52)
}

func TestRankFileToSquare(t *testing.T) {
	cases := [][3]int8{
		{0, 0, 0},
		{7, 7, 63},
		{7, 3, 59},
		{3, 7, 31},
	}

	for i := range cases {
		v := RankFileToSquare(cases[i][0], cases[i][1])
		if v != cases[i][2] {
			t.Errorf("Expected (%v,%v) to be %v, got %v", cases[i][0], cases[i][1], cases[i][2], v)
		}
	}
}

func TestSquareToRankFile(t *testing.T) {
	cases := [][3]int8{
		{0, 0, 0},
		{7, 7, 63},
		{7, 3, 59},
		{3, 7, 31},
	}

	for i := range cases {
		row, col := SquareToRankFile(cases[i][2])
		if row != cases[i][0] || col != cases[i][1] {
			t.Errorf("Expected %v to be (%v,%v), got (%v,%v)", cases[i][2], cases[i][0], cases[i][1], row, col)
		}
	}
}

func TestFlipBit(t *testing.T) {
	var v uint64

	v = 0b101

	v = FlipBit(v, 1)

	if v != 0b111 {
		t.Errorf("Expected 101 ^ 010 = 111 (7)")
	}

	v = FlipBit(v, 1)

	if v != 0b101 {
		t.Errorf("Expected 111 ^ 010 = 101 (5)")
	}

}

func TestSetBit(t *testing.T) {
	var v uint64
	v = SetBit(v, 0)
	if v != 1 {
		t.Errorf("%d should have been 1", v)
	}
	v = SetBit(v, 1)
	if v != 3 {
		t.Errorf("%d should have been 3", v)
	}
}

func TestClearBit(t *testing.T) {
	var v uint64 = 7
	v = ClearBit(v, 0)
	if v != 6 {
		t.Errorf("%d should have been 6", v)
	}
	v = ClearBit(v, 1)
	if v != 4 {
		t.Errorf("%d should have been 4", v)
	}
}

func TestTestBit(t *testing.T) {
	var v uint64 = 5
	if !TestBit(v, 0) {
		t.Errorf("Expected first bit to be true")
	}
	if TestBit(v, 1) {
		t.Errorf("Expected first bit to be false")
	}
	if !TestBit(v, 2) {
		t.Errorf("Expected first bit to be true")
	}

}
