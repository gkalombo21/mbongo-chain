package tokenomics

import (
	"fmt"
	"math"
)

// Global supply constants expressed in MBG units.
const (
	TotalSupply         uint64  = 10_000_000_000
	InitialCirculation  uint64  = 2_000_000_000
	BlockTimeMinutes             = 5
	HalvingIntervalYears         = 5
	MaxYears                     = 100
	initialBlockReward   float64 = 50.0
)

// RewardDistribution defines relative splits across incentive targets.
type RewardDistribution struct {
	PoSValidators   float64
	GPUContributors float64
	DevFund         float64
	Foundation      float64
}

// DefaultDistribution reflects the agreed MBG allocation strategy.
var DefaultDistribution = RewardDistribution{
	PoSValidators:   0.50,
	GPUContributors: 0.30,
	DevFund:         0.15,
	Foundation:      0.05,
}

// CurrentReward returns the block reward (MBG) for a given network year.
// Rewards halve every HalvingIntervalYears until MaxYears.
func CurrentReward(year int) float64 {
	if year < 0 {
		year = 0
	}

	halvingPeriods := year / HalvingIntervalYears
	reward := initialBlockReward / math.Pow(2, float64(halvingPeriods))

	if year >= MaxYears {
		return 0
	}

	return reward
}

// DistributeReward splits a total reward according to the distribution ratios.
func DistributeReward(totalReward float64, dist RewardDistribution) map[string]float64 {
	return map[string]float64{
		"pos_validators":   totalReward * dist.PoSValidators,
		"gpu_contributors": totalReward * dist.GPUContributors,
		"development_fund": totalReward * dist.DevFund,
		"foundation":       totalReward * dist.Foundation,
	}
}

// TotalMintedUpToYear estimates cumulative issuance up to (but not including) the given year.
func TotalMintedUpToYear(year int) float64 {
	if year <= 0 {
		return 0
	}

	total := 0.0
	for y := 0; y < year && y < MaxYears; y++ {
		reward := CurrentReward(y)
		blocksPerYear := float64((60 / BlockTimeMinutes) * 24 * 365)
		total += reward * blocksPerYear
	}

	if total > float64(TotalSupply) {
		return float64(TotalSupply)
	}

	return total
}

func init() {
	fmt.Println("📈 Mbongo Tokenomics Simulation (first 30 years)")
	for year := 0; year <= 30; year += 5 {
		reward := CurrentReward(year)
		fmt.Printf("Year %02d -> Block Reward: %.4f MBG, Total Minted: %.0f MBG\n", year, reward, TotalMintedUpToYear(year))
	}
}
