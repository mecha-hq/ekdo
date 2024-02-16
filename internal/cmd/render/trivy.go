package render

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mecha-ci/ekdo/internal/app"
	"github.com/mecha-ci/ekdo/internal/scn/trivy"
	cobrax "github.com/mecha-ci/ekdo/internal/x/cobra"
)

var (
	ErrCannotCompleteTrivyCommand = fmt.Errorf("cannot complete scan command")
)

type TrivyCommandFlags struct {
	DrawLayout bool
	OutputDir  string
	Publish    bool
}

// TODO: this command can probably be refactored into its parent, RenderCommand.
func NewTrivyCommand(_ *app.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trivy",
		Short: "render trivy scans results",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			flags, err := getTrivyCommandFlags(cmd)
			if err != nil {
				return err
			}

			input := args[0]

			sr, err := trivy.NewScanRenderer(input, flags.OutputDir)
			if err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteTrivyCommand, err)
			}

			if err := sr.Render(flags.DrawLayout); err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteTrivyCommand, err)
			}

			if err := sr.PublishAssets(flags.OutputDir); err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteTrivyCommand, err)
			}

			return nil
		},
	}

	return cmd
}

func getTrivyCommandFlags(cmd *cobra.Command) (TrivyCommandFlags, error) {
	drawLayout, err := cmd.Flags().GetBool("draw-layout")
	if err != nil {
		return TrivyCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "draw-layout", err)
	}

	outputDir, err := cmd.Flags().GetString("output-dir")
	if err != nil {
		return TrivyCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "output", err)
	}

	publish, err := cmd.Flags().GetBool("publish")
	if err != nil {
		return TrivyCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "publish", err)
	}

	return TrivyCommandFlags{
		DrawLayout: drawLayout,
		OutputDir:  outputDir,
		Publish:    publish,
	}, nil
}
