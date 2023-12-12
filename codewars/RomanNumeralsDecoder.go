// https://www.codewars.com/kata/51b6249c4612257ac0000005/train/go
package kata

func Decode(roman string) int {
	translator := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	total := 0
	for i, _ := range roman {
		currentVal := translator[string(roman[i])]
		nextExist := i+1 < len(roman)
		nextVal := 0
		if nextExist {
			nextVal = translator[string(roman[i+1])]
		}
		if currentVal < nextVal {
			total -= currentVal
		} else {
			total += currentVal
		}
	}
	return total
}
