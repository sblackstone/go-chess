package treesearch

import (
  "testing"
  "github.com/sblackstone/go-chess/boardstate"

)


func TestBestSuccessor(t *testing.T) {
    b := boardstate.Initial()
    m1 := BestSuccessor(b,5)
    m1.Print(127)

    m1.PlayTurn(49,33, boardstate.EMPTY)
    m1.Print(127)
    m2 := BestSuccessor(m1,5)
    m2.Print(127)



}
