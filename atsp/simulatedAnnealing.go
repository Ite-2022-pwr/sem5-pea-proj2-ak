package atsp

import (
	"math"
	"math/rand"
	"pea2/graph"
)

type SimulatedAnnealingSolver struct {
	graph graph.Graph

	coolingRate        float64
	minTemperature     float64
	initialTemperature float64
	epochs             int
}

func NewSimulatedAnnealingSolver(G graph.Graph, coolingRate, minTemp, initTemp float64, epochs int) *SimulatedAnnealingSolver {
	return &SimulatedAnnealingSolver{
		graph:              G,
		minTemperature:     minTemp,
		initialTemperature: initTemp,
		coolingRate:        coolingRate,
		epochs:             epochs,
	}
}

func (sa *SimulatedAnnealingSolver) GetGraph() graph.Graph {
	return sa.graph
}

func (sa *SimulatedAnnealingSolver) Solve(startVertex int) (int, []int) {
	return sa.SimulatedAnnealing(startVertex)
}

func (sa *SimulatedAnnealingSolver) generateInitialPermutation(startVertex int) []int {
	greedySolver := NewGreedySolver(sa.graph)
	_, perm := greedySolver.Solve(startVertex)
	return perm
}

func (sa *SimulatedAnnealingSolver) calculateAcceptanceProbability(delta int, temperature float64) float64 {
	//return math.Exp(float64(delta) / temperature)
	return 1 / (1 + math.Exp(float64(delta)/temperature))
}

func (sa *SimulatedAnnealingSolver) generateNeighbor2(path []int) []int {
	neigh := append([]int{}, path...)
	idx1, idx2 := rand.Intn(len(path)-1)+1, rand.Intn(len(path)-1)+1
	for idx1 == idx2 {
		idx2 = rand.Intn(len(path)-1) + 1
	}
	neigh[idx1], neigh[idx2] = neigh[idx2], neigh[idx1]
	return neigh
}

func (sa *SimulatedAnnealingSolver) generateNeighbor(path []int) []int {
	neigh := append([]int{}, path...)
	idx1, idx2 := rand.Intn(len(path)-1)+1, rand.Intn(len(path)-1)+1
	for idx1 == idx2 {
		idx2 = rand.Intn(len(path)-1) + 1
	}

	if rand.Int()%2 == 0 {
		// odwróć ścieżkę
		if idx1 > idx2 {
			idx1, idx2 = idx2, idx1
		}

		for i, j := 1, len(neigh)-1; i < j; i, j = i+1, j-1 {
			neigh[i], neigh[j] = neigh[j], neigh[i]
		}
	} else {
		neigh[idx1], neigh[idx2] = neigh[idx2], neigh[idx1]
	}
	return neigh
}

func (sa *SimulatedAnnealingSolver) SimulatedAnnealing(startVertex int) (int, []int) {
	currentPath := sa.generateInitialPermutation(startVertex)
	bestPath := append([]int{}, currentPath...)
	currentCost := sa.graph.CalculatePathCost(currentPath)
	bestCost := currentCost

	temperature := sa.initialTemperature
	for temperature > sa.minTemperature {
		for epoch := 0; epoch < sa.epochs; epoch++ {
			neighbor := sa.generateNeighbor(currentPath)
			neighborCost := sa.graph.CalculatePathCost(neighbor)
			delta := neighborCost - currentCost
			//fmt.Println(neighborCost, neighbor)
			if delta < 0 || rand.Float64() < sa.calculateAcceptanceProbability(delta, temperature) {
				//fmt.Println(neighborCost, neighbor)
				currentPath = neighbor
				currentCost = neighborCost
			}

			// sprawdzamy, czy znalezione lokalne rozwiązanie jest lepsze od globalnego
			if currentCost < bestCost {
				copy(bestPath, currentPath)
				bestCost = currentCost
			}
		}

		temperature *= sa.coolingRate
	}

	//gr := sa.generateInitialPermutation(startVertex)
	//fmt.Println("Greedy:", sa.graph.CalculatePathCost(gr), gr)
	return bestCost, bestPath
}
