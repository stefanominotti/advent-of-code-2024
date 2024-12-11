package utils

func IsOutbound[T any](input [][]T, i int, j int) bool {
	return i < 0 || j < 0 || i >= len(input) || j >= len(input[i])
}
