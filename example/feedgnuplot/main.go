package main

import (
	"fmt"
	"math/rand"
	"time"

	noise "github.com/palsivertsen/noise"
)

// Run:
// go run example/feedgnuplot/main.go | feedgnuplot --stream 0.05 --exit --xlen 5000 --lines --ymin 0 --ymax 1

func main() {
	n := noise.Smooth{
		Size: 6000,
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())), // Optional
	}

	for {
		fmt.Println(n.Next())
		time.Sleep(time.Millisecond * 10)
	}
}
