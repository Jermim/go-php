package main

func Array_all[U any](arr []U, fn func(item U, index int) bool) bool {

	for i, item := range arr {
		if !fn(item, i) {
			return false
		}
	}

	return true
}

func Array_any[U any](arr []U, fn func(item U, index int) bool) bool {

	for i, item := range arr {
		if fn(item, i) {
			return true
		}
	}

	return false
}

func Array_chunk[T any](arr []T, length int) [][]T {
	var result [][]T

	end := len(arr)

	for i := 0; i < end; i += length {
		stop := min(i+length, end)
		result = append(result, arr[i:stop])
	}

	return result
}

func Array_map[U, T any](arr []U, fn func(item U) T) []T {

	elems := make([]T, len(arr))

	for i, item := range arr {
		elems[i] = fn(item)
	}

	return elems
}
