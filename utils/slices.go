package utils

func SlicesMap[T any, U any](s []T, f func(T) U) []U {
	res := make([]U, 0, len(s))
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

func Deduplicate[T comparable](s []T) []T {
	keys := make(map[T]struct{})
	list := make([]T, 0, len(s))

	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = struct{}{}
			list = append(list, entry)
		}
	}

	return list
}
