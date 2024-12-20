package scan

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"

	sprig "github.com/go-task/slim-sprig"

	osx "github.com/mecha-hq/ekdo/internal/x/os"
)

var (
	ErrCannotCreateRenderer = errors.New("cannot create renderer")
	ErrUnknownToolName      = errors.New("unknown tool name")

	ErrCannotRender        = fmt.Errorf("cannot render scan template")
	ErrCannotLoadReport    = fmt.Errorf("cannot load report")
	ErrCannotLoadTemplate  = fmt.Errorf("cannot load template")
	ErrCannotPublishAssets = fmt.Errorf("cannot publish assets")
)

type Renderer interface {
	Render(drawLayout bool) error
	PublishAssets(path string) error
}

type RendererConstructor[T any] func(inputFile, outputDir string) (Renderer, error)

func NewRendererFactory() *RendererFactory {
	return &RendererFactory{}
}

type RendererFactory struct {
	ctrs map[string]RendererConstructor[any]
}

func (rf *RendererFactory) Register(name string, ctr RendererConstructor[any]) {
	if rf.ctrs == nil {
		rf.ctrs = make(map[string]RendererConstructor[any], 0)
	}

	rf.ctrs[name] = ctr
}

func (rf *RendererFactory) Create(toolName, inputFile, outputDir string) (Renderer, error) {
	ctr, ok := rf.ctrs[toolName]
	if !ok {
		return nil, fmt.Errorf("%w: %w '%s'", ErrCannotCreateRenderer, ErrUnknownToolName, toolName)
	}

	return ctr(inputFile, outputDir)
}

func NewDefaultRenderer[T any](name string, rl ReportLoader[T], w io.Writer, fs embed.FS) *DefaultRenderer[T] {
	return &DefaultRenderer[T]{
		n:  name,
		rl: rl,
		w:  w,
		fs: fs,
	}
}

type DefaultRenderer[T any] struct {
	n  string
	rl ReportLoader[T]
	w  io.Writer
	fs embed.FS
}

func (r *DefaultRenderer[T]) Name() string {
	return r.n
}

func (r *DefaultRenderer[T]) Render(drawLayout bool) error {
	report, err := r.rl.Load()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotRender, err)
	}

	tpl, err := r.loadTemplate()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotRender, err)
	}

	if err := tpl.Execute(r.w, report); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotRender, err)
	}

	return nil
}

func (r *DefaultRenderer[T]) PublishAssets(path string) error {
	if err := osx.EnsureDirExists(path); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotPublishAssets, err)
	}

	afs, err := fs.Sub(r.fs, "assets")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotPublishAssets, err)
	}

	if err := osx.CopyRecursive(afs, path); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotPublishAssets, err)
	}

	return nil
}

func (r *DefaultRenderer[T]) loadTemplate() (*template.Template, error) {
	data, err := r.fs.ReadFile("scan.html.tpl")
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotLoadTemplate, err)
	}

	t, err := template.New(r.Name()).Funcs(sprig.FuncMap()).Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotLoadTemplate, err)
	}

	return t, nil
}
