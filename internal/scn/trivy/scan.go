package trivy

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/aquasecurity/trivy/pkg/types"
	sprig "github.com/go-task/slim-sprig"

	iox "github.com/mecha-ci/ekdo/internal/x/io"
	osx "github.com/mecha-ci/ekdo/internal/x/os"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new trivy scan renderer")
	ErrCannotRender             = fmt.Errorf("cannot render trivy scan template")
	ErrCannotLoadReport         = fmt.Errorf("cannot load trivy report")
	ErrCannotLoadTemplate       = fmt.Errorf("cannot load trivy template")
	ErrCannotPublishAssets      = fmt.Errorf("cannot publish trivy assets")
)

func NewScanRenderer(input, outputDir string) (*ScanRenderer, error) {
	r, err := iox.NewInputReader(input)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "trivy-report.html"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	return &ScanRenderer{
		r: r,
		w: w,
	}, nil
}

type ScanRenderer struct {
	r io.Reader
	w io.Writer
}

func (r *ScanRenderer) Render(drawLayout bool) error {
	report, err := loadReport(r.r)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotRender, err)
	}

	tpl, err := loadTemplate()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotRender, err)
	}

	if err := tpl.Execute(r.w, report); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotRender, err)
	}

	return nil
}

func (r *ScanRenderer) PublishAssets(path string) error {
	if err := osx.EnsureDirExists(path); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotPublishAssets, err)
	}

	afs, err := fs.Sub(emfs, "assets")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotPublishAssets, err)
	}

	if err := osx.CopyRecursive(afs, path); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotPublishAssets, err)
	}

	return nil
}

func loadTemplate() (*template.Template, error) {
	data, err := emfs.ReadFile("scan.html.tpl")
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotLoadTemplate, err)
	}

	t, err := template.New("trivy").Funcs(sprig.FuncMap()).Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotLoadTemplate, err)
	}

	return t, nil
}

func loadReport(input io.Reader) (types.Report, error) {
	content, err := io.ReadAll(input)
	if err != nil {
		return types.Report{}, fmt.Errorf("%w: %w", ErrCannotLoadReport, err)
	}

	report := types.Report{}
	if err := json.Unmarshal(content, &report); err != nil {
		return types.Report{}, fmt.Errorf("%w: %w", ErrCannotLoadReport, err)
	}

	return report, nil
}
