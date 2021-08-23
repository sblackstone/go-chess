package boardstate


import (
	"testing"
//  "fmt"
)

func TestToggleTurn(t *testing.T) {

  b := Initial()

  if (b.getTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }

  b.toggleTurn()

  if (b.getTurn() != BLACK) {
    t.Errorf("Expected it to be black's turn")
  }

  b.toggleTurn()

  if (b.getTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }



}

func TestSetGetTurn(t *testing.T) {
  b := Initial()

  if (b.getTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }

  b.setTurn(BLACK)
  if (b.getTurn() != BLACK) {
    t.Errorf("Expected it to be black's turn")
  }


  b.setTurn(WHITE)
  if (b.getTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }
}
