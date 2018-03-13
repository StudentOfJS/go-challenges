// Package acronym takes a string and returns a string
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	strArray := strings.Fields(s)
	var a string
	f := func(r rune) bool {
		return r == '-'
	}
	for i := range strArray {
		h := strings.IndexFunc(strArray[i], f)
		if h != -1 {
			a += string(strArray[i][0]) + string(strArray[i][h+1])
		} else {
			a += string(strArray[i][0])
		}
	}
	return strings.ToUpper(a)
}
