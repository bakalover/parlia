package paxos

import (
	"bakalover/parlia/paxos/node"
	"fmt"
	"math/rand"
	"time"
)

var stepSeed = 10

type FaultyRunner struct {
	Config *InitConfig
	Id     int
	Slave  node.NodeBase
}

func PickTime() time.Duration {
	return time.Duration(rand.Intn(stepSeed)) * time.Second
}

func (runner *FaultyRunner) Run() {
	timer := time.NewTimer(runner.Config.SimulationTime)
	var kIter, kDeath uint64 = 1, 0
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
	fmt.Printf("Runner Id: %d, Iterations: %d, Node death count: %d", runner.Id, kIter, kDeath)
}

func Acceptor(registry *Registry, config *InitConfig) {
	// runner := FaultyRunner{node.Acceptor{}}
	// runner.Run()
}

func Proposer(registry *Registry, config *InitConfig) {
	// TODO
}

func Learner(registry *Registry, config *InitConfig) {
	// TODO
}
