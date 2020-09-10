// Package twofer implementing twofer
package twofer

import "fmt"

// ShareWith returns the result of twofer
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
