package mongox

func (s *SingleResult) Decode(val interface{}) error {
	return s.singleResult.Decode(val)
}
