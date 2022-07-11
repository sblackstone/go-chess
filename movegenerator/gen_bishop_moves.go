package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

func genSingleBishopMovesGeneric(b *boardstate.BoardState, bishopPos int8, updateFunc func(int8)) {
	file := bitopts.FileOfSquare(bishopPos)

	for r := bishopPos + 9; r < 64 && bitopts.FileOfSquare(r) > file; r += 9 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos + 7; r < 64 && bitopts.FileOfSquare(r) < file; r += 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 7; r >= 0 && bitopts.FileOfSquare(r) > file; r -= 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 9; r >= 0 && bitopts.FileOfSquare(r) < file; r -= 9 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}
}

// This will be almost identical everywhere.
func genSingleBishopMovesBitboard(b *boardstate.BoardState, piecePos int8) uint64 {
	var result uint64

	updateFunc := func(dst int8) {
		result = bitopts.SetBit(result, dst)
	}

	genSingleBishopMovesGeneric(b, piecePos, updateFunc)

	return result
}

// This will be almost identical everywhere.
func genSingleBishopMoves(b *boardstate.BoardState, piecePos int8) []*boardstate.Move {
	var result []*boardstate.Move

	updateFunc := func(dst int8) {
		result = append(result, &boardstate.Move{Src: piecePos, Dst: dst, PromotePiece: boardstate.EMPTY})
	}

	genSingleBishopMovesGeneric(b, piecePos, updateFunc)

	return result
}

func genAllBishopAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64

	updateFunc := func(dst int8) {
		result = bitopts.SetBit(result, dst)
	}
	bishopPositions := b.FindPieces(color, boardstate.BISHOP)
	for _, bp := range bishopPositions {
		genSingleBishopMovesGeneric(b, bp, updateFunc)
	}
	return result
}

func genAllBishopMoves(b *boardstate.BoardState, color int8) []*boardstate.Move {
	var result []*boardstate.Move
	bishopPositions := b.FindPieces(color, boardstate.BISHOP)
	for i := 0; i < len(bishopPositions); i++ {
		result = append(result, genSingleBishopMoves(b, bishopPositions[i])...)
	}
	return result
}

func genBishopSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	color := b.GetTurn()
	var result []*boardstate.BoardState
	bishopPositions := b.FindPieces(color, boardstate.BISHOP)

	for _, pos := range bishopPositions {
		updateFunc := func(dst int8) {
			result = append(result, b.CopyPlayTurn(pos, dst, boardstate.EMPTY))
		}
		genSingleBishopMovesGeneric(b, pos, updateFunc)
	}
	return result
}
