package boardstate

import (
	"testing"
//  "fmt"
)


func TestCopyPlayMove(t *testing.T) {
	b1 := Initial()
	b2 := b1.CopyPlayTurn(1, 18, EMPTY)
	if (b1 == b2) {
		t.Errorf("Expected b1 to be different than b2")
	}

	if (b2.PieceOfSquare(18) != KNIGHT || b2.ColorOfSquare(18) != WHITE) {
		t.Errorf("square 18 isn't a white knight")
	}

	if (b2.PieceOfSquare(1) != EMPTY || b2.ColorOfSquare(1) != EMPTY) {
		t.Errorf("square 1 isn't empty")
	}

	if (b1.PieceOfSquare(1) == EMPTY || b1.ColorOfSquare(1) == EMPTY) {
		t.Errorf("square 1 on b1 shouldn't be empty")
	}

	if (b1.PieceOfSquare(18) == KNIGHT || b1.ColorOfSquare(18) == WHITE) {
		t.Errorf("square 18 on b1 should be empty")
	}



}

func TestEnpassantAsWhite(t *testing.T) {
	b := Initial()
	if (b.IsEnpassant(2)) {
		t.Errorf("Did not expect enpassant to be set to 2")
	}
	b.PlayTurn(10,26, EMPTY)
	if (!b.IsEnpassant(2)) {
		t.Errorf("Expected enpassant to be set to 2")
	}
	b.PlayTurn(57,42, EMPTY)
	if (b.IsEnpassant(2)) {
		t.Errorf("Expected enpassant to not be set to 2")
	}

}

func TestEnpassantAsBlack(t *testing.T) {
	b := Initial()

	b.PlayTurn(10,26, EMPTY)
	if (b.IsEnpassant(5)) {
		t.Errorf("Did not expect enpassant to be set to 2")
	}
	b.PlayTurn(53,37, EMPTY)
	if (!b.IsEnpassant(5)) {
		t.Errorf("Expected enpassant to be set to 2")
	}
	b.PlayTurn(9,25, EMPTY)
	if (b.IsEnpassant(5)) {
		t.Errorf("Expected enpassant to not be set to 2")
	}

}



func TestEmptySquare(t *testing.T) {
	b := Initial()
	if (b.EmptySquare(10)) {
		t.Errorf("expected square 10 to not be empty\n")
	}
	if (b.EmptySquare(51)) {
		t.Errorf("expected square 51 to not be empty\n")
	}

	if (!b.EmptySquare(27)) {
		t.Errorf("expected square 27 to be empty\n")
	}

}

func TestEmptyOrEnemyOccupiedSquare(t *testing.T) {
	b := Initial()
	if (b.EmptyOrEnemyOccupiedSquare(10)) {
		t.Errorf("expected square 10 to not be empty/enemy occupiped\n")
	}
	if (!b.EmptyOrEnemyOccupiedSquare(18)) {
		t.Errorf("expected square 18 to be empty/enemy occupiped\n")
	}

	if (!b.EmptyOrEnemyOccupiedSquare(50)) {
		t.Errorf("expected square 50 to be empty/enemy occupiped\n")
	}

	b.ToggleTurn()

	if (b.EmptyOrEnemyOccupiedSquare(51)) {
		t.Errorf("expected square 10 to not be empty/enemy occupiped\n")
	}
	if (!b.EmptyOrEnemyOccupiedSquare(43)) {
		t.Errorf("expected square 43 to be empty/enemy occupiped\n")
	}

	if (!b.EmptyOrEnemyOccupiedSquare(10)) {
		t.Errorf("expected square 50 to be empty/enemy occupiped\n")
	}

}

func TestEnemyOccupriedSquare(t *testing.T) {
	b := Initial()
	if (b.EnemyOccupiedSquare(10)) {
		t.Errorf("expected square 10 to not be enemy occupiped\n")
	}
	if (b.EnemyOccupiedSquare(18)) {
		t.Errorf("expected square 18 to not be enemy occupiped\n")
	}

	if (!b.EnemyOccupiedSquare(50)) {
		t.Errorf("expected square 50 to be enemy occupiped\n")
	}

	b.ToggleTurn()

	if (b.EnemyOccupiedSquare(51)) {
		t.Errorf("expected square 10 to not be enemy occupiped\n")
	}
	if (b.EnemyOccupiedSquare(43)) {
		t.Errorf("expected square 43 to not be enemy occupiped\n")
	}

	if (!b.EnemyOccupiedSquare(10)) {
		t.Errorf("expected square 50 to be enemy occupiped\n")
	}



}

