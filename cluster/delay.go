package paxos

import (
	"math/rand"
	"time"
)

func MajorDelay() {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
}
