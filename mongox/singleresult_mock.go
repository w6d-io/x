package mongox

// MockSingleResult is the internal mock single result
type MockSingleResult struct {
	SingleResultAPI
	ErrorSingleResultDecode error
}

// Decode is an internal mock method
func (s *MockSingleResult) Decode(v interface{}) error {
	return s.ErrorSingleResultDecode
}
