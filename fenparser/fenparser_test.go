package fenparser


import (
  "testing"
  "fmt"
)


func TestFENParserDefaultBoard(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b, err := FromFEN(testStr)
  if err != nil {
    fmt.Printf("%v\n", err)
  } else {
    b.Print(125)
  }
}

func TestFENParserBigXBoard(t *testing.T) {
  testStr := "p6P/1p4P1/2p2P2/3pP3/3Pp3/2P2p2/1P4p1/P6p w - - 0 1"
  b, err := FromFEN(testStr)
  if err != nil {
    fmt.Printf("%v\n", err)
  } else {
    b.Print(125)
  }
}
