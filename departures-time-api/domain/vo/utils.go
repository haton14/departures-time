package vo

func pointer[T any](v T) *T {
	return &v
}
