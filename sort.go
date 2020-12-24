package main

func BubbleSort(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func SelectionSort(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		var temp int = 0
		for j := 1; j <= i; j++ {
			if a[j] > a[temp] {
				temp = j
			}
		}
		a[temp], a[i] = a[i], a[temp]
	}
	return a
}

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		newElement := a[i]
		var j int
		for j = i; j > 0 && a[i-1] > newElement; j-- {
			a[i] = a[i-1]
		}
		a[j] = newElement
	}
	return a
}

func MergeSort(input []int) []int {
	return split(input)
}

func split(input []int) []int {
	if len(input) < 2 {
		return input
	}
	mid := len(input) / 2
	return merge(split(input[:mid]), split(input[mid:]))
}

func merge(A, B []int) []int {
	arr := make([]int, len(A)+len(B))
	size := len(A) + len(B)
	j, k := 0, 0
	for i := 0; i < size; i++ {
		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		} else if A[j] > B[k] {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}
	return arr
}
