package main

import (
	"fmt"
	"os"

	"github.com/ToshihitoKon/nekonomi"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Store new key and value",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}
		// TODO: auguments validation
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			key   = args[0]
			value = args[1]
		)
		set(key, value)
	},
}

func set(key, value string) {
	opts := []nekonomi.Option{}
	nc, err := nekonomi.New("./", "nekonomi-cli", opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return
	}
	defer nc.Close()

	if _, err := nc.Write(key, value); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
