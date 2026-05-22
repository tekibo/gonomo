package main

import (
	"fmt"
	"os"

	"gonomo/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "build":
		cli.RunBuild(args)
	case "dev":
		cli.RunDev(args)
	case "version":
		fmt.Println("gonomo v1.0.0")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`gonomo - Webview App Builder

Usage:
  gonomo <command> [arguments]

Commands:
  build      Build the final .exe
  dev        Run in dev mode (starts dev server + webview)
  version    Print version`)
}
