package benchmark

import (
	"fmt"
	"log"
	"path/filepath"
	"pea2/atsp"
	"pea2/utils"
)

var iterations = []int{1000, 2000, 3000}
var moveTypes = []atsp.MovingMethod{atsp.MovingSwap, atsp.MovingInsert}
var tenures = []int{10, 15, 20}

func TestTabuSearchIterations() {
	prompt := "Tabu Search"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie liczby iteracji
		for it := 0; it < len(iterations); it++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("TS iterations: %v, test: %d/%d", iterations[it], r+1, Rounds)))
				tsp := atsp.NewTabuSearchSolver(G, tenures[0], iterations[it], moveTypes[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					fmt.Sprintf("%d", iterations[it]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("TS_IT_%v_%d.csv", filename, iterations[it])), result)
			}
		}
	}
}

func TestTabuSearchMoveTypes() {
	prompt := "Tabu Search"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie typu ruchu
		for mt := 0; mt < len(moveTypes); mt++ {
			var result [][]string
			moveTypeStr := ""
			if moveTypes[mt] == atsp.MovingSwap {
				moveTypeStr = "swap"
			} else {
				moveTypeStr = "insert"
			}
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("TS move type: %v, test: %d/%d", moveTypeStr, r+1, Rounds)))
				tsp := atsp.NewTabuSearchSolver(G, tenures[0], iterations[0], moveTypes[mt])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					moveTypeStr,
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("TS_MT_%v_%s.csv", filename, moveTypeStr)), result)
			}
		}
	}
}

func TestTabuSearchTenures() {
	prompt := "Tabu Search"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie kadencji
		for tn := 0; tn < len(tenures); tn++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("TS tenure: %v, test: %d/%d", tenures[tn], r+1, Rounds)))
				tsp := atsp.NewTabuSearchSolver(G, tenures[tn], iterations[0], moveTypes[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					fmt.Sprintf("%d", tenures[tn]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("TS_TN_%v_%d.csv", filename, tenures[tn])), result)
			}
		}
	}
}
