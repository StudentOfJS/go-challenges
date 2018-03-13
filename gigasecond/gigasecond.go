// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import "time"

// AddGigasecond gives the date a person will reach the age of 1 gigasecond
func AddGigasecond(t time.Time) time.Time {
	total := t.Unix() + 1000000000
	t = time.Unix(total, 0)

	return t
}
