package fenparser

import (
  //"strings"
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
	"regexp"
	"strings"
  "errors"
  "fmt"
	"strconv"
)

func applyBoardString(b *boardstate.BoardState, boardStr string) error {
	var rank,file int8
  var j int

	rank = 7


	addPiece := func(r int8, f int8, color int8, piece int8) {
		b.SetSquare(bitopts.RankFileToSquare(r,f), color, piece)
		file += 1
	}

	for j = range(boardStr) {
		char := string(boardStr[j])
		//fmt.Printf("%v %v %v\n", rank, file, char)

		if (char == "p") {
			addPiece(rank, file, boardstate.BLACK, boardstate.PAWN)
		}
		if (char == "r") {
			addPiece(rank, file, boardstate.BLACK, boardstate.ROOK)
		}

		if (char == "n") {
			addPiece(rank, file, boardstate.BLACK, boardstate.KNIGHT)
		}
		if (char == "b") {
			addPiece(rank, file, boardstate.BLACK, boardstate.BISHOP)
		}
		if (char == "q") {
			addPiece(rank, file, boardstate.BLACK, boardstate.QUEEN)
		}
		if (char == "k") {
			addPiece(rank, file, boardstate.BLACK, boardstate.KING)
		}

		if (char == "P") {
			addPiece(rank, file, boardstate.WHITE, boardstate.PAWN)
		}
		if (char == "R") {
			addPiece(rank, file, boardstate.WHITE, boardstate.ROOK)
		}

		if (char == "N") {
			addPiece(rank, file, boardstate.WHITE, boardstate.KNIGHT)
		}
		if (char == "B") {
			addPiece(rank, file, boardstate.WHITE, boardstate.BISHOP)
		}
		if (char == "Q") {
			addPiece(rank, file, boardstate.WHITE, boardstate.QUEEN)
		}
		if (char == "K") {
			addPiece(rank, file, boardstate.WHITE, boardstate.KING)
		}

		if (char == "1") {
				file += 1
		}
		if (char == "2") {
				file += 2
		}
		if (char == "3") {
				file += 3
		}
		if (char == "4") {
				file += 4
		}
		if (char == "5") {
				file += 5
		}
		if (char == "6") {
				file += 6
		}
		if (char == "7") {
				file += 7
		}
		if (char == "8") {
				file = 8
		}
		if char == "/" {
			rank -= 1
			file = 0
		}
	}

	if (rank != 0 || file != 8) {
		return errors.New("Invalid FEN when parsing board string")
	}
	return nil

}

func applyTurnString(b *boardstate.BoardState, turnString string) error {
	if turnString == "w" {
    b.SetTurn(boardstate.WHITE)
  } else if turnString == "b" {
    b.SetTurn(boardstate.BLACK)
  } else {
    return errors.New("Expected value for turn: " + turnString)
  }
	return nil
}

func applyCastlingString(b *boardstate.BoardState, castlingString string) error {
	var j int
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)
	for j = range(castlingString) {
		char := string(castlingString[j])
		if char == "k" {
			b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT, false)
		} else if char == "q" {
			b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG, false)
		} else if char == "K" {
			b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)
		} else if char == "Q" {
			b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
		} else if char != "-" {
			return errors.New("Unknown character in casting string: " + char)
		}
	}
	return nil
}

func applyEnpassantString(b *boardstate.BoardState, enpassantString string) error {
	if enpassantString == "-" {
		return nil
	}
	val, err := bitopts.AlgebraicToSquare(enpassantString)
	if err != nil {
		return err
	}
	b.SetEnpassant(val)
	return nil

}


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
		result += "/"
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

func FromFEN(fenString string) (*boardstate.BoardState, error) {
  b := boardstate.Blank()
	fmt.Println(fenString)

	re := regexp.MustCompile("([^\\s]+)([\\s]{1})([wb]+)([\\s]{1})([-KQkq]+)([\\s]{1})([-a-z0-9]+)([\\s]{1})([0-9]*)([\\s]{1})([0-9]*)")
	m := re.FindStringSubmatch(fenString)

	if (len(m) != 12) {
		return nil, errors.New("Invalid FEN: " + fenString )
	}
	boardStr := m[1]
	turnStr := m[3]
	castlingString := m[5]
	enpassantString := m[7]

	err := applyBoardString(b, boardStr)

	if err != nil {
		return nil, err
	}

	err = applyTurnString(b, turnStr)
	if err != nil {
		return nil, err
	}

	err = applyCastlingString(b, castlingString)
	if err != nil {
		return nil, err
	}

	err = applyEnpassantString(b, enpassantString)
	if err != nil {
		return nil, err
	}



	//enpassantSquare := m[7]
  halfMoveClock := m[9]
	fullMoveNumber := m[11]

	halfVal, err := strconv.Atoi(halfMoveClock)
	if err != nil {
		return nil, err
	}
	b.SetHalfMoves(halfVal)

	fullVal, err := strconv.Atoi(fullMoveNumber)
	if err != nil {
		return nil, err
	}
	b.SetFullMoves(fullVal)





	//for i := range(m) {
	//	fmt.Printf("%v: %v\n", i, m[i])
	//}

  return b, nil
}
