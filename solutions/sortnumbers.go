// https://www.codewars.com/kata/5174a4c0f2769dd8b1000003
package solutions

func SortNumbers(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		lowestNumberIndex := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}
		if lowestNumberIndex != i {
			temp := numbers[i]
			numbers[i] = numbers[lowestNumberIndex]
			numbers[lowestNumberIndex] = temp
		}
	}
	return numbers
}
