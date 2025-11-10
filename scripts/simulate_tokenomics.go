package main

import (
	"fmt"

	"github.com/gkalombo21/mbongo-chain/internal/tokenomics"
)

func main() {
	fmt.Println("📊 Mbongo Tokenomics Simulation (Years 1-100)")
	fmt.Println("Year | Block Reward | Annual Minted | Cumulative Minted | PoS | PoUW | Dev | Foundation")
	fmt.Println("------------------------------------------------------------------------------------------")

	for year := 1; year <= tokenomics.MaxYears; year++ {
		reward := tokenomics.CurrentReward(year)
		breakdown := tokenomics.DistributeReward(reward, tokenomics.DefaultDistribution)
		cumulative := tokenomics.TotalMintedUpToYear(year)
		annualMinted := tokenomics.TotalMintedUpToYear(year) - tokenomics.TotalMintedUpToYear(year-1)

		fmt.Printf(
			"%4d | %12.4f | %13.0f | %16.0f | %6.2f | %6.2f | %6.2f | %10.2f\n",
			year,
			reward,
			annualMinted,
			cumulative,
			breakdown["pos_validators"],
			breakdown["gpu_contributors"],
			breakdown["development_fund"],
			breakdown["foundation"],
		)
	}
}
