package boardstate

import (
	"fmt"

	"github.com/sblackstone/go-chess/bitopts"
)

type Move struct {
	Src          int8
	Dst          int8
	PromotePiece int8
}

func MoveFromUCI(uciStr string) (*Move, error) {
	if len(uciStr) > 5 || len(uciStr) < 4 {
		return nil, fmt.Errorf("invalid UCI str: %s", uciStr)
	}
	src, err := bitopts.AlgebraicToSquare(uciStr[0:2])
	if err != nil {
		return nil, err
	}
	dst, err := bitopts.AlgebraicToSquare(uciStr[2:4])
	if err != nil {
		return nil, err
	}

	if src == dst {
		return nil, fmt.Errorf("src and dst were both %v", src)
	}
	promotePiece := int8(EMPTY)
	if len(uciStr) == 5 {
		switch uciStr[4:] {
		case "r":
			promotePiece = ROOK
		case "n":
			promotePiece = KNIGHT
		case "b":
			promotePiece = BISHOP
		case "q":
			promotePiece = QUEEN
		default:
			return nil, fmt.Errorf("invalid UCI str: %s", uciStr)
		}
	}

	return &Move{Src: src, Dst: dst, PromotePiece: promotePiece}, nil

}
