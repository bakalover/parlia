package main

import (
	"bakalover/parlia/paxos"
	"bakalover/parlia/paxos/node"
	"bakalover/parlia/paxos/run"
	"fmt"
	"sync"
)

func main() {
	Init()
	RunSimulation()
}

func Init() {
	paxos.GenerateInitConfig()
	paxos.MakeNetwork(paxos.GetConfig())
}

func RunSimulation() {

	fmt.Println("-----------------------------Running-----------------------------")

	config := paxos.GetConfig()

	var wg sync.WaitGroup
	wg.Add(config.KAcceptors + config.KLearners + config.KProposers)

	for i := 0; i < config.KAcceptors; i++ {
		go func() {
			run.Node(node.TypeAcceptor, run.SimpleMode)
			wg.Done()
		}()
	}

	for i := 0; i < config.KLearners; i++ {
		go func() {
			run.Node(node.TypeLearner, run.SimpleMode)
			wg.Done()
		}()
	}

	for i := 0; i < config.KProposers; i++ {
		go func() {
			run.Node(node.TypeProposer, run.SimpleMode)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("-----------------------------Done-----------------------------")
}
