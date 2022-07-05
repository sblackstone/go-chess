package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
	//	"fmt"
)

func genSingleRookMovesGeneric(b *boardstate.BoardState, rookPos int8, updateFunc func(int8)) {
	for r := rookPos + 8; r < 64; r += 8 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos - 8; r >= 0; r -= 8 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos + 1; bitopts.FileOfSquare(r) > 0; r += 1 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos - 1; r >= 0 && bitopts.FileOfSquare(r) < 7; r -= 1 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}
}

func genSingleRookMovesBitboard(b *boardstate.BoardState, piecePos int8) uint64 {
	var result uint64

	updateFunc := func(dst int8) {
		result = bitopts.SetBit(result, dst)
	}

	genSingleRookMovesGeneric(b, piecePos, updateFunc)

	return result
}

// This will be almost identical everywhere.
func genSingleRookMoves(b *boardstate.BoardState, piecePos int8) []*boardstate.Move {
	var result []*boardstate.Move

	updateFunc := func(dst int8) {
		result = append(result, boardstate.CreateMove(piecePos, dst, boardstate.EMPTY))
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

func genAllRookMoves(b *boardstate.BoardState, color int8) []*boardstate.Move {
	var result []*boardstate.Move
	rookPositions := b.FindPieces(color, boardstate.ROOK)
	for i := 0; i < len(rookPositions); i++ {
		result = append(result, genSingleRookMoves(b, rookPositions[i])...)
	}
	return result
}

func genRookSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllRookMoves(b, b.GetTurn()))
}
