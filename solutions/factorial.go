// https://www.codewars.com/kata/57a049e253ba33ac5e000212
package solutions

func Factorial(n int) int {
	product := 1
	for i := 1; i <= n; i++ {
		product *= i
	}
	return product
}
