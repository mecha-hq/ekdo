package trivy

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/aquasecurity/trivy/pkg/types"

	"github.com/mecha-ci/ekdo/internal/scn"
	iox "github.com/mecha-ci/ekdo/internal/x/io"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new trivy scan renderer")
)

func NewScanRenderer(inputFile, outputDir string) (scn.Renderer, error) {
	r, err := iox.NewInputReader(inputFile)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "trivy-report.html"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	return scn.NewDefaultRenderer[types.Report]("trivy", r, w, emfs), nil
}
