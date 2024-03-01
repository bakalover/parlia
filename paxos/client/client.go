package client

import (
	"log"
	"net/rpc"
	"time"

	"github.com/bakalover/parlia/paxos"
	"github.com/bakalover/parlia/paxos/command"
	"github.com/bakalover/tate"
)

type Client struct {
	port      *paxos.Generator
	rpcClient *rpc.Client
	backoff   Backoff
}

func (client *Client) SendCommand() {
	paxos.InjectDelay()
	err := client.rpcClient.Call("ProxyService.Apply", command.RandCommand(2), nil)
	if err != nil {
		log.Println(err)
		time.Sleep(client.backoff.Next())
	} else {
		client.backoff.Reset()
	}
}

func (client *Client) Run() {
	cl, err := rpc.DialHTTP("tcp", "todo: search port")

	if err != nil {
		log.Fatal(err)
	}

	client.rpcClient = cl

	client.backoff = Backoff{Init: 500 * time.Millisecond, Mult: 2, RandWindow: 30 * time.Millisecond}

	rp := tate.NewRepeater()

	// DDOS =)
	rp.Repeat(func() {
		client.SendCommand()
	}).Repeat(func() {
		client.SendCommand()
	}).Repeat(func() {
		client.SendCommand()
	})

	time.AfterFunc(paxos.SimulationTime, func() { rp.Join() })
}
