package kafkax

import "context"

type MockClient struct {
	Event  Event
	Format string
}

func (c *MockClient) Consumer(ctx context.Context, opts ...Option) (<-chan Event, error) {
	messages := make(chan Event)
	go func() {
		messages <- c.Event
	}()
	return messages, nil
}

func (c *MockClient) Producer(key string, value interface{}, opts ...Option) error {
	return nil
}

var (
	_ IClient = &MockClient{}
)
