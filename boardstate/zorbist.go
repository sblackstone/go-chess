package boardstate

import "math/rand"

var zorbistPieces [2][6][64]uint64

var zorbistTurns [2]uint64
var zorbistCastling [2][2]uint64
var zorbistEnpassant [64]uint64

func initZorbistPieces() {
	var color, piece, square int8

	for color = WHITE; color <= BLACK; color++ {
		for piece = ROOK; piece <= PAWN; piece++ {
			for square = 0; square < 64; square++ {
				zorbistPieces[color][piece][square] = rand.Uint64()
			}
		}
	}
}

func initZorbistTurns() {
	zorbistTurns[WHITE] = rand.Uint64()
	zorbistTurns[BLACK] = rand.Uint64()
}

func initZorbistCastling() {
	zorbistCastling[WHITE][CASTLE_LONG] = rand.Uint64()
	zorbistCastling[WHITE][CASTLE_SHORT] = rand.Uint64()
	zorbistCastling[BLACK][CASTLE_LONG] = rand.Uint64()
	zorbistCastling[BLACK][CASTLE_SHORT] = rand.Uint64()
}

func initZorbistEnpassant() {
	var i int8
	for i = 0; i < 64; i++ {
		zorbistEnpassant[i] = rand.Uint64()
	}
}

func init() {
	initZorbistPieces()
	initZorbistTurns()
	initZorbistCastling()
	initZorbistEnpassant()
}
