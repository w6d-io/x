package mongox

import "go.mongodb.org/mongo-driver/bson/bsonrw"

// MockbsonWriter structure used by codec
type MockbsonWriter struct {
	bsonrw.ValueWriter
}

// WriteDateTime API
func (b *MockbsonWriter) WriteDateTime(dt int64) error {
	return nil
}

// WriteDateTime API
func (b *MockbsonWriter) WriteString(string) error {
	return nil
}

// MockbsonReader structure used by codec
type MockbsonReader struct {
	bsonrw.ValueReader
	ErrReadDateTime error
}

// ReadDateTime API
func (b *MockbsonReader) ReadDateTime() (int64, error) {
	return 10, b.ErrReadDateTime
}
