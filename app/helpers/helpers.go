package helpers

func Filter[T any](items []T, shouldInclude func(int, T) bool) []T {
	filtered := make([]T, 0, len(items))

	for i, item := range items {
		if shouldInclude(i, item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
