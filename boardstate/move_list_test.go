package boardstate

import "testing"

func TestMoveList(t *testing.T) {
	var ml MoveList
	ml.AddMove(&Move{Src: 1, Dst: 2, PromotePiece: EMPTY})
	ml.AddMove(&Move{Src: 3, Dst: 4, PromotePiece: EMPTY})
	ml.AddMove(&Move{Src: 5, Dst: 6, PromotePiece: EMPTY})

	sum := int8(0)

	for mle := ml.Head; mle != nil; mle = mle.Next {
		sum += mle.Move.Src
		sum += mle.Move.Dst
	}

	if sum != 21 {
		t.Errorf("Expected move list sum to be 21, got %v", sum)
	}

}
