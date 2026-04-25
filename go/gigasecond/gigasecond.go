// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond calculates the time one billion seconds after the given time.
func AddGigasecond(t time.Time) time.Time {
	d, err := time.ParseDuration("1000000000s")
	if err != nil {
		return t
	}
	return t.Add(d)
}
