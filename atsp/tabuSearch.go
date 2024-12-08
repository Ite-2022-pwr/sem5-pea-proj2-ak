package atsp

import (
	"math"
	"math/rand"
	"pea2/graph"
)

type MovingMethod int

const (
	MovingSwap   MovingMethod = iota
	MovingInsert MovingMethod = iota
)

type TabuElement struct {
	Vertex1 int
	Vertex2 int
}

type TabuSearchSolver struct {
	graph graph.Graph

	tabuTenure    int
	maxIterations int
	moving        MovingMethod
	bestPath      []int
	bestCost      int
	tabuList      map[TabuElement]int
}

func NewTabuSearchSolver(G graph.Graph, tenure, maxIterations int, moving MovingMethod) *TabuSearchSolver {
	return &TabuSearchSolver{
		graph:         G,
		tabuTenure:    tenure,
		maxIterations: maxIterations,
		moving:        moving,
		tabuList:      make(map[TabuElement]int),
	}
}

func (ts *TabuSearchSolver) generateInitialPermutation(startVertex int) []int {
	path := make([]int, ts.graph.GetVerticesCount())
	for i := 0; i < ts.graph.GetVerticesCount(); i++ {
		path[i] = i
	}

	path[startVertex], path[0] = path[0], path[startVertex]

	for i := 1; i < ts.graph.GetVerticesCount(); i++ {
		randIdx := rand.Intn(ts.graph.GetVerticesCount()-1) + 1
		path[i], path[randIdx] = path[randIdx], path[i]
	}

	return path
}

func (ts *TabuSearchSolver) GetGraph() graph.Graph {
	return ts.graph
}

func (ts *TabuSearchSolver) Solve(startVertex int) (int, []int) {
	return ts.TabuSearch(startVertex)
}

func swap(path []int, i, j int) []int {
	newPath := make([]int, len(path))
	copy(newPath, path)
	newPath[i], newPath[j] = newPath[j], newPath[i]
	return newPath
}

func insert(path []int, i, j int) []int {
	newPath := make([]int, 0, len(path))
	for k := 0; k < len(path); k++ {
		if k == i {
			continue
		}

		if k == j {
			newPath = append(newPath, path[i])
		}
		newPath = append(newPath, path[k])
	}
	return newPath
}

func (ts *TabuSearchSolver) isTabu(i, j, iter int) bool {
	if tabu, ok := ts.tabuList[TabuElement{i, j}]; ok {
		if tabu < iter {
			delete(ts.tabuList, TabuElement{i, j})
			return false
		} else {
			return true
		}
	}
	return false
}

func (ts *TabuSearchSolver) TabuSearch(startVertex int) (int, []int) {
	currentPath := ts.generateInitialPermutation(startVertex)
	bestPath := append([]int{}, currentPath...)
	bestCost := ts.graph.CalculatePathCost(bestPath)

	for iteration := 0; iteration < ts.maxIterations; iteration++ {
		bestNeighbor := []int{}
		bestNeighborCost := math.MaxInt

		// Generowanie sąsiedztwa
		for i := 1; i < ts.graph.GetVerticesCount(); i++ {
			for j := i + 1; j < ts.graph.GetVerticesCount(); j++ {
				if ts.isTabu(i, j, iteration) {
					continue
				}

				neighbor := []int{}
				if ts.moving == MovingSwap {
					neighbor = swap(currentPath, i, i)
				} else {
					neighbor = insert(currentPath, i, j)
				}
				neighborCost := ts.graph.CalculatePathCost(neighbor)

				if neighborCost < bestNeighborCost {
					bestNeighbor = neighbor
					bestNeighborCost = neighborCost
				}
			}
		}

		if bestNeighborCost < bestCost {
			copy(bestPath, bestNeighbor)
			bestCost = bestNeighborCost
		}

		// Aktualizacja listy ruchów zakazanych
		for i := 0; i < ts.graph.GetVerticesCount(); i++ {
			for j := i + 1; j < ts.graph.GetVerticesCount(); j++ {
				if currentPath[i] != bestPath[i] && currentPath[j] != bestPath[j] {
					ts.tabuList[TabuElement{i, j}] = iteration + ts.tabuTenure
				}
			}
		}

		copy(currentPath, bestNeighbor)
	}

	return bestCost, bestPath
}
