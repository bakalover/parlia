package cluster

import (
	"errors"
	"log"
	"sync"
)

type AddrRegistry struct {
	mutex  sync.Mutex
	Logger *log.Logger
	addrs  map[string]bool
}

func (c *AddrRegistry) InvalidateAddr(addr string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.addrs[addr] {
		c.Logger.Fatalf("Double free address:%v", addr)
	}
	c.addrs[addr] = true
}

func (c *AddrRegistry) PickAddr() (string, error) {
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

func (c *AddrRegistry) InitAddrRegistry(configs []ReplicaConfig) {
	c.addrs = make(map[string]bool)
	for _, config := range configs {
		c.addrs[config.Addr] = true
	}
}
