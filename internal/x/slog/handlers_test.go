package slog_test

import (
	"bytes"
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	slogx "github.com/mecha-hq/ekdo/internal/x/slog"
)

func TestStacktraceReplaceAttr(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		message string
		params  []any
		want    string
	}{
		{
			desc:    "log has no error attribute",
			message: "foo bar",
			want:    "level=ERROR msg=\"foo bar\"\n",
		},
		{
			desc:    "log has an error attribute",
			message: "foo bar",
			params:  []any{"error", fmt.Errorf("quux")},
			want:    "level=ERROR msg=\"foo bar\" error=quux\n",
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			logger, bbuf := newLogger()

			logger.Error(tC.message, tC.params...)

			assert.Equal(t, bbuf.String(), tC.want)
		})
	}
}

func newLogger() (*slog.Logger, *bytes.Buffer) {
	bbuf := bytes.NewBuffer(make([]byte, 0))

	logger := slog.New(
		slog.NewTextHandler(
			bbuf,
			&slog.HandlerOptions{
				ReplaceAttr: slogx.NoTimeReplaceAttr,
			},
		),
	)

	return logger, bbuf
}
