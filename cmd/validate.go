package cmd

import (
	"os"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/manifest"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate [manifest]",
	Short: "Validate a manifest file",
	Long: dedent.Dedent(`
		Validate a manifest file and print the result to stdout.
	`),
	RunE: runValidate,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

func runValidate(cmd *cobra.Command, args []string) error {
	m_file := args[0]

	ctx := context.With(cmd.Context())
	if cwd, err := os.Getwd(); err != nil {
		return err
	} else {
		ctx = ctx.WithCwd(cwd)
	}

	mf := manifest.Manifest{}
	return manifest.Unmarshal(ctx, m_file, &mf)
}
