package boardstate

import "fmt"

type Move struct {
	Src          int8
	Dst          int8
	PromotePiece int8
}

func (m *Move) Print() {
	fmt.Printf("%s %s %d", SquareToAlgebraic(m.Src), SquareToAlgebraic(m.Dst), m.PromotePiece)
}

func (m *Move) ToString() string {
	return fmt.Sprintf("%s %s %d", SquareToAlgebraic(m.Src), SquareToAlgebraic(m.Dst), m.PromotePiece)
}
