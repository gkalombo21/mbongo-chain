package tests

import (
	"testing"
	"time"

	network "github.com/gkalombo21/mbongo-chain/internal/ai/network"
)

func TestNetworkRegistration(t *testing.T) {
	net := network.NewNetwork()
	net.RegisterNode(&network.Node{ID: "node-alpha", Type: "gpu", Address: "127.0.0.1:9001"})

	nodes := net.ListNodes()
	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}

	if nodes[0].ID != "node-alpha" {
		t.Fatalf("expected node id node-alpha, got %s", nodes[0].ID)
	}
}

func TestManagerBroadcastAndResult(t *testing.T) {
	net := network.NewNetwork()
	net.RegisterNode(&network.Node{ID: "node-beta", Type: "cpu", Address: "127.0.0.1:9002"})

	manager := network.NewManager(net)
	manager.Start()

	defer manager.Stop()

	req := network.JobRequest{JobID: "job-broadcast-1", Model: "demo", PayloadURI: "demo://payload"}
	manager.BroadcastJobRequest(req)

	manager.SubmitResult(network.JobResult{JobID: req.JobID, NodeID: "node-beta", OutputHash: "hash", Accuracy: 0.8})

	time.Sleep(20 * time.Millisecond)
}
