package logx

import "context"

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
