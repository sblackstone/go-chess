package evaluator

import (
	"fmt"
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

//"github.com/sblackstone/go-chess/boardstate"
//"fmt"

// func TestEvaluateBoard(t *testing.T) {
// 	t.Errorf("not implemented")
// }

func TestIntuition(t *testing.T) {

	example := squarePieceMapsStart[boardstate.KING]

	// WHITE_KING_POS := 4
	// BLACK_KING_POS := 60

	fmt.Printf("White castle long = %f\n", example[2])
	fmt.Printf("White castle short = %f\n", example[6])

	fmt.Printf("Black castle long = %f\n", example[58^56])
	fmt.Printf("Black castle short = %f\n", example[62^56])

	//t.Error(1)
}

func TestIntuitionPawns(t *testing.T) {

	example := squarePieceMapsStart[boardstate.PAWN]

	// WHITE_KING_POS := 4
	// BLACK_KING_POS := 60

	fmt.Printf("White castle long = %f\n", example[27])
	fmt.Printf("White castle short = %f\n", example[28])

	fmt.Printf("Black castle long = %f\n", example[35^56])
	fmt.Printf("Black castle short = %f\n", example[36^56])

	t.Error(1)
}
