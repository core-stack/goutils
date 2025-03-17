package sliceutils

func Find[T any](slice []*T, predicate func(T) bool) (*T, bool) {
	for _, item := range slice {
		if predicate(*item) {
			return item, true
		}
	}
	var zero T
	return &zero, false
}

func FindIndex[T any](slice []T, predicate func(T) bool) int {
	for i, item := range slice {
		if predicate(item) {
			return i
		}
	}
	return -1
}
func Distinct[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	var list []T
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Map[T any, U any](slice []T, f func(T) U) []U {
	var list []U
	for _, item := range slice {
		list = append(list, f(item))
	}
	return list
}
