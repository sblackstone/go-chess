package movegenerator

import (
	"testing"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
//	"sort"

//  "fmt"
//	"github.com/sblackstone/go-chess/bitopts"

)

func TestGenBishopMovesMiddleOfBoard(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(5, boardstate.WHITE, boardstate.BISHOP)
  expected := []int8{12, 14, 19, 23, 26, 33, 40}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}

func TestGenBishopMovesSq42(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(42, boardstate.WHITE, boardstate.BISHOP)
  expected := []int8{7, 14, 21, 24, 28, 33, 35, 49, 51, 56, 60}
  testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}

func TestGenBishopMovesSq37(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(37, boardstate.WHITE, boardstate.BISHOP)
  expected := []int8{1, 10, 19, 23, 28, 30, 44, 46, 51, 55, 58}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}

func TestGenBishopMovesSqNWCorner(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(56, boardstate.WHITE, boardstate.BISHOP)
	expected := []int8{7, 14, 21, 28,35,42,49}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}

func TestGenBishopMovesSqSECorner(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(7, boardstate.WHITE, boardstate.BISHOP)
	expected := []int8{14, 21, 28,35,42,49,56}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}

func TestTwoBishopsOnBoard(t *testing.T) {
	b := boardstate.Blank()
  b.SetSquare(7, boardstate.WHITE, boardstate.BISHOP)
	b.SetSquare(63, boardstate.WHITE, boardstate.BISHOP)
	expected := []int8{0, 7, 7, 7, 7, 7, 7, 7, 9, 14, 18, 21, 27, 28, 35, 36, 42, 45, 49, 54, 56, 63, 63, 63, 63, 63, 63, 63}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}

func TestGenBishopMovesKnowsAboutTurns(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(7, boardstate.WHITE, boardstate.BISHOP)
	b.SetSquare(63, boardstate.BLACK, boardstate.BISHOP)
	expectedWhite := []int8{14, 21, 28,35,42,49,56}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expectedWhite)

	b.ToggleTurn()

	expectedBlack := []int8{0,9,18,27,36,45,54}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expectedBlack)
}

func TestGenBishopMovesBlockedByOwnPieces(t *testing.T) {
	b := boardstate.Blank()
  b.SetSquare(35, boardstate.WHITE, boardstate.BISHOP)

	b.SetSquare(53, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(49, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(17, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(21, boardstate.WHITE, boardstate.PAWN)
	expected := []int8{26,28,42,44}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}

func TestGenBishopMovesStopsAtCaptures(t *testing.T) {
	b := boardstate.Blank()
  b.SetSquare(35, boardstate.WHITE, boardstate.BISHOP)
	b.SetSquare(53, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(49, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(17, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(21, boardstate.BLACK, boardstate.PAWN)
	expected := []int8{17,21,26,28,42,44,49,53}
	testSuccessorsHelper(t, b, boardstate.BISHOP, expected)
}


func TestGenBishopAttacksBitboard(t *testing.T) {
	t.Errorf("Not implemented")
}
