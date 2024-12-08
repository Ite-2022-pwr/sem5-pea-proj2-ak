// benchmark jest odpowiedzialny za testowanie algorytmów
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

// All - testuje oba algorytmy dla różnych ilości miast
func All() {
	promptBB := utils.BlueColor("Branch and Bound")
	promptDP := utils.BlueColor("Dynamic programming")

	totalTimeBB, totalTimeDP := 0.0, 0.0 // zmienne do przechowywania czasu rozwiązania
	var resultBB, resultDP [][]string    // wyniki

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Algorytmów"))
	for numOfCities := MinVertices; numOfCities <= MaxVertices; numOfCities++ {
		for i := 0; i < NumberOfGraphs; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, test: %d/%d", numOfCities, i+1, NumberOfGraphs)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
			var tsp atsp.ATSP
			tsp = atsp.NewBranchAndBoundSolver(G)
			totalTimeBB += MeasureSolveTime(tsp, promptBB)

			tsp = atsp.NewDynamicProgrammingSolver(G)
			totalTimeDP += MeasureSolveTime(tsp, promptDP)
			debug.FreeOSMemory()
		}

		// średni czas dla każdego z algorytmów
		avgTimeBB := totalTimeBB / float64(NumberOfGraphs)
		avgTimeDP := totalTimeDP / float64(NumberOfGraphs)
		resultBB = append(resultBB, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%d", int64(avgTimeBB))})
		resultDP = append(resultDP, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%d", int64(avgTimeDP))})
		utils.SaveCSV(filepath.Join(OutputDirectory, "branch_and_bound3.csv"), resultBB)
		utils.SaveCSV(filepath.Join(OutputDirectory, "dynamic_programming3.csv"), resultDP)
		totalTimeBB, totalTimeDP = 0.0, 0.0
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "branch_and_bound3.csv"), resultBB)
	utils.SaveCSV(filepath.Join(OutputDirectory, "dynamic_programming3.csv"), resultDP)
}
