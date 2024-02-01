package roles

import "time"

type NodeId uint8

var kDeathTime time.Duration = 300 * time.Millisecond
var deathChance float64 = 0.001

type BallotNumber struct {
	KLocal uint64
	Owner  NodeId
}

type Ballot struct {
	Number BallotNumber
	Value  uint64 // Change to some Enum
}

type NetBase interface {
	SendRequest(to *NetBase) // Incapsulates delay
}

type FaultyBase interface {
	// Die with some smaaaaaaall probability
	// Death = context drop + goroutine sleep
	// Bool indicates whether goroutine "has died"
	MaybeDie() bool
}
