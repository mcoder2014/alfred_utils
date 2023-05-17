package utils

func GPtr[T any](t T) *T {
	return &t
}
