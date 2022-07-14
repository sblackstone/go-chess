package boardstate

import (
	"github.com/sblackstone/go-chess/bitops"
)

// BoardState contains the state of the Board
type BoardState struct {
	colors          [2]uint64
	pieces          [6]uint64
	castleData      [2][2]bool
	enpassantSquare int8
	turn            int8
	halfMoves       int
	fullMoves       int
	moveStack       *MoveStackData
	colorList       [64]int8
	pieceList       [64]int8
	kingPos         [2]int8
	pieceLocations  PieceLocations
}

type MoveStackData struct {
	src             int8
	dst             int8
	srcPiece        int8
	dstPiece        int8
	enpassantSquare int8
	halfMoves       int
	castleData      [2][2]bool
	prev            *MoveStackData
}

func (b *BoardState) GetKingPos(color int8) int8 {
	return b.kingPos[color]
}

// Blank returns a blank board with no pieces on it
func Blank() *BoardState {
	b := BoardState{}
	b.EnableAllCastling()
	b.turn = WHITE
	b.colors = [2]uint64{0, 0}
	b.pieces = [6]uint64{0, 0, 0, 0, 0, 0}
	b.enpassantSquare = NO_ENPASSANT
	for i := 0; i < 64; i++ {
		b.colorList[i] = EMPTY
		b.pieceList[i] = EMPTY
	}
	b.kingPos[WHITE] = NO_KING
	b.kingPos[BLACK] = NO_KING
	return &b
}

// // Initial returns a board with the initial setup.
// func Initial() *BoardState {
// 	b := Blank()
// 	// These constants are pre-calculated using InitialManual (see below)...
// 	b.colors = [2]uint64{65535, 18446462598732840960}
// 	b.pieces = [6]uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
// 	b.fullMoves = 1
// 	b.EnableAllCastling()
// 	return b
// }

// Copy returns a copy of a BoardState
func (b *BoardState) Copy() *BoardState {
	boardCopy := BoardState{
		colors:          b.colors,
		pieces:          b.pieces,
		enpassantSquare: b.enpassantSquare,
		turn:            b.turn,
		halfMoves:       b.halfMoves,
		fullMoves:       b.fullMoves,
		castleData:      b.castleData,
		moveStack:       b.moveStack,
		colorList:       b.colorList,
		kingPos:         b.kingPos,
		pieceList:       b.pieceList,
		pieceLocations:  b.pieceLocations.Copy(),
	}

	return &boardCopy
}

// initialManual sets up the board manually, only used to calculate the constants for the fast version Initial.
func Initial() *BoardState {
	var j int8

	b := Blank()

	backFile := []int8{ROOK, KNIGHT, BISHOP, QUEEN, KING, BISHOP, KNIGHT, ROOK}
	for j = 0; j < 8; j++ {
		b.SetSquareRankFile(7, j, BLACK, backFile[j])
		b.SetSquareRankFile(0, j, WHITE, backFile[j])

		b.SetSquareRankFile(6, j, BLACK, PAWN)
		b.SetSquareRankFile(1, j, WHITE, PAWN)
	}
	b.EnableAllCastling()

	b.fullMoves = 1
	// b.kingPos[WHITE] = 4
	// b.kingPos[BLACK] = 60

	return b
}

func (b *BoardState) GetOccupiedBitboard() uint64 {
	return b.colors[BLACK] | b.colors[WHITE]
}

func (b *BoardState) GetColorBitboard(color int8) uint64 {
	return b.colors[color]
}

func (b *BoardState) GetPieceBitboard(color int8, piece int8) uint64 {
	return b.pieces[piece] & b.colors[color]
}

func (b *BoardState) EnemyOccupiedSquare(n int8) bool {
	c := b.ColorOfSquare(n)
	return c != EMPTY && c != b.GetTurn()
}

func (b *BoardState) EmptySquare(n int8) bool {
	return b.colorList[n] == EMPTY
}

func (b *BoardState) EmptyOrEnemyOccupiedSquare(n int8) bool {
	c := b.ColorOfSquare(n)
	return c != b.GetTurn()
}

func (b *BoardState) GenerateSuccessors(moves []*Move) []*BoardState {
	var result []*BoardState
	for _, move := range moves {
		result = append(result, b.CopyPlayTurnFromMove(move))
	}
	return result
}

func (b *BoardState) MovePiece(src int8, dst int8) {
	color := b.ColorOfSquare(src)
	piece := b.PieceOfSquare(src)
	b.SetSquare(src, EMPTY, EMPTY)
	b.SetSquare(dst, color, piece)
}

// Returns an array of positions for a given set of pieces.
func (b *BoardState) FindPieces(color int8, pieceType int8) []int8 {
	return b.pieceLocations.GetLocations(color, pieceType)
}

// ColorOfSquare returns WHITE,BLACK, or EMPTY
func (b *BoardState) ColorOfSquare(n int8) int8 {
	return b.colorList[n]
}

// PieceOfSquare t
func (b *BoardState) PieceOfSquare(n int8) int8 {
	return b.pieceList[n]
}

// SetSquare removes any existing piece and sets the square to the new piece/color.
func (b *BoardState) SetSquare(n int8, color int8, piece int8) {
	origPiece := b.PieceOfSquare(n)
	origColor := b.ColorOfSquare(n)
	if origColor != EMPTY {
		b.pieces[origPiece] = bitops.ClearBit(b.pieces[origPiece], n)
		b.colors[origColor] = bitops.ClearBit(b.colors[origColor], n)
		b.pieceLocations.RemovePieceLocation(origColor, origPiece, n)
		if origPiece == KING {
			b.kingPos[origColor] = NO_KING
		}
	}
	b.colorList[n] = color
	b.pieceList[n] = piece

	if piece == KING {
		b.kingPos[color] = n
	}
	if color != EMPTY {
		b.pieceLocations.AddPieceLocation(color, piece, n)
		b.colors[color] = bitops.SetBit(b.colors[color], n)
		b.pieces[piece] = bitops.SetBit(b.pieces[piece], n)
	}
}

// SetSquareRankFile removes any existing piece and sets the square to the new piece/color with (x,y) coordinates.
func (b *BoardState) SetSquareRankFile(rank int8, file int8, color int8, piece int8) {
	b.SetSquare(bitops.RankFileToSquare(rank, file), color, piece)
}
