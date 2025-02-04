package helpers

import "strings"

func Filter[T any](items []T, shouldInclude func(int, T) bool) []T {
	filtered := make([]T, 0, len(items))

	for i, item := range items {
		if shouldInclude(i, item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func JoinWithAnd(fields []string) string {
	switch len(fields) {
	case 0:
		return ""
	case 1:
		return fields[0]
	case 2:
		return fields[0] + " e " + fields[1]
	default:
		return strings.Join(fields[:len(fields)-1], ", ") + " e " + fields[len(fields)-1]
	}
}
