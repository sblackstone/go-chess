package evaluator

import (
	"math/bits"

	"github.com/sblackstone/go-chess/arrops"
	"github.com/sblackstone/go-chess/boardstate"
)

var squarePieceMapsStart [6][64]float64
var squarePieceMapsEndGame [6][64]float64
var defaultPieceValues = [6]float64{500, 300, 350, 900, 100000, 1}

func init() {
	//// ROOK //////////////////////////////////////////////
	squarePieceMapsStart[boardstate.ROOK] = [64]float64{
		-5, -5, -5, -5, -5, -5, -5, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, -5, -5, -5, -5, -5, -5, -5,
	}
	squarePieceMapsEndGame[boardstate.ROOK] = [64]float64{
		0, 0, 0, 0, 0, 0, 0, 0,
		45, 50, 50, 50, 50, 50, 50, 45,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	//// KNIGHT //////////////////////////////////////////////
	squarePieceMapsStart[boardstate.KNIGHT] = [64]float64{
		-5, -5, -5, -5, -5, -5, -5, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, -10, -5, -5, -5, -5, -10, -5,
	}
	squarePieceMapsEndGame[boardstate.KNIGHT] = [64]float64{
		-5, -5, -5, -5, -5, -5, -5, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, -5, -5, -5, -5, -5, -5, -5,
	}

	//// BISHOP //////////////////////////////////////////////
	squarePieceMapsStart[boardstate.BISHOP] = [64]float64{
		-5, -5, -5, -5, -5, -5, -5, -5,
		-5, 10, 0, 0, 0, 0, 10, -5,
		-5, 0, 10, 0, 0, 10, 0, -5,
		-5, 0, 0, 10, 10, 0, 0, -5,
		-5, 0, 0, 10, 10, 0, 0, -5,
		-5, 0, 10, 0, 0, 10, 0, -5,
		-5, 10, 0, 0, 0, 0, 10, -5,
		-5, -5, -5, -5, -5, -5, -5, -5,
	}
	squarePieceMapsEndGame[boardstate.BISHOP] = [64]float64{
		-5, -5, -5, -5, -5, -5, -5, -5,
		-5, 10, 0, 0, 0, 0, 10, -5,
		-5, 0, 10, 0, 0, 10, 0, -5,
		-5, 0, 0, 10, 10, 0, 0, -5,
		-5, 0, 0, 10, 10, 0, 0, -5,
		-5, 0, 10, 0, 0, 10, 0, -5,
		-5, 10, 0, 0, 0, 0, 10, -5,
		-5, -5, -5, -5, -5, -5, -5, -5,
	}

	//// QUEEN //////////////////////////////////////////////
	squarePieceMapsStart[boardstate.QUEEN] = [64]float64{
		-5, -5, -5, -5, -5, -5, -5, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, -5, -5, -5, -5, -5, -5, -5,
	}
	squarePieceMapsEndGame[boardstate.QUEEN] = [64]float64{
		-5, -5, -5, -5, -5, -5, -5, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, -5, -5, -5, -5, -5, -5, -5,
	}

	//// KING //////////////////////////////////////////////
	// Encoruage castling but discourage single king move to the left or right to get to
	// castling square.

	squarePieceMapsStart[boardstate.KING] = [64]float64{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 50, -50, 0, -50, 50, 0,
	}
	squarePieceMapsEndGame[boardstate.KING] = [64]float64{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	//// PAWN //////////////////////////////////////////////
	squarePieceMapsStart[boardstate.PAWN] = [64]float64{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		-5, 0, 0, 20, 20, 0, 0, -5,
		-5, 0, 0, 10, 10, 0, 0, -5,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	squarePieceMapsEndGame[boardstate.PAWN] = [64]float64{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	// Everything above is drawn from the perspective of white.
	// But the arrays are reversed..   So put the arrays in the correct order.
	for piece := boardstate.ROOK; piece <= boardstate.PAWN; piece++ {
		arrops.FlipFloat64(&squarePieceMapsStart[piece])
		arrops.FlipFloat64(&squarePieceMapsEndGame[piece])
	}

}

func PercentTowardsEndgame(b *boardstate.BoardState) float64 {
	nonPawnPieces := b.GetOccupiedBitboard() ^ b.GetPieceBitboard(boardstate.PAWN)
	count := bits.OnesCount64(nonPawnPieces)
	return 1 - (float64(count) / 16.0)
}

func EvaluateBoard(b *boardstate.BoardState) float64 {
	var value float64
	var piece int8
	currentTurn := b.GetTurn()
	enemyColor := b.EnemyColor()
	pctEndgame := PercentTowardsEndgame(b)
	for piece = boardstate.ROOK; piece <= boardstate.PAWN; piece++ {
		b.PieceLocations.EachLocation(currentTurn, piece, func(pos int8) {
			if currentTurn == boardstate.BLACK {
				pos = pos ^ 56
			}
			value += defaultPieceValues[piece]
			value += (1-pctEndgame)*squarePieceMapsStart[piece][pos] + pctEndgame*squarePieceMapsEndGame[piece][pos]
		})
		b.PieceLocations.EachLocation(enemyColor, piece, func(pos int8) {
			if enemyColor == boardstate.BLACK {
				pos = pos ^ 56
			}
			value -= defaultPieceValues[piece]
			value -= (1-pctEndgame)*squarePieceMapsStart[piece][pos] + pctEndgame*squarePieceMapsEndGame[piece][pos]
		})

	}

	return value
}
