// https://www.codewars.com/kata/57f780909f7e8e3183000078
package solutions

func Grow(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result *= arr[i]
	}
	return result
}