func TestCopy(t *testing.T) {
	b1 := Initial()
	b1.PlayTurn(1, 18, EMPTY)

	b2 := b1.Copy()

	if (b2.PieceOfSquare(18) != KNIGHT || b2.ColorOfSquare(18) != WHITE) {
		t.Errorf("square 18 isn't a white knight")
	}

	if (b2.PieceOfSquare(1) != EMPTY || b2.ColorOfSquare(1) != EMPTY) {
		t.Errorf("square 1 isn't empty")
	}

	if b2.GetTurn() != BLACK {
		t.Errorf("Expected turn after PlayTurn to be black");
	}

	b2.PlayTurn(57, 42, EMPTY)

	if b1.PieceOfSquare(42) != EMPTY || b1.ColorOfSquare(42) != EMPTY || b1.GetTurn() != BLACK {
		t.Errorf("copy is effecting the original")
	}

}

func TestPlayTurnPromote(t *testing.T) {
	b := Initial()
	b.PlayTurn(1, 18, QUEEN)

	if (b.PieceOfSquare(18) != QUEEN || b.ColorOfSquare(18) != WHITE) {
		t.Errorf("square 18 isn't a white queen")
	}

}

func TestPlayTurn(t *testing.T) {
	b := Initial()
	if b.GetTurn() != WHITE {
		t.Errorf("Expected initial turn to be white");
	}

	b.PlayTurn(1, 18, EMPTY)

	if (b.PieceOfSquare(18) != KNIGHT || b.ColorOfSquare(18) != WHITE) {
		t.Errorf("square 18 isn't a white knight")
	}

	if (b.PieceOfSquare(1) != EMPTY || b.ColorOfSquare(1) != EMPTY) {
		t.Errorf("square 1 isn't empty")
	}


	if b.GetTurn() != BLACK {
		t.Errorf("Expected turn after PlayTurn to be black");
	}


}

func TestFindPieces(t *testing.T) {
	b := Initial()

	res := b.FindPieces(BLACK, QUEEN)
	if (len(res) != 1 || res[0] != 59) {
		t.Errorf("Expected %v to be [59]\n", res)
	}

	res2 := b.FindPieces(WHITE, QUEEN)
	if (len(res2) != 1 || res2[0] != 3) {
		t.Errorf("Expected %v to be [3]\n", res2)
	}

	res3 := b.FindPieces(BLACK, KNIGHT)
	if (len(res3) != 2 || res3[0] != 57 || res3[1] != 62) {
		t.Errorf("Expected %v to be [57,62]\n", res3)
	}

	res4 := b.FindPieces(WHITE, KNIGHT)
	if (len(res4) != 2 || res4[0] != 1 || res4[1] != 6) {
		t.Errorf("Expected %v to be [1,6]\n", res4)
	}


	res5 := b.FindPieces(WHITE, PAWN)
	if (len(res5) != 8 || res5[0] != 8 || res5[7] != 15) {
		t.Errorf("Expected %v to be [8,9,10,11,12,13,14,15]\n", res5)
	}

	res6 := b.FindPieces(BLACK, PAWN)
	if (len(res6) != 8 || res6[0] != 48 || res6[7] != 55) {
		t.Errorf("Expected %v to be [48,49,50,51,52,53,54,55]\n", res6)
	}

	b2 := Blank()
	res7 := b2.FindPieces(BLACK, PAWN)
	if (len(res7) != 0) {
		t.Errorf("Expected %v to be []\n", res7)
	}

	res8 := b2.FindPieces(WHITE, QUEEN)
	if (len(res8) != 0) {
		t.Errorf("Expected %v to be []\n", res8)
	}


	b3 := Initial()
	b3.SetSquare(57, EMPTY, EMPTY)
	res9 := b3.FindPieces(BLACK, KNIGHT)
	if (len(res9) != 1 || res9[0] != 62) {
		t.Errorf("Expected %v to be [62]\n", res9)
	}


	b4 := Blank()
	b4.SetSquare(55, BLACK, PAWN)
	res10 := b4.FindPieces(BLACK, PAWN)
	if (len(res10) != 1 || res10[0] != 55) {
		t.Errorf("Expected %v to be [55]\n", res10)
	}
}

func TestCreateBlankBoard(t *testing.T) {
  b := Blank();
  if (len(b.colors) != 2) {
    t.Errorf("Colors are not 0,0")
  }
  if (len(b.pieces) != 6) {
    t.Errorf("pieces are not 0,0")
  }
}

func TestMovePieceWhiteKnight(t *testing.T) {
	b := Initial()

	if b.PieceOfSquare(1) != KNIGHT {
		t.Errorf("Exepcted piece in square 1 to be KNIGHT")
	}
	if b.ColorOfSquare(1) != WHITE {
		t.Errorf("Exepcted color in square 1 to be WHITE")
	}

	if b.PieceOfSquare(18) != EMPTY {
		t.Errorf("Exepcted piece in square 18 to be empty")
	}
	if b.ColorOfSquare(18) != EMPTY {
		t.Errorf("Exepcted color in square 18 to be empty")
	}

	b.MovePiece(1, 18)

	if b.PieceOfSquare(1) != EMPTY {
		t.Errorf("Exepcted piece in square 1 to be empty")
	}
	if b.ColorOfSquare(1) != EMPTY {
		t.Errorf("Exepcted color in square 1 to be empty")
	}

	if b.PieceOfSquare(18) != KNIGHT {
		t.Errorf("Exepcted piece in square 18 to be KNIGHT")
	}
	if b.ColorOfSquare(18) != WHITE {
		t.Errorf("Exepcted color in square 18 to be WHITE")
	}
}

