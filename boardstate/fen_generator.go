package boardstate

import (
	//"strings"
	"fmt"
	"strings"

	"github.com/sblackstone/go-chess/bitops"
)

func (b *BoardState) ToFEN(withMoveCounts bool) (string, error) {
	var rank, file int8
	result := ""
	for rank = 7; rank >= 0; rank-- {
		emptyCount := 0
		for file = 0; file < 8; file++ {
			pos := bitops.RankFileToSquare(rank, file)
			color := b.ColorOfSquare(pos)
			piece := b.PieceOfSquare(pos)
			if color == EMPTY {
				emptyCount += 1
			} else {
				sqStr := ""
				switch piece {
				case PAWN:
					sqStr = "p"
				case ROOK:
					sqStr = "r"
				case KNIGHT:
					sqStr = "n"
				case BISHOP:
					sqStr = "b"
				case QUEEN:
					sqStr = "q"
				case KING:
					sqStr = "k"
				}
				if color == WHITE {
					sqStr = strings.ToUpper(sqStr)
				}
				if emptyCount > 0 {
					result += fmt.Sprint(emptyCount)
					emptyCount = 0
				}
				result += sqStr
			}
		}
		if emptyCount > 0 {
			result += fmt.Sprint(emptyCount)
		}
		if rank != 0 {
			result += "/"
		}
	}

	result += " "
	if b.GetTurn() == WHITE {
		result += "w"
	} else {
		result += "b"
	}

	result += " "

	/// Add castling rights
	wlong := b.HasCastleRights(WHITE, CASTLE_LONG)
	wshort := b.HasCastleRights(WHITE, CASTLE_SHORT)
	blong := b.HasCastleRights(BLACK, CASTLE_LONG)
	bshort := b.HasCastleRights(BLACK, CASTLE_SHORT)

	if !wlong && !wshort && !blong && !bshort {
		result += "-"
	} else {
		if wshort {
			result += "K"
		}
		if wlong {
			result += "Q"
		}
		if bshort {
			result += "k"
		}
		if blong {
			result += "q"
		}
	}

	result += " "

	/// Add Enpassant
	enpassantSq := b.GetEnpassant()
	if enpassantSq == NO_ENPASSANT {
		result += "-"
	} else {
		result += SquareToAlgebraic(enpassantSq)
	}

	if withMoveCounts {
		result += " "
		/// Add Clocks
		result += fmt.Sprint(b.GetHalfMoves())
		result += " "
		result += fmt.Sprint(b.GetFullMoves())
	}

	return result, nil

}
