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

func AllLocalSearch() {
	promptTS := utils.BlueColor("Tabu Search")
	promptSA := utils.BlueColor("Simulated Annealing")

	totalTimeTS, totalTimeSA := 0.0, 0.0 // zmienne do przechowywania czasu rozwiązania
	var resultTS, resultSA [][]string    // wyniki

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Algorytmów"))
	for numOfCities := 50; numOfCities <= 250; numOfCities += 25 {
		for i := 0; i < NumberOfGraphs; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, test: %d/%d", numOfCities, i+1, NumberOfGraphs)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
			var tsp atsp.ATSP
			tsp = atsp.NewTabuSearchSolver(G, 20, 1000, atsp.MovingInsert)

			totalTimeTS += MeasureSolveTime(tsp, promptTS)

			tsp = atsp.NewSimulatedAnnealingSolver(G, 0.975, 1e-12, 9000, 4000)
			totalTimeSA += MeasureSolveTime(tsp, promptSA)
			debug.FreeOSMemory()
		}

		// średni czas dla każdego z algorytmów
		avgTimeTS := totalTimeTS / float64(NumberOfGraphs)
		avgTimeSA := totalTimeSA / float64(NumberOfGraphs)
		resultTS = append(resultTS, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%.3f", avgTimeTS/1000000000.0)})
		resultSA = append(resultSA, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%.3f", avgTimeSA/1000000000.0)})
		utils.SaveCSV(filepath.Join(OutputDirectory, "tabu_search2.csv"), resultTS)
		utils.SaveCSV(filepath.Join(OutputDirectory, "simulated_annealing2.csv"), resultSA)
		totalTimeTS, totalTimeSA = 0.0, 0.0
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "tabu_search2.csv"), resultTS)
	utils.SaveCSV(filepath.Join(OutputDirectory, "simulated_annealing2.csv"), resultSA)
}
