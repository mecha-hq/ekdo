package os

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

const (
	Perm600 = 0o600
	Perm644 = 0o644
	Perm750 = 0o750
	Perm755 = 0o755
)

var (
	ErrCannotCopyFile            = errors.New("cannot copy file")
	ErrCannotRecurivelyCopyFiles = errors.New("cannot recursively copy files")
	ErrNotRegularFile            = errors.New("is not a regular file")
	ErrCannotEnsureDirExists     = errors.New("cannot ensure directory exists")
)

func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotCopyFile, err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%w: %s %w", ErrCannotCopyFile, src, ErrNotRegularFile)
	}

	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotCopyFile, err)
	}

	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotCopyFile, err)
	}

	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotCopyFile, err)
	}

	return nil
}

func CopyRecursive(src fs.FS, dest string) error {
	stuff, err := fs.ReadDir(src, ".")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotRecurivelyCopyFiles, err)
	}

	for _, file := range stuff {
		if file.IsDir() {
			sub, err := fs.Sub(src, file.Name())
			if err != nil {
				return fmt.Errorf("%w: %w", ErrCannotRecurivelyCopyFiles, err)
			}

			if err := os.Mkdir(path.Join(dest, file.Name()), Perm755); err != nil && !os.IsExist(err) {
				return fmt.Errorf("%w: %w", ErrCannotRecurivelyCopyFiles, err)
			}

			if err := CopyRecursive(sub, path.Join(dest, file.Name())); err != nil {
				return err
			}

			continue
		}

		fileContent, err := fs.ReadFile(src, file.Name())
		if err != nil {
			return fmt.Errorf("%w: %w", ErrCannotRecurivelyCopyFiles, err)
		}

		if err := EnsureDirExists(path.Join(dest, file.Name())); err != nil {
			return fmt.Errorf("%w: %w", ErrCannotRecurivelyCopyFiles, err)
		}

		if err := os.WriteFile(path.Join(dest, file.Name()), fileContent, Perm644); err != nil {
			return fmt.Errorf("%w: %w", ErrCannotRecurivelyCopyFiles, err)
		}
	}

	return nil
}

func EnsureDirExists(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		if !os.IsNotExist(serr) {
			return fmt.Errorf("%w: %w", ErrCannotEnsureDirExists, serr)
		}

		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			return fmt.Errorf("%w: %w", ErrCannotEnsureDirExists, err)
		}
	}

	return nil
}
