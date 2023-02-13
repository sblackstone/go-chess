package bitops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSetBitsGeneric(t *testing.T) {
	t.Parallel()
	var result []int8
	var testCase uint64

	f := func(n int8) {
		result = append(result, n)
	}

	testCase = SetBit(testCase, 5)
	testCase = SetBit(testCase, 7)
	testCase = SetBit(testCase, 63)

	FindSetBitsGeneric(testCase, f)
	expected := []int8{5, 7, 63}
	assert.Equal(t, result, expected)

}

func TestInternalMask(t *testing.T) {
	t.Parallel()
	var expected uint64
	var i, j int8
	for i = 1; i < 7; i++ {
		for j = 1; j < 7; j++ {
			expected = SetBit(expected, RankFileToSquare(i, j))
		}
	}

	im := InternalMask()
	assert.Equal(t, im, expected)
}
func TestPerimeterMask(t *testing.T) {
	t.Parallel()
	expected := RankMask(0) | RankMask(7) | FileMask(0) | FileMask(7)
	assert.Equal(t, expected, PerimeterMask())
}
func TestRankMasks(t *testing.T) {
	t.Parallel()
	var i int8
	for i = 0; i < 8; i++ {
		expected := uint64(255) << (i * 8)
		assert.Equal(t, expected, RankMask(i))
	}
}

func TestFileMasks(t *testing.T) {
	t.Parallel()
	var i int8
	for i = 0; i < 8; i++ {
		expected := Rotate90Clockwise(RankMask(i))
		assert.Equal(t, expected, FileMask(i))
	}
}
func TestMask(t *testing.T) {
	t.Parallel()
	val := Mask(1)
	assert.Equal(t, uint64(2), val)
}

func TestFindSetBits(t *testing.T) {
	var v uint64
	t.Parallel()
	v = SetBit(v, 2)
	v = SetBit(v, 4)
	v = SetBit(v, 6)
	result := FindSetBits(v)
	expected := []int8{2, 4, 6}
	assert.Equal(t, expected, result)
}

func TestFindSetBitsNone(t *testing.T) {
	t.Parallel()
	var v uint64
	result := FindSetBits(v)
	var expected []int8
	assert.Equal(t, expected, result)
}

func TestFindSetBitsExtrema(t *testing.T) {
	t.Parallel()
	var v uint64
	v = SetBit(v, 0)
	v = SetBit(v, 63)
	result := FindSetBits(v)
	expected := []int8{0, 63}
	assert.Equal(t, expected, result)
}

func TestRankFileToSquare(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		rank     int8
		file     int8
		expected int8
	}{
		{0, 0, 0},
		{7, 7, 63},
		{7, 3, 59},
		{3, 7, 31},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("square mapping %d", i), func(t *testing.T) {
			t.Parallel()
			v := RankFileToSquare(tc.rank, tc.file)
			assert.Equal(t, tc.expected, v)
		})
	}

}

func TestSquareToRankFile(t *testing.T) {
	testCases := []struct {
		rank   int8
		file   int8
		square int8
	}{
		{0, 0, 0},
		{7, 7, 63},
		{7, 3, 59},
		{3, 7, 31},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("square to rank file %d", i), func(t *testing.T) {
			t.Parallel()
			rank, file := SquareToRankFile(tc.square)
			assert.Equal(t, tc.rank, rank, "expected rank to match")
			assert.Equal(t, tc.file, file, "expected file to match")
		})
	}
}

func TestFlipBit(t *testing.T) {
	t.Parallel()
	var v uint64
	v = 0b101
	v = FlipBit(v, 1)
	assert.Equal(t, uint64(0b111), v, "Expected 101 ^ 010 = 111 (7)")
	v = FlipBit(v, 1)
	assert.Equal(t, uint64(0b101), v, "Expected 111 ^ 010 = 101 (5)")
}

func TestSetBit(t *testing.T) {
	t.Parallel()
	var v uint64
	v = SetBit(v, 0)
	assert.Equal(t, uint64(1), v)
	v = SetBit(v, 1)
	assert.Equal(t, uint64(3), v)
}

func TestClearBit(t *testing.T) {
	t.Parallel()
	var v uint64 = 7
	v = ClearBit(v, 0)
	assert.Equal(t, uint64(6), v)
	v = ClearBit(v, 1)
	assert.Equal(t, uint64(4), v)
}

func TestTestBit(t *testing.T) {
	t.Parallel()
	var v uint64 = 5
	assert.True(t, TestBit(v, 0))
	assert.False(t, TestBit(v, 1))
	assert.True(t, TestBit(v, 2))

}

func TestRankOfSquare(t *testing.T) {
	assert.Equal(t, int8(0), RankOfSquare(0))
	assert.Equal(t, int8(3), RankOfSquare(25))
	assert.Equal(t, int8(4), RankOfSquare(37))
	assert.Equal(t, int8(7), RankOfSquare(63))
}

func TestFileOfSquare(t *testing.T) {
	assert.Equal(t, int8(0), FileOfSquare(0))
	assert.Equal(t, int8(1), FileOfSquare(25))
	assert.Equal(t, int8(5), FileOfSquare(37))
	assert.Equal(t, int8(7), FileOfSquare(63))
}
