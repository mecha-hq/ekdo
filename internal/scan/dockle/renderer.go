package dockle

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/goodwithtech/dockle/pkg/report"
	"github.com/mecha-ci/ekdo/internal/scan"
	iox "github.com/mecha-ci/ekdo/internal/x/io"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new dockle scan renderer")
)

func NewScanRenderer(inputFile, outputDir string) (scan.Renderer, error) {
	r, err := iox.NewInputReader(inputFile)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "dockle-report.html"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	return scan.NewDefaultRenderer[report.JsonOutputFormat]("dockle", r, w, emfs), nil
}
