package utils

import "math/rand"

func GenerateOTP() string {
	min, max := 10000, 99999
	return string(rand.Intn(max-min) + min)
}
