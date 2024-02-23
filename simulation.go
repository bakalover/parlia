package main

import (
	"github.com/bakalover/parlia/client"
	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/parlia/paxos/run"
	"fmt"

	"github.com/bakalover/tate"
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

	var sim tate.Nursery

	for i := 0; i < config.Kclients; i++ {
		sim.Add(func() {
			client.Client()
		})
	}

	// TODO: split number of faulty replicas
	for i := 0; i < config.Kreplicas; i++ {
		sim.Add(func() {
			run.Replica(run.SimpleMode)
		})
	}

	sim.Join()
	fmt.Println("-----------------------------Done-----------------------------")
}
