package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

var rookMasks [64]uint64

// For future use.
func init() {
	var pos int8
	for pos = 0; pos < 64; pos++ {
		b := boardstate.Blank()
		b.SetSquare(pos, boardstate.WHITE, boardstate.ROOK)
		rookMasks[pos] = genAllRookAttacks(b, boardstate.WHITE)
		bitopts.Print(rookMasks[pos], pos)
	}

}

func genSingleRookMovesGeneric(b *boardstate.BoardState, rookPos int8, updateFunc func(int8, int8)) {
	for r := rookPos + 8; r < 64; r += 8 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos - 8; r >= 0; r -= 8 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos + 1; r <= 63 && bitopts.FileOfSquare(r) > 0; r += 1 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos - 1; r >= 0 && bitopts.FileOfSquare(r) < 7; r -= 1 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}
}

func genAllRookMovesGeneric(b *boardstate.BoardState, color int8, updatefunc func(int8, int8)) {
	rookPositions := b.FindPieces(color, boardstate.ROOK)
	for _, rookPos := range rookPositions {
		genSingleRookMovesGeneric(b, rookPos, updatefunc)
	}
}

func genAllRookAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	updateFunc := func(src, dst int8) {
		result = bitopts.SetBit(result, dst)
	}
	genAllRookMovesGeneric(b, color, updateFunc)
	return result
}

func genRookSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState
	updateFunc := func(src, dst int8) {
		result = append(result, b.CopyPlayTurn(src, dst, boardstate.EMPTY))
	}
	genAllRookMovesGeneric(b, b.GetTurn(), updateFunc)
	return result
}
