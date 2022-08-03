package bitops

import (
	"reflect"
	"testing"
)

func TestFindSetBitsGeneric(t *testing.T) {
	var result []int8
	var testCase uint64

	f := func(n int8) {
		result = append(result, n)
	}

	testCase = SetBit(testCase, 5)
	testCase = SetBit(testCase, 7)
	testCase = SetBit(testCase, 63)

	FindSetBitsGeneric(testCase, f)
	expected := []int8{5, 7, 63}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to be %v", result, expected)
	}

}

func TestInternalMask(t *testing.T) {
	var expected uint64
	var i, j int8
	for i = 1; i < 7; i++ {
		for j = 1; j < 7; j++ {
			expected = SetBit(expected, RankFileToSquare(i, j))
		}
	}

	im := InternalMask()
	if im != expected {
		t.Errorf("Expected %b to be %b", im, expected)
	}
}
func TestPerimeterMask(t *testing.T) {
	expected := RankMask(0) | RankMask(7) | FileMask(0) | FileMask(7)
	if PerimeterMask() != expected {
		t.Errorf("Expected %b to be %b", PerimeterMask(), expected)
	}
}
func TestRankMasks(t *testing.T) {
	var i int8
	for i = 0; i < 8; i++ {
		expected := uint64(255) << (i * 8)
		if RankMask(i) != expected {
			t.Errorf("Expected %b to be %b", RankMask(i), expected)
		}
	}
}

func TestFileMasks(t *testing.T) {
	var i int8
	for i = 0; i < 8; i++ {
		expected := Rotate90Clockwise(RankMask(i))
		if FileMask(i) != expected {
			t.Errorf("Expected %b to be %b", FileMask(i), expected)
		}
	}
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
