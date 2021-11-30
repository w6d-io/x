package logx

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Field int

const (
	CorrelationId Field = iota + 1
	Kind
	Method
	IpAddress
	URI
)

var FieldString = map[string]Field{
	"correlation_id": CorrelationId,
	"kind":           Kind,
	"method":         Method,
	"ipaddress":      IpAddress,
	"uri":            URI,
}

//func (f Field) String() string {
//	//return FieldString[f]
//}

// GetLogValues get values from context and return a key/value interface
func GetLogValues(ctx context.Context) []interface{} {
	var values []interface{}
	if ctx == nil {
		return values
	}

	for text, key := range FieldString {
		if ctx.Value(key) != nil {
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
	if ctx == nil {
		return ""
	}
	value := ctx.Value(CorrelationId)
	if value == nil {
		return ""
	}
	return value.(string)
}
