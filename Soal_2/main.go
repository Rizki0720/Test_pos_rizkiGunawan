package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func countAndSortLetters(texts []string) string {
	count := make(map[rune]int)

	// Counting letters
	for _, text := range texts {
		for _, letter := range text {
			count[letter]++
		}
	}

	// Sorting letters by frequency and case
	keys := make([]rune, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if count[keys[i]] == count[keys[j]] {
			if unicode.IsUpper(keys[i]) && unicode.IsLower(keys[j]) {
				return false
			} else if unicode.IsLower(keys[i]) && unicode.IsUpper(keys[j]) {
				return true
			} else {
				return keys[i] < keys[j]
			}
		}
		return count[keys[i]] > count[keys[j]]
	})

	// Constructing the sorted string
	var result strings.Builder
	for _, key := range keys {
		for i := 0; i < count[key]; i++ {
			result.WriteRune(key)
		}
	}

	return result.String()
}

func main() {
	input := []string{"pendanaan", "Terproteksi"}
	output := countAndSortLetters(input)
	fmt.Println(output)
}
