package arrops

import "fmt"

// Flip takes a 8x8 array and flips it along its horizontal axis.
func Flip[T any](a *[64]T) {
	for i := 0; i < 32; i++ {
		t := a[i]
		a[i] = a[i^56]
		a[i^56] = t
	}
}

// PrintWhiteBottom takes an array and prints its values from the perspective of white.
func PrintWhiteBottom[T any](a *[64]T) {
	for i := 7; i >= 0; i-- {
		for j := 0; j < 8; j++ {
			fmt.Printf(" %v ", a[i*8+j])
		}
		fmt.Println()
	}

}
