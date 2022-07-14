package movegenerator

import (
	"fmt"
	"testing"

	"github.com/sblackstone/go-chess/bitops"
)

func TestRookOccuppancyMasksForSquare(t *testing.T) {
	bbs := RookOccuppancyMasksForSquare(17)
	for _, bb := range bbs {
		bitops.Print(bb, 17)
		fmt.Printf("\n\n")
	}
	t.Errorf("not implemented")
}
