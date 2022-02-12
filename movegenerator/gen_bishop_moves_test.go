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
  locations := genSortedBoardLocationsBishops(b)
  expected := []uint8{12, 14, 19, 23, 26, 33, 40}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}



func TestGenBishopMovesSq42(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(42, boardstate.WHITE, boardstate.BISHOP)
  locations := genSortedBoardLocationsBishops(b)
  expected := []uint8{7, 14, 21, 24, 28, 33, 35, 49, 51, 56, 60}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestGenBishopMovesSq37(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(37, boardstate.WHITE, boardstate.BISHOP)
  locations := genSortedBoardLocationsBishops(b)
  expected := []uint8{1, 10, 19, 23, 28, 30, 44, 46, 51, 55, 58}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestGenBishopMovesSqNWCorner(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(56, boardstate.WHITE, boardstate.BISHOP)
  locations := genSortedBoardLocationsBishops(b)
  expected := []uint8{7, 14, 21, 28,35,42,49}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestGenBishopMovesSqSECorner(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(7, boardstate.WHITE, boardstate.BISHOP)
  locations := genSortedBoardLocationsBishops(b)
	expected := []uint8{14, 21, 28,35,42,49,56}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestTwoBishopsOnBoard(t *testing.T) {
	t.Errorf("TODO\n");
}

func TestGenBishopMovesKnowsAboutTurns(t *testing.T) {
	t.Errorf("TODO\n")
}

func TestGenBishopMovesBlockedByOwnPieces(t *testing.T) {
	b := boardstate.Blank()
  b.SetSquare(35, boardstate.WHITE, boardstate.BISHOP)

	b.SetSquare(53, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(49, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(17, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(21, boardstate.WHITE, boardstate.PAWN)

  locations := genSortedBoardLocationsBishops(b)
	expected := []uint8{26,28,42,44}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestGenBishopMovesStopsAtCaptures(t *testing.T) {
	b := boardstate.Blank()
  b.SetSquare(35, boardstate.WHITE, boardstate.BISHOP)

	b.SetSquare(53, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(49, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(17, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(21, boardstate.BLACK, boardstate.PAWN)

  locations := genSortedBoardLocationsBishops(b)
	expected := []uint8{17,21,26,28,42,44,49,53}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}
