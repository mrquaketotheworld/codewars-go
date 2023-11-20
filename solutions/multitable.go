// https://www.codewars.com/kata/5a2fd38b55519ed98f0000ce
package solutions

import "fmt"

func MultiTable(number int) string {
	table := ""
	for i := 1; i < 11; i++ {
		table += fmt.Sprintf("%d * %d = %d\n", i, number, number*i)
	}
	return table[:len(table)-1]
}
