package treesearch

import (
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

func TestBestSuccessor1(t *testing.T) {
	b := boardstate.Initial()
	b.PlayTurn(8, 24, boardstate.EMPTY)
	b.PlayTurn(48, 32, boardstate.EMPTY)
	b.Print(127)
	bs := BestSuccessor(b, 2)
	bs.Print(127)
}

// Old test for when no king on board caused a crash.
// func TestCrash(t *testing.T) {

// 	b := uci.BoardFromUCIPosition("position fen r5rk/5p1p/5R2/4B3/8/8/7P/7K w - - 0 1 moves f6f7 g8g7 f7g7 a8a5 g7g5 a5e5 g5e5 h7h6 h1g1 h8g7 e5a5 g7h7 a5b5 h7g6 g1f2")
// 	fen, _ := b.ToFEN()
// 	fmt.Printf("%s\n", fen)
// 	bm := BestMove(b, 2)
// 	bs := BestSuccessor(b, 2)
// 	bs.Print(127)
// 	t.Errorf("%v\n", uci.MoveToUCI(bm))

// }

// func TestPlayItself(t *testing.T) {
// 	b := boardstate.Initial()

// 	for {
// 		b = BestSuccessor(b, 3)
// 		b.Print(127)
// 	}
// 	// m1 := BestSuccessor(b,5)
// 	// m1.Print(127)
// 	//
// 	// m1.PlayTurn(49,33, boardstate.EMPTY)
// 	// m1.Print(127)
// 	// m2 := BestSuccessor(m1,5)
// 	// m2.Print(127)

// }
