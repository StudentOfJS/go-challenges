package isogram

import "strings"

// IsIsogram takes a string and returns true if it is an isogram and false if it is not
func IsIsogram(rawWord string) bool {
	if rawWord == "" {
		return true
	}
	wordBytes := []byte(strings.ToLower(rawWord))
	var iso = true

	for i, char := range wordBytes {
		if char < 97 || char > 122 {
			continue
		}
		for _, c := range wordBytes[i+1:] {
			if char == c {
				iso = false
				break
			}
		}
	}
	return iso
}
