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
	var i, j uint8
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			color := b.ColorOfSquare(bitopts.GridToLinear(i, j))
			if color == EMPTY {
				fmt.Printf(" - ")
			} else {
				piece := b.PieceOfSquare(bitopts.GridToLinear(i, j))
				fmt.Printf(" %s ", pieces[color][piece])
			}
		}
		fmt.Println()
	}
}
