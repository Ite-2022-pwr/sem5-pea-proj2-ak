package benchmark

import (
	"fmt"
	"log"
	"math"
	"pea2/atsp"
	"pea2/utils"
	"time"
)

func MeasureSolveTime(tsp atsp.ATSP, prompt string) float64 {
	log.Println("[*] Rozpoczynanie:", prompt)
	start := time.Now()

	cost, path := tsp.Solve(0)
	solveTime := utils.PrintTimeElapsed(start, prompt)

	log.Println(utils.YellowColor(fmt.Sprintf("[+] Koszt: %d", cost)))
	log.Println(utils.YellowColor(fmt.Sprintf("[+] Ścieżka: %v", path)))

	return solveTime
}

func MeasureSolveTimeWithCost(tsp atsp.ATSP, prompt string) (float64, int) {
	log.Println("[*] Rozpoczynanie:", prompt)
	start := time.Now()

	cost, path := tsp.Solve(0)
	solveTime := utils.PrintTimeElapsed(start, prompt)

	log.Println(utils.YellowColor(fmt.Sprintf("[+] Koszt: %d", cost)))
	log.Println(utils.YellowColor(fmt.Sprintf("[+] Ścieżka: %v", path)))

	return solveTime, cost
}

func CalculateError(solution, optimalSolution int) int {
	return int(math.Round(math.Abs(float64(solution-optimalSolution)) * 100 / float64(optimalSolution)))
}
