package sample

import (
	"math/rand"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomBoolean() bool {
	return randomInt(0, 1) == 1
}

func randomStringFromArray(array []string) string {
	return array[randomInt(0, len(array)-1)]
}
