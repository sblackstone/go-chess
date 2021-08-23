package boardstate


import (
	"testing"
//  "fmt"
)

func TestToggleTurn(t *testing.T) {

  b := Initial()

  if (b.GetTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }

  b.ToggleTurn()

  if (b.GetTurn() != BLACK) {
    t.Errorf("Expected it to be black's turn")
  }

  b.ToggleTurn()

  if (b.GetTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }



}

func TestSetGetTurn(t *testing.T) {
  b := Initial()

  if (b.GetTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }

  b.SetTurn(BLACK)
  if (b.GetTurn() != BLACK) {
    t.Errorf("Expected it to be black's turn")
  }


  b.SetTurn(WHITE)
  if (b.GetTurn() != WHITE) {
    t.Errorf("Expected it to be white's turn")
  }
}
