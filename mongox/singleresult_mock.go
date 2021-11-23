package mongox

type MockSingleResult struct {
	SingleResultAPI
	ErrorSingleResultDecode error
}

func (s *MockSingleResult) Decode(v interface{}) error {
	return s.ErrorSingleResultDecode
}
