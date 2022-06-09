package pointers

func Ptr[T any](v T) *T {
	return &v
}

func BoolToInt(b *bool) *int {
	if b == nil {
		return nil
	}

	if *b {
		return Ptr[int](1)
	}

	return Ptr[int](0)
}
