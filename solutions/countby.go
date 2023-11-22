// https://www.codewars.com/kata/5513795bd3fafb56c200049e
package solutions

func CountBy(x, n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i != 0 {
			result[i] = result[i-1] + x
		} else {
			result[i] = x
		}
	}
	return result
}
