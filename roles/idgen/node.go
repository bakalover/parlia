package idgen

import (
	"bakalover/parlia/roles"
)

var globalNodeSeed uint8 = 0

func Increase() {
	globalBallotSeed += 1
}

func GenerateNodeId() roles.NodeId {
	defer Increase() //++
	return roles.NodeId(globalNodeSeed)
}
