package utils

import (
	"strconv"
	"strings"
)

// FormatRupiah formats an integer to a Rupiah currency string
func FormatRupiah(number int64) string {
	// Convert the number to a string
	str := strconv.Itoa(int(number))

	// Reverse the string to facilitate adding commas
	reversed := reverseString(str)

	// Add commas every three digits
	var result strings.Builder
	for i, r := range reversed {
		if i > 0 && i%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(r)
	}

	// Reverse again to restore original order
	formatted := reverseString(result.String())

	// Add the "Rp" prefix
	return "Rp " + formatted
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}