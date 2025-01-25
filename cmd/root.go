package cmd

import (
	"os"

	"github.com/nodearch/cli/cmd/create"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nodearch",
	Short: "A CLI for building scalable Node.js backends with TypeScript and DI",
	Long: `A command-line tool for building scalable Node.js backends with TypeScript 
and dependency injection. Use nodearch help for available commands.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Welcome to NodeArch CLI")
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(create.Command)
	RegisterCommands(rootCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().String("loadMode", "", "load mode, either js or ts")
}
