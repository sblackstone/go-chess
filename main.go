package main

import (
	//  "github.com/sblackstone/go-chess/boardstate"
	//"github.com/sblackstone/go-chess/bitopts"
	"bufio"
	"fmt"
	"os"
)

func main() {
	logFile, _ := os.Create("/tmp/go-chess-log")
	buf := bufio.NewScanner(os.Stdin)

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
	}
}
