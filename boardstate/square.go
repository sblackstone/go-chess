package boardstate

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sblackstone/go-chess/bitops"
)

func SquareToAlgebraic(pos int8) string {
	rank, file := bitops.SquareToRankFile(pos)
	ranks := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	files := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return files[int(file)] + ranks[int(rank)]
}

func AlgebraicToSquare(algrbraicSquare string) (int8, error) {
	var result int8

	if len(algrbraicSquare) != 2 {
		return -1, errors.New("algrbraicSquare must be a string of len 2 " + fmt.Sprint(result))
	}

	algrbraicSquare = strings.ToLower(algrbraicSquare)
	var parts = strings.Split(algrbraicSquare, "")
	result = int8(algrbraicSquare[0]) - int8('a')

	if result < 0 || result > 7 {
		return -1, errors.New("Could not convert square value " + algrbraicSquare + ": Invalid file: " + fmt.Sprint(result))
	}

	rank, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, err
	}
	rank -= 1

	if rank > 7 {
		return -1, errors.New("Could not convert square value " + algrbraicSquare + ": Invalid Rank: " + fmt.Sprint(rank))
	}

	result += int8(rank) * 8
	return result, nil
}
