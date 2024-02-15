package paxos

import (
	"bakalover/parlia/paxos/node"
	"math/rand"
	"time"
)

var stepSeed = 10

type FaultyRunner struct {
	Slave node.NodeBase
}

func PickTime() time.Duration {
	return time.Duration(rand.Intn(stepSeed)) * time.Second
}

func (runner *FaultyRunner) Run() {
	runner.Slave.Step(PickTime())
	runner.Slave.MaybeDie()
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
