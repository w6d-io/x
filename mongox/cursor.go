package mongox

import "context"

func (c *Cursor) Next(ctx context.Context) bool {
	return c.cursor.Next(ctx)
}

func (c *Cursor) All(ctx context.Context, results interface{}) error {
	return c.cursor.All(ctx, results)
}

func (c *Cursor) Decode(val interface{}) error {
	return c.cursor.Decode(val)
}
