package main

import "pea2/benchmark"

func main() {
	//menu.RunMenu()

	//benchmark.TestSimulatedAnnealingInitialTemperatures()
	//benchmark.TestSimulatedAnnealingMinimalTemperatures()
	//benchmark.TestSimulatedAnnealingEpochs()
	//benchmark.TestSimulatedAnnealingCoolingRates()

	benchmark.TestTabuSearchIterations()
	benchmark.TestTabuSearchMoveTypes()
	benchmark.TestTabuSearchTenures()
}
