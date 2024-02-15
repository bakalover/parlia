package node

import (
	"time"
)

type Proposer struct {
	
}

func (p *Proposer) Propose() {

}

func (p *Proposer) Accept() {

}

func (p *Proposer) MaybeDie() bool {
	// if rand.Float64() < deathChance {
	// 	p.Ctx.ClearLog()
	// 	time.Sleep(kDeathTime)
	// 	return true
	// }
	return false
}

func (p *Proposer) Step(stepTime time.Duration) {
	for {
		// TODO
	}
}
