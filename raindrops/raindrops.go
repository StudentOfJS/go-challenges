package raindrops

import (
	"strconv"
)

// Convert takes an int and returns a string related to it's factors
func Convert(num int) string {
	var song string

	if num%3 == 0 {
		song = "Pling"
	}
	if num%5 == 0 {
		song = song + "Plang"
	}
	if num%7 == 0 {
		song = song + "Plong"
	}
	if len(song) == 0 {
		return strconv.Itoa(num)
	}
	return song
}
