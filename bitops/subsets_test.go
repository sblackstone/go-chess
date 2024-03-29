package bitops

import (
	"fmt"
	"testing"
)

func TestSNOOB(t *testing.T) {
	var testCases = []struct {
		input    uint64
		expected uint64
	}{
		{1, 2},
		{2, 4},
		{4, 8},
		{8, 16},
		{3, 5},
		{5, 6},
		{6, 9},
	}

	for _, tc := range testCases {
		v := SNOOB(tc.input)
		if v != tc.expected {
			t.Errorf("Expected %v to be %v", v, tc.expected)
		}
	}

}

func TestSubsets(t *testing.T) {

	var testCases = []struct {
		input uint64
		count uint64
	}{
		{1, 1},
		{2, 1},
		{3, 3},
		{4, 1},
		{7, 7},
		{6, 3},
		{14, 7},
	}

	var counter uint64
	f := func(n uint64) {
		fmt.Printf("%0.8b\n", n)
		counter += 1
	}

	for _, tc := range testCases {
		counter = 0
		Subsets(tc.input, f)
		if counter != tc.count {
			t.Errorf("Expected %v to be %v", counter, tc.count)
		}
	}

	// fmt.Printf("Input %0.8b\n", 255)
	// Subsets(255, f)
	// t.Errorf("1")
}
