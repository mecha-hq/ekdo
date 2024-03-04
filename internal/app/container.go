package app

import (
	"errors"
	"log/slog"

	"github.com/mecha-ci/ekdo/internal/scan"
	"github.com/mecha-ci/ekdo/internal/scan/grype"
	"github.com/mecha-ci/ekdo/internal/scan/snyk"
	"github.com/mecha-ci/ekdo/internal/scan/trivy"
)

var ErrCannotCreateContainer = errors.New("cannot create container")

type ContainerFactoryFunc func() (*Container, error)

func NewDefaultParameters() Parameters {
	return Parameters{}
}

type Parameters struct {
	Versions map[string]string
	LogLevel slog.Level
	Debug    bool
}

type services struct {
	scanRendererFactory *scan.RendererFactory
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

		c.ScanRendererFactory().Register("grype", grype.NewScanRenderer)
		c.ScanRendererFactory().Register("trivy", trivy.NewScanRenderer)
		c.ScanRendererFactory().Register("snyk", snyk.NewScanRenderer)
	}

	return c.scanRendererFactory
}
