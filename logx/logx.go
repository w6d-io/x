package logx

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Field is key for logs in context
type Field int

const (
	// CorrelationID correlation ID key
	CorrelationID Field = iota + 1
	// Kind kind key
	Kind
	// Method http method key
	Method
	// IPAddress ip address key
	IPAddress
	// URI http uri
	URI
)

// FieldString match string and Field
var FieldString = map[string]Field{
	"correlation_id": CorrelationID,
	"kind":           Kind,
	"method":         Method,
	"ipaddress":      IPAddress,
	"uri":            URI,
}

//func (f Field) String() string {
//	//return FieldString[f]
//}

// GetLogValues get values from context and return a key/value interface
func GetLogValues(ctx context.Context) []interface{} {
	var values []interface{}

	for text, key := range FieldString {
		if ctx != nil && ctx.Value(key) != nil {
			values = append(values, text)
			values = append(values, ctx.Value(key))
		}
	}
	return values
}

// WithName provides a new logger with the name appended and values from context
func WithName(ctx context.Context, name string) logr.Logger {
	return ctrl.Log.WithName(name).WithValues(GetLogValues(ctx)...)
}

// GetCorrelationID get the correlation id from the context or return an empty string
func GetCorrelationID(ctx context.Context) string {
	if ctx != nil && ctx.Value(CorrelationID) != nil {
		return ctx.Value(CorrelationID).(string)
	}
	return ""
}
