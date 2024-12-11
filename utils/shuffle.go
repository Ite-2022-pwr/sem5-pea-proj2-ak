package utils

import "math/rand"

func Shuffle(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		randIdx := 1 + rand.Intn(n-1)
		array[i], array[randIdx] = array[randIdx], array[i]
	}
}
