package node

import (
	"sync/atomic"
	"time"
)

type NodeId uint64
type NodeType uint8

const (
	TypeAcceptor NodeType = 0
	TypeProposer NodeType = 1
	TypeLearner  NodeType = 2
)

var globalNodeId uint64 = 0

func GenerateNodeId() NodeId {
	return NodeId(atomic.AddUint64(&globalNodeId, 1))
}

type NodeBase interface {
	// Die with some smaaaaaaall probability
	// Death = context drop + goroutine sleep
	// Bool indicates whether goroutine "has died"
	MaybeDie() bool

	// Perform some activity during some time
	Step(stepTime time.Duration)
}
