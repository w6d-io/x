package mongox

import (
	"context"
)

// MockCursor is the internal mock cursor
type MockCursor struct {
	CursorAPI
	ErrorCursorAll error
}

// Next is an internal mock method
func (c *MockCursor) Next(ctx context.Context) bool {
	return true
}

// All is an internal mock method
func (c *MockCursor) All(ctx context.Context, results interface{}) error {
	return c.ErrorCursorAll
}

// Decode is an internal mock method
func (c *MockCursor) Decode(v interface{}) error {
	return nil
}
