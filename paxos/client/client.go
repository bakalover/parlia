package client

import (
	"log"
	"math/rand"
	"net/rpc"
	"time"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/tate"
)

var commands = map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F"}

func PickCommand() string {
	return commands[rand.Intn(len(commands))]
}

func ApplyCommand(cl *rpc.Client) {
	err := cl.Call("ProxyService.Apply", PickCommand(), nil)
	if err != nil {
		log.Println(err)
	}
}

func Client() {

	cl, err := rpc.DialHTTP("tcp", ":todoport")

	if err != nil {
		log.Fatal(err)
	}

	rp := tate.NewRepeater()

	// DDOS =)
	rp.Repeat(func() {
		ApplyCommand(cl)
	}).Repeat(func() {
		ApplyCommand(cl)
	}).Repeat(func() {
		ApplyCommand(cl)
	})

	time.AfterFunc(paxos.SimulationTime, func() { rp.Join() })
}