package cluster

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/bakalover/parlia/proto"
	"github.com/bakalover/tate"
	"google.golang.org/grpc"
)

type Replica struct {
	pb.UnimplementedReplicaServer
	log     []string
	server  *grpc.Server
	addr    string
	Cluster *Cluster
	Logger  *log.Logger
}

func (r *Replica) Kill() {
	r.server.GracefulStop()
	r.log = nil
}

func (r *Replica) Step(stepTime time.Duration) {
	addr, err := r.Cluster.PickAddr()
	if err != nil {
		r.Logger.Fatalf("Cannot pick addr")
	}

	r.addr = addr

	//Setting up grpc server
	l, err := net.Listen("tcp", r.addr)
	if err != nil {
		r.Logger.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	r.server = grpcServer
	pb.RegisterReplicaServer(grpcServer, r)

	// Launch server then gracefull shutdown
	// New Step ONLY after end of shutdown -> await on Kill() -> Join() await on Kill
	var serverRoutine tate.Nursery
	serverRoutine.Add(func() {
		grpcServer.Serve(l)
		r.Cluster.InvalidateAddr(r.addr)
	}).Add(func() {
		await := make(chan bool)
		time.AfterFunc(stepTime, func() {
			r.Kill()
			await <- true
		})
		<-await
	})

	serverRoutine.Join()
}

//===============================RPC Service===============================

func (p *Replica) Apply(ctx context.Context, command *pb.Command) (*pb.Empty, error) {
	MaybeMajorDelay()
	return nil, nil
}

func (p *Replica) Prepare(ctx context.Context, req *pb.PrepareRequest) (*pb.Promise, error) {
	MaybeMajorDelay()
	return nil, nil
}

func (p *Replica) Accept(ctx context.Context, req *pb.AcceptRequest) (*pb.Accepted, error) {
	MaybeMajorDelay()
	return nil, nil
}

//===============================RPC Service================================
