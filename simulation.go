package main

import (
	"bakalover/parlia/paxos"
	"sync"
)

func main() {
	initConfig := paxos.GenerateInitConfig()
	RunSimulation(initConfig)
}

func RunSimulation(initConfig paxos.InitConfig) {

	registry := paxos.MakeRegistry(&initConfig)

	var wg sync.WaitGroup
	wg.Add(initConfig.KAcceptors + initConfig.KLearners + initConfig.KProposers)

	for i := 0; i < initConfig.KAcceptors; i++ {
		go func() {
			// Acceptor(&registry, &initConfig)
			wg.Done()
		}()
	}

	for i := 0; i < initConfig.KLearners; i++ {
		go func() {
			// Learner(&registry, &initConfig)
			wg.Done()
		}()
	}

	for i := 0; i < initConfig.KProposers; i++ {
		go func() {
			// Proposer(&registry, &initConfig)
			wg.Done()
		}()
	}

	wg.Wait()
}
