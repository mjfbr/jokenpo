package utils

import (
	"math/rand"
	"time"
)

func RandomChoice(choices []string) string {
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Intn(len(choices))]
}
