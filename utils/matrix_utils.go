package utils

func IsOutbound[T any](input [][]T, i int, j int) bool {
	return i < 0 || j < 0 || i >= len(input) || j >= len(input[i])
}

func CopyMatrix[T any](input [][]T) [][]T {
	duplicate := make([][]T, len(input))
	for i := range input {
		duplicate[i] = make([]T, len(input[i]))
		copy(duplicate[i], input[i])
	}
	return duplicate
}
