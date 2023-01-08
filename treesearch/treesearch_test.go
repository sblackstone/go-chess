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
