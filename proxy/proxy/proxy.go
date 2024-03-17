package proxy

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sync"

	proto "github.com/bakalover/parlia/proto"
	"github.com/bakalover/tate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Proxy struct {
	proto.UnimplementedProxyServer
	Logger            *log.Logger
	AvailableReplicas []string
	Addr              string
	rpcServer         *grpc.Server
	rpcClient         proto.ReplicaClient
	mutex             sync.Mutex
	targetAddr        string
}

// ===============================RPC Service===============================
func (p *Proxy) Apply(ctx context.Context, command *proto.Command) (*proto.Empty, error) {

	// Context???
	// TODO: batch window 10ms ~ rpc stream

	p.mutex.Lock()
	defer p.mutex.Unlock()

	resp, err := p.rpcClient.Apply(ctx, command)
	if err != nil {
		p.Logger.Println("Replica is unavailable")
		p.ConnToCluster() // Change target Replica
	}

	return resp, err // If err => client retries
}

//===============================RPC Service===============================

func (p *Proxy) Run() {
	// Proxy as client
	p.ConnToCluster()

	// Proxy as server
	l, err := net.Listen("tcp", p.Addr)
	if err != nil {
		p.Logger.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	p.rpcServer = grpcServer
	proto.RegisterProxyServer(grpcServer, p)

	var routine tate.Nursery
	routine.Add(func() {
		grpcServer.Serve(l)
	})
	routine.Join()
}

func (p *Proxy) ShutDown() {
	p.rpcServer.GracefulStop()
}

func (p *Proxy) AvailableReplica() string {
	return p.AvailableReplicas[rand.Intn(len(p.AvailableReplicas))]
}

func (p *Proxy) ConnToCluster() {
	replicaAddr := p.AvailableReplica()
	conn, err := grpc.Dial(replicaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		p.Logger.Printf("Fail to dial: %v", err)
	} else {
		p.targetAddr = replicaAddr
		p.rpcClient = proto.NewReplicaClient(conn)
		log.Printf("Established connection:\nProxy:%v <-> Replica%v", p.Addr, p.targetAddr)
	}
}
