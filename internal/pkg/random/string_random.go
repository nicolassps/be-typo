package random

import "math/rand"

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomLetter() string {
	index := rand.Intn(len(letterBytes))

	return string(letterBytes[index])
}
