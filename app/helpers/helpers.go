package helpers

func FindMissingIDs[T any](requestedIDs []string, foundItems []T, getID func(T) string) []string {
	idMap := make(map[string]struct{})

	for _, item := range foundItems {
		id := getID(item)
		idMap[id] = struct{}{}
	}

	var missingIDs []string
	for _, id := range requestedIDs {
		if _, exists := idMap[id]; !exists {
			missingIDs = append(missingIDs, id)
		}
	}

	return missingIDs
}
