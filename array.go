package main

import "cmp"

func return_zero[T any]() T {
	var zero T
	return zero
}

func Array_all[U any](arr []U, fn func(item U, index int) bool) bool {

	for i, item := range arr {
		if !fn(item, i) {
			return false
		}
	}

	return true
}

func Array_Pop[T any](arr *[]T) T {
	length := len(*arr)

	if length == 0 {
		return return_zero[T]()
	}

	v := (*arr)[length-1]
	*arr = (*arr)[:length-2]

	return v
}

func Array_shift[T any](arr *[]T) T {
	length := len(*arr)

	if length == 0 {
		return return_zero[T]()
	}

	v := (*arr)[0]
	*arr = (*arr)[1:]

	return v
}

func Array_Push[T any](arr *[]T, elements ...T) {
	(*arr) = append((*arr), elements...)
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

func Array_fill[T any](start_index, count int, value T) []T {

	result := make([]T, count+start_index)

	for i := start_index; i < len(result); i++ {
		result[i] = value
	}

	return result
}

func Array_filter[T any](arr []T, fn func(T) bool) []T {
	var result []T

	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func Array_find[T any](arr []T, fn func(T) bool) T {
	for _, v := range arr {
		if fn(v) {
			return v
		}
	}

	return return_zero[T]()
}

func Array_flip[T any](arr []T) []T {

	result := make([]T, len(arr))

	for i := len(arr); i > 0; i-- {
		result[i] = arr[i]
	}

	return result
}

func Array_search[T cmp.Ordered](arr []T, search T) int {
	for i, item := range arr {
		if item == search {
			return i
		}
	}

	return -1
}

func Array_map[U, T any](arr []U, fn func(item U) T) []T {

	elems := make([]T, len(arr))

	for i, item := range arr {
		elems[i] = fn(item)
	}

	return elems
}
