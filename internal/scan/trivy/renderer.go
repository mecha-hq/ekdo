package trivy

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/aquasecurity/trivy/pkg/types"

	"github.com/mecha-hq/ekdo/internal/scan"
	iox "github.com/mecha-hq/ekdo/internal/x/io"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new trivy scan renderer")
)

type Report = types.Report

func NewScanRenderer(inputFile, outputDir string) (scan.Renderer, error) {
	r, err := iox.NewInputReader(inputFile)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	rl := scan.NewDefaultReportLoader[Report](r)

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "trivy-report.html"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	return scan.NewDefaultRenderer("trivy", rl, w, emfs), nil
}
