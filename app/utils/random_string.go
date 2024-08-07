package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func randomAlphabets(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func GenerateFormattedString() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random alphabets
	sk := randomAlphabets(2)
	lske := randomAlphabets(4)

	// Get the current date and time
	now := time.Now()
	date := now.Format("0602150405")

	// Format the final string
	formattedString := fmt.Sprintf("%s-%s%s", sk, lske, date)

	return formattedString
}
