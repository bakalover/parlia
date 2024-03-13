package main

import (
	"log"
	"net/rpc"
	"time"

	"github.com/bakalover/tate"
)

type Client struct {
	addrGen   *paxos.Generator
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
	cl, err := rpc.DialHTTP("tcp", client.addrGen.GenerateAddr())

	if err != nil {
		log.Fatal(err)
	}

	client.rpcClient = cl

	client.backoff = DefaultBackoff

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
