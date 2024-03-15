package main

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
	cluster *Cluster
	logger  *log.Logger
	server  *grpc.Server
	addr    string
}

func (r *Replica) Kill() {
	r.server.GracefulStop()
	r.log = nil
}

func (r *Replica) Step(stepTime time.Duration) {
	addr, err := r.cluster.PickAddr()
	if err != nil {
		r.logger.Fatalf("Cannot pick addr")
	}

	r.addr = addr

	//Setting up grpc server
	l, err := net.Listen("tcp", r.addr)
	if err != nil {
		r.logger.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	r.server = grpcServer
	pb.RegisterReplicaServer(grpcServer, r)

	//Launch server
	var serverRoutine tate.Nursery
	serverRoutine.Add(func() {
		grpcServer.Serve(l)
	})

	time.AfterFunc(stepTime, func() {
		r.Kill()
	})

	serverRoutine.Join()
}

//===============================RPC Service===============================

func (p *Replica) Apply(ctx context.Context, command *pb.Command) (*pb.Empty, error) {
	return nil, nil
}

func (p *Replica) Prepare(ctx context.Context, req *pb.PrepareRequest) (*pb.Promise, error) {
	return nil, nil
}

func (p *Replica) Accept(ctx context.Context, req *pb.AcceptRequest) (*pb.Accepted, error) {
	return nil, nil
}

//===============================RPC Service================================
