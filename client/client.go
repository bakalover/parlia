package client

import (
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/tate"
)

var commands = map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F"}

func PickCommand() string {
	return commands[rand.Intn(len(commands))]
}

func SendCommand(net *paxos.Network) {
	net.Broadcast([]byte(PickCommand()), paxos.Proposers, paxos.Init)
}

// Await either signal from terminal or simulation
func AwaitSim() {
	cfg := paxos.GetConfig()
	ch := make(chan os.Signal)
	time.AfterFunc(cfg.SimulationTime, func() { ch <- os.Interrupt })
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func Client() {
	rp := tate.NewRepeater()
	net := paxos.GetWWW()

	// DDOS =)
	rp.Repeat(func() {
		SendCommand(net)
	}).Repeat(func() {
		SendCommand(net)
	}).Repeat(func() {
		SendCommand(net)
	})
	
	AwaitSim()
	rp.Join()
}
