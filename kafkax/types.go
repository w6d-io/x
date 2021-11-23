package kafkax

import (
	"time"
)

// Header Payload
type Header struct {
	Key   string // Header name (utf-8 string)
	Value []byte // Header value (nil, empty, or binary)
}

// Client Event
type Event struct {
	Key, Value []byte
	Topic      string
	Partition  int32
	Offset     int64
	Headers    []Header
	Timestamp  time.Time
}
