package main

import (
	"bakalover/parlia/paxos"
	"bakalover/parlia/paxos/node"
	"bakalover/parlia/paxos/run"
	"fmt"
	"sync"
)

func main() {
	initConfig := paxos.GenerateInitConfig()
	RunSimulation(initConfig)
}

func RunSimulation(initConfig *paxos.InitConfig) {

	fmt.Println("-----------------------------Running-----------------------------")

	registry := paxos.MakeRegistry(initConfig)

	var wg sync.WaitGroup
	wg.Add(initConfig.KAcceptors + initConfig.KLearners + initConfig.KProposers)

	for i := 0; i < initConfig.KAcceptors; i++ {
		go func() {
			run.Node(registry, initConfig, node.NodeAcceptor, run.SimpleMode)
			wg.Done()
		}()
	}

	for i := 0; i < initConfig.KLearners; i++ {
		go func() {
			run.Node(registry, initConfig, node.NodeLearner, run.SimpleMode)
			wg.Done()
		}()
	}

	for i := 0; i < initConfig.KProposers; i++ {
		go func() {
			run.Node(registry, initConfig, node.NodeProposer, run.SimpleMode)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("-----------------------------Done-----------------------------")
}
