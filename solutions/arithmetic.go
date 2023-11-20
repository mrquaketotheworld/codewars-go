// https://www.codewars.com/kata/583f158ea20cfcbeb400000a
package solutions

func Arithmetic(a int, b int, operator string) int {
	return map[string]func(a, b int) int{
		"add":      func(a, b int) int { return a + b },
		"multiply": func(a, b int) int { return a * b },
		"divide":   func(a, b int) int { return a / b },
		"subtract": func(a, b int) int { return a - b },
	}[operator](a, b)
}
