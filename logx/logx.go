package logx

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

// GetLogValues get values from context and return a key/value interface
func GetLogValues(ctx context.Context) []interface{} {
	var values []interface{}
	keys := []string{
		"correlation_id",
		"kind",
		"ipaddress",
	}
	for _, key := range keys {
		if ctx.Value(key) != nil {
			values = append(values, key)
			values = append(values, ctx.Value(key))
		}
	}
	return values
}

// WithName provides a new logger with the name appended and values from context
func WithName(ctx context.Context, name string) logr.Logger {
	return ctrl.Log.WithName(name).WithValues(GetLogValues(ctx)...)
}
