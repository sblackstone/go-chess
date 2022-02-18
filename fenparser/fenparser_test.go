package fenparser


import (
  "testing"
  "fmt"
  "github.com/sblackstone/go-chess/boardstate"

)


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
