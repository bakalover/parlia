package run

import (
	"sync/atomic"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/parlia/paxos/replica"
)

type RunMode uint8
type RunnerId uint32

var globalRunnerId uint32 = 0

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
	Network *paxos.Network
	Config  *paxos.InitConfig
	Id      RunnerId
	Slave   replica.Replica
}

func Replica(mode RunMode) {
	var runner Runner
	base := RunnerBase{paxos.GetWWW(), paxos.GetConfig(), IdGenRunner(), replica.SimpleReplica{}}
	if mode == FaultMode {
		runner = FaultyRunner{base}
	} else {
		runner = SimpleRunner{base}
	}
	runner.Run()
}
