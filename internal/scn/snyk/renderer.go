package snyk

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/mecha-ci/ekdo/internal/scn"
	iox "github.com/mecha-ci/ekdo/internal/x/io"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new snyk scan renderer")
)

// TODO: set the type of the report
func NewScanRenderer(inputFile, outputDir string) (scn.Renderer, error) {
	r, err := iox.NewInputReader(inputFile)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "snyk-report.html"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	return scn.NewDefaultRenderer[any]("snyk", r, w, emfs), nil
}
