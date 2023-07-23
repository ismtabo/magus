/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionString = "0.0.0"
	// versionCmd represents the version command
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Magus version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(versionString)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
