// The twofer package helps determine what to say when sharing your two-for-one deal
package twofer

import "fmt"

// What to say when sharing your two-for-one deal with someone else
func ShareWith(name string) string {
	nameOrDefault := "you"
	if name != "" {
		nameOrDefault = name
	}
	return fmt.Sprintf("One for %s, one for me.", nameOrDefault)
}
