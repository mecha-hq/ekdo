package app

import (
	"errors"
	"log/slog"
	"os"

	"github.com/mecha-hq/ekdo/internal/scan"
	"github.com/mecha-hq/ekdo/internal/scan/dockle"
	"github.com/mecha-hq/ekdo/internal/scan/grype"
	"github.com/mecha-hq/ekdo/internal/scan/snyk"
	"github.com/mecha-hq/ekdo/internal/scan/trivy"
)

var ErrCannotCreateContainer = errors.New("cannot create container")

type ContainerFactoryFunc func() (*Container, error)

func NewDefaultParameters() Parameters {
	return Parameters{
		LogLevel: slog.LevelInfo,
	}
}

type Parameters struct {
	Versions map[string]string
	LogLevel slog.Level
}

type services struct {
	scanRendererFactory *scan.RendererFactory
	logger              *slog.Logger
}

func NewContainer() *Container {
	return &Container{
		Parameters: NewDefaultParameters(),
	}
}

type Container struct {
	Parameters
	services
}

func (c *Container) ScanRendererFactory() *scan.RendererFactory {
	if c.scanRendererFactory == nil {
		c.scanRendererFactory = scan.NewRendererFactory()

		c.ScanRendererFactory().Register("dockle", dockle.NewScanRenderer)
		c.ScanRendererFactory().Register("grype", grype.NewScanRenderer)
		c.ScanRendererFactory().Register("trivy", trivy.NewScanRenderer)
		c.ScanRendererFactory().Register("snyk", snyk.NewScanRenderer)
	}

	return c.scanRendererFactory
}

func (c *Container) Logger() *slog.Logger {
	if c.logger == nil {
		c.logger = slog.New(
			slog.NewTextHandler(
				os.Stderr,
				&slog.HandlerOptions{
					AddSource: true,
					Level:     c.LogLevel,
				},
			),
		)
	}

	return c.logger
}
