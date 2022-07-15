package boardstate

import (
	"reflect"
	"testing"
)

func TestCanRemoveSinglePiece(t *testing.T) {
	var pl PieceLocations
	pl.AddPieceLocation(WHITE, PAWN, 2)
	pl.RemovePieceLocation(WHITE, PAWN, 2)
	wpawns := pl.GetLocations(WHITE, PAWN)
	expected := make([]int8, 0)
	if !reflect.DeepEqual(wpawns, expected) {
		t.Errorf("Expected %v to be %v", wpawns, expected)
	}

}

func TestCopyPieceLocations(t *testing.T) {
	var pl PieceLocations
	pl.AddPieceLocation(WHITE, PAWN, 1)
	pl.AddPieceLocation(WHITE, PAWN, 2)
	pl.AddPieceLocation(WHITE, PAWN, 3)
	pl.AddPieceLocation(WHITE, ROOK, 4)

	pl.AddPieceLocation(BLACK, PAWN, 8)
	pl.AddPieceLocation(BLACK, PAWN, 9)
	pl.AddPieceLocation(BLACK, PAWN, 10)
	pl.AddPieceLocation(BLACK, ROOK, 12)

	plCopy := pl.Copy()

	var color, piece int8
	for color = WHITE; color <= BLACK; color++ {
		for piece = ROOK; piece <= PAWN; piece++ {
			orig := pl.GetLocations(color, piece)
			latest := plCopy.GetLocations(color, piece)
			if (len(orig) != 0 || len(latest) != 0) && !reflect.DeepEqual(orig, latest) {
				t.Errorf("Expected %v to be %v for %v %v", latest, orig, color, piece)
			}
		}
	}

	wpawns := plCopy.GetLocations(WHITE, PAWN)
	wpawnsExpected := pl.GetLocations(WHITE, PAWN)

	bpawns := plCopy.GetLocations(BLACK, PAWN)
	bpawnsExpected := pl.GetLocations(BLACK, PAWN)

	if !reflect.DeepEqual(wpawns, wpawnsExpected) {
		t.Errorf("Expected %v to be %v", wpawns, wpawnsExpected)
	}
	if !reflect.DeepEqual(bpawns, bpawnsExpected) {
		t.Errorf("Expected %v to be %v", bpawns, bpawnsExpected)
	}

}
func TestPieceLocationsSeparateColors(t *testing.T) {
	var pl PieceLocations
	pl.AddPieceLocation(WHITE, PAWN, 1)
	pl.AddPieceLocation(WHITE, PAWN, 2)
	pl.AddPieceLocation(WHITE, PAWN, 3)
	pl.AddPieceLocation(WHITE, ROOK, 4)

	pl.AddPieceLocation(BLACK, PAWN, 8)
	pl.AddPieceLocation(BLACK, PAWN, 9)
	pl.AddPieceLocation(BLACK, PAWN, 10)
	pl.AddPieceLocation(BLACK, ROOK, 12)

	expectedWhitePawns := []int8{1, 2, 3}
	expectedBlackPawns := []int8{8, 9, 10}
	expectedWhiteRooks := []int8{4}
	expectedBlackRooks := []int8{12}

	actualWhitePawns := pl.GetLocations(WHITE, PAWN)
	actualBlackPawns := pl.GetLocations(BLACK, PAWN)

	actualWhiteRooks := pl.GetLocations(WHITE, ROOK)
	actualBlackRooks := pl.GetLocations(BLACK, ROOK)

	if !reflect.DeepEqual(expectedWhitePawns, actualWhitePawns) {
		t.Errorf("Expected %v to be %v", actualWhitePawns, expectedWhitePawns)
	}

	if !reflect.DeepEqual(expectedBlackPawns, actualBlackPawns) {
		t.Errorf("Expected %v to be %v", actualBlackPawns, expectedBlackPawns)
	}

	if !reflect.DeepEqual(expectedWhiteRooks, actualWhiteRooks) {
		t.Errorf("Expected %v to be %v", actualWhiteRooks, expectedWhiteRooks)
	}

	if !reflect.DeepEqual(expectedBlackRooks, actualBlackRooks) {
		t.Errorf("Expected %v to be %v", actualBlackRooks, expectedBlackRooks)
	}

}

func TestPieceLocationRemoval(t *testing.T) {
	var pl PieceLocations
	pl.AddPieceLocation(WHITE, PAWN, 1)
	pl.AddPieceLocation(WHITE, PAWN, 2)
	pl.AddPieceLocation(WHITE, PAWN, 3)

	pl.AddPieceLocation(BLACK, PAWN, 8)
	pl.AddPieceLocation(BLACK, PAWN, 9)
	pl.AddPieceLocation(BLACK, PAWN, 10)

	pl.AddPieceLocation(BLACK, ROOK, 11)
	pl.AddPieceLocation(WHITE, ROOK, 4)

	pl.RemovePieceLocation(WHITE, PAWN, 3)
	pl.RemovePieceLocation(BLACK, PAWN, 10)
	pl.RemovePieceLocation(WHITE, ROOK, 4)
	pl.RemovePieceLocation(BLACK, ROOK, 11)

	expectedWhitePawns := []int8{1, 2}
	expectedBlackPawns := []int8{8, 9}
	expectedBlackRooks := []int8{}
	expectedWhiteRooks := []int8{}

	actualWhitePawns := pl.GetLocations(WHITE, PAWN)
	actualBlackPawns := pl.GetLocations(BLACK, PAWN)
	actualBlackRooks := pl.GetLocations(BLACK, ROOK)
	actualWhiteRooks := pl.GetLocations(WHITE, ROOK)

	if !reflect.DeepEqual(expectedWhitePawns, actualWhitePawns) {
		t.Errorf("Expected %v to be %v", actualWhitePawns, expectedWhitePawns)
	}

	if !reflect.DeepEqual(expectedBlackPawns, actualBlackPawns) {
		t.Errorf("Expected %v to be %v", actualBlackPawns, expectedBlackPawns)
	}
	if !reflect.DeepEqual(expectedWhiteRooks, actualWhiteRooks) {
		t.Errorf("Expected %v to be %v", actualWhiteRooks, expectedWhiteRooks)
	}

	if !reflect.DeepEqual(expectedBlackRooks, actualBlackRooks) {
		t.Errorf("Expected %v to be %v", actualBlackRooks, expectedBlackRooks)
	}
}
