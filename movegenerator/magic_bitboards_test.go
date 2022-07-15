package movegenerator

import (
	"fmt"
	"testing"

	"github.com/sblackstone/go-chess/bitops"
)

func TestRookBlockerMasksForSquare(t *testing.T) {
	bbs := RookBlockerMasksForSquare(17)
	for _, bb := range bbs {
		bitops.Print(bb, 17)
		fmt.Printf("\n\n")
	}
	t.Errorf("not implemented")
}

func TestRookAttackSetForOccupancy(t *testing.T) {
	// bbs := RookBlockerMasksForSquare(0)
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

func TestRookFindMagic(t *testing.T) {
	var i int8
	var result [64]*MagicDefinition
	for i = 0; i < 64; i++ {
		result[i] = RookFindMagic(i)
	}
	t.Errorf("Done")
}

func TestGenerateRookMagicBitboards(t *testing.T) {
	result := GenerateRookMagicBitboards()
	fmt.Printf("%v", result)
}
