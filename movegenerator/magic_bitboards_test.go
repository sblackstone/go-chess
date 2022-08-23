package movegenerator

import (
	"testing"

	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

// No test coverage for findMagic or GenerateMagicBitboards
// Tests would be redundant with the move generation code already in place
// for rooks and bishops.

func TestBlockersForSquareRook(t *testing.T) {
	rookBlockers := blockersForSquare(0, boardstate.ROOK)

	expected := 4096 // (edges dont count, 6 vertical, 6 horizontal, 2**12 = 4096)
	if len(rookBlockers) != expected {
		t.Errorf("Expected %v to be %v", len(rookBlockers), expected)
	}
}

func TestBlockersForSquareBishop(t *testing.T) {
	rookBlockers := blockersForSquare(0, boardstate.BISHOP)

	expected := 64 // (edges dont count, 6 along diagnol, 2**6 = 64)
	if len(rookBlockers) != expected {
		t.Errorf("Expected %v to be %v", len(rookBlockers), expected)
	}
}

func TestAttackSetForBlockersRook(t *testing.T) {
	var blockerBoard uint64 = bitops.Mask(0) | bitops.Mask(2) | bitops.Mask(16)
	var expected uint64 = bitops.Mask(1) | bitops.Mask(2) | bitops.Mask(8) | bitops.Mask(16)

	rookAttacks := attackSetForBlockers(0, blockerBoard, boardstate.ROOK)

	bitops.Print(rookAttacks, 0)
	bitops.Print(expected, 0)

	if rookAttacks != expected {
		t.Errorf("expected %v to be %v", rookAttacks, expected)
	}

}

func TestAttackSetForBlockersBishop(t *testing.T) {
	var blockerBoard uint64
	blockerBoard = bitops.SetBit(blockerBoard, 0)
	blockerBoard = bitops.SetBit(blockerBoard, 18)

	var expected uint64 = bitops.Mask(9) | bitops.Mask(18)

	bishopAttacks := attackSetForBlockers(0, blockerBoard, boardstate.BISHOP)

	if bishopAttacks != expected {
		t.Errorf("expected %v to be %v", bishopAttacks, expected)
	}

}

func TestPreMaskRook(t *testing.T) {
	var expected uint64
	var i int8
	for i = 0; i < 7; i += 1 {
		expected = bitops.SetBit(expected, i*8)
		expected = bitops.SetBit(expected, i)
	}

	actual := preMask(0, boardstate.ROOK)
	bitops.Print(actual, 0)
	bitops.Print(expected, 0)

	if actual != expected {
		t.Errorf("expected %v to be %v", actual, expected)
	}

}

func TestPreMaskBishop(t *testing.T) {
	var expected uint64
	var i int8
	for i = 0; i < 7; i += 1 {
		expected = bitops.SetBit(expected, i*9)
	}

	actual := preMask(0, boardstate.BISHOP)
	bitops.Print(actual, 0)
	bitops.Print(expected, 0)

	if actual != expected {
		t.Errorf("expected %v to be %v", actual, expected)
	}

}
