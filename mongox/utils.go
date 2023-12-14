package mongox

// GetLogLevel returns log level from data value
func GetLogLevel(data interface{}) int {
	if data == nil {
		return 3
	}
	return 2
}
