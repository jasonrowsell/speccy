package random

import (
	"math/rand"

	"github.com/google/uuid"
)

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
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

func GenerateUUID() string {
	return uuid.New().String()
}
