package utils

import (
	"fmt"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("cant get pwd\n")
	}

	filePath := fmt.Sprintf("%s/testdata/testFile.txt", pwd)
	ok, _ := GetOpenFile(filePath)
	assert.True(t, ok)

	filePath = fmt.Sprintf("%s/testdata/testFile2.txt", pwd)
	ok, _ = GetOpenFile(filePath)
	assert.False(t, ok)
}

func TestFilterByExcept(t *testing.T) {
	const (
		EXCEPT_1 = "a"
		EXCEPT_2 = "b"
		EXCEPT_3 = "c"
	)

	data := []string{"z", EXCEPT_1, "y", EXCEPT_2, EXCEPT_3, "yy", "zz", EXCEPT_1}
	result := FilterByExcept(data)

	for _, v := range result {
		if v == EXCEPT_1 || v == EXCEPT_2 || v == EXCEPT_3 {
			t.Errorf("result of filterByExcept should not contain element %q\n", v)
		}
	}
}

func TestContains(t *testing.T) {
	data := []string{"a", "b", "c"}
	r := Contains(data, "c")
	if r != true {
		t.Errorf("expected data contains %q\n but result is %t\n", "c", true)
	}

	r = Contains(data, "d")
	if r != false {
		t.Errorf("expected data not contains %q\n but result is %t\n", "d", false)
	}
}

func TestSwapKeys(t *testing.T) {
	data := map[string]string{
		"key":  "val",
		"key2": "val2",
		"key3": "val3",
		"key5": "val4",
	}

	result := SwapKeys(data)
	for i, v := range result {
		if data[v] != i {
			t.Errorf("expected key of result must be %q\n got %q\n", data[v], i)
		}
	}
}

func TestConvertToMap(t *testing.T) {
	data := []entity.Roadmap{
		*entity.NewRoadMap(1, 1, "link"),
		*entity.NewRoadMap(2, 2, "link2"),
		*entity.NewRoadMap(3, 3, "link3"),
		*entity.NewRoadMap(4, 4, "link4"),
	}
	result, err := ConvertToMap(data)
	if err != nil {
		t.Errorf("error while convert to map, err: %s", err)
	}

	for i, v := range result {
		if i != v.Id {
			t.Errorf("index and id after convert must be identical\n")
		}
	}

	wrongData := []entity.Roadmap{
		*entity.NewRoadMap(1, 1, "link"),
		*entity.NewRoadMap(1, 2, "link2"),
	}

	result, err = ConvertToMap(wrongData)
	if err == nil {
		t.Fatal("when entity has duplicate id, must be error")
	}
	expectedError := "Duplicate Id 1"
	if expectedError != err.Error() {
		t.Errorf("error must be %q got %q", expectedError, err.Error())
	}
	if len(result) != 1 {
		t.Errorf("length of data when duplicate must expected 1")
	}
}

func TestSplitToBulks(t *testing.T) {
	data := map[int][]entity.Roadmap{
		1: {
			*entity.NewRoadMap(1, 1, "link"),
			*entity.NewRoadMap(2, 2, "link2"),
			*entity.NewRoadMap(3, 3, "link3"),
			*entity.NewRoadMap(4, 4, "link4"),
		},
		2: {
			*entity.NewRoadMap(1, 1, "link"),
			*entity.NewRoadMap(2, 2, "link2"),
			*entity.NewRoadMap(3, 3, "link3"),
			*entity.NewRoadMap(4, 4, "link4"),
		},
	}

	for i, v := range data {
		res := SplitToBulks(v, uint(i))
		countOfBatches := math.Ceil(float64(len(v)) / float64(i))
		if len(res) != int(countOfBatches) {
			t.Errorf("expected count of batcher is: %q\n got %q", int(countOfBatches), len(res))
		}
	}
}

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
