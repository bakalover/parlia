package cluster

import (
	"math/rand"
	"time"
)

// const rebornDelay = time.Duration(float64(stepSeed) * 1.5)

func MaybeMajorDelay() {
	if rand.Intn(11) == 2 {
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}

// func RebornDelay() {
// 	time.Sleep(rebornDelay)
// }
