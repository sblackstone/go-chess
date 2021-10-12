package movegenerator

import (
	"testing"
  "github.com/sblackstone/go-chess/bitopts"
  "fmt"
)

func TestBlarg(t *testing.T) {
  boards := genKnightMoveBitBoards();
  var pos uint8;
  for pos = 0; pos < 64; pos++ {
    if (boards[pos] > 0) {
      fmt.Println(pos);
      bitopts.Print(boards[pos], pos);
      fmt.Println()
    }
  }
  fmt.Printf("%v\n", boards)
}
