package boardstate

import (
  "github.com/sblackstone/go-chess/bitopts"
	"fmt"
)

// Print outputs a debug display of the current board.
func (b *BoardState) Print() {
	pieces := make([][]string, 2)
	pieces[BLACK] = []string{"♖", "♘", "♗", "♕", "♔", "♙"};
	pieces[WHITE] = []string{"♜", "♞", "♝", "♛", "♚", "♟"};
	var rank, file uint8
	for rank = 0; rank < 8; rank++ {
		for file = 0; file < 8; file++ {
			color := b.ColorOfSquare(bitopts.RankFileToSquare(rank, file))
			if color == EMPTY {
				fmt.Printf(" - ")
			} else {
				piece := b.PieceOfSquare(bitopts.RankFileToSquare(rank, file))
				fmt.Printf(" %s ", pieces[color][piece])
			}
		}
		fmt.Println()
	}
}
