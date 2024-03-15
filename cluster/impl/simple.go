package cluster

type SimpleRunner struct {
	RunnerBase
}

func (runner SimpleRunner) Run() {
	runner.Slave.Step(simTime + safeMargin)
	runner.Logger.Printf("Runner Id: %d, Steps: %d, Mode: Trust\n", runner.Id, 1)
}
