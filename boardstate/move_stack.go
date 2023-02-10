package boardstate

type MoveStackData struct {
	src             int8
	dst             int8
	srcPiece        int8
	dstPiece        int8
	enpassantSquare int8
	halfMoves       int
	castleData      [2][2]bool
}

type MoveStack struct {
	stack []MoveStackData
}

func (ms *MoveStack) Push(msd MoveStackData) {
	ms.stack = append(ms.stack, msd)
}

func (ms *MoveStack) Pop() MoveStackData {
	msd := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	return msd
}

func (ms *MoveStack) Copy() *MoveStack {
	result := &MoveStack{}
	result.stack = append(result.stack, ms.stack...)
	return result
}
