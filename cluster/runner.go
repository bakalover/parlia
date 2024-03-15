package main

import (
	"log"
	"sync/atomic"
	"time"
)

type RunMode string
type RunnerId uint32

var globalRunnerId uint32 = 0

const (
	safeMargin = 2 * time.Second
	simTime    = 50 * time.Second
)

type ReplicaConfig struct {
	Addr string
	Mode RunMode
}

func IdGenRunner() RunnerId {
	return RunnerId(atomic.AddUint32(&globalRunnerId, 1))
}

const (
	TrustMode = "Trust"
	FaultMode = "Faulty"
)

type Runner interface {
	Run()
}

type RunnerBase struct {
	Id     RunnerId
	Logger *log.Logger
	Slave  Replica
}

func RunReplica(c *Cluster, mode RunMode) {
	var runner Runner

	base := RunnerBase{
		IdGenRunner(),
		c.GetLogger(),
		Replica{cluster: c, logger: c.GetLogger()},
	}

	if mode == FaultMode {
		runner = FaultyRunner{base}
	} else {
		runner = SimpleRunner{base}
	}

	runner.Run()
}
