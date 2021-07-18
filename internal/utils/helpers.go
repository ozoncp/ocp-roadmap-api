package utils

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
