package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/divultion/nevern/cli/commands"
)

var rootCmd = &cobra.Command{
	Use:   "nevern-cli@divultion",
	Short: "Nevern is multi-client reverse shell",
	Long:  `This is CLI tool for nevern@divultion. Nevern is a multi-client, multi-threaded reverse shell written in Go. We are aiming for backcompitability with all existing NetCat reverse shells.`,
}

func main() {
	rootCmd.PersistentFlags().String("shost", "0.0.0.0", "Set up Host where the program will be running")
	rootCmd.PersistentFlags().String("sp", "3303", "Set up Port where the program will be running")
	rootCmd.AddCommand(commands.ListConnectionsCommand)
	rootCmd.AddCommand(commands.SelectCommand)
	rootCmd.AddCommand(commands.DisconnectCommand)
	rootCmd.AddCommand(commands.ForgetCommand)
	rootCmd.AddCommand(commands.ClearCommand)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
