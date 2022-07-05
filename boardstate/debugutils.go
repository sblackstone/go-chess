package boardstate

import (
	"fmt"
	"github.com/sblackstone/go-chess/bitopts"
)

// Print outputs a debug display of the current board.
func (b *BoardState) Print(highlight int8) {
	pieces := make([][]string, 2)
	pieces[BLACK] = []string{"♖", "♘", "♗", "♕", "♔", "♙"}
	pieces[WHITE] = []string{"♜", "♞", "♝", "♛", "♚", "♟"}
	var rank, file int8
	// This next line is correct because the indexes are int8 so 0 - 1 = 255.
	for rank = 7; rank >= 0; rank-- {
		for file = 0; file < 8; file++ {
			pos := bitopts.RankFileToSquare(rank, file)
			color := b.ColorOfSquare(pos)
			//fmt.Printf(" %v ", bitopts.RankFileToSquare(rank, file))
			if pos == highlight {
				fmt.Printf(" * ")
			} else if color == EMPTY {
				fmt.Printf(" - ")
			} else {
				piece := b.PieceOfSquare(bitopts.RankFileToSquare(rank, file))
				fmt.Printf(" %s ", pieces[color][piece])
			}
		}
		fmt.Println()
	}
	fmt.Printf("\n----------------------\n\n")

}
