package movegenerator

import (
	"fmt"

	"github.com/sblackstone/go-chess/bitops"
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
func GenMovesCount(b *boardstate.BoardState) int {
	var result int

	updateFunc := func(src, dst, promotePiece int8) {
		result += 1
	}
	GenAllMovesGeneric(b, updateFunc)
	return result
}

func GenAllMovesGeneric(b *boardstate.BoardState, updateFunc func(int8, int8, int8)) {
	turn := b.GetTurn()
	genAllQueenMovesGeneric(b, turn, updateFunc)
	genAllBishopMovesGeneric(b, turn, updateFunc)
	genAllKnightMovesGeneric(b, turn, updateFunc)
	genAllRookMovesGeneric(b, turn, updateFunc)
	genAllKingMovesGeneric(b, turn, false, updateFunc)
	genAllPawnMovesGeneric(b, turn, false, updateFunc)
}

func GenMoves(b *boardstate.BoardState) []*boardstate.Move {
	var result []*boardstate.Move

	updateFunc := func(src, dst, promotePiece int8) {
		result = append(result, &boardstate.Move{Src: src, Dst: dst, PromotePiece: promotePiece})
	}
	GenAllMovesGeneric(b, updateFunc)
	return result
}

func GenMovesInto(b *boardstate.BoardState, target *[]*boardstate.Move) {
	updateFunc := func(src, dst, promotePiece int8) {
		*target = append(*target, &boardstate.Move{Src: src, Dst: dst, PromotePiece: promotePiece})
	}
	GenAllMovesGeneric(b, updateFunc)
}

func LegalMoveCount(b *boardstate.BoardState) uint64 {
	var result uint64
	currentTurn := b.GetTurn()
	updateFunc := func(src, dst, promotePiece int8) {
		b.PlayTurn(src, dst, promotePiece)
		if !IsInCheck(b, currentTurn) {
			result += 1
		}
		b.UnplayTurn()
	}
	GenAllMovesGeneric(b, updateFunc)
	return result
}

func HasLegalMove(b *boardstate.BoardState) bool {
	var result bool
	currentTurn := b.GetTurn()
	updateFunc := func(src, dst, promotePiece int8) {
		if !result {
			b.PlayTurn(src, dst, promotePiece)
			if !IsInCheck(b, currentTurn) {
				result = true
			}
			b.UnplayTurn()
		}
	}
	GenAllMovesGeneric(b, updateFunc)
	return result
}

// func GenLegalMoves(b *boardstate.BoardState) []*boardstate.Move {
// 	var legalMoves []*boardstate.Move
// 	currentTurn := b.GetTurn()
// 	for _, move := range GenMoves(b) {
// 		b.PlayTurnFromMove(move)
// 		if !IsInCheck(b, currentTurn) {
// 			legalMoves = append(legalMoves, move)
// 		}
// 		b.UnplayTurn()
// 	}
// 	return legalMoves
// }

func IsInCheck(b *boardstate.BoardState, color int8) bool {
	kingPos := b.GetKingPos(color)
	if kingPos == boardstate.NO_KING {
		return false
	}
	// Find all the checked squares for the opposing side
	attacks := GenAllCheckedSquares(b, color^1)
	return bitops.TestBit(attacks, kingPos)
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
	if HasLegalMove(b) {
		return GAME_STATE_PLAYING
	}

	currentTurnColor := b.GetTurn()
	oppTurnColor := b.EnemyColor()

	kingPos := b.GetKingPos(currentTurnColor)
	if kingPos == boardstate.NO_KING {
		return GAME_STATE_CHECKMATE
	}
	attacks := GenAllCheckedSquares(b, oppTurnColor)
	if bitops.TestBit(attacks, kingPos) {
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
