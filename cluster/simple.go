package main

type SimpleRunner struct {
	RunnerBase
}

func (runner SimpleRunner) Run() {
	runner.Slave.Step(simTime + safeMargin)
	runner.Logger.Printf("Runner Id: %d, Steps: %d, Reborn count: %d, Mode: Simple\n", runner.Id, 1, 0)
}
