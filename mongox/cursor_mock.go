package mongox

import (
	"context"
)

type MockCursor struct {
	CursorAPI
	ErrorCursorAll error
}

func (c *MockCursor) Next(ctx context.Context) bool {
	return true
}

func (c *MockCursor) All(ctx context.Context, results interface{}) error {
	return c.ErrorCursorAll
}

func (c *MockCursor) Decode(v interface{}) error {
	return nil
}
