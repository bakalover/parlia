package paxos

import (
	"math/rand"
	"time"
)

var initPSeed, initASeed, initLSeed = 61, 17, 29
var deathDuration = 300 * time.Millisecond
var deathChance float64 = 0.001
var SimulationTime = 1 * time.Minute

type InitConfig struct {
	KProposers     int
	KAcceptors     int
	KLearners      int
	SimulationTime time.Duration
	DeathDuration  time.Duration
	DeathChance    float64
}

var GlobalInitConfig InitConfig

func GenerateInitConfig() {

	kProposers := rand.Intn(initPSeed)*2 + 3
	kAcceptors := rand.Intn(initASeed)*2 + 1 // Prefferable to be odd (Fault-tolerance points)
	kLearners := rand.Intn(initLSeed)*2 + 3

	GlobalInitConfig = InitConfig{
		kProposers,
		kAcceptors,
		kLearners,
		SimulationTime,
		deathDuration,
		deathChance,
	}
}

func GetConfig() *InitConfig {
	return &GlobalInitConfig
}
