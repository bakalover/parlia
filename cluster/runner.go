package main

import (
	"sync/atomic"
	"time"
)

type RunMode uint8
type RunnerId uint32

var globalRunnerId uint32 = 0

const (
	safeMargin = 2 * time.Second
	simTime    = 50 * time.Second
)

func IdGenRunner() RunnerId {
	return RunnerId(atomic.AddUint32(&globalRunnerId, 1))
}

const (
	SimpleMode = 0
	FaultMode  = 1
)

type Runner interface {
	Run()
}

type RunnerBase struct {
	Id    RunnerId
	Slave Replica
}

func RunReplica( c *Cluster, mode RunMode) {
	var runner Runner

	base := RunnerBase{
		IdGenRunner(),
		Replica{cluster: c, logger: c.GetLogger()},
	}

	if mode == FaultMode {
		runner = FaultyRunner{base}
	} else {
		runner = SimpleRunner{base}
	}

	runner.Run()
}
