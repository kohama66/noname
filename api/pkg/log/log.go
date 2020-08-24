package log

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/myapp/noname/api/pkg/log/spancontext"
)

func Debugf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "DEBUG", format, a...)
}

func Infof(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "INFO", format, a...)
}

func Warningf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "WARNING", format, a...)
}

func Errorf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "ERROR", format, a...)
}

func Criticalf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "CRITICAL", format, a...)
}

type jsonPayload struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Trace    string `json:"logging.googleapis.com/trace"`
	SpanID   string `json:"logging.googleapis.com/spanId"`
}

func out(ctx context.Context, severity, format string, a ...interface{}) {
	sc := spancontext.Get(ctx)
	payload := &jsonPayload{
		Severity: severity,
		Message:  fmt.Sprintf(format, a...),
		Trace:    fmt.Sprintf("projects/%s/traces/%s", "dev", sc.TraceID),
		SpanID:   sc.SpanID,
	}
	json.NewEncoder(os.Stdout).Encode(payload)
}
