package atsp

import (
	"math"
	"pea2/graph"
)

type GreedySolver struct {
	graph graph.Graph
}

func NewGreedySolver(G graph.Graph) *GreedySolver {
	return &GreedySolver{
		graph: G,
	}
}

func (atsp *GreedySolver) GetGraph() graph.Graph {
	return atsp.graph
}

func (atsp *GreedySolver) Solve(startVertex int) (int, []int) {
	return atsp.Greedy(startVertex)
}

func (atsp *GreedySolver) findMinNeighbor(costs []int, visited []bool, vertex int) int {
	minVal, minVertex := math.MaxInt, 0
	for i := 0; i < len(costs); i++ {
		if i == vertex {
			continue
		}

		if visited[i] {
			continue
		}

		if costs[i] < minVal {
			minVal = costs[i]
			minVertex = i
		}
	}

	return minVertex
}

func (atsp *GreedySolver) Greedy(startVertex int) (int, []int) {
	bestPath := make([]int, atsp.graph.GetVerticesCount())
	bestPath[0] = startVertex
	visited := make([]bool, atsp.graph.GetVerticesCount())
	visited[startVertex] = true

	for i := 1; i < atsp.graph.GetVerticesCount(); i++ {
		nextVertex := atsp.findMinNeighbor(atsp.GetGraph().AsMatrix()[bestPath[i-1]], visited, bestPath[i-1])
		bestPath[i] = nextVertex
		visited[nextVertex] = true
	}

	bestCost := atsp.graph.CalculatePathCost(bestPath)
	return bestCost, bestPath
}
