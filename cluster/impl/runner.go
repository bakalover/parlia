package cluster

import (
	"log"
	"sync/atomic"
	"time"
)

type RunMode string
type RunnerId uint32

var globalRunnerId uint32 = 0

const (
	safeMargin = 1 * time.Second
	simTime    = 5 * time.Second
)

type ReplicaConfig struct {
	Addr string
	Mode RunMode
}

func IdGenRunner() RunnerId {
	return RunnerId(atomic.AddUint32(&globalRunnerId, 1))
}

const (
	TrustMode = RunMode("Trust")
	FaultMode = RunMode("Faulty")
)

type Runner interface {
	Run()
}

type RunnerBase struct {
	Id     RunnerId
	Logger *log.Logger
	Slave  *Replica
}

func RunReplica(c *Cluster, mode RunMode) {
	var runner Runner

	base := RunnerBase{
		IdGenRunner(),
		c.GetLogger(),
		&Replica{Cluster: c, Logger: c.GetLogger()},
	}
	if mode == FaultMode {
		runner = FaultyRunner{base}
	} else {
		runner = SimpleRunner{base}
	}

	runner.Run()
}
