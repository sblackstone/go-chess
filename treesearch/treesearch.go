package treesearch

import (
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/evaluator"
  "github.com/sblackstone/go-chess/movegenerator"
	"math"
)

const INFINITY = float64(9999999999999)

func alphaBeta(b *boardstate.BoardState, depth int8, alpha float64, beta float64, maximizingPlayer bool) float64 {
	// Making this not a variable seems to be a performance boost?  not getting compiled away?
	gameState := movegenerator.CheckEndOfGame(b)
  if depth == 0 || gameState > movegenerator.GAME_STATE_PLAYING {
		return evaluator.EvaluateBoard(b)
	}

	if maximizingPlayer {
			value := -INFINITY
			for _, succ := range(movegenerator.GenLegalSuccessors(b)) {
				value = math.Max(value, alphaBeta(succ, depth - 1, alpha, beta, false))
				if value >= beta {
					break
				}
				alpha = math.Max(alpha, value)
			}
			return value

	} else {
		value := INFINITY
		for _, succ := range(movegenerator.GenLegalSuccessors(b)) {
			value = math.Min(value, alphaBeta(succ, depth - 1, alpha, beta, true))
			if value <= alpha {
				break
			}
			beta = math.Min(beta, value)
		}
		return value
	}
}


func BestSuccessor(b *boardstate.BoardState, depth int8) *boardstate.BoardState {
	var bestValue float64
	var bestSuccessor *boardstate.BoardState
	bestValue = -INFINITY
	for _, succ := range(movegenerator.GenLegalSuccessors(b)) {
		value := alphaBeta(succ, depth, -INFINITY, INFINITY, false) // FALSE!
		if value > bestValue {
			bestValue = value
			bestSuccessor = succ
		}
	}
	return bestSuccessor
}
