package cli

import (
	"flag"
	"fmt"
	"os"

	"gonomo/internal/builder"
)

func RunBuild(args []string) {
	fs := flag.NewFlagSet("build", flag.ExitOnError)
	skipFrontend := fs.Bool("skip-frontend", false, "Skip the frontend build step")
	verbose := fs.Bool("verbose", false, "Show detailed build output")
	clean := fs.Bool("clean", false, "Clean .gonomo/ before building")

	fs.Parse(args)

	opts := builder.BuildOptions{
		SkipFrontend: *skipFrontend,
		Verbose:      *verbose,
		Clean:        *clean,
	}

	fmt.Println("Starting gonomo build...")
	err := builder.Build(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Build failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Build completed successfully!")
}
