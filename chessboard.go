package chessboard

// File declares a type named File which stores if a piece occupies a square - this will be a slice of bools
type File []bool

// Chessboard declares a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard within the given file
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, v := range cb[file] {
		if v {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard, within the given rank
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	count := 0
	for _, v := range cb {
		if v[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	for _, _ = range cb {
		return len(cb) * 8
	}
	return 0
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	total := 0
	for i, _ := range cb {
		total += CountInFile(cb, i)
	}

	return total
}
