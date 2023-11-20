// https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad
package solutions

func Invert(arr []int) []int {
	inverted := make([]int, len(arr))
	for i, value := range arr {
		inverted[i] = value * -1
	}
	return inverted
}
