package utils

type Direction int

const (
	TopDirection    Direction = 1
	RightDirection  Direction = 2
	BottomDirection Direction = 3
	LeftDirection   Direction = 4
)

func GetDirections() [4]Direction {
	return [4]Direction{TopDirection, BottomDirection, RightDirection, LeftDirection}
}

func GetDirectionMoves(direction Direction) []int {
	switch direction {
	case TopDirection:
		return []int{-1, 0}
	case RightDirection:
		return []int{0, 1}
	case BottomDirection:
		return []int{1, 0}
	case LeftDirection:
		return []int{0, -1}
	}
	panic("Invalid direction")
}
