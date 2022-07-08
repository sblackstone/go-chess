package main

import (
	//  "github.com/sblackstone/go-chess/boardstate"
	//"github.com/sblackstone/go-chess/bitopts"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/fen"
	"github.com/sblackstone/go-chess/treesearch"
)

func main() {
	logFile, _ := os.OpenFile("/tmp/go-chess-log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	buf := bufio.NewScanner(os.Stdin)

	var board *boardstate.BoardState
	var err error

	sendReply := func(line string) {
		logFile.WriteString(line + "\n")
		fmt.Println(line)
		logFile.Sync()
	}

	for {
		buf.Scan()
		command := buf.Text()
		logFile.WriteString(command + "\n")
		logFile.Sync()

		if command == "uci" {
			sendReply("id name StephenChess 1.0")
			sendReply("id author Stephen Blackstone")
			sendReply("uciok")
		}

		if command == "isready" {
			sendReply("readyok")
		}

		if strings.HasPrefix(command, "go ") {
			move := treesearch.BestSuccessor(board, 3)
			logFile.WriteString(fmt.Sprintf("Best move: %v\n", move))
		}

		if strings.HasPrefix(command, "position") {
			command = command[9:]
			fenStr, moves, hasMoves := strings.Cut(command, "moves")
			fenStr = strings.Trim(fenStr, " ")
			moves = strings.Trim(moves, " ")
			if fenStr == "startpos" {
				board = boardstate.Initial()
			} else {
				board, err = fen.FromFEN(fenStr)
				if err != nil {
					logFile.WriteString(fmt.Sprintf("%v\n", err))
				}
			}
			if hasMoves {
				moveList := strings.Split(strings.ToLower(moves), " ")
				for _, moveStr := range moveList {
					src, _ := bitopts.AlgebraicToSquare(moveStr[0:2])
					dst, _ := bitopts.AlgebraicToSquare(moveStr[2:4])
					logFile.WriteString(fmt.Sprintf("moveStr: %v src=%v dst=%v\n", moveStr, src, dst))
					board.PlayTurn(src, dst, boardstate.EMPTY) // TODO: Promotion
				}
			}
		}
	}
}
