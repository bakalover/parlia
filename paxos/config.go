package paxos

import (
	"math/rand"
	"time"
)

var kReplicaSeed, kClientSeed = 61, 11
var deathDuration = 300 * time.Millisecond
var deathChance float64 = 0.001
var SimulationTime = 1 * time.Minute

type InitConfig struct {
	Kreplicas      int
	Kclients       int
	SimulationTime time.Duration
	DeathDuration  time.Duration
	DeathChance    float64
}

var GlobalInitConfig InitConfig

func GenerateInitConfig() {

	Kreplicas := rand.Intn(kReplicaSeed)*2 + 3
	Kclients := rand.Intn(kClientSeed)*2 + 3

	GlobalInitConfig = InitConfig{
		Kclients,
		Kreplicas,
		SimulationTime,
		deathDuration,
		deathChance,
	}
}

func GetConfig() *InitConfig {
	return &GlobalInitConfig
}
