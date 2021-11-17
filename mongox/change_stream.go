package mongox

import "context"

func (s *ChangeStream) Next(ctx context.Context) bool {
	return s.changeStream.Next(ctx)
}

func (s *ChangeStream) Current() interface{} {
	return s.changeStream.Current
}
