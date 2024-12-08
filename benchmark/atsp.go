package benchmark

import (
	"fmt"
	"log"
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
