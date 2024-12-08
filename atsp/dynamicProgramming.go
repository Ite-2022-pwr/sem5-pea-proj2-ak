package atsp

import (
	"math"
	"pea2/graph"
)

// DynamicProgrammingSolver to struktura implementująca algorytm rozwiązujący problem ATSP za pomocą programowania dynamicznego
type DynamicProgrammingSolver struct {
	graph graph.Graph

	visitedAll int     // maska wskazująca, że wszystkie wierzchołki zostały odwiedzone
	memoCosts  [][]int // tablica przechowująca koszty ścieżek
	memoPaths  [][]int // tablica przechowująca kolejne wierzchołki na ścieżce
	start      int
}

// GetGraph zwraca graf, na którym działa algorytm
func (atsp *DynamicProgrammingSolver) GetGraph() graph.Graph {
	return atsp.graph
}

// NewDynamicProgrammingSolver tworzy nowy obiekt DynamicProgrammingSolver
func NewDynamicProgrammingSolver(g graph.Graph) *DynamicProgrammingSolver {
	visitedAll := (1 << uint(g.GetVerticesCount())) - 1

	dp := make([][]int, 1<<uint(g.GetVerticesCount()))
	memo := make([][]int, 1<<uint(g.GetVerticesCount()))
	for i := 0; i < 1<<uint(g.GetVerticesCount()); i++ {
		dp[i] = make([]int, g.GetVerticesCount())
		memo[i] = make([]int, g.GetVerticesCount())
	}

	return &DynamicProgrammingSolver{
		graph:      g,
		visitedAll: visitedAll,
		memoCosts:  dp,
		memoPaths:  memo}
}

// Solve rozwiązuje problem ATSP na grafie g, zaczynając od wierzchołka startVertex
func (atsp *DynamicProgrammingSolver) Solve(startVertex int) (int, []int) {
	return atsp.DynamicProgramming(startVertex)
}

// DynamicProgramming to funkcja rozwiązująca problem ATSP za pomocą programowania dynamicznego
func (atsp *DynamicProgrammingSolver) DynamicProgramming(startVertex int) (int, []int) {
	for i := 0; i < 1<<uint(atsp.graph.GetVerticesCount()); i++ {
		for j := 0; j < atsp.graph.GetVerticesCount(); j++ {
			atsp.memoCosts[i][j] = -1
		}
	}
	atsp.start = startVertex

	cost := atsp.dynamicProgrammingRecursive(1<<uint(startVertex), startVertex)
	path := atsp.lookupPath()

	return cost, path
}

func (atsp *DynamicProgrammingSolver) dynamicProgrammingRecursive(mask, position int) int {
	// Jeśli odwiedzono wszystkie wierzchołki, to zwróć koszt przejścia z ostatniego wierzchołka do wierzchołka startowego
	if mask == atsp.visitedAll {
		cost, _ := atsp.graph.GetEdge(position, atsp.start)
		return cost
	}

	// Jeśli koszt dla danej maski i pozycji został już obliczony, to zwróć go
	if atsp.memoCosts[mask][position] != -1 {
		return atsp.memoCosts[mask][position]
	}

	answer := math.MaxInt
	answerCity := -1

	// Dla każdego miasta, które nie zostało jeszcze odwiedzone, oblicz koszt przejścia z obecnego miasta do tego miasta
	for city := 0; city < atsp.graph.GetVerticesCount(); city++ {
		if mask&(1<<uint(city)) == 0 {
			cost, _ := atsp.graph.GetEdge(position, city)
			newAnswer := cost + atsp.dynamicProgrammingRecursive(mask|(1<<uint(city)), city) // obliczenie kosztu dla kolejnego miasta

			// Jeśli koszt jest mniejszy od aktualnego kosztu, to zaktualizuj koszt i miasto
			if newAnswer < answer {
				answer = newAnswer
				answerCity = city
			}
		}
	}

	// Zapisz koszt i następne miasto dla danej maski i pozycji
	atsp.memoCosts[mask][position] = answer
	atsp.memoPaths[mask][position] = answerCity

	return answer
}

// lookupPath to funkcja zwracająca ścieżkę w grafie
func (atsp *DynamicProgrammingSolver) lookupPath() []int {
	mask := 1 << uint(atsp.start)

	path := make([]int, 0, atsp.graph.GetVerticesCount())
	path = append(path, atsp.start)
	lastCity := atsp.start // ostatnie miasto na ścieżce

	for len(path) < atsp.graph.GetVerticesCount() {
		nextCity := atsp.memoPaths[mask][lastCity] // następne miasto na ścieżce
		mask |= 1 << uint(nextCity)                // aktualizacja maski odwiedzonych miast
		path = append(path, nextCity)              // dodanie miasta do ścieżki
		lastCity = nextCity
	}

	return path
}
