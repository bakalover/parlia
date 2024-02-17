package replica

import (
	"bakalover/parlia/paxos"
	"sync/atomic"
	"time"
)

type NodeId uint64

var globalNodeId uint64 = 0

type ReplicaChans struct {
	requestChan  *chan paxos.NetRequest
	acceptorChan *chan paxos.NetRequest
	proposerChan *chan paxos.NetRequest
}

func GenerateNodeId() NodeId {
	return NodeId(atomic.AddUint64(&globalNodeId, 1))
}

func AssociatedChans(id NodeId) ReplicaChans {
	net := paxos.GetNetwork()
	return ReplicaChans{&net.RequestChans[id], &net.AcceptorChans[id], &net.ProposerChans[id]}
}

type Replica interface {
	// Die with some smaaaaaaall probability
	// Death = context drop + goroutine sleep
	// Bool indicates whether goroutine "has died"
	MaybeDie() bool

	// Perform some activity during some time
	Step(stepTime time.Duration)
}

type SimpleReplica struct {
	Id NodeId
}

// --------------------------Replica---------------------------
func (r SimpleReplica) MaybeDie() bool {
	return false
}

func (r SimpleReplica) Step(stepTime time.Duration) {
	chans := AssociatedChans(r.Id)

	// TODO Syncronize all pools and inject timer fault sig

	// Client Pool
	go func() {
		for req := range *chans.requestChan {
			go func() {
				// Process req
			}()
		}
	}()

	// Acceptor Pool
	go func() {
		for req := range *chans.acceptorChan {
			go func() {
				// Process req
			}()
		}
	}()

	// Proposer Pool
	go func() {
		for req := range *chans.proposerChan {
			go func() {
				// Process req
			}()
		}
	}()
}

// --------------------------Replica---------------------------

// --------------------------Acceptor---------------------------

// --------------------------Acceptor---------------------------

// --------------------------Proposer---------------------------

// --------------------------Proposer---------------------------
