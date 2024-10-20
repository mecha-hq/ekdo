package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/mecha-hq/ekdo/internal/app"
	cobrax "github.com/mecha-hq/ekdo/internal/x/cobra"
	slogx "github.com/mecha-hq/ekdo/internal/x/slog"
)

var ErrCannotExecutePersistentPreRun = fmt.Errorf("cannot execute persistent pre-run")

type RootCommandFlags struct {
	LogLevel slog.Level
	Debug    bool
}

func NewRootCommand(ctr *app.Container) *cobra.Command {
	root := &cobra.Command{
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := cobrax.InitializeConfig(cmd, "ekdo", "EKDO"); err != nil {
				return fmt.Errorf("%w: %w", ErrCannotExecutePersistentPreRun, err)
			}

			flags, err := getRootCommandFlags(cmd)
			if err != nil {
				return fmt.Errorf("%w: %w", ErrCannotExecutePersistentPreRun, err)
			}

			setupContainerParameters(ctr, flags)

			setupGlobals(ctr)

			return nil
		},
		Use:           "ekdo",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	setupRootCommandFlags(root)

	root.AddCommand(NewVersionCommand(ctr))
	root.AddCommand(NewRenderCommand(ctr))

	return root
}

func setupRootCommandFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(
		"log-level",
		slog.LevelInfo.String(),
		"Set the log level. Allowed values are: DEBUG, INFO, WARN and ERROR.",
	)
}

func getRootCommandFlags(cmd *cobra.Command) (RootCommandFlags, error) {
	logLevel, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return RootCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "log-level", err)
	}

	slogLevel, err := slogx.FromString(logLevel)
	if err != nil {
		return RootCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "log-level", err)
	}

	return RootCommandFlags{
		LogLevel: slogLevel,
	}, nil
}

func setupContainerParameters(ctr *app.Container, flags RootCommandFlags) {
	ctr.LogLevel = flags.LogLevel
}

func setupGlobals(ctr *app.Container) {
	slog.SetDefault(ctr.Logger())
}
