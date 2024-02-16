package run

import (
	"fmt"
)

type SimpleRunner struct {
	RunnerBase
}

func (runner SimpleRunner) Run() {
	runner.Slave.Step(runner.Config.SimulationTime)
	fmt.Printf("Runner Id: %d, Steps: %d, Reborn count: %d, Mode: Simple", runner.Id, 1, 0)
}
