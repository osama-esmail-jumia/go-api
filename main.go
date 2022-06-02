package main

import (
	"github.com/spf13/cobra"
	"go-api/cmd"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "app",
		Short: "Run App Commands",
	}
	
	rootCmd.AddCommand(
		cmd.API(),
	)
	
	cobra.CheckErr(rootCmd.Execute())
}
