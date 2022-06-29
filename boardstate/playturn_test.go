package boardstate

// Some tests for this code live in boardstate_test.go
import (
  "testing"
  "reflect"
  "errors"
  "fmt"
)


/*
// BoardState contains the state of the Board
type BoardState struct {
	colors [2]uint64
	pieces [6]uint64
	enpassantSquare int8
	meta   uint64
	turn int8
	halfMoves int
	fullMoves int
	movesData []MoveData;
}
*/

func Equal(b1 *BoardState, b2 *BoardState) error {
    if !reflect.DeepEqual(b1.colors, b2.colors) {
      return errors.New("Color mismatch")
    }

    if !reflect.DeepEqual(b1.pieces, b2.pieces) {
      return errors.New("Pieces mismatch")
    }

    if !reflect.DeepEqual(b1.enpassantSquare, b2.enpassantSquare) {
      return errors.New("enpassantSquare mismatch")
    }

    if !reflect.DeepEqual(b1.meta, b2.meta) {
      return errors.New("meta mismatch")
    }

    if !reflect.DeepEqual(b1.turn, b2.turn) {
      return errors.New("turn mismatch")
    }

    if !reflect.DeepEqual(b1.halfMoves, b2.halfMoves) {
      return errors.New("halfMoves mismatch")
    }

    if !reflect.DeepEqual(b1.fullMoves, b2.fullMoves) {
      return errors.New("fullMoves mismatch")
    }

    if len(b1.movesData) != len(b2.movesData) {
      return errors.New("movesData mismatch (count)")
    }

    for i := range(b1.movesData) {
      if !reflect.DeepEqual(b1.movesData[i], b2.movesData[i]) {
        return errors.New("movesData mismatch")
      }
    }

    return nil
}


func TestPopTurnEnpassant(t *testing.T) {
  t.Errorf("not implemented");
}

func TestPopTurnPromote(t *testing.T) {
  t.Errorf("not implemented");

}

func TestPopTurnCapture(t *testing.T) {
  t.Errorf("not implemented");
  
}


func TestPopTurn(t *testing.T) {
   b1 := Initial()
   b2 := b1.CopyPlayTurnFromMove(&Move{src: 8, dst: 24, promotePiece: EMPTY})
   b3 := b2.CopyPlayTurnFromMove(&Move{src: 48, dst: 32, promotePiece: EMPTY})

   fmt.Printf("b2 movesData %p\n", &(b2.movesData[0]))
   fmt.Printf("b3 movesData %p\n", &(b3.movesData[0]))


   b3.PopTurn()

   err := Equal(b3, b2)

   if err != nil {
     t.Errorf("%v", err)
     t.Errorf("Expected b3: \n%v to be \n%v", b3, b2)
   }


   b2.PopTurn()

   err = Equal(b2, b1)

   if err != nil {
     t.Errorf("%v", err)
     t.Errorf("Expected b2: \n%v to be \n%v", b2, b1)
   }




}
