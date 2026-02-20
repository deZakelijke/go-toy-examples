package sorting

import (
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nrCases := 10
	caseLength := 1000
	testcases := generateTestcases(nrCases, caseLength)

	for _, orig := range testcases {
		sortedOrig := BubbleSort(orig)

		if len(orig) != len(sortedOrig) {
			t.Errorf("Slice changed length during sorting. Before: %v, After: %v", orig, sortedOrig)
		}

		if !IsSorted(sortedOrig) {
			t.Errorf("Slice is not sorted correctly. Before: %v, After: %v", orig, sortedOrig)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	nrCases := 10
	caseLength := 1000
	testcases := generateTestcases(nrCases, caseLength)

	for _, orig := range testcases {
		sortedOrig := InsertionSort(orig)

		if len(orig) != len(sortedOrig) {
			t.Errorf("Slice changed length during sorting. Before: %v, After: %v", orig, sortedOrig)
		}

		if !IsSorted(sortedOrig) {
			t.Errorf("Slice is not sorted correctly. Before: %v, After: %v", orig, sortedOrig)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	nrCases := 10
	caseLength := 1000
	testcases := generateTestcases(nrCases, caseLength)

	for _, orig := range testcases {
		sortedOrig := SelectionSort(orig)

		if len(orig) != len(sortedOrig) {
			t.Errorf("Slice changed length during sorting. Before: %v, After: %v", orig, sortedOrig)
		}

		if !IsSorted(sortedOrig) {
			t.Errorf("Slice is not sorted correctly. Before: %v, After: %v", orig, sortedOrig)
		}
	}
}

func TestMergeSort(t *testing.T) {
	nrCases := 10
	caseLength := 1000
	testcases := generateTestcases(nrCases, caseLength)

	for _, orig := range testcases {
		sortedOrig := MergeSort(orig)

		if len(orig) != len(sortedOrig) {
			t.Errorf("Slice changed length during sorting. Before: %v, After: %v", orig, sortedOrig)
		}

		if !IsSorted(sortedOrig) {
			t.Errorf("Slice is not sorted correctly. Before: %v, After: %v", orig, sortedOrig)
		}
	}
}

func TestQuickSort(t *testing.T) {
	nrCases := 10
	caseLength := 100000
	testcases := generateTestcasesUInt(nrCases, caseLength)

	for _, orig := range testcases {
		sortedOrig := QuickSort(orig)

		if len(orig) != len(sortedOrig) {
			t.Errorf("Slice changed length during sorting. Before: %v, After: %v", orig, sortedOrig)
		}

		if !IsSorted(sortedOrig) {
			t.Errorf("Slice is not sorted correctly. Before: %v, After: %v", orig, sortedOrig)
		}
	}
}
func TestRadixSort(t *testing.T) {
	nrCases := 10
	caseLength := 100000
	testcases := generateTestcasesUInt(nrCases, caseLength)

	for _, orig := range testcases {
		sortedOrig := RadixSort(orig)

		if len(orig) != len(sortedOrig) {
			t.Errorf("Slice changed length during sorting. Before: %v, After: %v", orig, sortedOrig)
		}

		if !IsSorted(sortedOrig) {
			t.Errorf("Slice is not sorted correctly. Before: %v, After: %v", orig, sortedOrig)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	testcase := generateTestcasesUInt(1, 100000)[0]
	for b.Loop() {
		BubbleSort(testcase)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	testcase := generateTestcasesUInt(1, 100000)[0]
	for b.Loop() {
		MergeSort(testcase)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	testcase := generateTestcasesUInt(1, 100000)[0]
	for b.Loop() {
		QuickSort(testcase)
	}
}

func BenchmarkRadixSort(b *testing.B) {
	testcase := generateTestcasesUInt(1, 100000)[0]
	for b.Loop() {
		RadixSort(testcase)
	}
}

func generateTestcases(nrCases int, caseLength int) [][]float64 {
	var testcases [][]float64

	for i := 0; i < nrCases; i += 1 {
		var testcase []float64
		for j := 0; j < caseLength; j += 1 {
			testcase = append(testcase, rand.Float64()*1000_000-500_000)
		}
		testcases = append(testcases, testcase)
	}
	return testcases
}
func generateTestcasesUInt(nrCases int, caseLength int) [][]uint64 {
	var testcases [][]uint64

	for i := 0; i < nrCases; i += 1 {
		var testcase []uint64
		for j := 0; j < caseLength; j += 1 {
			testcase = append(testcase, uint64(rand.Float64()*1000_000-500_000))
		}
		testcases = append(testcases, testcase)
	}
	return testcases
}

func IsSorted[T Number](values []T) bool {
	for i := 1; i < len(values); i += 1 {
		if values[i] < values[i-1] {
			return false
		}
	}
	return true
}
