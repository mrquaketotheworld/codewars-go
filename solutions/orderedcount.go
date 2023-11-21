// https://www.codewars.com/kata/57a6633153ba33189e000074
package solutions

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	m := map[rune]int{}
	var charsCount int
	var keys []rune
	for _, value := range text {
		if _, ok := m[value]; !ok {
			charsCount++
			keys = append(keys, value)
		}
		m[value]++
	}
	chars := make([]Tuple, 0, charsCount)
	for _, value := range keys {
		chars = append(chars, Tuple{value, m[value]})
	}
	return chars
}
