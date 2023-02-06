package boardstate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanRemoveSinglePiece(t *testing.T) {
	var pl PieceLocations
	pl.AddPieceLocation(WHITE, PAWN, 2)
	fmt.Printf("%+v\n", pl.pieces[WHITE][PAWN])
	pl.RemovePieceLocation(WHITE, PAWN, 2)
	wpawns := pl.GetLocations(WHITE, PAWN)
	var expected []int8
	assert.ElementsMatch(t, expected, wpawns)
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
			assert.ElementsMatch(t, orig, latest)
		}
	}

	wpawns := plCopy.GetLocations(WHITE, PAWN)
	wpawnsExpected := pl.GetLocations(WHITE, PAWN)

	bpawns := plCopy.GetLocations(BLACK, PAWN)
	bpawnsExpected := pl.GetLocations(BLACK, PAWN)

	assert.ElementsMatch(t, wpawnsExpected, wpawns)
	assert.ElementsMatch(t, bpawnsExpected, bpawns)

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

	assert.ElementsMatch(t, expectedWhitePawns, actualWhitePawns)
	assert.ElementsMatch(t, expectedBlackPawns, actualBlackPawns)
	assert.ElementsMatch(t, expectedWhiteRooks, actualWhiteRooks)
	assert.ElementsMatch(t, expectedBlackRooks, actualBlackRooks)
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
	assert.ElementsMatch(t, expectedWhitePawns, actualWhitePawns)
	assert.ElementsMatch(t, expectedBlackPawns, actualBlackPawns)
	assert.ElementsMatch(t, expectedWhiteRooks, actualWhiteRooks)
	assert.ElementsMatch(t, expectedBlackRooks, actualBlackRooks)

}

func TestRemoveLastPiece(t *testing.T) {
	var pl PieceLocations
	pl.AddPieceLocation(WHITE, PAWN, 1)
	pl.AddPieceLocation(WHITE, PAWN, 2)
	pl.RemovePieceLocation(WHITE, PAWN, 1)
}
