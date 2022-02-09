package boardstate

import (
	"testing"
//  "fmt"
)


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
