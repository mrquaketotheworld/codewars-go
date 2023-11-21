// https://www.codewars.com/kata/544675c6f971f7399a000e79
package solutions

import "strconv"

func StringToNumber(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}
