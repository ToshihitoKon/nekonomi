package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nekonomi",
	Short: "nekonomi is simple KVS command line tool based SQLite3",
}

func main() {
	rootCmd.AddCommand(SetCmd, GetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
