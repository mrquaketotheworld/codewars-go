// https://www.codewars.com/kata/57eb8fcdf670e99d9b000272
package solutions

import "strings"

func High(s string) string {
	points := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
	}
	score := map[string]int{}
	splitted := strings.Split(s, " ")
	for _, word := range splitted {
		sum := 0
		for _, char := range word {
			sum += points[char]
		}
		score[word] = sum
	}
	var maxWord string
	var maxWordScore int
	for _, word := range splitted {
		if maxWordScore < score[word] {
			maxWord = word
			maxWordScore = score[word]
		}
	}
	for _, word := range splitted {
		if word == maxWord {
			return word
		}
	}
	return ""
}
