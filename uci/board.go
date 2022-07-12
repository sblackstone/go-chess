package uci

import (
	"fmt"
	"strings"

	"github.com/sblackstone/go-chess/boardstate"
)

func RemovePositionPrefix(positionStr string) string {
	return strings.Trim(strings.TrimPrefix(positionStr, "position"), " ")
}

func BoardFromUCIPosition(positionStr string) *boardstate.BoardState {
	positionStr = RemovePositionPrefix(positionStr)
	fenStr, moves, hasMoves := strings.Cut(positionStr, "moves")
	var board *boardstate.BoardState
	var err error
	fenStr = strings.Trim(fenStr, " ")
	moves = strings.Trim(moves, " ")
	if fenStr == "startpos" {
		board = boardstate.Initial()
	} else {
		board, err = boardstate.FromFEN(fenStr)
		if err != nil {
			panic(fmt.Sprintf("%v\n", err))
		}
	}
	if hasMoves {
		moveList := strings.Split(strings.ToLower(moves), " ")
		for _, moveStr := range moveList {
			move, err := MoveFromUCI(moveStr)
			if err != nil {
				panic(fmt.Sprintf("Unknown moveStr: %s", moveStr))
			}
			board.PlayTurnFromMove(move)
		}
	}

	return board
}
