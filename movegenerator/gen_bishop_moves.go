package movegenerator

import (
	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

// For future use.
var bishopAttackMasks [64]uint64

func genBishopAttackMasks() {
	var pos int8
	for pos = 0; pos < 64; pos++ {
		b := boardstate.Blank()
		b.SetSquare(pos, boardstate.WHITE, boardstate.BISHOP)
		bishopAttackMasks[pos] = genAllBishopAttacks(b, boardstate.WHITE)
		// bitops.Print(bishopAttackMasks[pos], pos)
		// fmt.Printf("***\n\n")
	}
}

func init() {
	genBishopAttackMasks()
}

func genSingleBishopMovesGeneric(b *boardstate.BoardState, bishopPos int8, updateFunc func(int8, int8)) {
	file := bitops.FileOfSquare(bishopPos)

	for r := bishopPos + 9; r < 64 && bitops.FileOfSquare(r) > file; r += 9 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos + 7; r < 64 && bitops.FileOfSquare(r) < file; r += 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 7; r >= 0 && bitops.FileOfSquare(r) > file; r -= 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 9; r >= 0 && bitops.FileOfSquare(r) < file; r -= 9 {
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
		result = bitops.SetBit(result, dst)
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
