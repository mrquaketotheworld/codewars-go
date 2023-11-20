// https://www.codewars.com/kata/55d24f55d7dd296eb9000030
package solutions

func Summation(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
