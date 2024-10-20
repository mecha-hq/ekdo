package cmd

import (
	"fmt"

	"github.com/mecha-hq/ekdo/internal/app"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func NewVersionCommand(ctr *app.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number and build information.",
		Run: func(_ *cobra.Command, _ []string) {
			keys := maps.Keys(ctr.Versions)

			slices.Sort(keys)

			for _, k := range keys {
				fmt.Printf("%s: %s\n", k, ctr.Versions[k])
			}
		},
	}
}
