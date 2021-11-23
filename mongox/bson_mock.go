package mongox

import "go.mongodb.org/mongo-driver/bson/bsonrw"

type MockbsonWriter struct {
	bsonrw.ValueWriter
}

func (b *MockbsonWriter) WriteDateTime(dt int64) error {
	return nil
}

type MockbsonRead struct {
	bsonrw.ValueReader
	ErrReadDateTime error
}

func (b *MockbsonRead) ReadDateTime() (int64, error) {
	return 10, b.ErrReadDateTime
}
