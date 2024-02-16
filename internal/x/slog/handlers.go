package slog

import (
	"fmt"
	"log/slog"
	"strings"
)

var ErrUnknownLevel = fmt.Errorf("unknown level")

type ReplaceAttrFn func(groups []string, a slog.Attr) slog.Attr

func FromString(level string) (slog.Level, error) {
	switch strings.ToUpper(strings.TrimSpace(level)) {
	case "DEBUG":
		return slog.LevelDebug, nil

	case "INFO":
		return slog.LevelInfo, nil

	case "WARN":
		return slog.LevelWarn, nil

	case "ERROR":
		return slog.LevelError, nil

	default:
		return slog.LevelInfo, fmt.Errorf("%w: %s", ErrUnknownLevel, level)
	}
}

// ReplaceAttrs returns a function that applies all the given functions to the attribute.
// It is useful to compose multiple functions that replace attributes, allowing to combine them on a per-need basis.
func ReplaceAttrs(fns ...ReplaceAttrFn) ReplaceAttrFn {
	return func(groups []string, a slog.Attr) slog.Attr {
		for _, fn := range fns {
			a = fn(groups, a)
		}

		return a
	}
}

// NoTimeReplaceAttr is a ReplaceAttrFn that removes the time attribute from the log, if present.
// This is useful in tests to increase predictability of the output.
func NoTimeReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}

	return a
}
