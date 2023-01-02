package main

import (
	"fmt"
	"os"

	"github.com/ToshihitoKon/nekonomi"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Read value",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		// TODO: auguments validation
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			key = args[0]
		)
		get(key)
	},
}

func get(key string) {
	opts := []nekonomi.Option{}
	nc, err := nekonomi.New("./", "nekonomi-cli", opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return
	}
	defer nc.Close()

	value, err := nc.Read(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(value)
}
