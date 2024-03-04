package app

import (
	"errors"
	"log/slog"

	"github.com/mecha-ci/ekdo/internal/scn"
	"github.com/mecha-ci/ekdo/internal/scn/grype"
	"github.com/mecha-ci/ekdo/internal/scn/snyk"
	"github.com/mecha-ci/ekdo/internal/scn/trivy"
)

var ErrCannotCreateContainer = errors.New("cannot create container")

type ContainerFactoryFunc func() (*Container, error)

func NewDefaultParameters() Parameters {
	return Parameters{}
}

type Parameters struct {
	Versions  map[string]string
	LogLevel  slog.Level
	Debug     bool
	InputFile string
	OutputDir string
}

type services struct {
	scanRendererFactory *scn.RendererFactory
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

func (c *Container) ScanRendererFactory() *scn.RendererFactory {
	if c.scanRendererFactory == nil {
		c.scanRendererFactory = scn.NewRendererFactory()

		c.ScanRendererFactory().Register("grype", grype.NewScanRenderer)
		c.ScanRendererFactory().Register("trivy", trivy.NewScanRenderer)
		c.ScanRendererFactory().Register("snyk", snyk.NewScanRenderer)
	}

	return c.scanRendererFactory
}
