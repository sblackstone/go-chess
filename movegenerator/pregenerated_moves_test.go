package movegenerator

import (
	"testing"
  "fmt"
	"github.com/sblackstone/go-chess/bitopts"

)

func TestBlarg(t *testing.T) {
  boards := genAllKnightMoves();
  var pos uint8;
	var tmp uint64
	var i int
  for pos = 0; pos < 64; pos++ {
		tmp = 0
		for i = range(boards[pos]) {
			tmp = bitopts.SetBit(tmp, boards[pos][i])
		}
		fmt.Printf("%v: %v\n", pos, boards[pos]);
		bitopts.Print(tmp, pos)
  }
}
