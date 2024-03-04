package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mecha-ci/ekdo/internal/app"
	cobrax "github.com/mecha-ci/ekdo/internal/x/cobra"
)

var (
	ErrCannotCompleteRenderCommand = fmt.Errorf("cannot complete render command")
)

type RenderCommandFlags struct {
	DrawLayout bool
	OutputDir  string
	Publish    bool
}

func NewRenderCommand(ctr *app.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "render",
		Short: "render resources",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			flags, err := getRenderCommandFlags(cmd)
			if err != nil {
				return err
			}

			toolName := args[0]
			inputFile := args[1]

			sr, err := ctr.ScanRendererFactory().Create(toolName, inputFile, flags.OutputDir)
			if err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteRenderCommand, err)
			}

			if err := sr.Render(flags.DrawLayout); err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteRenderCommand, err)
			}

			if err := sr.PublishAssets(flags.OutputDir); err != nil {
				return fmt.Errorf("%w: %w", ErrCannotCompleteRenderCommand, err)
			}

			return nil
		},
	}

	setupRenderCommandFlags(cmd)

	return cmd
}

func setupRenderCommandFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(
		"draw-layout",
		"l",
		true,
		"Render the outer layout of the report, including the header, footer, and navbar.",
	)

	cmd.PersistentFlags().StringP(
		"output-dir",
		"o",
		".",
		"Output dir where to render the scan results, defaults the folder ekdo is launched.",
	)

	cmd.PersistentFlags().BoolP(
		"publish",
		"p",
		false,
		"Publish the files to the configured remote destination.",
	)

	// cmd.PersistentFlags().BoolP(
	// 	"publish",
	// 	"p",
	// 	false,
	// 	"Publish the files to the configured remote destination.",
	// )
}

func getRenderCommandFlags(cmd *cobra.Command) (RenderCommandFlags, error) {
	drawLayout, err := cmd.Flags().GetBool("draw-layout")
	if err != nil {
		return RenderCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "draw-layout", err)
	}

	outputDir, err := cmd.Flags().GetString("output-dir")
	if err != nil {
		return RenderCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "output", err)
	}

	publish, err := cmd.Flags().GetBool("publish")
	if err != nil {
		return RenderCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "publish", err)
	}

	return RenderCommandFlags{
		DrawLayout: drawLayout,
		OutputDir:  outputDir,
		Publish:    publish,
	}, nil
}
