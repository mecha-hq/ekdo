package main

import (
	"log"

	"github.com/mecha-ci/ekdo/internal/app"
	"github.com/mecha-ci/ekdo/internal/cmd"
)

var (
	version   = "unknown"
	gitCommit = "unknown"
	buildTime = "unknown"
	goVersion = "unknown"
	osArch    = "unknown"
)

func main() {
	ctr := app.NewContainer()
	ctr.Versions = map[string]string{
		"version":   version,
		"gitCommit": gitCommit,
		"buildTime": buildTime,
		"goVersion": goVersion,
		"osArch":    osArch,
	}

	if err := cmd.NewRootCommand(ctr).Execute(); err != nil {
		log.Fatal(err)
	}
}
