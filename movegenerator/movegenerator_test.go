package movegenerator

import (
  "testing"
	"github.com/sblackstone/go-chess/boardstate"
  "reflect"
)



func TestGenGenSuccessorsInitialPosition(t *testing.T) {
  b := boardstate.Initial()
  successors := GenSuccessors(b)

  if len(successors) != 20 {
    t.Errorf("Expected initial successors to be 20, got %v", len(successors))
  }

  b.PlayTurn(8, 24, boardstate.EMPTY)


  successors2 := GenSuccessors(b)

  if len(successors2) != 20 {
    t.Errorf("Expected initial successors to be 20, got %v", len(successors2))
    for i := range(successors2) {
      successors2[i].Print(127)
    }
  }


}


func TestGenLegalSuccessorsOpposition(t *testing.T) {
  b := boardstate.Blank()
  b.ClearCastling()
	b.SetSquare(43, boardstate.WHITE, boardstate.KING)
  b.SetSquare(27, boardstate.BLACK, boardstate.KING)

  legalWhite := genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KING, GenLegalSuccessors(b))

  expectedWhite := []int8{42, 44, 50, 51, 52}
  if !reflect.DeepEqual(legalWhite, expectedWhite) {
    t.Errorf("Expected %v to be %v", legalWhite, expectedWhite)
  }

  b.SetTurn(boardstate.BLACK)

  legalBlack := genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KING, GenLegalSuccessors(b))
  expectedBlack := []int8{18, 19, 20, 26, 28}
  if !reflect.DeepEqual(legalBlack, expectedBlack) {
    t.Errorf("Expected %v to be %v", legalBlack, expectedBlack)
  }
  
}
