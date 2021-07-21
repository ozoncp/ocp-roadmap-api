package utils

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
