// Package greeting stakes a string and returns a string greeting

package greeting

import "fmt"

const testVersion = 3

// HelloWorld self explantory really
func HelloWorld(name string) string {
	if name != "" {
		return fmt.Sprint("Hello, ", name, "!")
	} else {
		return "Hello, World!"
	}
}
