package movegenerator

import (
  "testing"
	"github.com/sblackstone/go-chess/boardstate"
  //"fmt"
)



func TestGenSucessorsInitialPosition(t *testing.T) {
  b := boardstate.Initial()
  successors := GenSucessors(b)

  if len(successors) != 20 {
    t.Errorf("Expected initial successors to be 20, got %v", len(successors))
  }

  b.PlayTurn(8, 24, boardstate.EMPTY)


  successors2 := GenSucessors(b)

  if len(successors2) != 20 {
    t.Errorf("Expected initial successors to be 20, got %v", len(successors2))
    for i := range(successors2) {
      successors2[i].Print(127)
    }
  }


}
/*
func TestGenAllCheckedSquaresInitialPosition(t *testing.T) {
  b := boardstate.Initial()
  squares := genSortedCheckedSquares(b, boardstate.WHITE)

  fmt.Printf("%v\n", squares)

}
*/
/*

func TestGenAllMovesInitialPosition(t *testing.T) {
    b := boardstate.Initial()
    moves1 := GenAllMoves(b)

    if len(moves1) != 20 {
      t.Errorf("Expected initial successors to be 20, got %v", len(moves1))
    }

    b.PlayTurn(8, 24, boardstate.EMPTY)


    moves2 := GenAllMoves(b)

    if len(moves2) != 20 {
      t.Errorf("Expected initial successors to be 20, got %v", len(moves2))
      for i := range(moves2) {
        t.Errorf("Move %v", moves2[i])
      }
    }

}
*/
