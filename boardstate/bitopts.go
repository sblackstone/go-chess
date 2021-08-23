package boardstate

func setBit(n uint64, pos uint8) uint64 {
  n |= (1 << pos)
  return n
}

func clearBit(n uint64, pos uint8) uint64 {
    var mask uint64 = ^(1 << pos)
    n &= mask
    return n
}

func testBit(n uint64, pos uint8) bool {
    return n & (1 << pos) > 0;
}

func flipBit(n uint64, pos uint8) uint64 {
  n ^= (1 << pos)
  return n
}
