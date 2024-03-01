package client

import (
	"math/rand"
	"time"
)

type Backoff struct {
	Init       time.Duration
	curr       time.Duration
	Mult       int
	RandWindow time.Duration
}

func (b *Backoff) Next() time.Duration {
	b.curr = b.curr*time.Duration(b.Mult) + time.Duration(rand.Intn(int(b.RandWindow)))
	return b.curr
}

func (b *Backoff) Reset() {
	b.curr = b.Init
}
