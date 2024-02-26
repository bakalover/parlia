package paxos

import (
	"math/rand"
	"time"
)

func InjectDelay() {
	time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
}
