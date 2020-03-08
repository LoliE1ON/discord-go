package SliceHelper

// Search through each element in a slice
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, result := set[item]
	return result
}
