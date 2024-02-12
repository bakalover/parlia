package roles

import (
	"math/rand"
	"time"
)

type ProposerCtx struct {
	Log []int
}

func (ctx *ProposerCtx) Clear() {
	ctx.Log = ctx.Log[:0]
}

type Proposer struct {
	Ctx ProposerCtx
	// All communications represented as a buffered channel
	// Communication pattern: MPSC
	NetWires chan int
}

func (p *Proposer) MaybeDie() bool {
	if rand.Float64() < deathChance {
		p.Ctx.Clear()
		time.Sleep(kDeathTime)

		return true
	}
	return false
}

func RunProposer() {
	// MaybeDie()
	// Prepare()
	// MaybeDie()
	// GatherQuorum()
	// MaybeDie()
	// Accept()

}
