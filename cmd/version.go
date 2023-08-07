package cmd

import (
	"fmt"

	"github.com/ismtabo/magus/config"
	"github.com/spf13/cobra"
)

var (
	// versionCmd represents the version command
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Magus version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(config.GetVersion())
			if build {
				fmt.Println(config.GetBuildTime())
			}
			if _os {
				fmt.Println(config.GetOS())
			}
		},
	}
	build bool
	_os   bool
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&build, "build", "b", false, "Show build time")
	versionCmd.Flags().BoolVarP(&_os, "os", "o", false, "Show OS")
}
