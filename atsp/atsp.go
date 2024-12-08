// atsp to pakiet, który zawiera implementacje algorytmów rozwiązujących problem ATSP (Asymmetric Travelling Salesman Problem).
package atsp

import "pea2/graph"

// ATSP to interfejs, który muszą implementować wszystkie algorytmy rozwiązujące problem ATSP
type ATSP interface {
	Solve(startVertex int) (int, []int)
	GetGraph() graph.Graph
}
