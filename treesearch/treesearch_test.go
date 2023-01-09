package treesearch

import (
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanFindMateInOneForWhite(t *testing.T) {
	b, err := boardstate.FromFEN("r1bqkbnr/1ppp1ppp/p1n5/4p3/2B1P3/5Q2/PPPP1PPP/RNB1K1NR w KQkq - 0 4")
	require.Nil(t, err)
	bm := BestMove(b, 1)
	assert.Equal(t, bm.Src, int8(21))
	assert.Equal(t, bm.Dst, int8(53))
}

func TestCanFindMateInThreeForBlack(t *testing.T) {
	b, err := boardstate.FromFEN("2r3k1/p4p2/3Rp2p/1p2P1pK/8/1P4P1/P3Q2P/1q6 b - - 0 1")
	require.Nil(t, err)
	bm := BestMove(b, 4)
	assert.Equal(t, bm.Src, int8(1))
	assert.Equal(t, bm.Dst, int8(46))
}
