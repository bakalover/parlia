package run

import (
	"bakalover/parlia/paxos"
	"bakalover/parlia/paxos/node"
	"sync/atomic"
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
	Config   *paxos.InitConfig
	Id       RunnerId
	Slave    node.NodeBase
}

func Node(nodeType node.NodeType, mode RunMode) {
	runner := GetRunner(paxos.GetNetwork(), paxos.GetConfig(), nodeType, mode)
	runner.Run()
}

func GetNode(nodeType node.NodeType) node.NodeBase {
	switch nodeType {
	case node.TypeAcceptor:
		return node.Acceptor{}
	case node.TypeLearner:
		return node.Learner{}
	default:
		return node.Proposer{}
	}
}

func GetRunner(registry *paxos.Network, config *paxos.InitConfig, nodeType node.NodeType, mode RunMode) Runner {
	base := RunnerBase{registry, config, IdGenRunner(), GetNode(nodeType)}
	if mode == FaultMode {
		return FaultyRunner{base}
	} else {
		return SimpleRunner{base}
	}
}
