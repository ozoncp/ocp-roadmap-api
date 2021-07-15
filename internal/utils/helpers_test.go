package utils

import (
	"strings"
	"testing"
)

func TestSplitSliceToBatches(t *testing.T) {
	expectedData := map[int]string{
		2: "Some.Long.String For.Testing",
		3: "Some.Long String.For Testing",
		5: "Some Long String For Testing",
	}
	original := "Some.Long.String.For.Testing"
	data := strings.Split(original, ".")

	for i, v := range expectedData {
		result := SplitSliceToBatches(data, i)
		if len(result) != i {
			t.Errorf("length of %q element must be %q", i, i)
		}
		var r []string
		for _, val := range result {
			r = append(r, strings.Join(val, "."))
		}

		resString := strings.Join(r, " ")
		if v != resString {
			t.Errorf("\nexpected element:\n %q \ngot\n %q", v, resString)
		}
	}
}
