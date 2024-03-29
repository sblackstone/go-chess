package movegenerator

import (
	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

/*
00 01 02 03 04 05 06 07
08 09 10 11 12 13 14 15
16 17 18 19 20 21 22 23
24 25 26 27 28 29 30 31
32 33 34 35 36 37 38 39
40 41 42 43 44 45 46 47
48 49 50 51 52 53 54 55
56 57 58 59 60 61 62 63


 A      B    C    D   <-- These labels match code sections...

      -17  -15
-10              -6

+6               +10
      +15  +17

*/

var pregeneratedKnightMoves [64][]int8
var pregeneratedKnightMovesBitboard [64]uint64

func getPregeneratedKnightMoves() [64][]int8 {
	return pregeneratedKnightMoves
}

func init() {
	var rank, file int8
	for rank = 0; rank < 8; rank++ {
		for file = 0; file < 8; file++ {
			pos := bitops.RankFileToSquare(rank, file)
			pregeneratedKnightMovesBitboard[pos] = 0

			appendPos := func(dst int8) {
				pregeneratedKnightMovesBitboard[pos] = bitops.SetBit(pregeneratedKnightMovesBitboard[pos], dst)
				pregeneratedKnightMoves[pos] = append(pregeneratedKnightMoves[pos], dst)
			}

			// A
			if file >= 2 {
				if rank >= 1 {
					appendPos(pos - 10)
				}
				if rank <= 6 {
					appendPos(pos + 6)
				}
			}

			// B
			if file >= 1 {
				if rank >= 2 {
					appendPos(pos - 17)
				}
				if rank <= 5 {
					appendPos(pos + 15)
				}
			}

			// C
			if file <= 6 {
				if rank >= 2 {
					appendPos(pos - 15)
				}
				if rank <= 5 {
					appendPos(pos + 17)
				}
			}

			// D
			if file <= 5 {
				if rank >= 1 {
					appendPos(pos - 6)
				}
				if rank <= 6 {
					appendPos(pos + 10)
				}
			}

		}
	}
}

func genSingleKnightMovesGeneric(b *boardstate.BoardState, knightPos int8, updateFunc func(int8, int8, int8)) {
	knightColor := b.ColorOfSquare(knightPos)
	for _, move := range pregeneratedKnightMoves[knightPos] {
		if b.ColorOfSquare(move) != knightColor {
			updateFunc(knightPos, move, boardstate.EMPTY)
		}
	}
}

func genAllKnightMovesGeneric(b *boardstate.BoardState, color int8, updateFunc func(int8, int8, int8)) {
	b.PieceLocations.EachLocation(color, boardstate.KNIGHT, func(knightPos int8) {
		genSingleKnightMovesGeneric(b, knightPos, updateFunc)
	})
}

func genAllKnightAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	b.PieceLocations.EachLocation(color, boardstate.KNIGHT, func(knightPos int8) {
		result = result | ((pregeneratedKnightMovesBitboard[knightPos] ^ b.GetColorBitboard(b.ColorOfSquare(knightPos))) & pregeneratedKnightMovesBitboard[knightPos])
	})
	return result
}

func genKnightSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState

	updateFunc := func(src, dst, promotionPiece int8) {
		result = append(result, b.CopyPlayTurn(src, dst, promotionPiece))
	}
	genAllKnightMovesGeneric(b, b.GetTurn(), updateFunc)
	return result
}
