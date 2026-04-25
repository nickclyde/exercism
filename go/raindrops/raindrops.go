package raindrops

import "fmt"

func Convert(number int) string {
	var result string
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = fmt.Sprint(number)
	}
	return result
}
