package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bakalover/tate"
)

const (
	replicaConfigPath = "./config/replica_ports.txt"
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
		c.logger.Fatalln("Trying take occupied address:%v", addr)
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
	return "", errors.New("No available addr")
}

func (c *Cluster) InitPortRegistry(addrs []string) {
	for _, addr := range addrs {
		c.addrs[addr] = false
	}
}

func (c *Cluster) GetLogger() *log.Logger {
	return c.logger
}

func (c *Cluster) Run(addrs []string) {
	c.InitPortRegistry(addrs)

	var replicaRoutines tate.Nursery
	defer replicaRoutines.Join()

	for _, _ = range c.addrs {
		replicaRoutines.Add(func() {
			RunReplica(c, SimpleMode)
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
	var replicaAddrs []string
	for scanner.Scan() {
		replicaPort := scanner.Text()
		if len(replicaPort) <= 4 || len(replicaPort) >= 6 {
			logger.Fatalf("Parsed invalid replica port: %v", replicaPort)
		}
		addr := fmt.Sprintf("localhost:%v", replicaPort)
		replicaAddrs = append(replicaAddrs, addr)
	}

	c := &Cluster{logger: logger}
	c.Run(replicaAddrs)
}
