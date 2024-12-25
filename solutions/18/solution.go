package solution18

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"fmt"
	"math"
	"strings"
)

type Solution18 struct{}

type Node struct {
	i, j      int
}

func (s Solution18) PartA(lineIterator *utils.LineIterator) any {
	_, _, nodes, nodesMap, startingNode, finalNode := parseInput(lineIterator, 1024)
	distances := dijkstra(nodes, nodesMap, startingNode, finalNode)
	return distances[startingNode]
}

func (s Solution18) PartB(lineIterator *utils.LineIterator) any {
	input, corruptedPositions, nodes, nodesMap, startingNode, finalNode := parseInput(lineIterator, -1)
	// Try removing corrupted nodes from the end, when a valid path is found
	// (i.e. a path with distance from starting node less than infinite),
	// the current byte is the one which blocks the path if corrupted
	for idx := len(input) - 1; idx >= 0; idx-- {
		i, j := input[idx][0], input[idx][1]
		corruptedPositions[i][j] = false
		node := Node{i, j}
		coordsString := fmt.Sprintf("%d,%d", i, j)
		nodesMap[coordsString] = node
		nodes = append(nodes, node)
		distances := dijkstra(nodes, nodesMap, startingNode, finalNode)
		if distances[startingNode] < math.Inf(1) {
			return fmt.Sprintf("%d,%d", i, j)
		}
	}
	panic("Solution not found")
}

func dijkstra(nodes []Node, nodesMap map[string]Node, startingNode Node, finalNode Node) map[Node]float64 {
	distances := map[Node]float64{}
	visited := map[Node]bool{}

	// Init queue with end node
	queue := append([]Node{}, finalNode)

	// For non-final nodes init distance with infinite, 0 for final node
	for _, node := range nodes {
		if node == finalNode {
			distances[node] = 0
		} else {
			distances[node] = math.Inf(1)
		}
	}

	for len(queue) > 0 {
		// Remove current node from queue
		// First node of the queue is always the one with minimum distance from the final node(s)
		currNode := queue[0]
		queue = queue[1:]
		if visited[currNode] {
			continue
		}
		visited[currNode] = true

		for _, direction := range utils.GetDirections() {
			// Compute neighbor for each direction
			directionMoves := utils.GetDirectionMoves(direction)
			neighborI, neighborJ := currNode.i + directionMoves[0], currNode.j + directionMoves[1]
			coordsString := fmt.Sprintf("%d,%d", neighborI, neighborJ)

			neighbor, ok := nodesMap[coordsString]
			if !ok || visited[neighbor] {
				continue
			}

			// Compute distance of neighbor from distance of current node
			distance := distances[currNode] + 1

			// If current distance is less than minimum neighbor distance
			// update the minimum and add the neighbor to the queue
			// Also set the current node as the predecessor of the neighbor
			// in the shortest path(s)
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				queue = addToQueue(queue, distances, neighbor)
			}
		}
		// If start node is processed than end, we are only interested in the paths from
		// the start node
		if currNode == startingNode {
			break
		}
	}
	return distances
}

// Add a node to the queue keeping the queue ordered by distance
func addToQueue(queue []Node, distances map[Node]float64, node Node) []Node {
	added := false
	for idx, currNode := range queue {
		if distances[currNode] >= distances[node] {
			queue = append(queue[:idx], append([]Node{node}, queue[idx:]...)...)
			added = true
			break
		}
	}
	if !added {
		queue = append(queue, node)
	}
	return queue
}

func parseInput(lineIterator *utils.LineIterator, limit int) ([][2]int, [71][71]bool, []Node, map[string]Node, Node, Node) {
	input := [][2]int{}
	for lineIterator.Next() {
		line := lineIterator.Value()
		input = append(input, [2]int(utils.StringsToIntegers(strings.Split(line, ","))))
	}
	corruptedPositions := [71][71]bool{}
	if limit == -1 {
		limit = len(input)
	}
	for idx := range limit {
		corruptedPositions[input[idx][0]][input[idx][1]] = true
	}
	nodes := []Node{}
	nodesMap := map[string]Node{}
	var startingNode Node
	var finalNode Node
	for i, row := range corruptedPositions {
		for j, col := range row {
			if !col {
				node := Node{i, j}
				coordsString := fmt.Sprintf("%d,%d", i, j)
				nodesMap[coordsString] = node
				nodes = append(nodes, node)
				if i == 0 && j == 0 {
					startingNode = node
				} else if i == 70 && j == 70 {
					finalNode = node
				}
			}
		}
	}
	return input, corruptedPositions, nodes, nodesMap, startingNode, finalNode
}

func init() {
	solutions.RegisterSolution(18, Solution18{})
}
