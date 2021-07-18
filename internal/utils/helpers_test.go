package utils

import (
	"math"
	"testing"
)

func TestSplitSliceToBatches(t *testing.T) {
	data := map[int][]string{
		2:  {"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"},
		3:  {"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"},
		5:  {"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"},
		8:  {"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"},
		10: {"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"},
	}

	for i, v := range data {
		countOfBatches := math.Ceil(float64(len(v)) / float64(i))
		result := SplitSliceToBatches(v, i)
		if len(result) != int(countOfBatches) {
			t.Errorf("expected count of batcher is: %q\n got %q", int(countOfBatches), len(result))
		}

		lastElementIndex := len(result) - 1
		for j, r := range result {
			if len(r) != i && j != lastElementIndex {
				t.Errorf("expected batch size is %d\n got %d\n in %q", i, len(r), result)
			}
		}
	}
}
