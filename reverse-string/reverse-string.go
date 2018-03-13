package reverse

// String takes a string reverses it and returns the new string
func String(s string) string {
	r := []rune(s)
	var t []rune
	for i := len(r) - 1; i >= 0; i-- {
		t = append(t, r[i])
	}
	return string(t)
}
