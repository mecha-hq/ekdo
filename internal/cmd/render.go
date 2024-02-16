package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mecha-ci/ekdo/internal/app"
	"github.com/mecha-ci/ekdo/internal/cmd/render"
)

func NewRenderCommand(ctr *app.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "render",
		Short: "render resources",
	}

	cmd.AddCommand(render.NewTrivyCommand(ctr))

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
