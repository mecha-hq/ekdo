package io

import (
	"fmt"
	"io"
	"os"

	osx "github.com/mecha-hq/ekdo/internal/x/os"
)

var (
	ErrCannotGetOutputWriter = fmt.Errorf("cannot get output writer")
	ErrCannotGetInputReader  = fmt.Errorf("cannot get input reader")
)

func NewInputReader(input string) (io.Reader, error) {
	if input == "" {
		return nil, fmt.Errorf("%w: %s", ErrCannotGetInputReader, "input file cannot be empty")
	}

	if input == "-" {
		return os.Stdin, nil
	}

	r, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotGetInputReader, err)
	}

	return r, nil
}

func NewOutputWriter(output string) (io.Writer, error) {
	if output == "" {
		return nil, fmt.Errorf("%w: %s", ErrCannotGetOutputWriter, "output file cannot be empty")
	}

	if output == "-" {
		return os.Stdout, nil
	}

	if err := osx.EnsureDirExists(output); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotGetOutputWriter, err)
	}

	w, err := os.Create(output)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotGetOutputWriter, err)
	}

	return w, nil
}
