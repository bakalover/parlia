package main

import (
	"bakalover/parlia/client"
	"bakalover/parlia/paxos"
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
	wg.Add(config.Kreplicas + config.Kclients)

	for i := 0; i < config.Kclients; i++ {
		go func() {
			client.Client()
			wg.Done()
		}()
	}

	// TODO: split number of faulty replicas
	for i := 0; i < config.Kreplicas; i++ {
		go func() {
			run.Replica(run.SimpleMode)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("-----------------------------Done-----------------------------")
}
