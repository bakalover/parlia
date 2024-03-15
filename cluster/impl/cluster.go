package cluster

import (
	"log"

	"github.com/bakalover/tate"
)

type Cluster struct {
	AddrRegistry
	Logger *log.Logger
}

func (c *Cluster) GetLogger() *log.Logger {
	return c.Logger
}

func (c *Cluster) Run(configs []ReplicaConfig) {
	c.InitAddrRegistry(configs)

	var replicaRoutines tate.Nursery
	defer replicaRoutines.Join()

	for _, config := range configs {
		mode := config.Mode
		replicaRoutines.Add(func() {
			RunReplica(c, mode)
		})
	}
}
