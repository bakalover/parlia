package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bakalover/tate"
)

var kSteps = 7
var stepSeed = int(simTime) / (int(time.Second) * kSteps)

func PickStepTime() time.Duration {
	return time.Duration(rand.Intn(stepSeed)) * time.Second
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

	fmt.Printf("Runner Id: %d, Restarts: %d, Mode: Fault\n", runner.Id, kIter)
}
