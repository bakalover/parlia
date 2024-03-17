package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bakalover/parlia/proxy/proxy"
	"github.com/bakalover/tate"
)

const (
	proxyConfigPath   = "./config/proxy_ports.txt"
	replicaConfigPath = "./config/replica_config.txt"
)

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

	//=============================Files================================
	proxyConfig, err := os.Open(proxyConfigPath)
	if err != nil {
		logger.Fatalf("Error opening file: %v", err)
	}
	replicaConfig, err := os.Open(replicaConfigPath)
	if err != nil {
		logger.Fatalf("Error opening file: %v ", err)
	}
	defer proxyConfig.Close()
	defer replicaConfig.Close()
	//=============================Files================================

	replicaScanner := bufio.NewScanner(replicaConfig)
	var replicaAddrs []string

	for replicaScanner.Scan() {
		replicaPort := replicaScanner.Text()
		if len(replicaPort) <= 4 || len(replicaPort) >= 6 {
			logger.Fatalf("Parsed invalid replica port: %v", replicaPort)
		}
		replicaAddrs = append(replicaAddrs, fmt.Sprintf("localhost:%s", replicaPort))
	}

	var proxies []*proxy.Proxy

	proxyScanner := bufio.NewScanner(proxyConfig)
	for proxyScanner.Scan() {
		portProxy := proxyScanner.Text()

		if len(portProxy) <= 4 || len(portProxy) >= 6 {
			logger.Fatalf("Parsed invalid proxy port: %v", portProxy)
		}

		proxyAddr := fmt.Sprintf("localhost:%s", portProxy)

		proxy := &proxy.Proxy{
			AvailableReplicas: replicaAddrs,
			Addr:              proxyAddr,
			Logger:            logger,
		}
		proxies = append(proxies, proxy)

		var proxyRoutines tate.Nursery
		proxyRoutines.Add(func() {
			proxy.Run()
		})
	}

	//=================================Await===================================

	//=====================Manual Cancelling=====================
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	//=====================Manual Cancelling=====================

	var cancelRoutine tate.Nursery
	cancelRoutine.Add(func() {
		<-sigCh
		for _, p := range proxies {
			p.ShutDown()
		}
	})

	cancelRoutine.Join()
	//=================================Await===================================

}
