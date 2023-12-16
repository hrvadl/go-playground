package main

func Append[T any](slice []T, elements ...T) []T {
	if slice == nil {
		slice = make([]T, 0, len(elements))
	}

	var next int
	if len(slice) == 0 {
		next = 0
	} else {
		next = len(slice)
	}

	if cap(slice)-len(slice)-len(elements) > 0 {
		slice = slice[:len(slice)+len(elements)]
		for i := 0; i < len(elements); i++ {
			slice[next+i] = elements[i]
		}
		return slice
	}

	newLen := len(elements) + len(slice)
	newSlice := make([]T, newLen, cap(slice)*2)
	copy(newSlice, slice)

	for i, j := next, 0; i < newLen; i, j = i+1, j+1 {
		newSlice[i] = elements[j]
	}

	return newSlice
}
