// https://www.codewars.com/kata/5a00e05cc374cb34d100000d
package test

func ReverseSeq(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i
	}
	return arr
}
