package fenparser


import (
  "testing"
  "fmt"
  "github.com/sblackstone/go-chess/boardstate"

)

func TestImportTurn(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b, err := FromFEN(testStr)
  if (err != nil || b.GetTurn() != boardstate.WHITE) {
    t.Errorf("Expected turn to be WHITE after import: %v %v", b.GetTurn(), err)
  }

  testStr2 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1"
  b2, err2 := FromFEN(testStr2)
  if (err != nil || b2.GetTurn() != boardstate.BLACK) {
    t.Errorf("Expected turn to be WHITE after import: %v %v", b2.GetTurn(), err2)
  }
}

func TestImportEnpassant(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b, err := FromFEN(testStr)
  if (err != nil || b.GetEnpassant() != boardstate.NO_ENPASSANT) {
    t.Errorf("Expected GetEnpassant to be NO_ENPASSANT after import: %v %v", b.GetEnpassant(), err)
  }

  testStr2 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq c4 0 1"
  b2, err2 := FromFEN(testStr2)
  if (err2 != nil || b2.GetEnpassant() != 26) {
    t.Errorf("Expected GetEnpassant to be 26 after import: %v %v", b2.GetEnpassant(), err2)
  }
}




func TestsMissing(t *testing.T) {
  t.Errorf("caslting on import")
  t.Errorf("caslting on export")
  t.Errorf("enpassasnt on import")
  t.Errorf("enpassasnt on export")
  t.Errorf("halfmove on import")
  t.Errorf("halfmove on export")
  t.Errorf("turn on import")
  t.Errorf("turn on export")
  t.Errorf("complete fen parser tests")
}
func TestFENParserDefaultBoard(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b, err := FromFEN(testStr)
  t.Errorf("TODO: TESTME")

  if err != nil {
    fmt.Printf("%v\n", err)
  } else {
    b.Print(125)
  }



}

func TestFENParserBigXBoard(t *testing.T) {
  testStr := "p6P/1p4P1/2p2P2/3pP3/3Pp3/2P2p2/1P4p1/P6p w - - 25 26"
  b, err := FromFEN(testStr)
  t.Errorf("TODO: TESTME")
  if err != nil {
    fmt.Printf("%v\n", err)
  } else {
    b.Print(125)
  }
  if (b.GetHalfMoves() != 25) {
    t.Errorf("Expected half moves to be 25")
  }
  if (b.GetFullMoves() != 26) {
    t.Errorf("Expected half moves to be 26")
  }


}


func TestToFEN(t *testing.T) {
  b := boardstate.Initial()
  str, _ := ToFEN(b)
  t.Errorf(str)
}
