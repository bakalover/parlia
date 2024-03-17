package client

import (
	"context"
	"log"
	"math/rand"
	"time"

	proto "github.com/bakalover/parlia/proto"
	"github.com/bakalover/tate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	AvailableProxies []string
	Logger           *log.Logger
	Id               string
	SimTime          time.Duration
	Backoff          Backoff
	requestIndex     uint64
	rpcClient        proto.ProxyClient
	targetAddr       string
}

var Commands = []string{"Inc", "Dec", "Zero"}

func (cl *Client) PickCommand() string {
	return Commands[rand.Intn(len(Commands))]
}

func (cl *Client) AvailableProxy() string {
	return cl.AvailableProxies[rand.Intn(len(cl.AvailableProxies))]
}

func (cl *Client) NextRequestId() *proto.RequestId {
	defer func() { cl.requestIndex++ }()
	return &proto.RequestId{OwnerId: cl.Id, Index: cl.requestIndex}
}

func (cl *Client) SendCommand() {
	requestId := cl.NextRequestId()

	_, err := cl.rpcClient.Apply(
		context.Background(),
		&proto.Command{Type: cl.PickCommand(), Id: requestId},
	)

	if err != nil {
		cl.Logger.Printf("Cluster is unavailbale:\nClient:%v -> Proxy:%v", cl.Id, cl.targetAddr)
		time.Sleep(cl.Backoff.Next())
	} else {
		cl.Backoff.Reset()
	}
}

func (cl *Client) ConnToCluster() {
	proxyAddr := cl.AvailableProxy()
	conn, err := grpc.Dial(proxyAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		cl.Logger.Fatalf("Fail to dial: %v", err)
	} else {
		cl.targetAddr = proxyAddr // Proxy rotation
		cl.rpcClient = proto.NewProxyClient(conn)
	}
}

func (cl *Client) Run() {
	//Proxies are not faulty
	cl.ConnToCluster()

	// DDOS =)
	client := tate.NewRepeater()
	client.Repeat(func() {
		cl.SendCommand()
	})

	await := make(chan bool)
	time.AfterFunc(cl.SimTime, func() {
		client.Join()
		await <- true
	})
	<-await
}
