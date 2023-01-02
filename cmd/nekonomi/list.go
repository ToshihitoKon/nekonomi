package main

import (
	"fmt"
	"os"

	"github.com/ToshihitoKon/nekonomi"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored keys",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(0)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	opts := []nekonomi.Option{}
	nc, err := nekonomi.New("./", "nekonomi-cli", opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return
	}
	defer nc.Close()

	keys, err := nc.ListKeys()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, key := range keys {
		fmt.Println(key)
	}
}
