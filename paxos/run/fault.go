package run

import (
	"bakalover/parlia/paxos"
	"fmt"
	"math/rand"
	"time"
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
	var kIter, kDeath uint64 = 0, 0
	go func() {
		for {
			runner.Slave.Step(PickTime())
			if haveDied := runner.Slave.MaybeDie(); haveDied {
				kDeath++
			}
			select {
			case <-timer.C:
				return
			default:
				kIter++
			}
		}
	}()
	timer.Stop()
	fmt.Printf("Runner Id: %d, Steps: %d, Reborn count: %d Mode: Fault\n", runner.Id, kIter, kDeath)
}
