package solution15

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"fmt"
	"math"
	"strings"
)

type Solution16 struct{}

type Node struct {
	i, j      int
	direction utils.Direction
}

func (s Solution16) PartA(lineIterator *utils.LineIterator) int {
	_, _, startNode, _, distances, _ := findShortestPath(lineIterator)
	return int(distances[startNode])
}

func (s Solution16) PartB(lineIterator *utils.LineIterator) int {
	nodes, neighbors, startNode, endNodes, preComputedDistances, _ := findShortestPath(lineIterator)

	// Re-run Dijkstra with each of the four possible end nodes (one for each direction) using
	// the pre-compuited distances of the shortest path in order to find for which final node is
	// the shortest path
	minDistance := math.Inf(1)
	var minPath map[Node]map[Node]bool
	for _, endNode := range endNodes {
		distances, path := dijkstra(nodes, neighbors, startNode, []Node{endNode}, preComputedDistances)
		if distances[startNode] <= minDistance {
			minDistance = distances[startNode]
			minPath = path
		}
	}

	// Follow the path from start to end and collect the visited nodes
	visited := map[Node]bool{}
	collectVisitedNodes(minPath, startNode, visited)
	// Make a distinct of the visited nodes
	distinctVisited := map[string]bool{}
	for node := range visited {
		nodeString := fmt.Sprintf("%d,%d", node.i, node.j)
		distinctVisited[nodeString] = true
	}
	return len(distinctVisited)
}

func findShortestPath(lineIterator *utils.LineIterator) ([]Node, map[Node][]Node, Node, []Node, map[Node]float64, map[Node]map[Node]bool) {
	input := [][]string{}
	for lineIterator.Next() {
		line := lineIterator.Value()
		input = append(input, strings.Split(line, ""))
	}

	// Init nodes list and map (used for fast-computing neighbors)
	nodes := []Node{}
	nodesMap := map[string]Node{}
	for i, row := range input {
		for j, col := range row {
			if col == "#" {
				continue
			}
			for _, direction := range utils.GetDirections() {
				node := Node{i, j, direction}
				nodes = append(nodes, node)
				key := fmt.Sprintf("%d,%d,%d", node.i, node.j, node.direction)
				nodesMap[key] = node
			}
		}
	}

	// Initialize start node, end nodes and neighbors for each node
	var startNode Node
	endNodes := []Node{}
	neighbors := map[Node][]Node{}
	for _, node := range nodes {
		neighbors[node] = findNeighbors(node, nodesMap)
		if input[node.i][node.j] == "S" && node.direction == utils.RightDirection {
			startNode = node
		}
		if input[node.i][node.j] == "E" {
			endNodes = append(endNodes, node)
		}
	}

	distances, path := dijkstra(nodes, neighbors, startNode, endNodes, map[Node]float64{})
	return nodes, neighbors, startNode, endNodes, distances, path
}

func dijkstra(nodes []Node, neighbors map[Node][]Node, startNode Node, endNodes []Node, preComputedDistances map[Node]float64) (map[Node]float64, map[Node]map[Node]bool) {
	distances := map[Node]float64{}
	path := map[Node]map[Node]bool{}
	visited := map[Node]bool{}

	// Init queue with end nodes
	queue := append([]Node{}, endNodes...)

	// For non-final nodes init distance with infinite, 0 for final nodes
	for _, node := range nodes {
		if utils.IsInSlice(endNodes, node) {
			distances[node] = 0
		} else {
			distances[node] = math.Inf(1)
			path[node] = map[Node]bool{}
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

		for _, neighbor := range neighbors[currNode] {
			if visited[neighbor] {
				continue
			}

			// Compute distance of neighbor from distance of current node
			distance := distances[currNode]
			if currNode.i != neighbor.i || currNode.j != neighbor.j {
				distance += 1
			}
			if currNode.direction != neighbor.direction {
				distance += 1000
			}

			// If distances are pre-computed (i.e. we just want to follow the path) and
			// current neightbor distance is equal to the pre-computed one,
			// than the neighbor is on the best path so add it to the queue
			if preComputedDistance, ok := preComputedDistances[neighbor]; ok && preComputedDistance == distance {
				path[neighbor][currNode] = true
				distances[neighbor] = preComputedDistance
				queue = addToQueue(queue, distances, neighbor)
			} else if !ok { // If distance is not pre-comuted run standard Dijkstra
				if distance < distances[neighbor] {
					// If current distance is less than minimum neighbor distance
					// update the minimum and add the neighbor to the queue
					// Also set the current node as the predecessor of the neighbor
					// in the shortest path(s)
					distances[neighbor] = distance
					path[neighbor] = map[Node]bool{currNode: true}
					queue = addToQueue(queue, distances, neighbor)
				} else if distance == distances[neighbor] {
					// If current distance is equal to minimum neighbor distance
					// just add the current node within the predecessors of the neighbor
					// in the shortest path(s)
					path[neighbor][currNode] = true
				}
			}
		}
		// If start node is processed than end, we are only interested in the paths from
		// the start node
		if currNode == startNode {
			break
		}
	}
	return distances, path
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

func collectVisitedNodes(path map[Node]map[Node]bool, current Node, visited map[Node]bool) {
	if visited[current] {
		return
	}
	visited[current] = true
	for node := range path[current] {
		collectVisitedNodes(path, node, visited)
	}
}

func findNeighbors(node Node, nodeMap map[string]Node) []Node {
	neighbors := []Node{}
	for _, direction := range getNeighborDirections(node.direction) {
		directionMoves := utils.GetDirectionMoves(direction)

		nextI, nextJ := node.i-directionMoves[0], node.j-directionMoves[1]

		key := fmt.Sprintf("%d,%d,%d", nextI, nextJ, direction)
		if neighbor, exists := nodeMap[key]; exists {
			neighbors = append(neighbors, neighbor)
		}

		samePositionKey := fmt.Sprintf("%d,%d,%d", node.i, node.j, direction)
		if neighbor, exists := nodeMap[samePositionKey]; exists {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func getNeighborDirections(direction utils.Direction) [3]utils.Direction {
	nextDirections := []utils.Direction{}
	nextDirections = append(nextDirections, direction)
	switch direction {
	case utils.TopDirection, utils.BottomDirection:
		nextDirections = append(nextDirections, []utils.Direction{utils.LeftDirection, utils.RightDirection}...)
	case utils.RightDirection, utils.LeftDirection:
		nextDirections = append(nextDirections, []utils.Direction{utils.TopDirection, utils.BottomDirection}...)
	}
	return [3]utils.Direction(nextDirections)
}

func init() {
	solutions.RegisterSolution(16, Solution16{})
}
