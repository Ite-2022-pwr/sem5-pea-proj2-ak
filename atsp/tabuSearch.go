package atsp

import (
	"pea2/graph"
	"pea2/utils"
)

type MovingMethod int

const (
	MovingSwap   MovingMethod = iota
	MovingInsert MovingMethod = iota
)

type Move struct {
	Vertex1 int
	Vertex2 int
}

type Neighbor struct {
	Path []int
	Cost int
	Move Move
}

type TabuSearchSolver struct {
	graph graph.Graph

	tabuTenure    int
	maxIterations int
	moving        MovingMethod
	bestPath      []int
	bestCost      int
	tabuList      map[Move]int
}

func NewTabuSearchSolver(G graph.Graph, tenure, maxIterations int, moving MovingMethod) *TabuSearchSolver {
	return &TabuSearchSolver{
		graph:         G,
		tabuTenure:    tenure,
		maxIterations: maxIterations,
		moving:        moving,
		tabuList:      make(map[Move]int),
	}
}

func (ts *TabuSearchSolver) generateInitialPermutation(startVertex int) []int {
	greedySolver := NewGreedySolver(ts.graph)
	_, perm := greedySolver.Solve(startVertex)
	return perm
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
	newPath := append([]int{}, path...)
	for k := j; k > i; k-- {
		newPath[k] = newPath[k-1]
	}
	newPath[i] = path[j]
	return newPath
}

//func (ts *TabuSearchSolver) isTabu(i, j, iter int) bool {
//	if tabu, ok := ts.tabuList[Move{i, j}]; ok {
//		if tabu < iter {
//			delete(ts.tabuList, Move{i, j})
//			return false
//		} else {
//			return true
//		}
//	}
//	return false
//}

func (ts *TabuSearchSolver) isTabu(i, j, iter int) bool {
	if tabu, ok := ts.tabuList[Move{i, j}]; ok {
		if tabu < iter {
			delete(ts.tabuList, Move{i, j})
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
	currentCost := bestCost

	for iteration := 0; iteration < ts.maxIterations; iteration++ {
		neighbors := utils.NewPriorityQueue(func(a, b Neighbor) bool { return a.Cost < b.Cost })

		// Generowanie sąsiedztwa
		for i := 1; i < ts.graph.GetVerticesCount(); i++ {
			for j := i + 1; j < ts.graph.GetVerticesCount(); j++ {

				neighbor := []int{}
				if ts.moving == MovingSwap {
					neighbor = swap(currentPath, i, j)
				} else {
					neighbor = insert(currentPath, i, j)
				}
				neighborCost := ts.graph.CalculatePathCost(neighbor)

				neighbors.Push(Neighbor{Path: neighbor, Cost: neighborCost, Move: Move{currentPath[i], currentPath[j]}})
			}
		}

		moved := false
		for !neighbors.IsEmpty() {
			neighbor := neighbors.Pop()

			if !ts.isTabu(neighbor.Move.Vertex1, neighbor.Move.Vertex2, iteration) || neighbor.Cost < bestCost {
				currentPath = neighbor.Path
				currentCost = neighbor.Cost
				ts.tabuList[neighbor.Move] = iteration + ts.tabuTenure
				moved = true
				break
			}
		}

		if moved && currentCost < bestCost {
			copy(bestPath, currentPath)
			bestCost = currentCost
		}

		if !moved || iteration%100 == 0 {
			utils.Shuffle(currentPath)
		}
	}

	return bestCost, bestPath
}
