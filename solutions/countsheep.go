// https://www.codewars.com/kata/5b077ebdaf15be5c7f000077
package solutions

import "fmt"

func countSheep(num int) string {
	str := ""
	for i := 1; i <= num; i++ {
		str += fmt.Sprintf("%d sheep...", i)
	}
	return str
}
