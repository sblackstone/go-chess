package treesearch

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/evaluator"
	"github.com/sblackstone/go-chess/movegenerator"
)

const INFINITY = float64(9999999999999)

func DepthPrint(depth int8, formatStr string, args ...any) {
	for i := int8(0); i < depth; i++ {
		fmt.Print("%\t")
	}
	fmt.Printf(formatStr, args...)
}

// https://www.chessprogramming.org/Alpha-Beta#Negamax_Framework
// Alpha-Beta pruning NegaMax evaluation.
func alphaBeta(b *boardstate.BoardState, moves []*boardstate.Move, depth int8, alpha float64, beta float64) float64 {

	currentTurn := b.GetTurn()
	// Making this not a variable seems to be a performance boost?  not getting compiled away?
	gameState := movegenerator.CheckEndOfGame(b)
	if gameState == movegenerator.GAME_STATE_CHECKMATE {
		return -INFINITY - float64(depth)
	}

	if depth == 0 || gameState > movegenerator.GAME_STATE_PLAYING {
		return evaluator.EvaluateBoard(b)
	}

	movegenerator.GenMovesInto(b, &moves)
	for _, move := range moves {

		//DepthPrint(5-depth, "Playing %+v\n", move)
		b.PlayTurnFromMove(move)

		if !movegenerator.IsInCheck(b, currentTurn) {
			score := -alphaBeta(b, moves[len(moves):], depth-1, -beta, -alpha)
			if score >= beta {
				//DepthPrint(5-depth, "Unplaying %+v\n", move)
				b.UnplayTurn()
				return beta
			}

			if score > alpha {
				alpha = score
			}

		}
		//DepthPrint(5-depth, "Unplaying %+v\n\n", move)
		b.UnplayTurn()

	}
	return alpha
}

func BestMoveSmp(b *boardstate.BoardState, depth int8) *boardstate.Move {
	var bestValue float64
	var bestMoves []*boardstate.Move
	bestValue = -INFINITY
	currentTurn := b.GetTurn()
	var wg sync.WaitGroup
	var mux sync.Mutex

	moves := make([]*boardstate.Move, 0, 300)
	movegenerator.GenMovesInto(b, &moves)
	for _, move := range moves {
		wg.Add(1)
		go func(b *boardstate.BoardState, move *boardstate.Move) {
			defer wg.Done()
			b.PlayTurnFromMove(move)
			if !movegenerator.IsInCheck(b, currentTurn) {
				moves := make([]*boardstate.Move, 0, 1000)
				value := -alphaBeta(b, moves, depth, -INFINITY, INFINITY)
				mux.Lock()
				if value == bestValue {
					bestMoves = append(bestMoves, move)
				}
				if value > bestValue {
					bestValue = value
					bestMoves = make([]*boardstate.Move, 1)
					bestMoves[0] = move
				}
				mux.Unlock()
			}
			b.UnplayTurn()

		}(b.Copy(), move)
	}
	wg.Wait()
	randomIndex := rand.Intn(len(bestMoves))
	return bestMoves[randomIndex]
}

func BestMove(b *boardstate.BoardState, depth int8) *boardstate.Move {
	var bestValue float64
	var bestMoves []*boardstate.Move
	bestValue = -INFINITY
	currentTurn := b.GetTurn()
	moves := make([]*boardstate.Move, 0, 1000)
	movegenerator.GenMovesInto(b, &moves)
	for _, move := range moves {
		b.PlayTurnFromMove(move)
		if !movegenerator.IsInCheck(b, currentTurn) {
			value := -alphaBeta(b, moves[len(moves):], depth, -INFINITY, INFINITY)
			if value == bestValue {
				bestMoves = append(bestMoves, move)
			}
			if value > bestValue {
				bestValue = value
				bestMoves = make([]*boardstate.Move, 1)
				bestMoves[0] = move
			}
		}
		b.UnplayTurn()
	}
	randomIndex := rand.Intn(len(bestMoves))
	return bestMoves[randomIndex]
}
