package mongox

// GetLogLevel returns log level from data value
func GetLogLevel(data interface{}) int {
	if data == nil {
		return 2
	}
	return 1
}
