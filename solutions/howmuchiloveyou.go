// https://www.codewars.com/kata/57f24e6a18e9fad8eb000296
package solutions

func HowMuchILoveYou(j int) string {
	petal := 0
	for i := 0; i < j; i++ {
		petal++
		if petal > 6 {
			petal = 1
		}
	}
	return [...]string{
		"I love you",
		"a little",
		"a lot",
		"passionately",
		"madly",
		"not at all"}[petal-1]
}
