package replica

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/tate"
)

type NodeId uint64

var globalNodeId uint64 = 0

type ReplicaChans struct {
	requestChan  *chan []byte
	acceptorChan *chan paxos.NetRequest
	proposerChan *chan paxos.NetRequest
}

func GenerateNodeId() NodeId {
	return NodeId(atomic.AddUint64(&globalNodeId, 1))
}

func AssociatedChans(id NodeId) ReplicaChans {
	net := paxos.GetWWW()
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
	Id  NodeId
	Log []string
}

// --------------------------Replica----------------------------
func (r SimpleReplica) MaybeDie() bool {
	return false
}

func (r SimpleReplica) Step(stepTime time.Duration) {
	chans := AssociatedChans(r.Id)
	ctx, cancel := context.WithCancel(context.Background())
	net := paxos.GetWWW()

	var (
		replicaPool  tate.Nursery
		clientPool   tate.Nursery
		acceptorPool tate.Nursery
		proposerPool tate.Nursery
	)

	replicaPool.Add(func() {
		for req := range *chans.requestChan {
			select {
			case <-ctx.Done():
				break
			default:
				clientPool.Add(func() {
					// Self Broadcast
					net.Broadcast(req, paxos.Proposers, paxos.Init)
				})
			}
		}
		clientPool.Join()
	}).Add(func() {
		for req := range *chans.proposerChan {
			select {
			case <-ctx.Done():
				break
			default:
				proposerPool.Add(func() {
					// Receive Self broadcast message
					// Proposer Logic
				})
			}
		}
		proposerPool.Join()
	}).Add(func() {
		for req := range *chans.acceptorChan {
			select {
			case <-ctx.Done():
				break
			default:
				acceptorPool.Add(func() {
					// Acceptor Logic
				})
			}
		}
		acceptorPool.Join()
	})

	time.AfterFunc(stepTime, func() {
		cancel()
	})

	replicaPool.Join()
}

// --------------------------Replica----------------------------

// --------------------------Acceptor---------------------------

func (r *SimpleReplica) Promise(b paxos.Ballot) {

}

func (r *SimpleReplica) Accepted(n paxos.BallotNumber) {

}

// --------------------------Acceptor---------------------------

// --------------------------Proposer---------------------------

func (r *SimpleReplica) Prepare(b paxos.Ballot) {

}

func (r *SimpleReplica) Accept(n paxos.BallotNumber) {

}

// --------------------------Proposer---------------------------
