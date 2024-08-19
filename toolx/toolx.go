package toolx

import (
	"os"
)

func InArray[T comparable](val T, array []T) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}

func Contains[T comparable](s []T, val T) bool {
	for _, item := range s {
		if item == val {
			return true
		}
	}
	return false
}

func KeysMap[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Getenv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return value
}
