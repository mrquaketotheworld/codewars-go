// https://www.codewars.com/kata/5a34b80155519e1a00000009
package solutions

func multipleOfIndex(ints []int) []int {
	arr := []int{}
	for i, value := range ints {
		if i != 0 && value%i == 0 {
			arr = append(arr, value)
		}
	}
	return arr
}
