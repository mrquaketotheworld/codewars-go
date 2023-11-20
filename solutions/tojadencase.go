// https://www.codewars.com/kata/5390bac347d09b7da40006f6
package solutions

import "strings"

func ToJadenCase(str string) string {
	splitted := strings.Split(str, " ")
	titleized := ""
	for _, value := range splitted {
		titleized += strings.ToUpper(string(value[0])) + value[1:] + " "
	}
	return titleized[:len(titleized)-1]
}
