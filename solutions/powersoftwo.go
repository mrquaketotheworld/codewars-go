// https://www.codewars.com/kata/57a083a57cb1f31db7000028
package solutions

import "math"

func PowersOfTwo(n int) []uint64 {
	powers := make([]uint64, n+1)
	for i := 0; i <= n; i++ {
		powers[i] = uint64(math.Pow(2, float64(i)))
	}
	return powers
}
