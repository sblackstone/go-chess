package movegenerator

import (
	"testing"
	//"fmt"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"
)

func TestGenerateChecksOnlyBlack(t *testing.T) {
	b := boardstate.Initial()
	squares := GenAllCheckedSquares(b, boardstate.BLACK)
	expected := []int8{40, 41, 42, 43, 44, 45, 46, 47,}
	var actual []int8
	var i int8
	for i = 0; i < 64; i++ {
		if bitopts.TestBit(squares, i) {
			actual = append(actual, i)
		}
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}


func TestPushPawnBlack(t *testing.T) {
  b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(42, boardstate.BLACK, boardstate.PAWN)
	locations := genSortedBoardLocationsPawns(b)
  expected := []int8{34}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	// Obstructed by SELF
	b.SetSquare(34, boardstate.BLACK, boardstate.QUEEN)
	var expected2 []int8
	locations2 := genSortedBoardLocationsPawns(b)

  if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2, expected2)
  }

	/// Obstructed by ENEMY
	b.SetSquare(34, boardstate.WHITE, boardstate.QUEEN)
	var expected3 []int8
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
  expected := []int8{33,41}
	locations := genSortedBoardLocationsPawns(b)
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	positions := genPawnSuccessors(b)
	for i := range(positions) {
		if positions[i].PieceOfSquare(33) == boardstate.PAWN {
			if positions[i].GetEnpassant() != 41 {
			  t.Errorf("Expected 41 to be enpassant after double push, got %v", positions[i].GetEnpassant())
			}
		} else {
			if positions[i].GetEnpassant() == 41 {
				t.Errorf("Expected 41 to NOT be enpassant after double push, got %v", positions[i].GetEnpassant())
			}

		}
	}


	var expected2 []int8

	b.SetSquare(41, boardstate.WHITE, boardstate.QUEEN)
	locations2 := genSortedBoardLocationsPawns(b)


	if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2	, expected2)
  }

	b.SetSquare(41, boardstate.EMPTY, boardstate.EMPTY)
	b.SetSquare(33, boardstate.BLACK, boardstate.QUEEN)

	locations3 := genSortedBoardLocationsPawns(b)
	expected3 := []int8{41}

	if !reflect.DeepEqual(locations3, expected3) {
    t.Errorf("Expected %v to be %v", locations3	, expected3)
  }

}

func TestCaptureHigherFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(43, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{35,36}
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

	expected := []int8{34,35}
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
	boards := genPawnSuccessors(b)
	var sum int8
	for i := range(boards) {
		//boards[i].Print(255)
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
	boards := genPawnSuccessors(b)
	var sum int8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(1)
	}
	if (sum != 6) {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}



func TestCaptureNoWarpingCapturesHigherFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(31, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(22, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(24, boardstate.WHITE, boardstate.PAWN)

	expected := []int8{22,23}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}

func TestCaptureNoWarpingCapturesLowerFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(24, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(17, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(15, boardstate.WHITE, boardstate.PAWN)

	expected := []int8{16,17}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestCaptureNoSelfCapturesBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(36, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(27, boardstate.BLACK, boardstate.QUEEN)
	b.SetSquare(29, boardstate.BLACK, boardstate.QUEEN)
	expected := []int8{28}
	locations := genSortedBoardLocationsPawns(b)
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}



func TestPushPawnPromoteBlack(t *testing.T) {
  b := boardstate.Blank()
	b.ToggleTurn()
  b.SetSquare(11, boardstate.BLACK, boardstate.PAWN)
	boards := genPawnSuccessors(b)
	var sum int8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(3)
	}
	if (sum != 6) {
		t.Errorf("Expected square 3 to have rook,knight,bishop or queen")
	}
}



func TestEnPassantCaptureAsBlackLowerFile(t *testing.T) {

	b := boardstate.Blank()
	b.SetSquare(26, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(9, boardstate.WHITE, boardstate.PAWN)

	// White pushes two setting up enpassant
	b.PlayTurn(9, 25, boardstate.EMPTY)

	locations := genSortedBoardLocationsPawns(b)
	expected := []int8{17,18}
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestEnPassantCaptureAsBlackHigherFile(t *testing.T) {

	b := boardstate.Blank()
	b.SetSquare(29, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(14, boardstate.WHITE, boardstate.PAWN)

	// White pushes two setting up enpassant
	b.PlayTurn(14, 30, boardstate.EMPTY)

	locations := genSortedBoardLocationsPawns(b)
	expected := []int8{21,22}
	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations	, expected)
  }
}


func TestEnPassantCaptureAsBlackUnavailableAfterAdditionalMove(t *testing.T) {

	b := boardstate.Blank()
	b.SetSquare(30, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(13, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(9, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(57, boardstate.BLACK, boardstate.QUEEN)


	// White pushes two setting up enpassant
	b.PlayTurn(13, 29, boardstate.EMPTY)

	locations1 := genSortedBoardLocationsPawns(b)
	expected1 := []int8{21,22}
	if !reflect.DeepEqual(locations1, expected1) {
    t.Errorf("Expected %v to be %v", locations1	, expected1)
  }


	// Block pushes something else`
	b.PlayTurn(57, 58, boardstate.EMPTY)

	// White pushes something else
	b.PlayTurn(9, 10, boardstate.EMPTY)

	locations2 := genSortedBoardLocationsPawns(b)
	expected2 := []int8{22}
	if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2	, expected2)
  }


	// fmt.Println(b.GetEnpassant())
	// pawnMoves := genPawnSuccessors(b)
	// for i := range(pawnMoves) {
	// 	fmt.Println()
	// 	pawnMoves[i].Print(255)
	// }
	// t.Errorf("TODO")
}


func TestGenPawnAttacksBitboardBlack(t *testing.T) {
	t.Errorf("Not implemented")
}

func TestGenPawnAttacksBitboardBlackEnpassant(t *testing.T) {
	t.Errorf("Not implemented")
}
