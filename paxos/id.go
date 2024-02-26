package paxos

import (
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
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := strconv.Itoa(minPort + gen.Intn(maxPort-minPort+1))
	g.RegisterPort(port)
	g.idRegistry[port] = true
	return port
}

func (g *Generator) RegisterPort(port string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.idRegistry[port] = true
}

func (g *Generator) InvalidatePort(port string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.idRegistry[port] = false
}
