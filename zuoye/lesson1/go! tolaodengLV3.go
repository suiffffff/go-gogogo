package main

import "fmt"

var luoma = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		cur := string(s[i])
		if i < len(s)-1 {
			next := string(s[i+1])
			if luoma[cur] < luoma[next] {
				sum -= luoma[cur]
				continue
			}
		}
		sum += luoma[cur]
	}
	return sum
}
func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("LVII"))
	fmt.Println(romanToInt("IX"))
}
