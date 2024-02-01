package idgen

import (
	"bakalover/parlia/roles"
	"sync/atomic"
)

var globalBallotSeed uint64 = 0

// Contention point across all proposers
func GenerateBallotNumber(id roles.NodeId) roles.BallotNumber {
	return roles.BallotNumber{KLocal: atomic.AddUint64(&globalBallotSeed, 1), Owner: id}
}
