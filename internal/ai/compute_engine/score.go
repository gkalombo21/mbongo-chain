package computeengine

// Scorecard maintains operator reputation and performance indicators.
type Scorecard struct {
	OperatorID string
	Score      float64
	JobsTotal  uint64
	JobsFailed uint64
}

// Scorer exposes methods for updating operator scores.
type Scorer interface {
	RecordSuccess(operatorID string, result WorkResult) error
	RecordFailure(operatorID string, request WorkRequest, reason error) error
	GetScore(operatorID string) (Scorecard, bool)
}

// MemoryScorer provides a placeholder in-memory implementation.
type MemoryScorer struct {
	scores map[string]Scorecard
}

// NewMemoryScorer builds an empty scorer ready for use in tests or MVP flows.
func NewMemoryScorer() *MemoryScorer {
	return &MemoryScorer{
		scores: make(map[string]Scorecard),
	}
}

func (m *MemoryScorer) ensure(operatorID string) {
	if _, ok := m.scores[operatorID]; !ok {
		m.scores[operatorID] = Scorecard{OperatorID: operatorID}
	}
}

// RecordSuccess updates the scorecard for a successful job.
// TODO: refine scoring algorithm and decay strategies.
func (m *MemoryScorer) RecordSuccess(operatorID string, result WorkResult) error {
	m.ensure(operatorID)
	card := m.scores[operatorID]
	card.JobsTotal++
	card.Score++
	m.scores[operatorID] = card
	return nil
}

// RecordFailure updates the scorecard for a failed job.
// TODO: integrate penalty multipliers and configurable thresholds.
func (m *MemoryScorer) RecordFailure(operatorID string, request WorkRequest, reason error) error {
	m.ensure(operatorID)
	card := m.scores[operatorID]
	card.JobsTotal++
	card.JobsFailed++
	card.Score--
	m.scores[operatorID] = card
	return nil
}

// GetScore retrieves the current scorecard.
func (m *MemoryScorer) GetScore(operatorID string) (Scorecard, bool) {
	card, ok := m.scores[operatorID]
	return card, ok
}
