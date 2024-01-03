package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"task-management-with-gin/helpers"
)

// generateRandomString generates a random string of a specified length
func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result += string(charset[randomIndex.Int64()])
	}
	return result, nil
}

// generateRandomUsername generates a random username with a specified prefix
func GenerateRandomUsername(prefix string) string {
	randomString, err := generateRandomString(8) // Change the length as needed
	if err != nil {
		helpers.ErrorPanic(err)
		return ""
	}
	return fmt.Sprintf("%s_%s", prefix, randomString)
}
