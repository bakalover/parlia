package replica

import "github.com/bakalover/parlia/paxos"

type Proposer interface {
	Prepare(n paxos.BallotNumber) // + Slot N
	Accept(b paxos.Ballot)
}
