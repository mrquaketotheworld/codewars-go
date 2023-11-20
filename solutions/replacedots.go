// https://www.codewars.com/kata/596c6eb85b0f515834000049
package solutions

import (
	"regexp"
)

func ReplaceDots(str string) string {
	return regexp.MustCompile("\\.").ReplaceAllString(str, "-")
}
