package computeengine

import (
	"errors"
	"fmt"
)

// GenerateProof simulates creating a Proof of Useful Work.
func GenerateProof(jobID string, resultHash string) string {
	proof := fmt.Sprintf("proof-%s-%s", jobID, resultHash[:6])
	fmt.Printf("✅ Generated proof for job %s: %s\n", jobID, proof)
	return proof
}

// ProofVerifier describes the behaviour required to validate PoUW proofs.
type ProofVerifier interface {
	Verify(proof Proof) error
}

// simpleVerifier performs lightweight validation to keep the engine extensible.
type simpleVerifier struct{}

// NewSimpleVerifier constructs a proof verifier intended for local simulations.
func NewSimpleVerifier() ProofVerifier {
	return &simpleVerifier{}
}

// Verify checks basic invariants before delegating to future cryptographic modules.
func (v *simpleVerifier) Verify(proof Proof) error {
	if proof.JobID == "" {
		return errors.New("missing job id in proof")
	}

	if proof.NodeID == "" {
		return errors.New("missing node id in proof")
	}

	if proof.ProofHex == "" || proof.Hash == "" {
		return errors.New("proof payload incomplete")
	}

	fmt.Println("verify proof placeholder", proof.JobID)
	// TODO: Integrate with Cosmos SDK proof verification modules.
	return nil
}
