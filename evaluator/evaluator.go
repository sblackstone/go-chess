package evaluator

import (
	"math/bits"

	"github.com/sblackstone/go-chess/boardstate"
)

func EvaluateBoard(b *boardstate.BoardState) float64 {
	var value float64
	currentTurn := b.GetTurn()
	enemyColor := b.EnemyColor()
	value += float64(bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.PAWN))) * 1.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.ROOK))) * 5.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.KNIGHT))) * 3.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.BISHOP))) * 3.5
	value += float64(bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.QUEEN))) * 9.0
	value += float64(bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.KING))) * 100

	value -= float64(bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.PAWN))) * 1.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.ROOK))) * 5.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.KNIGHT))) * 3.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.BISHOP))) * 3.5
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.QUEEN))) * 9.0
	value -= float64(bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.KING))) * 100

	return value
}
