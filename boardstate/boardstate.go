package boardstate

import (
	"os"

	"github.com/sblackstone/go-chess/bitops"
)

// BoardState contains the state of the Board
type BoardState struct {
	colors           [2]uint64
	pieces           [6]uint64
	castleData       [2][2]bool
	enpassantSquare  int8
	turn             int8
	halfMoves        int
	fullMoves        int
	moveStack        []MoveStackData
	moveStackNextIdx int
	colorList        [64]int8
	pieceList        [64]int8
	kingPos          [2]int8
	PieceLocations   PieceLocations
	zorbistKey       uint64
}

type MoveStackData struct {
	src             int8
	dst             int8
	srcPiece        int8
	dstPiece        int8
	enpassantSquare int8
	halfMoves       int
	castleData      [2][2]bool
}

func (b *BoardState) GetKingPos(color int8) int8 {
	return b.kingPos[color]
}

// Blank returns a blank board with no pieces on it
func Blank() *BoardState {
	b := BoardState{}
	b.EnableAllCastling()
	b.SetTurn(WHITE)
	b.colors = [2]uint64{0, 0}
	b.pieces = [6]uint64{0, 0, 0, 0, 0, 0}
	b.SetEnpassant(NO_ENPASSANT)
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
		colors:           b.colors,
		pieces:           b.pieces,
		castleData:       b.castleData,
		enpassantSquare:  b.enpassantSquare,
		turn:             b.turn,
		halfMoves:        b.halfMoves,
		fullMoves:        b.fullMoves,
		moveStack:        b.moveStack,
		moveStackNextIdx: 0,
		colorList:        b.colorList,
		pieceList:        b.pieceList,
		kingPos:          b.kingPos,
		PieceLocations:   b.PieceLocations.Copy(),
		zorbistKey:       b.zorbistKey,
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

	b.fullMoves = 1
	return b
}

func (b *BoardState) GetOccupiedBitboard() uint64 {
	return b.colors[BLACK] | b.colors[WHITE]
}

func (b *BoardState) GetColorBitboard(color int8) uint64 {
	return b.colors[color]
}

func (b *BoardState) GetZorbistKey() uint64 {
	return b.zorbistKey
}

func (b *BoardState) UpdateZorbistKey(value uint64) {
	b.zorbistKey ^= value
}

func (b *BoardState) GetPieceBitboard(color int8, piece int8) uint64 {
	return b.pieces[piece] & b.colors[color]
}

func (b *BoardState) EnemyOccupiedSquare(n int8) bool {
	if n > 63 {
		b.Print(127)
		os.Exit(-1)
	}
	return b.colorList[n] == b.EnemyColor()
}

func (b *BoardState) EmptySquare(n int8) bool {
	return b.colorList[n] == EMPTY
}

func (b *BoardState) EmptyOrEnemyOccupiedSquare(n int8) bool {
	return b.colorList[n] != b.GetTurn()
}

func (b *BoardState) MovePiece(src int8, dst int8) {
	color := b.ColorOfSquare(src)
	piece := b.PieceOfSquare(src)
	b.SetSquare(src, EMPTY, EMPTY)
	b.SetSquare(dst, color, piece)
}

// Returns an array of positions for a given set of pieces.
func (b *BoardState) FindPieces(color int8, pieceType int8) []int8 {
	return b.PieceLocations.GetLocations(color, pieceType)
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
		b.UpdateZorbistKey(zorbistPieces[origColor][origPiece][n])
		b.pieces[origPiece] = bitops.ClearBit(b.pieces[origPiece], n)
		b.colors[origColor] = bitops.ClearBit(b.colors[origColor], n)
		b.PieceLocations.RemovePieceLocation(origColor, origPiece, n)
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
		b.UpdateZorbistKey(zorbistPieces[color][piece][n])
		b.PieceLocations.AddPieceLocation(color, piece, n)
		b.colors[color] = bitops.SetBit(b.colors[color], n)
		b.pieces[piece] = bitops.SetBit(b.pieces[piece], n)
	}
}

// SetSquareRankFile removes any existing piece and sets the square to the new piece/color with (x,y) coordinates.
func (b *BoardState) SetSquareRankFile(rank int8, file int8, color int8, piece int8) {
	b.SetSquare(bitops.RankFileToSquare(rank, file), color, piece)
}
