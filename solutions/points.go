// https://www.codewars.com/kata/5bb904724c47249b10000131
package solutions

func Points(games []string) int {
	var sum int
	for i := 0; i < len(games); i++ {
		our := games[i][0]
		their := games[i][2]
		if our > their {
			sum += 3
		}
		if our == their {
			sum += 1
		}
	}
	return sum
}
