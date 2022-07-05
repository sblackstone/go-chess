package treesearch

import (
  "testing"
  "github.com/sblackstone/go-chess/boardstate"
)

func TestBestSuccessor1(t *testing.T) {
  b := boardstate.Initial()
  b.PlayTurn(8, 24, boardstate.EMPTY)
  b.PlayTurn(48, 32, boardstate.EMPTY)
  b.Print(127)
  bs := BestSuccessor(b,2)
  bs.Print(127)
}


// func TestPlayItself(t *testing.T) {
//     b := boardstate.Initial()
//
//     for {
//       b = BestSuccessor(b,5)
//       b.Print(127)
//     }
//     // m1 := BestSuccessor(b,5)
//     // m1.Print(127)
//     //
//     // m1.PlayTurn(49,33, boardstate.EMPTY)
//     // m1.Print(127)
//     // m2 := BestSuccessor(m1,5)
//     // m2.Print(127)
//
//
//
// }
