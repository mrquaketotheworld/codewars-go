// https://www.codewars.com/kata/58ca658cc0d6401f2700045f
package solutions

func FindMultiples(integer, limit int) []int {
	nums := make([]int, 0, limit/integer)
	for i := integer; i <= limit; i += integer {
		nums = append(nums, i)
	}
	return nums
}
