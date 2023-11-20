// https://www.codewars.com/kata/55a2d7ebe362935a210000b2
package solutions

func SmallestIntegerFinder(numbers []int) int {
	var smallest int
	for i, value := range numbers {
		if i == 0 {
			smallest = value
		}
		if value < smallest {
			smallest = value
		}
	}
	return smallest
}
