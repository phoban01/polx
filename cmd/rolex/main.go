package main

import (
	"fmt"
	"os"

	"github.com/phoban01/rolex/pkg/ct"
	"github.com/spf13/cobra"
)

// TODO:
// add cobra with flags
// - lookback (minutes)
// - access-key-id
// - username
// - start-time
// - end-time
// - region
//
// are resources being handled correctly?

var (
	//Version holds current application version
	Version string
	//Build holds application build sha
	Build string
	//BuildDate holds application build date
	BuildDate string
)

func main() {
	cmd := &cobra.Command{
		Use:   "rolex",
		Short: "Rolex helps generate IAM policies",
	}

	cmd.AddCommand(version())
	cmd.AddCommand(ct.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func version() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Rolex helps generate IAM policies",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("rolex 🔑\nversion %s\nbuild: %s\nbuild date: %s\n", Version, Build, BuildDate)
		},
	}
	return cmd
}
