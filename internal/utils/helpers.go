package utils

import (
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"log"
	"os"
)

func FilterByExcept(data []string) []string {
	var result []string
	exceptElements := []string{"a", "b", "c"}

	for _, v := range data {
		if Contains(exceptElements, v) {
			continue
		}
		result = append(result, v)
	}

	return result
}

func Contains(data []string, needle string) bool {
	for _, v := range data {
		if needle == v {
			return true
		}
	}
	return false
}

func SwapKeys(data map[string]string) map[string]string {
	result := make(map[string]string, len(data))
	for i, v := range data {
		result[v] = i
	}

	return result
}

func SplitSliceToBatches(data []string, batchSize int) [][]string {
	var chunks [][]string
	for {
		if len(data) == 0 {
			break
		}

		if len(data) < batchSize {
			batchSize = len(data)
		}

		chunks = append(chunks, data[0:batchSize])
		data = data[batchSize:]
	}

	return chunks
}

func GetOpenFile(filePath string) (bool, error) {

	fileOpen := func(path string) error {
		file, err := os.Open(path)
		if err != nil {
			return errors.New(fmt.Sprintf("error while open file, err: %v", err))
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.Fatalf("error while close file %q, err: %v", path, err)
			}
		}()

		return nil
	}

	if fileErr := fileOpen(filePath); fileErr != nil {
		return false, fileErr
	}

	return true, nil
}

func SplitToBulks(entities []entity.Roadmap, butchSize uint) [][]entity.Roadmap {
	var chunks [][]entity.Roadmap

	for {
		if len(entities) == 0 {
			break
		}

		if uint(len(entities)) < butchSize {
			butchSize = uint(len(entities))
		}

		chunks = append(chunks, entities[0:butchSize])
		entities = entities[butchSize:]
	}

	return chunks
}

func ConvertToMap(entities []entity.Roadmap) (map[uint64]entity.Roadmap, error) {
	output := make(map[uint64]entity.Roadmap)
	for _, v := range entities {
		if _, ok := output[v.Id]; ok {
			return output, errors.New(fmt.Sprintf("Duplicate Id %d", v.Id))
		}
		output[v.Id] = v
	}

	return output, nil
}
