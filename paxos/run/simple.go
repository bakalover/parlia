package run

import (
	"fmt"
	"time"
)

const safeMargin = 2 * time.Second

type SimpleRunner struct {
	RunnerBase
}

func (runner SimpleRunner) Run() {
	runner.Slave.Step(runner.Config.SimulationTime + safeMargin)
	fmt.Printf("Runner Id: %d, Steps: %d, Reborn count: %d, Mode: Simple\n", runner.Id, 1, 0)
}
