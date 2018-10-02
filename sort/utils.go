package sort

import (
	"math/rand"
	"time"
)

func GenerateInput() []int {
	res := make([]int, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		res[i] = rand.Intn(100)
	}

	return res
}
