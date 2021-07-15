package utils

func SplitSliceToBatches(data []string, batchCount int) [][]string {
	var output = make([][]string, batchCount)

	lengthData := len(data)
	batchSize := (lengthData + batchCount - 1) / batchCount

	for i, j := 0, 0; i < lengthData; i, j = i+batchSize, j+1 {
		step := i + batchSize

		if step > lengthData {
			step = lengthData
		}

		output[j] = data[i:step]
	}

	return output
}
