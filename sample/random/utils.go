package random

import "math/rand"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomBoolean() bool {
	return RandomInt(0, 1) == 1
}

func RandomStringFromArray(array []string) string {
	return array[RandomInt(0, len(array)-1)]
}
