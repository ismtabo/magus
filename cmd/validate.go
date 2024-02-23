package cmd

import (
	"os"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/manifest"
	"github.com/ismtabo/magus/v2/validate"
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
	if err := manifest.Unmarshal(ctx, m_file, &mf); err != nil {
		return err
	}
	if err := validate.Version(ctx, mf); err != nil {
		return err
	}
	if err := validate.NoCycles(ctx, mf); err != nil {
		return err
	}
	return nil
}
