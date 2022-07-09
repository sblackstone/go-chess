package movegenerator

import (
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

func TestGenRookMovesUnderstandsTurn(t *testing.T) {
	b := boardstate.Blank()

	b.SetSquare(56, boardstate.WHITE, boardstate.ROOK)
	b.SetSquare(7, boardstate.BLACK, boardstate.ROOK)

	expected := []int8{0, 8, 16, 24, 32, 40, 48, 57, 58, 59, 60, 61, 62, 63}

	testSuccessorsHelper(t, b, boardstate.ROOK, expected)
	testAttacksHelper(t, b, boardstate.ROOK, expected)

	b.ToggleTurn()
	expectedBlack := []int8{0, 1, 2, 3, 4, 5, 6, 15, 23, 31, 39, 47, 55, 63}
	testSuccessorsHelper(t, b, boardstate.ROOK, expectedBlack)
	testAttacksHelper(t, b, boardstate.ROOK, expectedBlack)
}

func TestGenRookNWCorner(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(56, boardstate.WHITE, boardstate.ROOK)
	expected := []int8{0, 8, 16, 24, 32, 40, 48, 57, 58, 59, 60, 61, 62, 63}
	testSuccessorsHelper(t, b, boardstate.ROOK, expected)
	testAttacksHelper(t, b, boardstate.ROOK, expected)
}

func TestGenRookSECorner(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(7, boardstate.WHITE, boardstate.ROOK)
	expected := []int8{0, 1, 2, 3, 4, 5, 6, 15, 23, 31, 39, 47, 55, 63}
	testSuccessorsHelper(t, b, boardstate.ROOK, expected)
	testAttacksHelper(t, b, boardstate.ROOK, expected)
}

func TestGenRookBlockedBySelf(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.ROOK)
	b.SetSquare(25, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(30, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(51, boardstate.WHITE, boardstate.KNIGHT)
	b.SetSquare(11, boardstate.WHITE, boardstate.BISHOP)
	expected := []int8{19, 26, 28, 29, 35, 43}
	testSuccessorsHelper(t, b, boardstate.ROOK, expected)
	testAttacksHelper(t, b, boardstate.ROOK, expected)
}

func TestGenRookStopsAtCapture(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.ROOK)
	b.SetSquare(25, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(30, boardstate.BLACK, boardstate.QUEEN)
	b.SetSquare(51, boardstate.BLACK, boardstate.KNIGHT)
	b.SetSquare(11, boardstate.BLACK, boardstate.BISHOP)
	expected := []int8{11, 19, 25, 26, 28, 29, 30, 35, 43, 51}
	testSuccessorsHelper(t, b, boardstate.ROOK, expected)
	testAttacksHelper(t, b, boardstate.ROOK, expected)

}

func TestTwoRooksOnBoard(t *testing.T) {
	b := boardstate.Blank()

	b.SetSquare(56, boardstate.WHITE, boardstate.ROOK)
	b.SetSquare(7, boardstate.WHITE, boardstate.ROOK)

	// We expect 7 and 56 to repeat a bunch as the other rook is still on that square whent he other moves.
	expected := []int8{0, 0, 1, 2, 3, 4, 5, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 15, 16, 23, 24, 31, 32, 39, 40, 47, 48, 55, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 57, 58, 59, 60, 61, 62, 63, 63}
	testSuccessorsHelper(t, b, boardstate.ROOK, expected)

	expectedAttacks := []int8{0, 1, 2, 3, 4, 5, 6, 8, 15, 16, 23, 24, 31, 32, 39, 40, 47, 48, 55, 57, 58, 59, 60, 61, 62, 63}
	testAttacksHelper(t, b, boardstate.ROOK, expectedAttacks)
}

func TestGenRookMovesMiddleOfBoard(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.ROOK)
	expected := []int8{3, 11, 19, 24, 25, 26, 28, 29, 30, 31, 35, 43, 51, 59}
	testSuccessorsHelper(t, b, boardstate.ROOK, expected)
	testAttacksHelper(t, b, boardstate.ROOK, expected)

}
