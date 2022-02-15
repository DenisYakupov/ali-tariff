package utils

func GetChunkSlices(slice []int, sizeChunk int) [][]int {

	result := make([][]int, 0)
	key := 0
	for i := 0; i < len(slice); i++ {

		if i%sizeChunk == 0 {
			if len(slice) < key+sizeChunk {
				result = append(
					result,
					slice[key:],
				)
			} else {
				result = append(
					result,
					slice[key:key+sizeChunk],
				)
			}
		}
		key++
	}

	return result
}

func ConverterKeyValue(m map[string]int) map[int]string {
	result := make(map[int]string)

	for key, value := range m {
		result[value] = key
	}
	return result
}
