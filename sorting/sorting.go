package sorting

type Number interface {
	~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~uint16 | ~uint32 | ~uint64
}
type Sortable[T Number] interface {
	~[]T
}

func BubbleSort[S Sortable[N], N Number](values S) S {
	for i := len(values) - 1; i > 0; i -= 1 {
		for j := 0; j < i; j += 1 {
			if values[j] > values[i] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}

func SelectionSort[S Sortable[N], N Number](values S) S {
	for i := 0; i < len(values); i += 1 {
		smallest := i
		for j := i + 1; j < len(values); j += 1 {
			if values[smallest] > values[j] {
				smallest = j
			}
		}
		values[i], values[smallest] = values[smallest], values[i]
	}

	return values
}

func InsertionSort[S Sortable[N], N Number](values S) S {
	for i := 1; i < len(values); i += 1 {
		for j := i; j > 0; j -= 1 {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
			} else {
				break
			}
		}
	}
	return values
}

func MergeSort[S Sortable[N], N Number](values S) S {
	if len(values) <= 1 {
		return values
	}

	valuesFirstHalf := MergeSort(values[:len(values)/2])
	valuesSecondHalf := MergeSort(values[len(values)/2:])

	var newValues S
	var val N

	for len(valuesFirstHalf) > 0 && len(valuesSecondHalf) > 0 {
		if valuesFirstHalf[0] > valuesSecondHalf[0] {
			val, valuesSecondHalf = valuesSecondHalf[0], valuesSecondHalf[1:]
		} else {
			val, valuesFirstHalf = valuesFirstHalf[0], valuesFirstHalf[1:]
		}
		newValues = append(newValues, val)
	}

	if len(valuesFirstHalf) > 0 {
		newValues = append(newValues, valuesFirstHalf...)
	}
	if len(valuesSecondHalf) > 0 {
		newValues = append(newValues, valuesSecondHalf...)
	}

	return newValues
}

func QuickSort[S Sortable[N], N Number](values S) S {
	var pivot N
	if len(values) < 2 {
		return values
	}

	if len(values) > 5 {
		pivot_1 := values[0]
		pivot_2 := values[len(values)/2]
		pivot_3 := values[len(values)-1]
		if (pivot_1 > pivot_2 && pivot_1 < pivot_3) ||
			(pivot_1 < pivot_2 && pivot_1 > pivot_3) {
			pivot = pivot_1
		} else if (pivot_2 > pivot_1 && pivot_2 < pivot_3) ||
			(pivot_2 < pivot_1 && pivot_2 > pivot_3) {
			pivot = pivot_2
		} else {
			pivot = pivot_3
		}

	} else {
		pivot = values[0]
	}

	var firstSlice, pivotSlice, secondSlice S
	for i := 0; i < len(values); i += 1 {
		if values[i] < pivot {
			firstSlice = append(firstSlice, values[i])
		} else if values[i] > pivot {
			secondSlice = append(secondSlice, values[i])
		} else {
			pivotSlice = append(pivotSlice, values[i])
		}
	}
	firstSlice = QuickSort(firstSlice)
	secondSlice = QuickSort(secondSlice)
	values = append(firstSlice, pivotSlice...)
	values = append(values, secondSlice...)
	return values
}

func RadixSort(values []uint64) []uint64 {
	radix := 64
	for i := 0; i < radix; i += 1 {
		var firstHalf, secondHalf []uint64
		for _, val := range values {
			if val>>i&1 == 1 {
				secondHalf = append(secondHalf, val)
			} else {
				firstHalf = append(firstHalf, val)
			}
		}
		values = append(firstHalf, secondHalf...)
	}

	return values
}
