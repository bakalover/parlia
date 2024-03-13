package replica

import "github.com/bakalover/parlia/paxos"

type Acceptor interface {
	Promise(b paxos.Ballot)
	Accepted(n paxos.BallotNumber)
}
