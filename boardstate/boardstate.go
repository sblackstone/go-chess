package boardstate


type BoardState struct {
  white      int64
  black      int64
  pawns      int64
  knights    int64
  rooks      int64
  bishops    int64
  queens     int64
  kings      int64
  turn       int8
  wpassant   int8
  bpassant   int8
  wcastle    int8
  bcastle    int8
}

func (b *BoardState) GetWhiteSquares() int64 {
  return b.white;
}

func BlankBoard() *BoardState {
  return &BoardState{};
}
