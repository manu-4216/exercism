// Package hamming handles Hamming Distance
package hamming

import "fmt"

// Distance returns the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	runesA, runesB := []rune(a), []rune(b) // convert each string (slice of bytes) to rune slice

	if len(runesA) != len(runesB) {
		return 0, fmt.Errorf("strings should have the same length")
	}

	result := 0

	for i := range runesA {
		if runesA[i] != runesB[i] {
			result++
		}
	}

	return result, nil
}
