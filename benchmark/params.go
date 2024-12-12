package benchmark

// Parametry testowania algorytmów
const (
	NumberOfGraphs = 100
	MinVertices    = 7
	MaxVertices    = 8
	Rounds         = 10

	OutputDirectory = "data/output/"
	InputDirectory  = "data/input/"
)

var OptimalSolutions = map[string]int{
	"ftv33.txt": 1286,
	"ftv55.txt": 1608,
	"ftv64.txt": 1839,
}
