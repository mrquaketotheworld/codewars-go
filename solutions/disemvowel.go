// https://www.codewars.com/kata/52fba66badcd10859f00097e
package solutions

func Disemvowel(comment string) string {
	notVowels := []rune{}
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	for _, value := range comment {
		if !vowels[value] {
			notVowels = append(notVowels, value)
		}
	}
	return string(notVowels)
}
