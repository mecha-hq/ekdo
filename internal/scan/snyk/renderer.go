package snyk

import (
	"cmp"
	"embed"
	"fmt"
	"path/filepath"
	"slices"

	"github.com/mecha-hq/ekdo/internal/scan"
	iox "github.com/mecha-hq/ekdo/internal/x/io"
)

var (
	//go:embed scan.html.tpl assets/*
	emfs embed.FS

	ErrCannotCreateScanRenderer = fmt.Errorf("cannot create new snyk scan renderer")
)

type report interface {
	GetVulnerabilities() []*Vulnerability
	SetVulnerabilities(vulns []*Vulnerability)
}

type Report struct {
	Path            string           `json:"path"`
	Platform        string           `json:"platform"`
	ProjectName     string           `json:"projectName"`
	Vulnerabilities []*Vulnerability `json:"vulnerabilities"`
}

func (r *Report) GetVulnerabilities() []*Vulnerability {
	return r.Vulnerabilities
}

func (r *Report) SetVulnerabilities(vulns []*Vulnerability) {
	r.Vulnerabilities = vulns
}

type Vulnerability struct {
	ID         string `json:"id"`
	Identifier struct {
		CVE []string `json:"cve"`
	} `json:"identifiers"`
	References []struct {
		URL string `json:"url"`
	} `json:"references"`
	Severity    string `json:"severity"`
	PackageName string `json:"packageName"`
	Version     string `json:"version"`
	Title       string `json:"title"`
}

func NewScanRenderer(inputFile, outputDir string) (scan.Renderer, error) {
	r, err := iox.NewInputReader(inputFile)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	rl := NewDedupeLoader(scan.NewDefaultReportLoader[*Report](r))

	w, err := iox.NewOutputWriter(filepath.Join(outputDir, "snyk.html"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateScanRenderer, err)
	}

	return scan.NewDefaultRenderer("snyk", rl, w, emfs), nil
}

func NewDedupeLoader[T report](rl *scan.DefaultReportLoader[T]) *DedupeLoader[T] {
	return &DedupeLoader[T]{
		rl: rl,
	}
}

type DedupeLoader[T report] struct {
	rl *scan.DefaultReportLoader[T]
}

func (dl *DedupeLoader[T]) Load() (T, error) {
	var t T

	t, err := dl.rl.Load()
	if err != nil {
		return t, err
	}

	vulns := t.GetVulnerabilities()

	slices.SortFunc(vulns, func(a, b *Vulnerability) int {
		if a == nil && b == nil {
			return 0
		}

		if n := cmp.Compare(a.ID, b.ID); n != 0 {
			return n
		}

		return cmp.Compare(a.PackageName, b.PackageName)
	})

	vulns = slices.CompactFunc(vulns, func(a, b *Vulnerability) bool {
		if a == nil && b == nil {
			return true
		}

		return a.ID == b.ID && a.PackageName == b.PackageName
	})

	t.SetVulnerabilities(slices.Clip(vulns))

	return t, nil
}

var _ report = (*Report)(nil)
