package boardstate

import (
	"github.com/sblackstone/go-chess/bitopts"
	"math"
)


// BoardState contains the state of the Board
type BoardState struct {
	colors [2]uint64
	pieces [6]uint64
	enpassantCol uint8
	meta   uint64
}

// Blank returns a blank board with no pieces on it
func Blank() *BoardState {
	b := BoardState{}
	b.colors = [2]uint64{0, 0}
	b.pieces = [6]uint64{0, 0, 0, 0, 0, 0}
	b.enpassantCol = NO_ENPASSANT
	return &b
}

// Initial returns a board with the initial setup.
func Initial() *BoardState {
	b := BoardState{}
	// These constants are pre-calculated using InitialManual (see below)...
	b.colors = [2]uint64{65535, 18446462598732840960 }
	b.pieces = [6]uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
	b.enpassantCol = NO_ENPASSANT
	return &b
}

// Copy returns a copy of a BoardState
func (b *BoardState) Copy() *BoardState {
	boardCopy := BoardState{
		meta: b.meta,
		colors: b.colors,
		pieces: b.pieces,
		enpassantCol: b.enpassantCol,
	}

	return &boardCopy
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

func (b *BoardState) EnemyOccupiedSquare(n uint8) bool{
	c := b.ColorOfSquare(n)
	return c != EMPTY && c != b.GetTurn()
}

func (b *BoardState) EmptySquare(n uint8) bool {
	c := b.ColorOfSquare(n)
	return c == EMPTY
}

func (b *BoardState) EmptyOrEnemyOccupiedSquare(n uint8) bool{
	c := b.ColorOfSquare(n)
	return c != b.GetTurn()
}

func (b *BoardState) PlayTurn(src uint8, dst uint8, promotePiece uint8) {
		// TODO: We can make this only check the pawns for a few ops worth of savings.
		piece := b.PieceOfSquare(src)
		if (piece == PAWN) {
			diff := math.Abs(float64(dst) - float64(src))
			if (diff == 16) {
				b.SetEnpassant(bitopts.FileOfSquare(src))
			} else {
				b.ClearEnpassant()
			}
		} else {
			b.ClearEnpassant()
		}

		b.MovePiece(src, dst)
		// TODO: SET OR CLEAR ENPASSANT.
		// TODO: Castling rights


		if promotePiece != EMPTY {
			b.SetSquare(dst, b.GetTurn(), promotePiece)
		}
		b.ToggleTurn()

}

func (b *BoardState) CopyPlayTurn(src uint8, dst uint8, promotePiece uint8) *BoardState{
	bCopy := b.Copy()
	bCopy.PlayTurn(src, dst, promotePiece)
	return bCopy
}

func (b *BoardState) MovePiece(src uint8, dst uint8) {
	color := b.ColorOfSquare(src)
	piece := b.PieceOfSquare(src)
	b.pieces[piece]    = bitopts.ClearBit(b.pieces[piece], src)
	b.colors[color]    = bitopts.ClearBit(b.colors[color], src)
	b.colors[color]    = bitopts.SetBit(b.colors[color],   dst)
	b.pieces[piece]    = bitopts.SetBit(b.pieces[piece],   dst)
}

// Returns an array of positions for a given set of pieces.
func (b *BoardState) FindPieces(color uint8, pieceType uint8) []uint8 {
	pieceBitboard := b.colors[color] & b.pieces[pieceType]
	twoPiecePos := bitopts.FindTwoPiecePositions(pieceBitboard)

	if pieceType == PAWN {
		// [1] or []
		if len(twoPiecePos) < 2 {
			return twoPiecePos
		}
		// [8,15]
		var result []uint8
		var i uint8
		for i = twoPiecePos[0]; i <= twoPiecePos[1]; i++ {
			if bitopts.TestBit(pieceBitboard, i) {
				result = append(result, i)
			}
		}
		return result
	} else {
		return twoPiecePos
	}
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
	if (color != EMPTY) {
		b.colors[color]   = bitopts.SetBit(b.colors[color], n)
		b.pieces[piece]   = bitopts.SetBit(b.pieces[piece], n)
	}
}

// SetSquareRankFile removes any existing piece and sets the square to the new piece/color with (x,y) coordinates.
func (b *BoardState) SetSquareRankFile(rank uint8, file uint8, color uint8, piece uint8) {
	b.SetSquare(bitopts.RankFileToSquare(rank, file), color, piece);
}
