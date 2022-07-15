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

func TestRookAttackSetForOccupancy(t *testing.T) {
	// bbs := RookOccuppancyMasksForSquare(0)
	// fmt.Printf("len = %v\n", len(bbs))
	// for _, bb := range bbs {
	// 	attackSet := RookAttackSetForOccupancy(0, bb)
	// 	fmt.Printf("Blockers\n")
	// 	bitops.Print(bb, 0)
	// 	fmt.Printf("Attack Set\n")
	// 	bitops.Print(attackSet, 0)
	// 	fmt.Printf("----------\n")
	// }
	t.Errorf("not implemented")

}
