// [[file:../docs.org::Code/Source/cmd/root.go][Code/Source/cmd/root.go]]
/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godadjoke",
	Short: "Get dad jokes amazingly fast right in your terminal",
	Long: `godadjoke CLI empowers applications with dad jokes.
Generate dad jokes to quickly liven up your beloved terminal.`,
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
// Code/Source/cmd/root.go ends here
