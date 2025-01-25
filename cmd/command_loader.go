package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type CommandConfig struct {
	Command     string          `json:"command"`
	Description string          `json:"description"`
	Flags       map[string]Flag `json:"flags"`
}

type Flag struct {
	Description string `json:"description"`
	Argument    string `json:"argument"`
}

func RegisterCommands(rootCmd *cobra.Command) {
	// Load JSON
	commands, err := loadCommands("commands.json")

	if err != nil {
		fmt.Println("Error loading commands:", err)
		os.Exit(1)
	}

	// Dynamically add commands from JSON
	for _, cmdConfig := range commands {
		cmd := createCommandFromConfig(cmdConfig)
		rootCmd.AddCommand(cmd)
	}
}

func loadCommands(filePath string) ([]CommandConfig, error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode JSON into slice of CommandConfig
	var commands []CommandConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&commands)
	if err != nil {
		return nil, err
	}

	return commands, nil
}

func createCommandFromConfig(cmdConfig CommandConfig) *cobra.Command {
	// Create a new command
	cmd := &cobra.Command{
		Use:   cmdConfig.Command,
		Short: cmdConfig.Description,
		Run: func(cmd *cobra.Command, args []string) {
			var flagPairs []string
			cmd.Flags().VisitAll(func(flag *pflag.Flag) {
				flagPairs = append(flagPairs, fmt.Sprintf("--%s=%s", flag.Name, flag.Value.String()))
			})

			allArgs := append(flagPairs, args...)

			RunJs(allArgs)
		},
	}

	// Add flags to the command
	for flagName, flagConfig := range cmdConfig.Flags {
		flagType := flagConfig.Argument

		// Support string-type flags for now
		if flagType == "file" {
			cmd.Flags().StringP(flagName, "", "", flagConfig.Description)
		}
	}

	return cmd
}
