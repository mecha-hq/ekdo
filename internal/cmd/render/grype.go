package render

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mecha-ci/ekdo/internal/app"
	"github.com/mecha-ci/ekdo/internal/scn/grype"
	cobrax "github.com/mecha-ci/ekdo/internal/x/cobra"
)

var (
	ErrCannotCompleteGrypeCommand = fmt.Errorf("cannot complete scan command")
)

type GrypeCommandFlags struct {
	DrawLayout bool
	OutputDir  string
	Publish    bool
}

// TODO: this command can probably be refactored into its parent, RenderCommand.
func NewGrypeCommand(_ *app.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grype",
		Short: "render grype scans results",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			flags, err := getGrypeCommandFlags(cmd)
			if err != nil {
				return err
			}

			input := args[0]

			sr, err := grype.NewScanRenderer(input, flags.OutputDir)
			if err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteGrypeCommand, err)
			}

			if err := sr.Render(flags.DrawLayout); err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteGrypeCommand, err)
			}

			if err := sr.PublishAssets(flags.OutputDir); err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteGrypeCommand, err)
			}

			return nil
		},
	}

	return cmd
}

func getGrypeCommandFlags(cmd *cobra.Command) (GrypeCommandFlags, error) {
	drawLayout, err := cmd.Flags().GetBool("draw-layout")
	if err != nil {
		return GrypeCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "draw-layout", err)
	}

	outputDir, err := cmd.Flags().GetString("output-dir")
	if err != nil {
		return GrypeCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "output", err)
	}

	publish, err := cmd.Flags().GetBool("publish")
	if err != nil {
		return GrypeCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "publish", err)
	}

	return GrypeCommandFlags{
		DrawLayout: drawLayout,
		OutputDir:  outputDir,
		Publish:    publish,
	}, nil
}
