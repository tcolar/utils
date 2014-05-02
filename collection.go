// History: May 02 14 tcolar Creation

package utils

// HasString returns wether a slice contains the given string
func HasString(slice []string, value string) bool {
	return FindString(slice, value) != -1
}

// FindString returns the index of value in the slice
// Or -1 if not found
func FindString(slice []string, value string) int {
	for i, s := range slice {
		if s == value {
			return i
		}
	}
	return -1
}

// HasInt returns wether a slice contains the given int
func HasInt(slice []int, value int) bool {
	return FindInt(slice, value) != -1
}

// FindString returns the index of value in the slice
// Or -1 if not found
func FindInt(slice []int, value int) int {
	for i, nb := range slice {
		if nb == value {
			return i
		}
	}
	return -1
}
