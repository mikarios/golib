package slices

// Contains checks whether a list of comparables contains a target value.
func Contains[T comparable](lst []T, elem T) bool {
	for i := range lst {
		if lst[i] == elem {
			return true
		}
	}

	return false
}

// SubtractSlice returns a slice of the elements that appear in slice a and not in slice b.
func SubtractSlice[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(a))

	for _, x := range b {
		mb[x] = struct{}{}
	}

	diff := make([]T, 0)

	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}
