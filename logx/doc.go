// Package logx provides tools to handle log
//
// WithName provides a new logr.Logger:
//
// ctx := context.WithValue(c.Request.Context(), "correlation_id", correlationID)
// ctx = context.WithValue(ctx, "ipaddress", ip)
// ctx = context.WithValue(ctx, "kind", "http")
// log := logx.WithName(ctx, "Test")
//
// GetLogValues returns a keysAndValues model compatible with logrLogger.
//
// var log logr.Logger
// ctx := context.WithValue(c.Request.Context(), "correlation_id", correlationID)
// ctx = context.WithValue(ctx, "ipaddress", ip)
// ctx = context.WithValue(ctx, "kind", "http")
// kv := logx.GetLogValues(ctx)
// log = zapr.NewLogger(zapLog)
// log.Info("test", kv...)
//
package logx
