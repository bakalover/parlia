package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bakalover/parlia/client/client"
	"github.com/bakalover/tate"
)

const (
	proxyConfigPath  = "./config/proxy_ports.txt"
	clientConfigPath = "./config/client_config.txt"
	simTime          = 40 * time.Second
)

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

	//=============================Files================================
	proxyConfig, err := os.Open(proxyConfigPath)
	if err != nil {
		logger.Fatalf("Error opening file: %v", err)
		return
	}
	defer proxyConfig.Close()
	clientConfig, err := os.Open(clientConfigPath)
	if err != nil {
		logger.Fatalf("Error opening file: %v", err)
		return
	}
	defer clientConfig.Close()
	//=============================Files================================

	var proxyAddrs []string
	scanner := bufio.NewScanner(proxyConfig)
	for scanner.Scan() {
		proxyPort := scanner.Text()

		if len(proxyPort) <= 4 || len(proxyPort) >= 6 {
			logger.Fatalf("Parsed invalid replica port: %v", proxyPort)
		}

		addr := fmt.Sprintf("localhost:%s", proxyPort)
		proxyAddrs = append(proxyAddrs, addr)
	}

	logger.Printf("Available proxies: %v", proxyAddrs)

	var clientRoutines tate.Nursery
	scanner = bufio.NewScanner(clientConfig)
	for scanner.Scan() {
		clientId := scanner.Text()
		cl := &client.Client{
			Logger:           logger,
			AvailableProxies: proxyAddrs,
			Id:               clientId,
			SimTime:          simTime,
			Backoff:          client.DefaultBackoff,
		}
		clientRoutines.Add(func() {
			cl.Run()
		})
	}
	clientRoutines.Join()
}
