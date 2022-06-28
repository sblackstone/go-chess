package fen
import (
  //"strings"
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
	"strings"
  "fmt"
)


func ToFEN(b *boardstate.BoardState) (string, error) {
	var rank,file int8
	result := ""
	for rank = 7; rank >= 0; rank-- {
		emptyCount := 0
		for file = 0; file < 8; file++ {
			pos := bitopts.RankFileToSquare(rank, file)
			color := b.ColorOfSquare(pos)
			piece := b.PieceOfSquare(pos)
			if color == boardstate.EMPTY {
				emptyCount +=1;
			} else {
				sqStr := ""
				switch(piece) {
					case boardstate.PAWN:
						sqStr = "p"
					case boardstate.ROOK:
						sqStr = "r"
					case boardstate.KNIGHT:
						sqStr = "n"
					case boardstate.BISHOP:
						sqStr = "b"
					case boardstate.QUEEN:
						sqStr = "q"
					case boardstate.KING:
						sqStr = "k"
				}
				if color == boardstate.WHITE {
					sqStr = strings.ToUpper(sqStr)
				}
				if (emptyCount > 0) {
					result += fmt.Sprint(emptyCount);
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
	if b.GetTurn() == boardstate.WHITE {
		result += "w"
	} else {
		result += "b"
	}

	result += " "


	/// Add castling rights
  wlong  := b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG)
	wshort := b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG)
	blong  := b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG)
	bshort := b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG)

	if (!wlong && !wshort && !blong && !bshort) {
		result += "-"
	} else {
		if wshort { result += "K" }
		if wlong  { result += "Q" }
		if bshort { result += "k" }
		if blong  { result += "q" }
	}

	result += " "

	/// Add Enpassant
	enpassantSq := b.GetEnpassant()
	if enpassantSq == boardstate.NO_ENPASSANT {
		result += "-"
	} else {
		result += bitopts.SquareToAlgebraic(enpassantSq)
	}

	result += " "

	/// Add Clocks
	result += fmt.Sprint(b.GetHalfMoves())
	result += " "
	result += fmt.Sprint(b.GetFullMoves())

	return result, nil

}
