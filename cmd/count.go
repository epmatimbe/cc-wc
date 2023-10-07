package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(countCmd)
}

var countCmd = &cobra.Command{
	Use: "c",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Count Words")
	},
}
