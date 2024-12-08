package atsp

import (
	"math"
	"pea2/graph"
)

// BruteForceSolver to struktura implementująca algorytm rozwiązujący problem ATSP za pomocą siłowego przeglądu
type BruteForceSolver struct {
	graph graph.Graph
}

// GetGraph zwraca graf, na którym działa algorytm
func (atsp *BruteForceSolver) GetGraph() graph.Graph {
	return atsp.graph
}

// NewBruteForceSolver tworzy nowy obiekt BruteForceSolver
func NewBruteForceSolver(g graph.Graph) *BruteForceSolver {
	return &BruteForceSolver{graph: g}
}

// Solve rozwiązuje problem ATSP na grafie g, zaczynając od wierzchołka startVertex
func (atsp *BruteForceSolver) Solve(startVertex int) (int, []int) {
	return atsp.BruteForce(startVertex)
}

// BruteForce to funkcja rozwiązująca problem ATSP za pomocą siłowego przeglądu
func (atsp *BruteForceSolver) BruteForce(startVertex int) (int, []int) {
	visited := make([]bool, atsp.graph.GetVerticesCount())
	path := make([]int, 0, atsp.graph.GetVerticesCount())
	bestPath := make([]int, atsp.graph.GetVerticesCount())
	bestCost := math.MaxInt

	visited[startVertex] = true
	path = append(path, startVertex)

	bestCost, bestPath = atsp.bruteForceRecursive(startVertex, visited, startVertex, 0, bestCost, bestPath, path)

	return bestCost, bestPath
}

// bruteForceRecursive to funkcja rekurencyjna rozwiązująca problem ATSP za pomocą siłowego przeglądu
func (atsp *BruteForceSolver) bruteForceRecursive(startVertex int, visited []bool, currentVertex int, currentCost int, bestCost int, bestPath []int, path []int) (int, []int) {
	// Jeśli odwiedzono wszystkie wierzchołki, to zwróć koszt powrotu do wierzchołka startowego
	if len(path) == atsp.graph.GetVerticesCount() {
		cost, _ := atsp.graph.GetEdge(currentVertex, startVertex) // koszt powrotu do wierzchołka startowego
		currentCost += cost

		// Jeśli znaleziono lepsze rozwiązanie, to zaktualizuj najlepsze rozwiązanie
		if currentCost < bestCost {
			bestCost = currentCost
			copy(bestPath, path)
		}

		return bestCost, bestPath
	}

	// Rekurencyjnie sprawdź wszystkie możliwe ścieżki
	for i := 0; i < atsp.graph.GetVerticesCount(); i++ {
		if !visited[i] {
			visited[i] = true
			cost, _ := atsp.graph.GetEdge(currentVertex, i)
			currentCost += cost
			path = append(path, i)

			bestCost, bestPath = atsp.bruteForceRecursive(startVertex, visited, i, currentCost, bestCost, bestPath, path)

			visited[i] = false
			currentCost -= cost
			path = path[:len(path)-1]
		}
	}

	return bestCost, bestPath
}
