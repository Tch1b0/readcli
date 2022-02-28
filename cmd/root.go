package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readcli",
	Short: "Generate READMEs for projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		if Version {
			fmt.Printf("Version: %s", BuildVersion)
			return nil
		}
		if Help {
			return cmd.Help()
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

var (
	BuildVersion string = "0.0.0"
	Version      bool
	Help         bool
)

func init() {
	rootCmd.Flags().BoolVar(
		&Version,
		"version",
		false,
		"Print version",
	)
	rootCmd.Flags().BoolVar(
		&Help,
		"help",
		false,
		"Open help",
	)
}
