package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	cluster "github.com/bakalover/parlia/cluster/impl"
)

const (
	replicaConfigPath = "./config/replica_config.txt"
)

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

	// =============================Files================================
	replicaConfig, err := os.Open(replicaConfigPath)
	if err != nil {
		logger.Fatalf("Error opening file: %v ", err)
	}
	defer replicaConfig.Close()
	//=============================Files================================

	scanner := bufio.NewScanner(replicaConfig)
	var replicaConfigs []cluster.ReplicaConfig

	for scanner.Scan() {
		config := strings.Split(scanner.Text(), " ")
		replicaPort := config[0]
		replicaMode := config[1]

		if len(replicaPort) <= 4 || len(replicaPort) >= 6 {
			logger.Fatalf("Parsed invalid replica port: %v", replicaPort)
		}

		if replicaMode != "Trust" && replicaMode != "Faulty" {
			logger.Fatalf("Invalide replica mode: %v\n Avaliable:\n- Trust\n- Faulty", replicaMode)
		}
		addr := fmt.Sprintf("localhost:%v", replicaPort)
		replicaConfigs = append(replicaConfigs, cluster.ReplicaConfig{Addr: addr, Mode: cluster.RunMode(replicaMode)})
	}


	c := &cluster.Cluster{Logger: logger}
	c.Run(replicaConfigs)
}
