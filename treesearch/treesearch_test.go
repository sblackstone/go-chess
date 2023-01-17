package treesearch

import (
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestCanFindMateInThreeForBlack(t *testing.T) {
	b, err := boardstate.FromFEN("2r3k1/p4p2/3Rp2p/1p2P1pK/8/1P4P1/P3Q2P/1q6 b - - 0 1")
	require.Nil(t, err)
	bm := BestMove(b, 4)
	assert.Equal(t, bm.Src, int8(1))
	assert.Equal(t, bm.Dst, int8(46))
}
