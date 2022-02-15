package movegenerator

import (
	"testing"
	"fmt"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
	//"github.com/sblackstone/go-chess/bitopts"

)


func TestPushPawnBlack(t *testing.T) {
  b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(42, boardstate.BLACK, boardstate.PAWN)
	locations := genSortedBoardLocationsPawns(b)
  expected := []uint8{34}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	// Obstructed by SELF
	b.SetSquare(34, boardstate.BLACK, boardstate.QUEEN)
	var expected2 []uint8
	locations2 := genSortedBoardLocationsPawns(b)

  if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2, expected2)
  }

	/// Obstructed by ENEMY
	b.SetSquare(34, boardstate.WHITE, boardstate.QUEEN)
	var expected3 []uint8
	locations3 := genSortedBoardLocationsPawns(b)

  if !reflect.DeepEqual(locations3, expected3) {
    t.Errorf("Expected %v to be %v", locations3, expected3)
  }



}

func TestPushPawnTwoBlack(t *testing.T) {
	// Setup the inital board with a pawn on 8, expect to be pushable 1 or 2 squares.
  b := boardstate.Blank()
	b.ToggleTurn()
  b.SetSquare(49, boardstate.BLACK, boardstate.PAWN)
  expected := []uint8{33,41}
	locations := genSortedBoardLocationsPawns(b)
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	positions := genPawnMoves(b)
	for i := range(positions) {
		if positions[i].PieceOfSquare(33) == boardstate.PAWN {
			if !positions[i].IsEnpassant(1) {
			  t.Errorf("Expected 1 to be enpassant after double push, got %v", positions[i].GetEnpassant())
			}
		} else {
			if positions[i].IsEnpassant(1) {
				t.Errorf("Expected 1 to NOT be enpassant after double push, got %v", positions[i].GetEnpassant())
			}

		}
	}


	var expected2 []uint8

	b.SetSquare(41, boardstate.WHITE, boardstate.QUEEN)
	locations2 := genSortedBoardLocationsPawns(b)


	if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2	, expected2)
  }

	b.SetSquare(41, boardstate.EMPTY, boardstate.EMPTY)
	b.SetSquare(33, boardstate.BLACK, boardstate.QUEEN)

	locations3 := genSortedBoardLocationsPawns(b)
	expected3 := []uint8{41}

	if !reflect.DeepEqual(locations3, expected3) {
    t.Errorf("Expected %v to be %v", locations3	, expected3)
  }

}

func TestCaptureHigherFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(43, boardstate.BLACK, boardstate.PAWN)

	expected := []uint8{35,36}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}

func TestCaptureLowerFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(34, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(43, boardstate.BLACK, boardstate.PAWN)

	expected := []uint8{34,35}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}

func TestCaptureHigherFileBlackWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(14, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(7, boardstate.WHITE, boardstate.KING)
	b.SetSquare(6, boardstate.WHITE, boardstate.QUEEN)

	//b.SetSquare(62, boardstate.BLACK, boardstate.QUEEN)
	boards := genPawnMoves(b)
	var sum uint8
	for i := range(boards) {
		boards[i].Print(255)
		sum += boards[i].PieceOfSquare(7)
	}
	if (sum != 6) {
		t.Errorf("Expected square 7 to have rook,knight,bishop or queen, sum was %v", sum)
	}
}


func TestCaptureLowerFileBlackWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(10, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(1, boardstate.WHITE, boardstate.KING)

	b.SetSquare(2, boardstate.WHITE, boardstate.QUEEN)
	boards := genPawnMoves(b)
	var sum uint8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(1)
	}
	if (sum != 6) {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}



func TestCaptureNoWarpingCapturesHigherFileBlack(t *testing.T) {
	t.Errorf("TODO")

	b := boardstate.Blank()
	b.SetSquare(23, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(30, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(32, boardstate.BLACK, boardstate.PAWN)

	expected := []uint8{30,31}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}

func TestCaptureNoWarpingCapturesLowerFileBlack(t *testing.T) {
	t.Errorf("TODO")

	b := boardstate.Blank()
	b.SetSquare(16, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(23, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(25, boardstate.BLACK, boardstate.PAWN)

	expected := []uint8{24,25}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestCaptureNoSelfCapturesBlack(t *testing.T) {
	t.Errorf("TODO")

	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(34, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(36, boardstate.WHITE, boardstate.QUEEN)
	expected := []uint8{35}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}



func TestPushPawnPromoteBlack(t *testing.T) {
	t.Errorf("TODO")

  b := boardstate.Blank()
  b.SetSquare(49, boardstate.WHITE, boardstate.PAWN)
	boards := genPawnMoves(b)
	var sum uint8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(57)
	}
	if (sum != 6) {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}



func TestEnPassantCaptureAsBlackLowerFile(t *testing.T) {
	t.Errorf("TODO")

	b := boardstate.Blank()
	b.SetSquare(48, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(33, boardstate.WHITE, boardstate.PAWN)

	b.SetTurn(boardstate.BLACK)
	fmt.Println(b.GetEnpassant())
	// Black pushes two setting up enpassant
	b.PlayTurn(48, 32, boardstate.EMPTY)

	locations := genSortedBoardLocationsPawns(b)
	expected := []uint8{40,41}
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestEnPassantCaptureAsBlackHigherFile(t *testing.T) {
	t.Errorf("TODO")

	b := boardstate.Blank()
	b.SetSquare(53, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)

	b.SetTurn(boardstate.BLACK)
	fmt.Println(b.GetEnpassant())
	// Black pushes two setting up enpassant
	b.PlayTurn(53, 37, boardstate.EMPTY)

	locations := genSortedBoardLocationsPawns(b)
	expected := []uint8{44,45}
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestEnPassantCaptureAsBlackUnavailableAfterAdditionalMove(t *testing.T) {
	t.Errorf("TODO")

	b := boardstate.Blank()
	b.SetSquare(48, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(33, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(14, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(55, boardstate.BLACK, boardstate.QUEEN)

	b.SetTurn(boardstate.BLACK)
	fmt.Println(b.GetEnpassant())

	// Black pushes two setting up enpassant
	b.PlayTurn(48, 32, boardstate.EMPTY)

	locations1 := genSortedBoardLocationsPawns(b)
	expected1 := []uint8{40,41}
	if !reflect.DeepEqual(locations1, expected1) {
    t.Errorf("Expected %v to be %v", locations1	, expected1)
  }


	// White pushes something else`
	b.PlayTurn(14, 22, boardstate.EMPTY)

	// Black pushes something else
	b.PlayTurn(55, 47, boardstate.EMPTY)

	locations2 := genSortedBoardLocationsPawns(b)
	expected2 := []uint8{41}
	if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2	, expected2)
  }


	// fmt.Println(b.GetEnpassant())
	// pawnMoves := genPawnMoves(b)
	// for i := range(pawnMoves) {
	// 	fmt.Println()
	// 	pawnMoves[i].Print(255)
	// }
	// t.Errorf("TODO")
}
