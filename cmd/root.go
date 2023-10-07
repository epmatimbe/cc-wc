/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var FileNameC string
var FileNameL string
var FileNameW string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "A code challenge Word Count",
	Long:  `Word counter for Code Challenge`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {

		if FileNameC != "" {
			content, err := os.ReadFile(FileNameC)
			if err != nil {
				return err
			}
			fmt.Printf("%d %s\n", len(content), FileNameC)
		}
		if FileNameL != "" {
			lineCount := 0
			content, err := os.ReadFile(FileNameL)
			if err != nil {
				return err
			}
			strReader := strings.NewReader(string(content))
			reader := bufio.NewReader(strReader)
			for {
				_, _, err = reader.ReadLine()
				if err != nil {
					break
				}
				lineCount++
			}
			fmt.Printf("%d %s\n", lineCount, FileNameL)
		}
		if FileNameW != "" {
			content, err := os.ReadFile(FileNameW)
			if err != nil {
				return err
			}
			words, _, err := bufio.ScanWords(content, false)

			fmt.Printf("%d %s\n", words, FileNameW)
		}

		return nil
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&FileNameC, "c", "c", "", "")
	rootCmd.Flags().StringVarP(&FileNameL, "l", "l", "", "")
	rootCmd.Flags().StringVarP(&FileNameW, "w", "w", "", "")
}
