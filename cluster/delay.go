package main

import (
	"math/rand"
	"time"
)

func MaybeMajorDelay() {
	if rand.Intn(11) == 2 {
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}
