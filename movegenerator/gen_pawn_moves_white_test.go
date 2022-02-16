package movegenerator

import (
	"testing"
	//"fmt"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
	//"github.com/sblackstone/go-chess/bitopts"

)

func TestPushPawnWhite(t *testing.T) {
  b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)
	locations := genSortedBoardLocationsPawns(b)
  expected := []int8{35}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	// Obstructed by SELF
	b.SetSquare(35, boardstate.WHITE, boardstate.QUEEN)
	var expected2 []int8
	locations2 := genSortedBoardLocationsPawns(b)

  if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2, expected2)
  }

	/// Obstructed by ENEMY
	b.SetSquare(35, boardstate.BLACK, boardstate.QUEEN)
	var expected3 []int8
	locations3 := genSortedBoardLocationsPawns(b)

  if !reflect.DeepEqual(locations3, expected3) {
    t.Errorf("Expected %v to be %v", locations3, expected3)
  }



}

func TestPushPawnTwoWhite(t *testing.T) {
	// Setup the inital board with a pawn on 8, expect to be pushable 1 or 2 squares.
  b := boardstate.Blank()
  b.SetSquare(8, boardstate.WHITE, boardstate.PAWN)
  expected := []int8{16,24}
	locations := genSortedBoardLocationsPawns(b)
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	positions := genPawnMoves(b)
	for i := range(positions) {
		if positions[i].PieceOfSquare(24) == boardstate.PAWN {
			if positions[i].GetEnpassant() != 16 {
			  t.Errorf("Expected 16 to be enpassant after double push")
			}
		} else {
			if positions[i].GetEnpassant() == 16 {
			  t.Errorf("Expected 16 to NOT be enpassant after double push")
			}

		}
	}


	var expected2 []int8

	b.SetSquare(16, boardstate.WHITE, boardstate.QUEEN)
	locations2 := genSortedBoardLocationsPawns(b)


	if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2	, expected2)
  }

	b.SetSquare(16, boardstate.EMPTY, boardstate.EMPTY)
	b.SetSquare(24, boardstate.BLACK, boardstate.QUEEN)

	locations3 := genSortedBoardLocationsPawns(b)
	expected3 := []int8{16}

	if !reflect.DeepEqual(locations3, expected3) {
    t.Errorf("Expected %v to be %v", locations3	, expected3)
  }

}

func TestCaptureHigherFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(19, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(28, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{27,28}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}

func TestCaptureLowerFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(19, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(26, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{26,27}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}

func TestCaptureHigherFileWhiteWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(53, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(61, boardstate.BLACK, boardstate.KING)

	b.SetSquare(62, boardstate.BLACK, boardstate.QUEEN)
	boards := genPawnMoves(b)
	var sum int8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(62)
	}
	if (sum != 6) {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}

func TestCaptureLowerFileWhiteWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(53, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(61, boardstate.BLACK, boardstate.KING)

	b.SetSquare(60, boardstate.BLACK, boardstate.QUEEN)
	boards := genPawnMoves(b)
	var sum int8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(60)
	}
	if (sum != 6) {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}



func TestCaptureNoWarpingCapturesHigherFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(23, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(30, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(32, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{30,31}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}

func TestCaptureNoWarpingCapturesLowerFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(16, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(23, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(25, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{24,25}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestCaptureNoSelfCapturesWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(34, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(36, boardstate.WHITE, boardstate.QUEEN)
	expected := []int8{35}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}



func TestPushPawnPromoteWhite(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(49, boardstate.WHITE, boardstate.PAWN)
	boards := genPawnMoves(b)
	var sum int8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(57)
	}
	if (sum != 6) {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}



func TestEnPassantCaptureAsWhiteLowerFile(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(48, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(33, boardstate.WHITE, boardstate.PAWN)

	b.SetTurn(boardstate.BLACK)
	//fmt.Println(b.GetEnpassant())
	// Black pushes two setting up enpassant
	b.PlayTurn(48, 32, boardstate.EMPTY)

	locations := genSortedBoardLocationsPawns(b)
	expected := []int8{40,41}
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestEnPassantCaptureAsWhiteHigherFile(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(53, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)

	b.SetTurn(boardstate.BLACK)
	//fmt.Println(b.GetEnpassant())
	// Black pushes two setting up enpassant
	b.PlayTurn(53, 37, boardstate.EMPTY)

	locations := genSortedBoardLocationsPawns(b)
	expected := []int8{44,45}
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestEnPassantCaptureAsWhiteUnavailableAfterAdditionalMove(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(48, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(33, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(14, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(55, boardstate.BLACK, boardstate.QUEEN)

	b.SetTurn(boardstate.BLACK)
	//fmt.Println(b.GetEnpassant())

	// Black pushes two setting up enpassant
	b.PlayTurn(48, 32, boardstate.EMPTY)

	locations1 := genSortedBoardLocationsPawns(b)
	expected1 := []int8{40,41}
	if !reflect.DeepEqual(locations1, expected1) {
    t.Errorf("Expected %v to be %v", locations1	, expected1)
  }


	// White pushes something else`
	b.PlayTurn(14, 22, boardstate.EMPTY)

	// Black pushes something else
	b.PlayTurn(55, 47, boardstate.EMPTY)

	locations2 := genSortedBoardLocationsPawns(b)
	expected2 := []int8{41}
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
