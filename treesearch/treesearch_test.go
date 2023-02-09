package treesearch

import (
	"fmt"
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/movegenerator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkCanMateIn3Single(bench *testing.B) {
	b, _ := boardstate.FromFEN("r1b1kb1r/pppp1ppp/5q2/4n3/3KP3/2N3PN/PPP4P/R1BQ1B1R b kq - 0 1")
	for i := 0; i < bench.N; i++ {
		BestMove(b, 4)
	}
}

func BenchmarkCanMateIn3Multi(bench *testing.B) {
	b, _ := boardstate.FromFEN("r1b1kb1r/pppp1ppp/5q2/4n3/3KP3/2N3PN/PPP4P/R1BQ1B1R b kq - 0 1")
	for i := 0; i < bench.N; i++ {
		BestMoveSmp(b, 4)
	}
}

func TestCanFindMateInThree(t *testing.T) {
	b, err := boardstate.FromFEN("r1b1kb1r/pppp1ppp/5q2/4n3/3KP3/2N3PN/PPP4P/R1BQ1B1R b kq - 0 1")
	require.Nil(t, err)
	bm := BestMove(b, 4)
	assert.Equal(t, int8(61), bm.Src)
	assert.Equal(t, int8(34), bm.Dst)

}
func TestCanFindMateInOneForWhite(t *testing.T) {
	b, err := boardstate.FromFEN("r1bqkbnr/1ppp1ppp/p1n5/4p3/2B1P3/5Q2/PPPP1PPP/RNB1K1NR w KQkq - 0 4")
	require.Nil(t, err)
	bm := BestMove(b, 1)
	assert.Equal(t, int8(21), bm.Src)
	assert.Equal(t, int8(53), bm.Dst)
}

func BenchmarkFirstMove(b *testing.B) {
	board := boardstate.Initial()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		BestMove(board, 5)
	}
}

// func BenchmarkMateInNine(b *testing.B) {
// 	board, _ := boardstate.FromFEN("4nr1k/p1p1p1pp/bp1pn1r1/8/6QR/6RP/1BBq1PP1/6K1 w - - 0 1")
// 	fmt.Printf("%+v\n", BestMove(board, 10))
// }

func BenchmarkMateInSix(b *testing.B) {
	board, _ := boardstate.FromFEN("r1k4r/ppp1bq1p/2n1N3/6B1/3p2Q1/8/PPP2PPP/R5K1 w - - 0 1")
	fmt.Printf("%+v\n", BestMove(board, 7))
}

// This occasionally fails because there is more than one solution.
func TestCanFindMateInFourForWhite(t *testing.T) {
	board, _ := boardstate.FromFEN("1k6/2p3r1/p6p/1pQP4/3N2q1/8/P5P1/6K1 w - - 0 1")

	// NC6+
	move := BestMove(board, 5)
	assert.Equal(t, "d4 c6 99", move.ToString())
	board.PlayTurnFromMove(move)

	// Kb7
	board.PlayTurn(57, 49, 99)

	// Qa7+
	move = BestMove(board, 5)
	assert.Equal(t, "c5 a7 99", move.ToString())
	board.PlayTurnFromMove(move)

	// Kc8
	board.PlayTurn(49, 58, 99)

	// Qa8
	move = BestMove(board, 5)
	assert.Equal(t, "a7 a8 99", move.ToString())
	board.PlayTurnFromMove(move)

	// Kd7
	board.PlayTurn(58, 51, 99)

	move = BestMove(board, 5)
	fmt.Printf("%+v\n", move)
	assert.Equal(t, "a8 d8 99", move.ToString())
	board.PlayTurnFromMove(move)

	assert.Equal(t, int8(movegenerator.GAME_STATE_CHECKMATE), movegenerator.CheckEndOfGame(board))

}

func TestCanFindMateInThreeForBlack(t *testing.T) {
	b, err := boardstate.FromFEN("2r3k1/p4p2/3Rp2p/1p2P1pK/8/1P4P1/P3Q2P/1q6 b - - 0 1")
	require.Nil(t, err)
	bm := BestMove(b, 4)
	assert.Equal(t, bm.Src, int8(1))
	assert.Equal(t, bm.Dst, int8(46))
}
