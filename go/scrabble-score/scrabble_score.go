package scrabble

import (
	"slices"
	"unicode"
)

func Score(word string) int {
	scoreChart := map[int][]rune{
		1:  {'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't'},
		2:  {'d', 'g'},
		3:  {'b', 'c', 'm', 'p'},
		4:  {'f', 'h', 'v', 'w', 'y'},
		5:  {'k'},
		8:  {'j', 'x'},
		10: {'q', 'z'},
	}

	var score int
	for _, letter := range word {
		lowercaseLetter := unicode.ToLower(letter)
		for scoreToAdd, letters := range scoreChart {
			if slices.Contains(letters, lowercaseLetter) {
				score += scoreToAdd
				break
			}
		}
	}

	return score
}
