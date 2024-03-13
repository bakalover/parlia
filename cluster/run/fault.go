package run

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/tate"
)

var kSteps = 7
var stepSeed = int(paxos.SimulationTime) / (int(time.Second) * kSteps)

func PickTime() time.Duration {
	return time.Duration(rand.Intn(stepSeed)) * time.Second
}

type FaultyRunner struct {
	RunnerBase
}

func (runner FaultyRunner) Run() {
	timer := time.NewTimer(runner.Config.SimulationTime)
	var kIter uint64 = 0

	tate.NewRepeater().Repeat(func() {
		runner.Slave.Step(PickTime())
		select {
		case <-timer.C:
			return
		default:
			kIter++
		}
	}).Join()
	
	timer.Stop()
	fmt.Printf("Runner Id: %d, Restarts: %d, Mode: Fault\n", runner.Id, kIter)
}
