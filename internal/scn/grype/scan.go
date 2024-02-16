package grype

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/anchore/grype/grype/presenter/models"
	sprig "github.com/go-task/slim-sprig"

	iox "github.com/mecha-ci/ekdo/internal/x/io"
	osx "github.com/mecha-ci/ekdo/internal/x/os"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new grype scan renderer")
	ErrCannotRender             = fmt.Errorf("cannot render grype scan template")
	ErrCannotLoadReport         = fmt.Errorf("cannot load grype report")
	ErrCannotLoadTemplate       = fmt.Errorf("cannot load grype template")
	ErrCannotPublishAssets      = fmt.Errorf("cannot publish grype assets")
)

func NewScanRenderer(input, outputDir string) (*ScanRenderer, error) {
	r, err := iox.NewInputReader(input)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "grype-report.html"))
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

	t, err := template.New("grype").Funcs(sprig.FuncMap()).Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotLoadTemplate, err)
	}

	return t, nil
}

func loadReport(input io.Reader) (models.Document, error) {
	content, err := io.ReadAll(input)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w: %w", ErrCannotLoadReport, err)
	}

	report := models.Document{}
	if err := json.Unmarshal(content, &report); err != nil {
		return models.Document{}, fmt.Errorf("%w: %w", ErrCannotLoadReport, err)
	}

	return report, nil
}
