package boardstate

func setBit(n int64, pos uint8) int64 {
  n |= (1 << pos)
  return n
}

func clearBit(n int64, pos uint8) int64 {
    var mask int64 = ^(1 << pos)
    n &= mask
    return n
}

func testBit(n int64, pos uint8) bool {
    return n & (1 << pos) > 0;
}
