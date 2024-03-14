package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/bakalover/parlia/client/backoff"
	pb "github.com/bakalover/parlia/proto"
	"github.com/bakalover/tate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	proxyConfigPath = "./config/proxy_ports.txt"
	simTime         = 40 * time.Second
)

type Client struct {
	client           pb.ProxyClient
	backoff          backoff.Backoff
	logger           *log.Logger
	mutex            sync.Mutex
	availableProxies []string
	targetAddr       string
}

func (cl *Client) AvailableProxy() string {
	return cl.availableProxies[rand.Intn(len(cl.availableProxies))]
}

func (cl *Client) SendCommand() {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()

	_, err := cl.client.Apply(context.Background(), &pb.Command{Type: "hi"})

	if err != nil {
		cl.logger.Printf("Proxy with addr: %v unavailable", cl.targetAddr)
		time.Sleep(cl.backoff.Next())
	} else {
		cl.backoff.Reset()
	}

}

func (cl *Client) ConnToCluster() {
	proxyAddr := cl.AvailableProxy()
	conn, err := grpc.Dial(proxyAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		cl.logger.Fatalf("Fail to dial: %v", err)
	} else {
		cl.targetAddr = proxyAddr
		cl.client = pb.NewProxyClient(conn)
	}
}

func (cl *Client) Run() {
	//Proxies are not faulty
	cl.ConnToCluster()

	await := make(chan bool)

	// DDOS =)
	rp := tate.NewRepeater()
	rp.Repeat(func() {
		cl.SendCommand()
	})

	time.AfterFunc(simTime, func() {
		rp.Join()
		await <- true
	})

	<-await
}

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

	//=============================Files================================
	proxyConfig, err := os.Open(proxyConfigPath)
	if err != nil {
		logger.Fatalf("Error opening file: %v", err)
		return
	}
	defer proxyConfig.Close()
	//=============================Files================================

	var proxyAddrs []string

	scanner := bufio.NewScanner(proxyConfig)
	var clientRoutines tate.Nursery

	for scanner.Scan() {
		proxyPort := scanner.Text()
		if len(proxyPort) <= 4 || len(proxyPort) >= 6 {
			logger.Fatalf("Parsed invalid replica port: %v", proxyPort)
		}

		addr := fmt.Sprintf("localhost:%s", proxyPort)
		logger.Printf("Connecting to proxy with addr: %v", addr)
		proxyAddrs = append(proxyAddrs, addr)
		cl := &Client{backoff: backoff.DefaultBackoff, logger: logger, availableProxies: proxyAddrs}
		clientRoutines.Add(func() {
			cl.Run()
		})
	}

	clientRoutines.Join()

}
