package main

import (
	"fmt"
	"os"

	"github.com/phoban01/polx/pkg/ct"
	"github.com/spf13/cobra"
)

// TODO: check resources being handled correctly

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
		Use:   "polx",
		Short: "polx helps generate IAM policies",
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
		Short: "Polx helps generate IAM policies",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("polx 🔑\nversion %s\nbuild: %s\nbuild date: %s\n", Version, Build, BuildDate)
		},
	}
	return cmd
}
