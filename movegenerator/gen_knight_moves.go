package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
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
			pos := bitopts.RankFileToSquare(rank, file)
			pregeneratedKnightMovesBitboard[pos] = 0

			appendPos := func(dst int8) {
				pregeneratedKnightMovesBitboard[pos] = bitopts.SetBit(pregeneratedKnightMovesBitboard[pos], dst)
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

// All the base generators need to look like this.....
func genSingleKnightMovesGeneric(b *boardstate.BoardState, knightPos int8, updateFunc func(int8)) {
	for i := range pregeneratedKnightMoves[knightPos] {
		move := pregeneratedKnightMoves[knightPos][i]
		if b.ColorOfSquare(move) != b.ColorOfSquare(knightPos) {
			updateFunc(move)
		}
	}
}

// This will be almost identical everywhere.
func genAllKnightAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	knightPositions := b.FindPieces(color, boardstate.KNIGHT)
	for _, pos := range knightPositions {
		result = result | ((pregeneratedKnightMovesBitboard[pos] ^ b.GetColorBitboard(b.ColorOfSquare(pos))) & pregeneratedKnightMovesBitboard[pos])
	}
	return result
}

func genAllKnightMoves(b *boardstate.BoardState, color int8) []*boardstate.Move {
	var result []*boardstate.Move
	knightPositions := b.FindPieces(color, boardstate.KNIGHT)

	for _, pos := range knightPositions {
		updateFunc := func(dst int8) {
			result = append(result, &boardstate.Move{Src: pos, Dst: dst, PromotePiece: boardstate.EMPTY})
		}
		genSingleKnightMovesGeneric(b, pos, updateFunc)
	}
	return result
}

func genKnightSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	color := b.GetTurn()
	var result []*boardstate.BoardState
	knightPositions := b.FindPieces(color, boardstate.KNIGHT)

	for _, pos := range knightPositions {
		updateFunc := func(dst int8) {
			result = append(result, b.CopyPlayTurn(pos, dst, boardstate.EMPTY))
		}
		genSingleKnightMovesGeneric(b, pos, updateFunc)
	}
	return result
}
