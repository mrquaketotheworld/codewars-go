// https://www.codewars.com/kata/56747fd5cb988479af000028
package solutions

func GetMiddle(s string) string {
	half := len(s) / 2
	if len(s)%2 == 0 {
		return s[half-1 : half+1]
	}
	return s[half : half+1]
}
