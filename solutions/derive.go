// https://www.codewars.com/kata/5963c18ecb97be020b0000a2
package solutions

import "strconv"

func Derive(coefficient, exponent int) string {
	product := coefficient * exponent
	return strconv.Itoa(product) + "x^" + strconv.Itoa(exponent-1)
}
