// https://www.codewars.com/kata/545a4c5a61aa4c6916000755
package solutions

func Gimme(array [3]int) int {
	orig := array
	unsortedTill := len(array) - 1
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < unsortedTill; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
	}
	toFind := array[1]
	for i, value := range orig {
		if value == toFind {
			return i
		}
	}
	return 0
}
