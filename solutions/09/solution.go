package solution09

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"strings"
)

type Solution09 struct{}

func (s Solution09) PartA(lineIterator *utils.LineIterator) any {
	return runSolution(lineIterator, buildFilesystemPartA)
}

func (s Solution09) PartB(lineIterator *utils.LineIterator) any {
	return runSolution(lineIterator, buildFilesystemPartB)
}

func runSolution(lineIterator *utils.LineIterator, buildFilesystem func([]int) []int) int {
	lineIterator.Next()
	line := lineIterator.Value()
	blocks := utils.StringsToIntegers(strings.Split(line, ""))

	filesystem := buildFilesystem(blocks)

	result := 0
	for i, num := range filesystem {
		if num == -1 {
			continue
		}
		result += i * num
	}
	return result
}

func buildFilesystemPartA(blocks []int) []int {
	filesystem := []int{}
	for i, block := range blocks {
		for _ = range block {
			if i%2 == 0 {
				filesystem = append(filesystem, i/2)
			} else {
				filesystem = append(filesystem, -1)
			}
		}
	}

	// Put two pointers at start and end of blocks list
	// Whenever a -1 is found at the start put the first block from the end
	i := 0
	j := len(filesystem) - 1
	for i < j {
		if filesystem[i] != -1 {
			i++
			continue
		}
		if filesystem[j] == -1 {
			j--
			continue
		}
		filesystem[i] = filesystem[j]
		filesystem[j] = -1
	}
	return filesystem
}

func buildFilesystemPartB(blocks []int) []int {
	filesOrder := []int{}
	for i := 0; i <= len(blocks)/2; i++ {
		filesOrder = append(filesOrder, i)
	}

	// For each file try moving it back
	for i := len(filesOrder) - 1; i >= 0; i-- {
		blocks, filesOrder = moveBackFile(blocks, i, filesOrder)
	}

	filesystem := []int{}
	for i, block := range blocks {
		for range block {
			if i%2 == 0 {
				filesystem = append(filesystem, filesOrder[i/2])
			} else {
				filesystem = append(filesystem, -1)
			}
		}
	}
	return filesystem
}

func moveBackFile(blocks []int, fileNumber int, filesOrder []int) ([]int, []int) {
	// Detect file block index in the blocks list
	blocksOrderNumIdx := 0
	for idx, order := range filesOrder {
		if order == fileNumber {
			blocksOrderNumIdx = idx
			break
		}
	}
	i := blocksOrderNumIdx * 2

	// Find a block of free space in which the file can fit
	for j, freeSpace := range blocks {
		// The free space must be on the left of the file
		if i <= j {
			break
		}
		// If it is not free space or file does not fit go to next block
		if j%2 == 0 || freeSpace < blocks[i] {
			continue
		}

		// Free space and file block are not adjacent
		if i-1 > j {
			// Calculate total free space to put on orignal file position
			movedFreeSpace := blocks[i]
			prevI := i
			nextI := i
			if i-1 >= 0 {
				prevI = i - 1
				movedFreeSpace += blocks[prevI]
			}
			if i+1 < len(blocks) {
				nextI = i + 1
				movedFreeSpace += blocks[nextI]
			}

			// Calculate the blocks to put to substitute the free space
			// 0 (no free space), file block, extra free space
			update := []int{0, blocks[i], freeSpace - blocks[i]}

			// Make the swap
			blocks = append(blocks[:prevI], append([]int{movedFreeSpace}, blocks[nextI+1:]...)...)
			blocks = append(blocks[:j], append(update, blocks[j+1:]...)...)
		} else { // Free space and file block are adjacent
			// Add to the free space the free space next to the file
			nextI := i
			if i+1 < len(blocks) {
				nextI = i + 1
				freeSpace += blocks[nextI]
			}

			// Calculate the blocks to put to substitute the free space
			// 0 (no free space), file block, extra free space
			update := []int{0, blocks[i], freeSpace}

			// Make the swap
			blocks = append(blocks[:i], blocks[nextI+1:]...)
			blocks = append(blocks[:j], append(update, blocks[j+1:]...)...)
		}

		// Swap the file number in the ordered list of file numbers
		filesOrder = append(filesOrder[:blocksOrderNumIdx], filesOrder[blocksOrderNumIdx+1:]...)
		filesOrder = append(filesOrder[:(j+1)/2], append([]int{fileNumber}, filesOrder[(j+1)/2:]...)...)
		break
	}
	return blocks, filesOrder
}

func init() {
	solutions.RegisterSolution(9, Solution09{})
}
