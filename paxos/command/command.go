package command

import (
	"math/rand"
	"time"
)

const CommandCount = 3

const (
	SetComm = 0
	AddComm = 1
	MulComm = 2
)

type Applicable interface {
	Apply(val LogType) LogType
}

func RandCommand(val LogType) Applicable {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	commN := gen.Intn(3)
	switch commN {
	case SetComm:
		return Set{val}
	case MulComm:
		return Mul{val}
	default:
		return Add{val}
	}
}
