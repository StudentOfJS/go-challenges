package scrabble

import "strings"

// Score returns the scrabble word score of a single turn
func Score(word string) int {
	letters := strings.Split(strings.ToUpper(word), "")
	var totalScore int
	for _, letter := range letters {
		switch letter {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			totalScore++
		case "D", "G":
			totalScore = totalScore + 2
		case "B", "C", "M", "P":
			totalScore = totalScore + 3
		case "F", "H", "V", "W", "Y":
			totalScore = totalScore + 4
		case "K":
			totalScore = totalScore + 5
		case "J", "X":
			totalScore = totalScore + 8
		case "Q", "Z":
			totalScore = totalScore + 10
		default:
			continue
		}
	}
	return totalScore
}
