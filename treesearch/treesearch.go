package treesearch

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/evaluator"
	"github.com/sblackstone/go-chess/movegenerator"
	"math/rand"
	//"math"
)

const INFINITY = float64(9999999999999)

// https://www.chessprogramming.org/Alpha-Beta#Negamax_Framework
// Alpha-Beta pruning NegaMax evaluation.
func alphaBeta(b *boardstate.BoardState, depth int8, alpha float64, beta float64) float64 {
	// Making this not a variable seems to be a performance boost?  not getting compiled away?
	gameState := movegenerator.CheckEndOfGame(b)

	if depth == 0 || gameState > movegenerator.GAME_STATE_PLAYING {
		return evaluator.EvaluateBoard(b)
	}

	for _, succ := range movegenerator.GenLegalSuccessors(b) {
		score := -alphaBeta(succ, depth-1, -beta, -alpha)

		if score >= beta {
			return beta
		}

		if score > alpha {
			alpha = score
		}
	}
	return alpha
}

func BestSuccessor(b *boardstate.BoardState, depth int8) *boardstate.BoardState {
	var bestValue float64
	var bestSuccessors []*boardstate.BoardState
	bestValue = -INFINITY
	for _, succ := range movegenerator.GenLegalSuccessors(b) {
		value := -alphaBeta(succ, depth, -INFINITY, INFINITY)
		if value == bestValue {
			bestSuccessors = append(bestSuccessors, succ)
		}
		if value > bestValue {
			bestValue = value
			bestSuccessors = make([]*boardstate.BoardState, 1)
			bestSuccessors[0] = succ
		}
	}
	randomIndex := rand.Intn(len(bestSuccessors))
	return bestSuccessors[randomIndex]
}
