// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9
package solutions

import "strings"

func FindShort(s string) int {
	splits := strings.Split(s, " ")
	lowest := len(splits[0])
	for _, value := range splits {
		l := len(value)
		if lowest > l {
			lowest = l
		}
	}
	return lowest
}
