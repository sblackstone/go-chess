package evaluator

import (
	"math/bits"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/movegenerator"
)

func EvaluateBoard(b *boardstate.BoardState) int {
	var value int
	currentTurn := b.GetTurn()
	enemyColor := b.EnemyColor()
	value += bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.PAWN)) * 100
	value += bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.ROOK)) * 500
	value += bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.KNIGHT)) * 300
	value += bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.BISHOP)) * 350
	value += bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.QUEEN)) * 900
	value += bits.OnesCount64(b.GetPieceBitboard(currentTurn, boardstate.KING)) * 10000

	value -= bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.PAWN)) * 100
	value -= bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.ROOK)) * 500
	value -= bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.KNIGHT)) * 300
	value -= bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.BISHOP)) * 350
	value -= bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.QUEEN)) * 900
	value -= bits.OnesCount64(b.GetPieceBitboard(enemyColor, boardstate.KING)) * 10000

	value += int(movegenerator.GenMovesCount(b))

	return value
}
