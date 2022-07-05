package evaluator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"math/bits"
)

func EvaluateBoard(b *boardstate.BoardState) float64 {
	var value float64
	value += float64(bits.OnesCount64(b.GetPieceBitboard(b.GetTurn(), boardstate.PAWN))) * 1.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(b.GetTurn(), boardstate.ROOK))) * 5.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(b.GetTurn(), boardstate.KNIGHT))) * 3.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(b.GetTurn(), boardstate.BISHOP))) * 3.5
	value += float64(bits.OnesCount64(b.GetPieceBitboard(b.GetTurn(), boardstate.QUEEN))) * 9.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(b.GetTurn(), boardstate.KING))) * 100

	value -= float64(bits.OnesCount64(b.GetPieceBitboard(b.EnemyColor(), boardstate.PAWN))) * 1.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(b.EnemyColor(), boardstate.ROOK))) * 5.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(b.EnemyColor(), boardstate.KNIGHT))) * 3.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(b.EnemyColor(), boardstate.BISHOP))) * 3.5
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(b.EnemyColor(), boardstate.QUEEN))) * 9.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(b.EnemyColor(), boardstate.KING))) * 100

	return value
}
