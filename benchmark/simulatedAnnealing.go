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

var coolingRates = []float64{0.975, 0.95, 0.75}
var epochs = []int{2000, 4000, 6000}
var minTemperatures = []float64{1e-9, 1e-12, 1e-15}
var initTemperatures = []float64{3000, 6000, 9000}

func TestSimulatedAnnealingInitialTemperatures() {
	prompt := "Simulated Annealing"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie początkowej temperatury
		for it := 0; it < len(initTemperatures); it++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("SA init temp: %v, test: %d/%d", initTemperatures[it], r+1, Rounds)))
				tsp := atsp.NewSimulatedAnnealingSolver(G, coolingRates[0], minTemperatures[0], initTemperatures[it], epochs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					fmt.Sprintf("%d", int(initTemperatures[it])),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("SA_IT_%v_%d.csv", filename, int(initTemperatures[it]))), result)
			}
		}
	}
}

func TestSimulatedAnnealingMinimalTemperatures() {
	prompt := "Simulated Annealing"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie końcowej temperatury
		for mt := 0; mt < len(minTemperatures); mt++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("SA min temp: %v, test: %d/%d", minTemperatures[mt], r+1, Rounds)))
				tsp := atsp.NewSimulatedAnnealingSolver(G, coolingRates[0], minTemperatures[mt], initTemperatures[0], epochs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					fmt.Sprintf("%g", minTemperatures[mt]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("SA_MT_%v_%g.csv", filename, minTemperatures[mt])), result)
			}
		}
	}
}

func TestSimulatedAnnealingEpochs() {
	prompt := "Simulated Annealing"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie epok
		for ep := 0; ep < len(epochs); ep++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("SA epochs: %v, test: %d/%d", epochs[ep], r+1, Rounds)))
				tsp := atsp.NewSimulatedAnnealingSolver(G, coolingRates[0], minTemperatures[0], initTemperatures[0], epochs[ep])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					fmt.Sprintf("%d", epochs[ep]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("SA_EP_%v_%d.csv", filename, epochs[ep])), result)
			}
		}
	}
}

func TestSimulatedAnnealingCoolingRates() {
	prompt := "Simulated Annealing"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie końcowej temperatury
		for cr := 0; cr < len(coolingRates); cr++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("SA cooling rate: %v, test: %d/%d", coolingRates[cr], r+1, Rounds)))
				tsp := atsp.NewSimulatedAnnealingSolver(G, coolingRates[cr], minTemperatures[0], initTemperatures[0], epochs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					fmt.Sprintf("%v", coolingRates[cr]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("SA_CR_%v_%v.csv", filename, coolingRates[cr])), result)
			}
		}
	}
}

func TestSimulatedAnnealingBestParams() {
	prompt := "Simulated Annealing"

	for filename, optimalCost := range OptimalSolutions {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		var result [][]string
		for r := 0; r < Rounds; r++ {
			log.Println(utils.BlueColor(fmt.Sprintf("SA best, test: %d/%d", r+1, Rounds)))
			tsp := atsp.NewSimulatedAnnealingSolver(G, 0.975, 1e-12, 9000, 4000)
			elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
			result = append(result, []string{
				fmt.Sprintf("%d", cost),
				fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
				fmt.Sprintf("%.3f", elapsed/1000000000.0),
			})
			utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("SA_CR_%v_best.csv", filename)), result)
		}
	}
}

func TestSimulatedAnnealingLimit() {
	promptSA := utils.BlueColor("Simulated Annealing")

	totalTimeSA := 0.0      // zmienne do przechowywania czasu rozwiązania
	var resultSA [][]string // wyniki

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Algorytmów"))
	for numOfCities := 600; numOfCities <= 1000; numOfCities += 100 {
		for i := 0; i < NumberOfGraphs; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, test: %d/%d", numOfCities, i+1, NumberOfGraphs)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
			var tsp atsp.ATSP

			tsp = atsp.NewSimulatedAnnealingSolver(G, 0.975, 1e-12, 9000, 4000)
			totalTimeSA += MeasureSolveTime(tsp, promptSA)
			debug.FreeOSMemory()
		}

		// średni czas dla każdego z algorytmów
		avgTimeDP := totalTimeSA / float64(NumberOfGraphs)
		resultSA = append(resultSA, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%.3f", avgTimeDP/1000000000.0)})
		utils.SaveCSV(filepath.Join(OutputDirectory, "simulated_annealing4.csv"), resultSA)
		totalTimeSA = 0.0
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "simulated_annealing3.csv"), resultSA)
}
