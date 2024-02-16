package cmd

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/mecha-ci/ekdo/internal/app"
	cobrax "github.com/mecha-ci/ekdo/internal/x/cobra"
	slogx "github.com/mecha-ci/ekdo/internal/x/slog"
)

type RootCommandFlags struct {
	LogLevel slog.Level
	Debug    bool
}

func NewRootCommand(ctr *app.Container) *cobra.Command {
	const envPrefix = "ekdo"

	root := &cobra.Command{
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cobrax.BindFlags(cmd, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

			flags, err := getRootCommandFlags(cmd)
			if err != nil {
				return err
			}

			ctr.LogLevel = flags.LogLevel
			ctr.Debug = flags.Debug

			slog.SetDefault(
				slog.New(
					slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
						AddSource: true,
						Level:     ctr.LogLevel,
					}),
				),
			)

			return nil
		},
		Use:           "ekdo",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cobrax.BindFlags(root, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

	setupRootCommandFlags(root)

	root.AddCommand(NewVersionCommand(ctr))
	root.AddCommand(NewRenderCommand(ctr))

	return root
}

func setupRootCommandFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(
		"log-level",
		slog.LevelDebug.String(),
		"set the log level",
	)

	cmd.PersistentFlags().Bool(
		"debug",
		true,
		"activates some behaviors that facilitate debugging",
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

	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return RootCommandFlags{}, fmt.Errorf("%w '%s': %w", cobrax.ErrParsingFlag, "debug", err)
	}

	return RootCommandFlags{
		LogLevel: slogLevel,
		Debug:    debug,
	}, nil
}
