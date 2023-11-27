package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, random.Intn(100-10)+100)
	for i := range array {
		array[i] = random.Intn(100)
		fmt.Println(array[i])
	}

}

func Sum(a, b int64) int64 {
	return a - b
}
