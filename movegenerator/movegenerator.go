package movegenerator

import (
	"fmt"

	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

// Used in testing only.
func genSuccessorsForPiece(b *boardstate.BoardState, pieceType int8) []*boardstate.BoardState {
	switch pieceType {
	case boardstate.PAWN:
		return genPawnSuccessors(b)
	case boardstate.KING:
		return genKingSuccessors(b)
	case boardstate.QUEEN:
		return genQueenSuccessors(b)
	case boardstate.BISHOP:
		return genBishopSuccessors(b)
	case boardstate.KNIGHT:
		return genKnightSuccessors(b)
	case boardstate.ROOK:
		return genRookSuccessors(b)
	default:
		panic("GenSuccessorsForPiece: Unknown piece type " + fmt.Sprint(pieceType))
	}
}

// Used in testing only.
func genAttacksForPiece(b *boardstate.BoardState, color int8, pieceType int8) uint64 {
	switch pieceType {
	case boardstate.PAWN:
		return genAllPawnAttacks(b, color)
	case boardstate.KING:
		return genAllKingAttacks(b, color)
	case boardstate.QUEEN:
		return genAllQueenAttacks(b, color)
	case boardstate.BISHOP:
		return genAllBishopAttacks(b, color)
	case boardstate.KNIGHT:
		return genAllKnightAttacks(b, color)
	case boardstate.ROOK:
		return genAllRookAttacks(b, color)
	default:
		panic("GenSuccessorsForPiece: Unknown piece type " + fmt.Sprint(pieceType))
	}
}

func GenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState
	result = append(result, genPawnSuccessors(b)...)
	result = append(result, genKingSuccessors(b)...)
	result = append(result, genQueenSuccessors(b)...)
	result = append(result, genBishopSuccessors(b)...)
	result = append(result, genKnightSuccessors(b)...)
	result = append(result, genRookSuccessors(b)...)
	return result
}

// For testing at the moment, not tested.
func GenMoves(b *boardstate.BoardState) []*boardstate.Move {
	var result []*boardstate.Move
	result = append(result, genAllPawnMoves(b, b.GetTurn(), false)...)
	result = append(result, genAllKingMoves(b, b.GetTurn(), false)...)
	result = append(result, genAllQueenMoves(b, b.GetTurn())...)
	result = append(result, genAllBishopMoves(b, b.GetTurn())...)
	result = append(result, genAllKnightMoves(b, b.GetTurn())...)
	result = append(result, genAllRookMoves(b, b.GetTurn())...)
	return result
}

func IsInCheck(b *boardstate.BoardState, color int8) bool {
	kingPos := b.GetKingPos(color)
	if kingPos == boardstate.NO_KING {
		return false
	}
	// Find all the checked squares for the opposing side
	attacks := GenAllCheckedSquares(b, color^1)
	return bitopts.TestBit(attacks, kingPos)
}

// We can do much better here, naieve O(n^2) first attempt.
func GenLegalSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	// Find all the successors for B...
	successors := GenSuccessors(b)
	color := b.GetTurn()
	var result []*boardstate.BoardState

	// For each candidate successor
	for _, succ := range successors {
		if !IsInCheck(succ, color) {
			result = append(result, succ)
		}
	}
	return result
}

const (
	GAME_STATE_PLAYING = iota
	GAME_STATE_CHECKMATE
	GAME_STATE_STALEMATE
)

func CheckEndOfGame(b *boardstate.BoardState) int8 {
	successors := GenLegalSuccessors(b)
	if len(successors) > 0 {
		return GAME_STATE_PLAYING
	}

	currentTurnColor := b.GetTurn()
	oppTurnColor := b.EnemyColor()

	kingPos := b.FindPieces(currentTurnColor, boardstate.KING)
	attacks := GenAllCheckedSquares(b, oppTurnColor)
	if bitopts.TestBit(attacks, kingPos[0]) {
		return GAME_STATE_CHECKMATE
	}

	return GAME_STATE_STALEMATE

}

func GenAllCheckedSquares(b *boardstate.BoardState, color int8) uint64 {
	return genAllKnightAttacks(b, color) |
		genAllBishopAttacks(b, color) |
		genAllRookAttacks(b, color) |
		genAllQueenAttacks(b, color) |
		genAllKingAttacks(b, color) |
		genAllPawnAttacks(b, color)

}
