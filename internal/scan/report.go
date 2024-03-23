package scan

import (
	"encoding/json"
	"fmt"
	"io"
)

type ReportLoader[T any] interface {
	Load() (T, error)
}

func NewDefaultReportLoader[T any](r io.Reader) *DefaultReportLoader[T] {
	return &DefaultReportLoader[T]{
		r: r,
	}
}

type DefaultReportLoader[T any] struct {
	r io.Reader
}

func (l *DefaultReportLoader[T]) Load() (T, error) {
	var report T

	content, err := io.ReadAll(l.r)
	if err != nil {
		return report, fmt.Errorf("%w: %w", ErrCannotLoadReport, err)
	}

	if err := json.Unmarshal(content, &report); err != nil {
		return report, fmt.Errorf("%w: %w", ErrCannotLoadReport, err)
	}

	return report, nil
}
