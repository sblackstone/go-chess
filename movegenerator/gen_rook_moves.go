package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

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

func genSingleRookMovesBitboard(b *boardstate.BoardState, piecePos int8) uint64 {
	var result uint64

	updateFunc := func(src, dst int8) {
		result = bitopts.SetBit(result, dst)
	}

	genSingleRookMovesGeneric(b, piecePos, updateFunc)

	return result
}

func genAllRookAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	rookPositions := b.FindPieces(color, boardstate.ROOK)
	for i := 0; i < len(rookPositions); i++ {
		result = result | genSingleRookMovesBitboard(b, rookPositions[i])
	}
	return result
}

func genRookSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	color := b.GetTurn()
	var result []*boardstate.BoardState
	rookPositions := b.FindPieces(color, boardstate.ROOK)

	updateFunc := func(src, dst int8) {
		result = append(result, b.CopyPlayTurn(src, dst, boardstate.EMPTY))
	}

	for _, pos := range rookPositions {
		genSingleRookMovesGeneric(b, pos, updateFunc)
	}
	return result
}
