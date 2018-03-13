// Package twofer - two for one. One for you and one for me.
package twofer

// ShareWith create a sharing string. Takes a string, returns a string
func ShareWith(s string) string {
	if s == "" {
		return "One for you, one for me."
	}
	return "One for " + s + ", one for me."
}
