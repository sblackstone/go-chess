package boardstate

import (
	//"strings"
	"errors"
	"regexp"

	"github.com/sblackstone/go-chess/bitops"

	//"fmt"
	"strconv"
)

func applyBoardString(b *BoardState, boardStr string) error {
	var rank, file int8
	var j int

	rank = 7

	addPiece := func(r int8, f int8, color int8, piece int8) {
		b.SetSquare(bitops.RankFileToSquare(r, f), color, piece)
		file += 1
	}

	for j = range boardStr {
		char := string(boardStr[j])
		switch char {
		case "p":
			addPiece(rank, file, BLACK, PAWN)
		case "r":
			addPiece(rank, file, BLACK, ROOK)
		case "n":
			addPiece(rank, file, BLACK, KNIGHT)
		case "b":
			addPiece(rank, file, BLACK, BISHOP)
		case "q":
			addPiece(rank, file, BLACK, QUEEN)
		case "k":
			addPiece(rank, file, BLACK, KING)

		case "P":
			addPiece(rank, file, WHITE, PAWN)
		case "R":
			addPiece(rank, file, WHITE, ROOK)
		case "N":
			addPiece(rank, file, WHITE, KNIGHT)
		case "B":
			addPiece(rank, file, WHITE, BISHOP)
		case "Q":
			addPiece(rank, file, WHITE, QUEEN)
		case "K":
			addPiece(rank, file, WHITE, KING)
		case "1":
			file += 1
		case "2":
			file += 2
		case "3":
			file += 3
		case "4":
			file += 4
		case "5":
			file += 5
		case "6":
			file += 6
		case "7":
			file += 7
		case "8":
			file += 8
		case "/":
			rank -= 1
			file = 0
		default:
			return errors.New("unexpected character in FEN string: " + char)
		}
	}

	if rank != 0 || file != 8 {
		return errors.New("invalid FEN when parsing board string")
	}
	return nil

}

func applyTurnString(b *BoardState, turnString string) error {
	// regex ensures this is already a w or b.
	if turnString == "w" {
		b.SetTurn(WHITE)
	} else if turnString == "b" {
		b.SetTurn(BLACK)
	}
	return nil
}

func applyCastlingString(b *BoardState, castlingString string) error {
	b.SetCastleRights(BLACK, CASTLE_LONG, false)
	b.SetCastleRights(WHITE, CASTLE_LONG, false)
	b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	// We're sure its one of these cases via the REGEX
	for _, char := range castlingString {
		switch string(char) {
		case "-":
			return nil
		case "k":
			b.SetCastleRights(BLACK, CASTLE_SHORT, true)
		case "q":
			b.SetCastleRights(BLACK, CASTLE_LONG, true)
		case "K":
			b.SetCastleRights(WHITE, CASTLE_SHORT, true)
		case "Q":
			b.SetCastleRights(WHITE, CASTLE_LONG, true)
		}
	}
	return nil
}

func applyEnpassantString(b *BoardState, enpassantString string) error {
	if enpassantString == "-" {
		return nil
	}
	val, err := AlgebraicToSquare(enpassantString)
	if err != nil {
		return err
	}
	b.SetEnpassant(val)
	return nil

}

func FromFEN(fenString string) (*BoardState, error) {
	b := Blank()
	re := regexp.MustCompile(`([^\s]+)([\s]{1})([wb]+)([\s]{1})([-KQkq]+)([\s]{1})([-a-z0-9]+)([\s]{1})([0-9]*)([\s]{1})([0-9]*)`)
	m := re.FindStringSubmatch(fenString)

	if len(m) != 12 {
		return nil, errors.New("Invalid FEN: " + fenString)
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
