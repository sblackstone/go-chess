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
  empty      int64
  state      int32

  //wpassant   int8 // 8 bits needed
  //bpassant   int8 // 8 bits needed
  //turn       int8 // 1 bit needed
  //wcastle    int8 // 2 bits needed
  //bcastle    int8 // 2 bits needed
}

func planeToAffine(i int8, j int8) int8 {
  return i * 8 + j;
}

func BlankBoard() *BoardState {
  return &BoardState{};
}
