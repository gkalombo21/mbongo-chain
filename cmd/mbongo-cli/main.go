package main

import (
	"fmt"
	"os"
)

// TODO: integrate cobra or another CLI framework to manage subcommands.
func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "mbongo-cli error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	// TODO: dispatch commands for node control, job submission, and queries.
	fmt.Println("mbongo-cli placeholder - commands not yet implemented")
	return nil
}

