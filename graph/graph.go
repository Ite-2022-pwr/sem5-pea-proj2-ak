package graph

type Graph interface {
	PutEdge(sourceVertex, destinationVertex, weight int) error
	GetEdge(sourceVertex, destinationVertex int) (int, error)
	GetVerticesCount() int
	GetEdgesCount() int
	GetCopy() Graph
	ToString() string
	AsMatrix() [][]int
	CalculatePathCost(path []int) int
}
