package core_utils

func Contains[T comparable](array []T, element T) bool {
	for _, a := range array {
		if a == element {
			return true
		}
	}
	return false
}
