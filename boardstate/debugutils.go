package boardstate

import (
	"fmt"

	"github.com/sblackstone/go-chess/bitops"
)

func (b *BoardState) PrintFen(withMoveCounts bool) {
	f, _ := b.ToFEN(withMoveCounts)
	fmt.Printf("%s\n", f)
}

// Print outputs a debug display of the current board.
func (b *BoardState) Print(highlight int8) {
	if b.GetTurn() == WHITE {
		fmt.Printf("WHITE to move\n")
	} else {
		fmt.Printf("BLACK to move\n")
	}
	pieces := make([][]string, 2)
	pieces[BLACK] = []string{"♖", "♘", "♗", "♕", "♔", "♙"}
	pieces[WHITE] = []string{"♜", "♞", "♝", "♛", "♚", "♟"}
	var rank, file int8
	// This next line is correct because the indexes are int8 so 0 - 1 = 255.
	for rank = 7; rank >= 0; rank-- {
		for file = 0; file < 8; file++ {
			pos := bitops.RankFileToSquare(rank, file)
			color := b.ColorOfSquare(pos)
			//fmt.Printf(" %v ", bitops.RankFileToSquare(rank, file))
			if pos == highlight {
				fmt.Printf(" * ")
			} else if color == EMPTY {
				fmt.Printf(" - ")
			} else {
				piece := b.PieceOfSquare(bitops.RankFileToSquare(rank, file))
				fmt.Printf(" %s ", pieces[color][piece])
			}
		}
		fmt.Println()
	}
	fmt.Printf("\n----------------------\n\n")

}
