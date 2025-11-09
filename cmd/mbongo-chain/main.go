package main

import (
	"fmt"
	"time"

	computeengine "github.com/gkalombo21/mbongo-chain/internal/ai/compute_engine"
	network "github.com/gkalombo21/mbongo-chain/internal/ai/network"
)

func main() {
	fmt.Println("🚀 Mbongo Node Started")
	fmt.Println("🔗 Initializing AI Compute Engine...")

	// Start the AI compute engine in a separate goroutine
	go computeengine.StartEngine()

	fmt.Println("🌐 Bootstrapping AI network...")
	net := network.NewNetwork()
	net.RegisterNode(&network.Node{ID: "node-gpu-01", Type: "gpu", Address: "127.0.0.1:9100"})

	netManager := network.NewManager(net)
	netManager.Start()

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for range ticker.C {
			jobID := fmt.Sprintf("demo-%d", time.Now().Unix())
			req := network.JobRequest{JobID: jobID, Model: "demo-model", PayloadURI: "demo://payload", RewardHint: 0.25}
			netManager.BroadcastJobRequest(req)
		}
	}()

	for {
		time.Sleep(10 * time.Second)
		fmt.Println("⏳ Mbongo Node running... (heartbeat)")
	}
}
