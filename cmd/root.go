package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tch1b0/readcli/pkg/utility"
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
		var readme utility.Readme
		var err error
		if len(loadConfig) != 0 {
			content, err := os.ReadFile(loadConfig)
			if err != nil {
				return err
			}
			json.Unmarshal(content, &readme)
		} else {
			readme, err = CreateReadme()
			if err != nil {
				return err
			}
		}
		content, err := readme.Render()
		if err != nil {
			return err
		}

		return os.WriteFile(outPath, content, 0600)
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
	loadConfig   string
	outPath      string
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
	rootCmd.Flags().StringVar(
		&loadConfig,
		"load",
		"",
		"Load readcli config",
	)
	rootCmd.Flags().StringVar(
		&outPath,
		"out",
		"./README.md",
		"Path of the output file",
	)
}
