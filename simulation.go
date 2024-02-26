package main

import (
	"fmt"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/parlia/paxos/client"
	"github.com/bakalover/parlia/paxos/proxy"
	"github.com/bakalover/parlia/paxos/run"

	"github.com/bakalover/tate"
)

func main() {
	Init()
	RunSimulation()
}

func Init() {
	paxos.GenerateInitConfig()
}

func RunSimulation() {

	fmt.Println("-----------------------------Running-----------------------------")

	config := paxos.GetConfig()
	portGen := &paxos.Generator{}
	sim := tate.NewNursery(nil)

	for i := 0; i < config.Kclients; i++ {
		sim.Add(func(c *tate.Linker) {
			client.Client(portGen)
		})
	}

	for i := 0; i < config.Kproxy; i++ {
		sim.Add(func(c *tate.Linker) {
			proxy.Proxy(portGen)
		})
	}

	// TODO: split number of faulty replicas
	for i := 0; i < config.Kreplicas; i++ {
		sim.Add(func(c *tate.Linker) {
			run.Replica(portGen, run.SimpleMode)
		})
	}

	sim.Join()
	fmt.Println("-----------------------------Done-----------------------------")
}
