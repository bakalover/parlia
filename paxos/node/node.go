package node

import (
	"sync/atomic"
	"time"
)

type NodeId uint64

var globalNodeSeed uint64 = 0

func Increase() {
	globalNodeSeed += 1
}

func GenerateNodeId() NodeId {
	return NodeId(atomic.AddUint64(&globalNodeSeed, 1))
}

type NodeBase interface {
	// Die with some smaaaaaaall probability
	// Death = context drop + goroutine sleep
	// Bool indicates whether goroutine "has died"
	MaybeDie() bool

	// Perform some activity during some time
	Step(time.Duration)
}
