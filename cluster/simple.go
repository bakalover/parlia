package main

import (
	"fmt"
)

type SimpleRunner struct {
	RunnerBase
}

func (runner SimpleRunner) Run() {
	runner.Slave.Step(simTime + safeMargin)
	fmt.Printf("Runner Id: %d, Steps: %d, Reborn count: %d, Mode: Simple\n", runner.Id, 1, 0)
}
