package grype

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/anchore/grype/grype/presenter/models"
	"github.com/mecha-ci/ekdo/internal/scn"
	iox "github.com/mecha-ci/ekdo/internal/x/io"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new grype scan renderer")
)

func NewScanRenderer(inputFile, outputDir string) (scn.Renderer, error) {
	r, err := iox.NewInputReader(inputFile)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "grype-report.html"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	return scn.NewDefaultRenderer[models.Document]("grype", r, w, emfs), nil
}
