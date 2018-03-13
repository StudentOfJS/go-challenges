package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Valid takes a string representation of a number and determines whether or not it is valid per the Luhn formula.
func Valid(str string) bool {
	var space = regexp.MustCompile(`\s+`)
	var test = regexp.MustCompile(`\D+`)
	newStr := space.ReplaceAllString(str, "")

	if test.MatchString(newStr) || len(newStr) < 2 {
		return false
	}
	nums := strings.Split(newStr, "")
	even := len(nums)%2 == 0
	var sum int

	for i, n := range nums {
		number, _ := strconv.Atoi(n)
		if (even == true && i%2 == 0) || (even == false && i%2 != 0) {
			double := number * 2
			if double > 9 {
				sum += (double - 9)
			} else {
				sum += double
			}
		} else {
			sum += number
		}
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
