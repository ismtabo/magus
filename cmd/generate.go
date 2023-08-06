package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/domain"
	"github.com/ismtabo/magus/manifest"
	"github.com/ismtabo/magus/variable"
	"github.com/lithammer/dedent"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// generateCmd represents the generate command
	generateCmd = &cobra.Command{
		Use:   "generate [manifest]",
		Short: "Generate files from a manifest",
		Long: dedent.Dedent(`
			Generate files from a manifest
		`),
		Example: dedent.Dedent(`
		Given the following manifest at manifest.yaml:
		  ---
		  version: "1"
		  name: hello-world
		  root: .
		  casts:
		    hello-world:
		      to: ./hello-world.md
		      from: |
		        # Hello World
		        This is my first cast!

		When running:
		  magus generate manifest.yaml

		Then it will generate the following file:
		  ./hello-world.md

		With the content:
		  # Hello World
		  This is my first cast!
		`),
		Args: cobra.ExactArgs(1),
		RunE: runGenerate,
	}
	output_dir              = "."
	dry_run                 = false
	clean                   = false
	overwrite               = false
	variables      []string = []string{}
	variablesFiles []string = []string{}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVar(&output_dir, "dir", ".", "Output directory")
	generateCmd.Flags().BoolVar(&dry_run, "dry-run", false, "Dry run")
	generateCmd.Flags().BoolVar(&clean, "clean", false, "Clean output directory")
	generateCmd.Flags().BoolVarP(&overwrite, "overwrite", "w", false, "Overwrite existing files")
	generateCmd.Flags().StringSliceVar(&variables, "var", []string{}, `Comma separated variables (e.g. '--var foo=bar' or '--var foo=bar,baz=qux')`)
	generateCmd.Flags().StringSliceVar(&variablesFiles, "var-file", []string{}, `Comma separated variables files (e.g. '--var-file ./foo.yaml' or '--var-file ./foo.yaml,./bar.yaml')`)
}

func runGenerate(cmd *cobra.Command, args []string) error {
	m_file := args[0]
	out_dir, _ := filepath.Abs(output_dir)

	ctx := context.With(cmd.Context())
	if cwd, err := os.Getwd(); err != nil {
		return err
	} else {
		ctx = ctx.WithCwd(cwd)
	}

	mf := manifest.Manifest{}
	err := manifest.Unmarshal(ctx, m_file, &mf)
	if err != nil {
		return err
	}

	vars := variable.Variables{}
	for _, v := range variables {
		segments := strings.SplitN(v, "=", 2)
		if segments[0] == "" || segments[1] == "" {
			return errors.Errorf("invalid variable %s", v)
		}
		vars = append(vars, variable.NewLiteralVariable(segments[0], segments[1]))
	}
	for _, f := range variablesFiles {
		fileVars, err := variable.FromFile(ctx, f)
		if err != nil {
			return err
		}
		vars = append(vars, fileVars...)
	}

	files, err := domain.Generate(ctx, out_dir, mf, domain.GenerateOptions{
		DryRun:    dry_run,
		Clean:     clean,
		Overwrite: overwrite,
		Variables: vars,
	})
	if err != nil {
		return err
	}

	log.Printf("Generated files %d files\n", len(files))
	for _, f := range files {
		path, _ := f.Rel(ctx)
		log.Printf("File: %s\n", path)
	}

	return nil
}
