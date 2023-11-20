// https://www.codewars.com/kata/554b4ac871d6813a03000035
package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numbers := strings.Split(in, " ")
	first, _ := strconv.Atoi(numbers[0])
	low, high := first, first
	for _, value := range numbers {
		current, _ := strconv.Atoi(value)
		if low > current {
			low = current
		} else if high < current {
			high = current
		}
	}
	return fmt.Sprintf("%d %d", high, low)
}
