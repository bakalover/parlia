package paxos

import (
	"sync/atomic"
)

var globalBallotSeed uint64 = 0

type BallotNumber struct {
	KLocal uint64
	Owner  string
}

type Ballot struct {
	Number BallotNumber
	Value  uint64 // Change to some "Command" class aka Command-pattern
}

// Contention point across all proposers
func GenerateBallotNumber(id string) BallotNumber {
	return BallotNumber{KLocal: atomic.AddUint64(&globalBallotSeed, 1), Owner: id}
}
