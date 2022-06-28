package movegenerator

import (
  "testing"
	"github.com/sblackstone/go-chess/boardstate"
  //"reflect"
)



func genPerf(state *boardstate.BoardState, depth int) int {
  successors := GenLegalSuccessors(state)
  result := 0
  if (depth == 3) {
    return len(successors)
  } else {
    for _, succ := range(successors) {
      result += genPerf(succ, depth + 1)
    }
  }
  return result
}


func TestPerft(t *testing.T) {
  b := boardstate.Initial()
  t.Errorf("%v\n", genPerf(b, 0))

}
