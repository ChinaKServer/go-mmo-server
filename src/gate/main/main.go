package main

import (
	"github.com/jzyong/go-mmo-server/src/gate/config"
	"github.com/jzyong/go-mmo-server/src/gate/rpc"
	"log"
)

func main() {
	log.Printf("gate:%d starting", config.GateConfig.Id)

	rpc.GateToClusterClient = new(rpc.GateToCluster)
	rpc.GateToClusterClient.Start(config.GateConfig.ClusterRpcURL)
	rpc.RegisterToCluster()
}
