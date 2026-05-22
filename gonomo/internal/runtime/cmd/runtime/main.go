package main

import (
	"fmt"
	"os"

	"gonomo/runtime/internal/app"
	"gonomo/runtime/internal/config"
)

func main() {
	cfg := &config.AppConfig

	if err := app.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
