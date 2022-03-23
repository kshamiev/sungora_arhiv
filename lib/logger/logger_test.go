package logger

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	Init(&Config{
		Title:     "TestTitle",
		Output:    "stdout",
		Formatter: "json",
		Level:     logrus.TraceLevel,
		IsCaller:  false,
		Hooks:     Hooks{},
	})

	ctx := context.Background()
	lg := Get(ctx)
	ctx = WithLogger(ctx, lg.
		WithField("sql", "query").
		WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   10,
		}),
	)

	lg.Trace("Trace message")
	lg.Debug("Debug message")
	lg.Info("Info message")
	lg.Warning("Warning message")
	lg.Error("Error message")
	lg.WithError(errors.New("APP ERROR")).Error("Error message")

	fmt.Println()

	lgCtx := Get(ctx)
	lgCtx.Trace("Trace message")
	lgCtx.Debug("Debug message")
	lgCtx.Info("Info message")
	lgCtx.Warning("Warning message")
	lgCtx.Error("Error message")
	lgCtx.WithError(errors.New("APP ERROR")).Error("Error message")
}
