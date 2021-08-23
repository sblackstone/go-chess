package boardstate

import (
	"testing"
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
  b.Print()

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
