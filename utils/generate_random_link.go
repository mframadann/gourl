package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomShortenedLink() string {
	res := make([]rune, 6)
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i := 0; i < len(res); i++ {
		res[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(res)
}
