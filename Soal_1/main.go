package main

import (
	"fmt"
	"strings"
)

func main() {
	text := []string{"We Always Mekar"}

	for _, letter := range text {
		result := hitungHuruf(letter)
		textResult := ""
		for k, v := range result {
			textResult += fmt.Sprintf("%c=%d,", k, v)
		}

		textResult = strings.TrimSuffix(textResult, ", ")
		fmt.Println(textResult)

	}
}

func hitungHuruf(str string) map[rune]int {
	count := make(map[rune]int)
	str = strings.ToLower(str)

	for _, huruf := range str {
		if huruf >= 'a' && huruf <= 'z' {
			count[huruf]++
		}
	}
	return count
}
