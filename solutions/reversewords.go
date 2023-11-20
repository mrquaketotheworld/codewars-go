// https://www.codewars.com/kata/5259b20d6021e9e14c0010d4
package solutions

func ReverseWords(str string) string {
	sentence := ""
	word := ""
	for _, value := range str {
		if value == ' ' {
			sentence += word + " "
			word = ""
		} else {
			word = string(value) + word
		}
	}
	return sentence + word
}
