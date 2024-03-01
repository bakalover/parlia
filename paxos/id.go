package paxos

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	minPort = 10_000
	maxPort = 65_535
)

type Generator struct {
	mutex      sync.Mutex
	idRegistry map[string]bool
}

func (g *Generator) GeneratePort() string {
	var port string
	for {
		port = getRandPort()
		if g.TryRegisterPort(port) {
			break
		}
	}
	return port
}

func getRandPort() string {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(minPort + gen.Intn(maxPort-minPort+1))
}

func (g *Generator) TryRegisterPort(port string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.idRegistry[port] {
		return false
	}

	g.idRegistry[port] = true
	return true
}

func (g *Generator) InvalidatePort(port string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.idRegistry[port] = false
}

func (g *Generator) GenerateAddr() string {
	return fmt.Sprintf(":%s", g.GeneratePort())
}

func (g *Generator) InvalidateAddr(addr string) {
	g.InvalidatePort(addr[1:])
}
