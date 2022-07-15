package bitops

import (
	"testing"
)

func TestFlipDiagA1H8(t *testing.T) {

	var x, expected uint64
	x = SetBit(x, 14)
	x = SetBit(x, 15)
	x = SetBit(x, 6)
	x = SetBit(x, 7)

	r := FlipDiagA1H8(x)

	expected = SetBit(expected, 56)
	expected = SetBit(expected, 57)
	expected = SetBit(expected, 48)
	expected = SetBit(expected, 49)

	if r != expected {
		t.Errorf("Expected %b to be %b", r, expected)
	}
}

func TestFlipVertical(t *testing.T) {
	x := RankMask(5) | RankMask(6)
	expected := RankMask(1) | RankMask(2)
	r := FlipVertical(x)
	if r != expected {
		t.Errorf("Expected %b to be %b", r, expected)
	}

}

func TestRotate90Clockwise(t *testing.T) {
	var x, expected uint64

	x = SetBit(x, 56)
	x = SetBit(x, 49)

	expected = SetBit(expected, 63)
	expected = SetBit(expected, 54)

	r := Rotate90Clockwise(x)
	if r != expected {
		t.Errorf("Expected %b to be %b", r, expected)
	}

}

func TestRotate90AntiClockwise(t *testing.T) {
	var x, expected uint64

	x = SetBit(x, 63)
	x = SetBit(x, 54)

	expected = SetBit(expected, 56)
	expected = SetBit(expected, 49)

	r := Rotate90AntiClockwise(x)
	if r != expected {
		t.Errorf("Expected %b to be %b", r, expected)
	}
}
