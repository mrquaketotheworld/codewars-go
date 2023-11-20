// https://www.codewars.com/kata/555086d53eac039a2a000083
package solutions

func LoveFunc(flower1, flower2 int) bool {
	sum := flower1 + flower2
	if sum%2 == 0 {
		return false
	}
	return true
}
