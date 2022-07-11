package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

func genSingleBishopMovesGeneric(b *boardstate.BoardState, bishopPos int8, updateFunc func(int8, int8)) {
	file := bitopts.FileOfSquare(bishopPos)

	for r := bishopPos + 9; r < 64 && bitopts.FileOfSquare(r) > file; r += 9 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos + 7; r < 64 && bitopts.FileOfSquare(r) < file; r += 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 7; r >= 0 && bitopts.FileOfSquare(r) > file; r -= 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 9; r >= 0 && bitopts.FileOfSquare(r) < file; r -= 9 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}
}

func genAllBishopMovesGeneric(b *boardstate.BoardState, color int8, updateFunc func(int8, int8)) {
	bishopPositions := b.FindPieces(color, boardstate.BISHOP)
	for _, pos := range bishopPositions {
		genSingleBishopMovesGeneric(b, pos, updateFunc)
	}
}

func genAllBishopAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	updateFunc := func(src int8, dst int8) {
		result = bitopts.SetBit(result, dst)
	}
	genAllBishopMovesGeneric(b, color, updateFunc)
	return result
}

func genAllBishopMoves(b *boardstate.BoardState, color int8) []*boardstate.Move {
	var result []*boardstate.Move
	updateFunc := func(src int8, dst int8) {
		result = append(result, &boardstate.Move{Src: src, Dst: dst, PromotePiece: boardstate.EMPTY})
	}
	genAllBishopMovesGeneric(b, color, updateFunc)
	return result
}

func genBishopSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState
	updateFunc := func(src int8, dst int8) {
		result = append(result, b.CopyPlayTurn(src, dst, boardstate.EMPTY))
	}
	genAllBishopMovesGeneric(b, b.GetTurn(), updateFunc)
	return result
}
