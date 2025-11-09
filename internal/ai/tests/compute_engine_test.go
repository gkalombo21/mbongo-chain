package tests

import (
	"testing"
	"time"

	computeengine "github.com/gkalombo21/mbongo-chain/internal/ai/compute_engine"
)

func TestJobQueueLifecycle(t *testing.T) {
	tJob := computeengine.Job{ID: "job-queue-1", Model: "demo-model"}
	computeengine.AddJob(tJob)

	next := computeengine.GetNextJob()
	if next == nil {
		t.Fatal("expected job to be dequeued, got nil")
	}

	if next.ID != tJob.ID {
		t.Fatalf("expected job id %s, got %s", tJob.ID, next.ID)
	}
}

func TestEngineHandleResult(t *testing.T) {
	engine := computeengine.NewEngine(nil, nil, nil)

	result := computeengine.Result{
		JobID:      "job-result-1",
		NodeID:     "node-1",
		OutputHash: "hash123",
		Accuracy:   0.92,
		RuntimeMs:  1200,
		Metadata:   map[string]any{"tests": true},
	}

	proof := computeengine.Proof{
		JobID:    result.JobID,
		NodeID:   result.NodeID,
		ProofHex: computeengine.GenerateProof(result.JobID, result.OutputHash),
		Hash:     result.OutputHash,
		Accuracy: result.Accuracy,
	}

	reward, err := engine.HandleResult(result, proof)
	if err != nil {
		t.Fatalf("expected no error handling result, got %v", err)
	}

	if reward.RewardAmountMBG <= 0 {
		t.Fatalf("expected positive reward, got %f", reward.RewardAmountMBG)
	}

	// allow async logs to flush in the engine loop
	time.Sleep(10 * time.Millisecond)
}
