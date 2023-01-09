package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/treesearch"
	"github.com/sblackstone/go-chess/uci"
)

func main() {
	logFile, _ := os.OpenFile("/tmp/go-chess-log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	buf := bufio.NewScanner(os.Stdin)

	var board *boardstate.BoardState

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
			move := treesearch.BestMove(board, 5)
			sendReply(fmt.Sprintf("bestmove %s", uci.MoveToUCI(move)))
		}

		if strings.HasPrefix(command, "position") {
			board = uci.BoardFromUCIPosition(command)
		}
	}
}
