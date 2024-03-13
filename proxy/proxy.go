package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/bakalover/parlia/proto"
	"github.com/bakalover/tate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Proxy struct {
	pb.UnimplementedProxyServer
	client            pb.ReplicaClient
	logger            *log.Logger
	availableReplicas []string
	targetAddr        string
	myAddr            string
}

const (
	proxyConfigPath   = "./config/proxy_ports.txt"
	replicaConfigPath = "./config/replica_ports.txt"
)

// var dialOpts []grpc.DialOption

func (p *Proxy) Apply(ctx context.Context, command *pb.Command) (*pb.Empty, error) {

	//Context???

	resp, err := p.client.Apply(ctx, command)
	if err != nil {
		p.logger.Println("Replica unavailable")
		p.TryConnToCluster()
	}

	return resp, err
}

func (p *Proxy) AvailableReplica() string {
	return p.availableReplicas[rand.Intn(len(p.availableReplicas))]
}

func (p *Proxy) TryConnToCluster() bool {
	replicaAddr := p.AvailableReplica()
	conn, err := grpc.Dial(replicaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		p.logger.Printf("Fail to dial: %v", err)
		return false
	} else {
		p.targetAddr = replicaAddr
		p.client = pb.NewReplicaClient(conn)
		return true
	}
}

func main() {
	//=============================Files================================
	proxyConfig, err := os.Open(proxyConfigPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	replicaConfig, err := os.Open(replicaConfigPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer proxyConfig.Close()
	defer replicaConfig.Close()
	//=============================Files================================

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	replicaScanner := bufio.NewScanner(replicaConfig)
	var replicaAddrs []string

	for replicaScanner.Scan() {
		replicaPort := replicaScanner.Text()
		if len(replicaPort) <= 4 || len(replicaPort) >= 6 {
			logger.Fatalf("Parsed invalid replica port: %v", replicaPort)
		}
		replicaAddrs = append(replicaAddrs, fmt.Sprintf("localhost:%s", replicaPort))
	}

	var serverRoutines, cancels tate.Nursery

	//=====================Manual Cancelling=====================
	var serverHandles []*grpc.Server
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	//=====================Manual Cancelling=====================

	proxyScanner := bufio.NewScanner(proxyConfig)
	for proxyScanner.Scan() {
		portProxy := proxyScanner.Text()

		if len(portProxy) <= 4 || len(portProxy) >= 6 {
			logger.Fatalf("Parsed invalid proxy port: %v", portProxy)
		}

		proxyAddr := fmt.Sprintf("localhost:%s", portProxy)

		proxy := &Proxy{
			availableReplicas: replicaAddrs,
			myAddr:            proxyAddr,
			logger:            logger,
		}

		// Proxy as client
		proxy.TryConnToCluster()

		// Proxy as server
		l, err := net.Listen("tcp", proxyAddr)
		if err != nil {
			logger.Fatalf("Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		serverHandles = append(serverHandles, grpcServer)
		pb.RegisterProxyServer(grpcServer, proxy)
		serverRoutines.Add(func() {
			grpcServer.Serve(l)
		})

	}

	//=================================Await===================================
	cancels.Add(func() {
		<-sigCh
		for _, h := range serverHandles {
			h.GracefulStop()
		}
	})
	//=================================Await===================================

	cancels.Join()
	serverRoutines.Join()

}
