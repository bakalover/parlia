package backoff

import (
	"log"
	"time"
)

type Backoff struct {
	BaseDelay    time.Duration
	CurrentDelay time.Duration
	Multiplier   float64
	Jitter       float64
	MaxDelay     time.Duration
}

var DefaultBackoff = Backoff{
	BaseDelay:    1.0 * time.Second,
	CurrentDelay: 1.0 * time.Second,
	Multiplier:   1.6,
	Jitter:       0.3,
	MaxDelay:     10 * time.Second,
}

func (b *Backoff) Update() {
	b.CurrentDelay = time.Duration(float64(b.CurrentDelay)*b.Multiplier) + time.Duration(float64(b.BaseDelay)*b.Jitter)
}

func (b *Backoff) Next() time.Duration {
	log.Println(b.CurrentDelay)
	defer b.Update()
	return b.CurrentDelay
}

func (b *Backoff) Reset() {
	b.CurrentDelay = b.BaseDelay
}
