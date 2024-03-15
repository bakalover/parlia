package cluster

import (
	"math/rand"
	"time"

	"github.com/bakalover/tate"
)

const (
	kSteps   = 11277
	stepSeed = simTime / time.Duration(kSteps)
)

func PickStepTime() time.Duration {
	return time.Duration(rand.Intn(int(stepSeed)))
}

type FaultyRunner struct {
	RunnerBase
}

func (runner FaultyRunner) Run() {
	var kIter uint64 = 0
	rp := tate.NewRepeater()

	rp.Repeat(func() {
		runner.Slave.Step(PickStepTime())
		kIter++
	})

	await := make(chan bool)

	time.AfterFunc(simTime, func() {
		rp.Join()
		await <- true
	})

	<-await

	runner.Logger.Printf("Runner Id: %d, Steps: %d, Mode: Faulty\n", runner.Id, kIter)
}
