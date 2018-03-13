// Package bob provides speech for Bob the annoying teenager
package bob

import (
	"regexp"
)

// Hey is Bob's speech function
func Hey(remark string) string {
	if matched, _ := regexp.MatchString(`^[^a-z]*[A-Z]+[^a-z]*$`, remark); matched {
		return "Whoa, chill out!"
	}
	if matched, _ := regexp.MatchString(`(\?|\?\s+)$`, remark); matched {
		return "Sure."
	}
	if matched, _ := regexp.MatchString(`\w`, remark); matched {
		return "Whatever."
	}
	return "Fine. Be that way!"

}
