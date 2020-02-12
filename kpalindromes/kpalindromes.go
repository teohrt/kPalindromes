package kpalindromes

import "strings"

// Kpalindromes solution complexity:
// O(n) time | O(n) space
func Kpalindromes(k int, str string) bool {
	n := len(str)
	arr := strings.Split(str, "")
	charCounts := map[string]int{}
	for _, s := range arr {
		if _, ok := charCounts[s]; ok {
			charCounts[s]++
		} else {
			charCounts[s] = 1
		}
	}
	if len(charCounts) > (n / 2) {
		return false
	}
	oddCharacterCount := 0
	for key, val := range charCounts {
		if val%2 != 0 {
			oddCharacterCount++
		}
		charCounts[key] = charCounts[key] / 2
	}
	if oddCharacterCount > 1 {
		return false
	}
	evenInput := n%2 == 0
	if evenInput && oddCharacterCount > 0 {
		return false
	}
	if !evenInput && oddCharacterCount != 1 {
		return false
	}
	possiblePalendromeCount := 0
	for i := 0; i < n/2; i++ {
		count := len(charCounts) - i
		possiblePalendromeCount += count
	}
	return possiblePalendromeCount >= k
}
