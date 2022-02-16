package movegenerator

import (
  "testing"
	"github.com/sblackstone/go-chess/boardstate"
)



func TestMoveGeneratorInitialPosition(t *testing.T) {
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
