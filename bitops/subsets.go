package bitops

// https://www.chessprogramming.org/Traversing_Subsets_of_a_Set

// Enumerates all sets with the same
func SNOOB(x uint64) uint64 {
	smallest := x & -x
	ripple := x + smallest
	ones := x ^ ripple
	ones = (ones >> 2) / smallest
	return ripple | ones
}

// Enumerates all sets with the same bits as d
// will not include zero.
func Subsets(d uint64, callback func(uint64)) {
	var n uint64 = 0
	for {
		n = (n - d) & d
		if n == 0 {
			break
		}
		callback(n)
	}
}
