package boardstate

import (
	"testing"
  "fmt"
)


func TestCreateBlankBoard(t *testing.T) {
  b := Blank();
  if (len(b.colors) != 2) {
    t.Errorf("Colors are not 0,0")
  }
  if (len(b.pieces) != 6) {
    t.Errorf("pieces are not 0,0")
  }
}

func TestCreateInitialBoard(t *testing.T) {
  b := Initial()
  // 64bit int versions of the chess board assuming a correct inital state.
  expectedPieces := []uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
  expectedColors := []uint64{18446462598732840960, 65535};

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


func TestCreateInitialBoardManual(t *testing.T) {
	b := InitialManual()
  // 64bit int versions of the chess board assuming a correct inital state.
  expectedPieces := []uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
  expectedColors := []uint64{18446462598732840960, 65535};

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


func TestColorOfSquare(t *testing.T) {
  b := Initial()
  if (b.ColorOfSquare(63) != WHITE) {
    t.Errorf("Expected square 0 to be white")
  }
  if (b.ColorOfSquare(0) != BLACK) {
    t.Errorf("Expected square 63 to be black")
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

	if (b.ColorOfSquare(0) != BLACK) {
		t.Errorf("Expected color of square 0 to be BLACK (%v), got %v", BLACK, b.ColorOfSquare(0))
	}

	if (b.PieceOfSquare(0) != ROOK) {
		t.Errorf("Expected piece of square 0 to be ROOK (%v), got %v", ROOK, b.PieceOfSquare(0))
	}

	b.Print()
	fmt.Print("-----\n")
	b.SetSquareLinear(0,0, BLACK, PAWN)

	b.Print()

	if (b.ColorOfSquare(0) != BLACK) {
		t.Errorf("Expected color of square 0 to be BLACK")
	}
	if (b.PieceOfSquare(0) != PAWN) {
		t.Errorf("Expected piece of square 0 to be PAWN (%v), got %v", PAWN, b.PieceOfSquare(0))
	}


}
