package graph

import "fmt"

type AdjacencyMatrix struct {
	vertices int
	edges    int
	matrix   [][]int
}

func NewAdjacencyMatrix(vertices int) (*AdjacencyMatrix, error) {
	if vertices <= 0 {
		return nil, fmt.Errorf("vertices must be a positive number, given: %d", vertices)
	}

	am := &AdjacencyMatrix{vertices: vertices}
	am.matrix = make([][]int, am.vertices)
	for i := 0; i < am.vertices; i++ {
		am.matrix[i] = make([]int, am.vertices)
	}

	return am, nil
}

func (am *AdjacencyMatrix) PutEdge(startVertex, destinationVertex, weight int) error {
	if startVertex < 0 || startVertex >= am.vertices {
		return fmt.Errorf("wrong vertex value: %d", startVertex)
	}

	if destinationVertex < 0 || destinationVertex > am.vertices {
		return fmt.Errorf("wrong vertex value: %d", destinationVertex)
	}

	if weight < 1 {
		return fmt.Errorf("weight must be a positive number, given: %d", weight)
	}

	if am.matrix[startVertex][destinationVertex] < 1 {
		am.edges++ // new edge
	}

	am.matrix[startVertex][destinationVertex] = weight

	return nil
}

func (am *AdjacencyMatrix) GetEdge(startVertex, destinationVertex int) (int, error) {
	if startVertex < 0 || startVertex >= am.vertices {
		return 0, fmt.Errorf("wrong vertex value: %d", startVertex)
	}

	if destinationVertex < 0 || destinationVertex > am.vertices {
		return 0, fmt.Errorf("wrong vertex value: %d", destinationVertex)
	}

	return am.matrix[startVertex][destinationVertex], nil
}

func (am *AdjacencyMatrix) GetVerticesCount() int {
	return am.vertices
}

func (am *AdjacencyMatrix) GetEdgesCount() int {
	return am.edges
}

func (am *AdjacencyMatrix) GetCopy() Graph {
	amCopy := &AdjacencyMatrix{vertices: am.vertices, edges: am.edges}
	amCopy.matrix = make([][]int, am.vertices)
	for i := 0; i < am.vertices; i++ {
		amCopy.matrix[i] = make([]int, am.vertices)
		copy(amCopy.matrix[i], am.matrix[i])
	}

	return amCopy
}

func (am *AdjacencyMatrix) ToString() string {
	var str string
	for i := 0; i < am.vertices; i++ {
		for j := 0; j < am.vertices; j++ {
			str += fmt.Sprintf("%d ", am.matrix[i][j])
		}
		str += "\n"
	}

	return str
}

func (am *AdjacencyMatrix) AsMatrix() [][]int {
	return am.matrix
}

func (am *AdjacencyMatrix) CalculatePathCost(path []int) int {
	n := len(path)
	cost := 0
	for i := 0; i < n-1; i++ {
		cost += am.matrix[path[i]][path[i+1]]
	}
	cost += am.matrix[path[n-1]][path[0]]
	return cost
}