func TestMovePieceBlackKnight(t *testing.T) {
	b := Initial()

	if b.PieceOfSquare(57) != KNIGHT {
		t.Errorf("Exepcted piece in square 1 to be KNIGHT")
	}
	if b.ColorOfSquare(57) != BLACK {
		t.Errorf("Exepcted color in square 1 to be BLACK")
	}

	if b.PieceOfSquare(42) != EMPTY {
		t.Errorf("Exepcted piece in square 42 to be empty")
	}
	if b.ColorOfSquare(42) != EMPTY {
		t.Errorf("Exepcted color in square 42 to be empty")
	}

	b.MovePiece(57, 42)

	if b.PieceOfSquare(57) != EMPTY {
		t.Errorf("Exepcted piece in square 1 to be empty")
	}
	if b.ColorOfSquare(57) != EMPTY {
		t.Errorf("Exepcted color in square 1 to be empty")
	}

	if b.PieceOfSquare(42) != KNIGHT {
		t.Errorf("Exepcted piece in square 18 to be KNIGHT")
	}
	if b.ColorOfSquare(42) != BLACK {
		t.Errorf("Exepcted color in square 18 to be BLACK")
	}
}


func testInitialBoard(t *testing.T, b *BoardState) {
	// 64bit int versions of the chess board assuming a correct inital state.
  expectedPieces := []uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
  expectedColors := []uint64{65535,18446462598732840960};

  for i := range(b.pieces) {
    if (b.pieces[i] != expectedPieces[i]) {
      t.Errorf("Initial Board mismatch pieces");
    }
  }

  for i := range(b.colors) {
    if (b.colors[i] != expectedColors[i]) {
      t.Errorf("Initial Board mismatch colors");
    }
  }
}

func TestCreateInitialBoard(t *testing.T) {
	testInitialBoard(t, Initial())
}

func TestCreateInitialBoardManual(t *testing.T) {
	testInitialBoard(t, initialManual())
}


func TestColorOfSquare(t *testing.T) {
  b := Initial()
  if (b.ColorOfSquare(63) != BLACK) {
    t.Errorf("Expected square 63 to be BLACK")
  }
  if (b.ColorOfSquare(0) != WHITE) {
    t.Errorf("Expected square 0 to be WHITE")
  }
  if (b.ColorOfSquare(32) != EMPTY) {
    t.Errorf("Expected square 32 to be EMPTY")
  }
}


func TestPieceOfSquare(t *testing.T) {
  b := Initial()

  expected := [][]uint8{
    {0,ROOK},
    {1,KNIGHT},
    {2,BISHOP},
    {3,QUEEN},
    {4,KING},
    {8,PAWN},
		{32,EMPTY},
		{56,ROOK},
    {57,KNIGHT},
    {58,BISHOP},
    {59,QUEEN},
    {60,KING},
    {48,PAWN},
  }

  for i := range(expected) {
    if b.PieceOfSquare(expected[i][0]) != expected[i][1] {
        t.Errorf("Expected square %v to be %v", expected[i][0], expected[i][1])
    }
  }
}

func TestSetSquare(t *testing.T) {
	b := Initial()

	if (b.ColorOfSquare(0) != WHITE) {
		t.Errorf("Expected color of square 0 to be WHITE (%v), got %v", WHITE, b.ColorOfSquare(0))
	}

	if (b.PieceOfSquare(0) != ROOK) {
		t.Errorf("Expected piece of square 0 to be ROOK (%v), got %v", ROOK, b.PieceOfSquare(0))
	}

	b.SetSquareRankFile(0,0, BLACK, PAWN)

	if (b.ColorOfSquare(0) != BLACK) {
		t.Errorf("Expected color of square 0 to be BLACK")
	}
	if (b.PieceOfSquare(0) != PAWN) {
		t.Errorf("Expected piece of square 0 to be PAWN (%v), got %v", PAWN, b.PieceOfSquare(0))
	}

	b.SetSquareRankFile(0,0, EMPTY, EMPTY)
	if (b.ColorOfSquare(0) != EMPTY) {
		t.Errorf("Expected color of square 0 to be EMPTY")
	}
	if (b.PieceOfSquare(0) != EMPTY) {
		t.Errorf("Expected piece of square 0 to be EMPTY (%v), got %v", EMPTY, b.PieceOfSquare(0))
	}



}
