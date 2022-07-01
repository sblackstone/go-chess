package boardstate

import (
	"github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)

type MoveData struct {
	src int8
	dst int8
	srcColor int8
	srcPiece int8
	dstColor int8
	dstPiece int8
	preHalfMoves int
	preMoveEnpassantSquare int8
	preMoveMeta uint64
}

// BoardState contains the state of the Board
type BoardState struct {
	colors [2]uint64
	pieces [6]uint64
	enpassantSquare int8
	meta   uint64
	turn int8
	halfMoves int
	fullMoves int
	movesData []MoveData;
}

// Blank returns a blank board with no pieces on it
func Blank() *BoardState {
	b := BoardState{}
	b.turn = WHITE
	b.colors = [2]uint64{0, 0}
	b.pieces = [6]uint64{0, 0, 0, 0, 0, 0}
	b.enpassantSquare = NO_ENPASSANT
	return &b
}

// Initial returns a board with the initial setup.
func Initial() *BoardState {
	b := Blank()
	// These constants are pre-calculated using InitialManual (see below)...
	b.colors = [2]uint64{65535, 18446462598732840960 }
	b.pieces = [6]uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
	b.fullMoves = 1
	return b
}

// Copy returns a copy of a BoardState
func (b *BoardState) Copy() *BoardState {
	boardCopy := BoardState{
		meta: b.meta,
		colors: b.colors,
		pieces: b.pieces,
		enpassantSquare: b.enpassantSquare,
		turn: b.turn,
		halfMoves: b.halfMoves,
		fullMoves: b.fullMoves,
		movesData: b.movesData,
	}

	return &boardCopy
}

// initialManual sets up the board manually, only used to calculate the constants for the fast version Initial.
func initialManual() *BoardState {
	var j int8

	b := Blank()

	backFile := []int8{ROOK, KNIGHT, BISHOP, QUEEN, KING, BISHOP, KNIGHT, ROOK}
	for j = 0; j < 8; j++ {
		b.SetSquareRankFile(7, j, BLACK, backFile[j])
		b.SetSquareRankFile(0, j, WHITE, backFile[j])

		b.SetSquareRankFile(6, j, BLACK, PAWN)
		b.SetSquareRankFile(1, j, WHITE, PAWN)
	}

	return b
}

func (b *BoardState) GetColorBitboard(color int8) uint64 {
	return b.colors[color]
}

func (b *BoardState) EnemyOccupiedSquare(n int8) bool{
	c := b.ColorOfSquare(n)
	return c != EMPTY && c != b.GetTurn()
}

func (b *BoardState) EmptySquare(n int8) bool {
	c := b.ColorOfSquare(n)
	return c == EMPTY
}

func (b *BoardState) EmptyOrEnemyOccupiedSquare(n int8) bool{
	c := b.ColorOfSquare(n)
	return c != b.GetTurn()
}

func (b *BoardState) GenerateSuccessors(moves []*Move) []*BoardState {
	var result []*BoardState;
	for _, move := range(moves) {
		result = append(result, b.CopyPlayTurnFromMove(move))
	}
	return result
}

func (b *BoardState) ClearSquare(n int8) {
	b.pieces[b.PieceOfSquare(n)]    = bitopts.ClearBit(b.pieces[b.PieceOfSquare(n)], n)
	b.colors[b.ColorOfSquare(n)]    = bitopts.ClearBit(b.colors[b.ColorOfSquare(n)], n)
}

func (b *BoardState) MovePiece(src int8, dst int8) {
	color := b.ColorOfSquare(src)
	piece := b.PieceOfSquare(src)

  // Clear the source square.
	b.pieces[piece]    = bitopts.ClearBit(b.pieces[piece], src)
	b.colors[color]    = bitopts.ClearBit(b.colors[color], src)

	// Clear the destination square.
	if !b.EmptySquare(dst) {
		b.pieces[b.PieceOfSquare(dst)]    = bitopts.ClearBit(b.pieces[b.PieceOfSquare(dst)], dst)
		b.colors[b.ColorOfSquare(dst)]    = bitopts.ClearBit(b.colors[b.ColorOfSquare(dst)], dst)
	}

	// Set the new piece.
	b.colors[color]    = bitopts.SetBit(b.colors[color],   dst)
	b.pieces[piece]    = bitopts.SetBit(b.pieces[piece],   dst)
}


// Returns an array of positions for a given set of pieces.
func (b *BoardState) FindPieces(color int8, pieceType int8) []int8 {
	// TODO: We can't rely on using upper/lower if we aren't sure there hasn't been promotions
	// If we track promotions, we can make this go faster by skipping the loop between upper and lower.
	pieceBitboard := b.colors[color] & b.pieces[pieceType]
	twoPiecePos := bitopts.FindTwoPiecePositions(pieceBitboard)

	// [1] or []
	if len(twoPiecePos) < 2 {
		return twoPiecePos
	}
	// [8,15]
	var result []int8
	var i int8
	for i = twoPiecePos[0]; i <= twoPiecePos[1]; i++ {
		if bitopts.TestBit(pieceBitboard, i) {
			result = append(result, i)
		}
	}
	return result

}

// ColorOfSquare returns WHITE,BLACK, or EMPTY
func (b *BoardState) ColorOfSquare(n int8) int8 {
	if bitopts.TestBit(b.colors[WHITE], n) {
		return WHITE
	}
	if bitopts.TestBit(b.colors[BLACK], n) {
		return BLACK
	}
	return EMPTY
}

// PieceOfSquare t
func (b *BoardState) PieceOfSquare(n int8) int8 {
	var i int8
	for i = 0; i < 6; i++ {
		if bitopts.TestBit(b.pieces[i], n) {
			return i
		}
	}
	return EMPTY
}

// SetSquare removes any existing piece and sets the square to the new piece/color.
func (b *BoardState) SetSquare(n int8, color int8, piece int8) {
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
func (b *BoardState) SetSquareRankFile(rank int8, file int8, color int8, piece int8) {
	b.SetSquare(bitopts.RankFileToSquare(rank, file), color, piece);
}
