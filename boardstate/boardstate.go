package boardstate

import (
	"github.com/sblackstone/go-chess/bitopts"
)


// BoardState contains the state of the Board
type BoardState struct {
	colors []uint64
	pieces []uint64
	meta   uint64
}

// Blank returns a blank board with no pieces on it
func Blank() *BoardState {
	b := BoardState{}
	b.colors = []uint64{0, 0}
	b.pieces = []uint64{0, 0, 0, 0, 0, 0}
	return &b
}

// Initial returns a board with the initial setup.
func Initial() *BoardState {
	b := BoardState{}
	// These constants are pre-calculated using InitialManual (see below)...
	b.colors = []uint64{65535, 18446462598732840960 }
	b.pieces = []uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
	return &b
}

// initialManual sets up the board manually, only used to calculate the constants for the fast version Initial.
func initialManual() *BoardState {
	var j uint8

	b := Blank()

	backFile := []uint8{ROOK, KNIGHT, BISHOP, QUEEN, KING, BISHOP, KNIGHT, ROOK}
	for j = 0; j < 8; j++ {
		b.SetSquareRankFile(7, j, BLACK, backFile[j])
		b.SetSquareRankFile(0, j, WHITE, backFile[j])

		b.SetSquareRankFile(6, j, BLACK, PAWN)
		b.SetSquareRankFile(1, j, WHITE, PAWN)
	}

	return b
}

func (b *BoardState) MovePiece(src uint8, dst uint8) {

}

// ColorOfSquare returns WHITE,BLACK, or EMPTY
func (b *BoardState) ColorOfSquare(n uint8) uint8 {
	if bitopts.TestBit(b.colors[WHITE], n) {
		return WHITE
	}
	if bitopts.TestBit(b.colors[BLACK], n) {
		return BLACK
	}
	return EMPTY
}

// PieceOfSquare t
func (b *BoardState) PieceOfSquare(n uint8) uint8 {
	var i uint8
	for i = 0; i < 6; i++ {
		if bitopts.TestBit(b.pieces[i], n) {
			return i
		}
	}
	return EMPTY
}

// SetSquare removes any existing piece and sets the square to the new piece/color.
func (b *BoardState) SetSquare(n uint8, color uint8, piece uint8) {
	// Theres gotta be room for improvement here...
	// we really only need to update the bitboard that is currently set.
	b.pieces[ROOK]    = bitopts.ClearBit(b.pieces[ROOK],   n)
	b.pieces[BISHOP]  = bitopts.ClearBit(b.pieces[BISHOP], n)
	b.pieces[KNIGHT]  = bitopts.ClearBit(b.pieces[KNIGHT], n)
	b.pieces[QUEEN]   = bitopts.ClearBit(b.pieces[QUEEN],  n)
	b.pieces[KING]    = bitopts.ClearBit(b.pieces[KING],   n)
	b.pieces[PAWN]    = bitopts.ClearBit(b.pieces[PAWN],   n)
	b.colors[WHITE]   = bitopts.ClearBit(b.colors[WHITE],  n)
	b.colors[BLACK]   = bitopts.ClearBit(b.colors[BLACK],  n)
	b.colors[color]   = bitopts.SetBit(b.colors[color], n)
	b.pieces[piece]   = bitopts.SetBit(b.pieces[piece], n)
}

// SetSquareRankFile removes any existing piece and sets the square to the new piece/color with (x,y) coordinates.
func (b *BoardState) SetSquareRankFile(rank uint8, file uint8, color uint8, piece uint8) {
	b.SetSquare(bitopts.RankFileToSquare(rank, file), color, piece);
}
