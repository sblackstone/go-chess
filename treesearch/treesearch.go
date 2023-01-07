package treesearch

import (
	"math/rand"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/evaluator"
	"github.com/sblackstone/go-chess/movegenerator"
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

	for _, move := range movegenerator.GenMoves(b) {
		b.PlayTurnFromMove(move)
		score := -alphaBeta(b, depth-1, -beta, -alpha)
		b.UnplayTurn()

		if score >= beta {
			return beta
		}

		if score > alpha {
			alpha = score
		}
	}
	return alpha
}

func BestMove(b *boardstate.BoardState, depth int8) *boardstate.Move {
	var bestValue float64
	var bestMoves []*boardstate.Move
	bestValue = -INFINITY
	for _, move := range movegenerator.GenMoves(b) {
		//succ := b.CopyPlayTurnFromMove(move)
		b.PlayTurnFromMove(move)
		value := -alphaBeta(b, depth-1, -INFINITY, INFINITY)
		b.UnplayTurn()
		if value == bestValue {
			bestMoves = append(bestMoves, move)
		}
		if value > bestValue {
			bestValue = value
			bestMoves = make([]*boardstate.Move, 1)
			bestMoves[0] = move
		}
	}
	randomIndex := rand.Intn(len(bestMoves))
	return bestMoves[randomIndex]
}
