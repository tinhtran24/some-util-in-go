package common

import (
	"math/rand"
	"time"
)

func Randomize(additionalToNowTime ...int64) {
	seed := time.Now().UTC().UnixNano()
	if len(additionalToNowTime) == 1 {
		seed += additionalToNowTime[0]
	}
	rand.Seed(seed)
}

var (
	urlSafeLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandString(length int) string {
	line := make([]byte, length)
	for i := range line {
		line[i] = urlSafeLetters[rand.Intn(len(urlSafeLetters))]
	}
	return string(line)
}
