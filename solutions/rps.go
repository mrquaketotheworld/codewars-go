// https://www.codewars.com/kata/5672a98bdbdd995fad00000f
package solutions

func Rps(p1, p2 string) string {
	player1Won := "Player 1 won!"
	player2Won := "Player 2 won!"
	pack := map[string]string{
		"scissors": "paper",
		"rock":     "scissors",
		"paper":    "rock",
	}
	pack2 := map[string]string{
		"paper":    "scissors",
		"scissors": "rock",
		"rock":     "paper",
	}
	if pack[p1] == p2 {
		return player1Won
	}
	if pack[p2] == p1 {
		return player2Won
	}
	if pack2[p1] == p2 {
		return player1Won
	}
	if pack2[p2] == p1 {
		return player2Won
	}
	return "Draw!"
}
