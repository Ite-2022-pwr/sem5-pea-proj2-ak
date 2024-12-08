package benchmark

import (
	"fmt"
	"log"
	"path/filepath"
	"pea2/atsp"
	"pea2/generator"
	"pea2/utils"
	"runtime/debug"
)

func BruteForce() {
	promt := utils.BlueColor("Brute Force")
	totalTime := 0.0
	var result [][]string

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Brute Force"))
	for numOfCities := MinVertices; numOfCities <= MaxVertices; numOfCities++ {
		for i := 0; i < NumberOfGraphs; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, test: %d/%d", numOfCities, i+1, NumberOfGraphs)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
			tsp := atsp.NewBruteForceSolver(G)
			totalTime += MeasureSolveTime(tsp, promt)
			debug.FreeOSMemory()
		}
		avgTime := totalTime / float64(NumberOfGraphs)
		result = append(result, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%d", int64(avgTime))})
		utils.SaveCSV(filepath.Join(OutputDirectory, "brute_force3.csv"), result)
		totalTime = 0.0
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "brute_force.csv"), result)
}
