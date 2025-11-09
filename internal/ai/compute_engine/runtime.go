package computeengine

import (
	"fmt"
	"time"
)

// StartEngine boots the AI compute engine runtime loop.
func StartEngine() {
	fmt.Println("🚀 AI Compute Engine started — waiting for tasks...")
	go jobScheduler()
}

// jobScheduler simulates periodic job dispatching.
func jobScheduler() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("🧠 Dispatching queued AI jobs to available compute nodes...")
	}
}
