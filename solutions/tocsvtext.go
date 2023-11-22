// https://www.codewars.com/kata/5a34af40e1ce0eb1f5000036
package solutions

import "fmt"

func ToCsvText(array [][]int) string {
	csv := []byte{}
	for _, innerArray := range array {
		for _, value := range innerArray {
			csv = append(csv, []byte(fmt.Sprintf("%d,", value))...)
		}
		csv = csv[:len(csv)-1]
		csv = append(csv, []byte("\n")...)
	}
	return string(csv)[:len(csv)-1]
}
