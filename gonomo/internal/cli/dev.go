package cli

import (
	"fmt"
	"os"

	"gonomo/internal/builder"
)

func RunDev(args []string) {
	fmt.Println("Starting gonomo dev mode...")
	err := builder.Dev()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dev mode failed: %v\n", err)
		os.Exit(1)
	}
}
