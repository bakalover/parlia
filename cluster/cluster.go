package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/bakalover/tate"
)

const (
	replicaConfigPath = "./config/replica_config.txt"
)

type Cluster struct {
	mutex  sync.Mutex
	logger *log.Logger
	addrs  map[string]bool
}

func (c *Cluster) InvalidateAddr(addr string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.addrs[addr] {
		c.logger.Fatalf("Trying take occupied address:%v", addr)
	}
	c.addrs[addr] = true
}

func (c *Cluster) PickAddr() (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for addr, available := range c.addrs {
		if available {
			c.addrs[addr] = false
			return addr, nil
		}
	}
	return "", errors.New("no available addr")
}

func (c *Cluster) InitAddrRegistry(configs []ReplicaConfig) {
	for _, config := range configs {
		c.addrs[config.Addr] = false
	}
}

func (c *Cluster) GetLogger() *log.Logger {
	return c.logger
}

func (c *Cluster) Run(configs []ReplicaConfig) {
	c.InitAddrRegistry(configs)

	var replicaRoutines tate.Nursery
	defer replicaRoutines.Join()

	for _, config := range configs {
		replicaRoutines.Add(func() {
			RunReplica(c, config.Mode)
		})
	}
}

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

	// =============================Files================================
	replicaConfig, err := os.Open(replicaConfigPath)
	if err != nil {
		logger.Fatalf("Error opening file: %v ", err)
	}
	defer replicaConfig.Close()
	//=============================Files================================

	scanner := bufio.NewScanner(replicaConfig)
	var replicaConfigs []ReplicaConfig

	for scanner.Scan() {
		config := strings.Split(scanner.Text(), " ")
		replicaPort := config[0]
		replicaMode := config[1]

		if len(replicaPort) <= 4 || len(replicaPort) >= 6 {
			logger.Fatalf("Parsed invalid replica port: %v", replicaPort)
		}

		if replicaMode != "Trust" && replicaMode != "Faulty" {
			logger.Fatalf("Invalide replica mode: %v\n Avaliable:\n- Trust\n- Faulty", replicaMode)
		}
		addr := fmt.Sprintf("localhost:%v", replicaPort)
		replicaConfigs = append(replicaConfigs, ReplicaConfig{addr, RunMode(replicaMode)})
	}

	c := &Cluster{logger: logger}
	c.Run(replicaConfigs)
}
