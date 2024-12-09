package menu

import (
	"fmt"
	"log"
	"pea2/atsp"
	"pea2/benchmark"
	"pea2/generator"
	"pea2/graph"
	"pea2/utils"
)

type TabuOpts struct {
	Tenure        int
	MaxIterations int
	MoveType      atsp.MovingMethod
}

type AnnealingOpts struct {
	CoolingRate        float64
	Epochs             int
	MinimalTemperature float64
	InitialTemperature float64
}

type Options struct {
	Graph graph.Graph

	Tabu      TabuOpts
	Annealing AnnealingOpts
}

var opts = Options{}

func RunMenu() {
	for {
		PrintOptions()
		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			log.Println(utils.RedColor(err))
			continue
		}
		switch choice {
		case 0:
			return
		case 1:
			ReadGraph()
		case 2:
			if opts.Graph == nil {
				log.Println(utils.RedColor("[!!] Nie wczytano grafu"))
				continue
			}
			fmt.Println(opts.Graph.ToString())
		case 3:
			RunAlgorithm()
		case 4:
			GenerateGraph()
		case 5:
			SetTabuOptions()
		case 6:
			SetSimulatedAnnealingOptions()
		default:
			log.Println(utils.RedColor("[!!] Tylko opcje 0-6"))
		}
	}
}

func PrintOptions() {
	fmt.Println("Wybierz opcję:")
	fmt.Println("0. Wyjście")
	fmt.Println("1. Wczytaj graf z pliku")
	fmt.Println("2. Wyświetl graf")
	fmt.Println("3. Wykonaj algorytm rozwiązywania ATSP")
	fmt.Println("4. Wygeneruj graf")
	fmt.Println("5. Ustaw parametry Tabu Search")
	fmt.Println("6. Ustaw parametry Simulated Annealing")
	fmt.Print("> ")
}

func ReadGraph() {
	var filename string
	fmt.Print("Podaj ścieżkę do pliku: ")
	var err error
	if _, err = fmt.Scanln(&filename); err != nil {
		log.Println(err)
		return
	}

	opts.Graph, err = utils.ReadGraphFromFile(filename)
	if err != nil {
		log.Println(err)
	}
}

func RunAlgorithm() {
	if opts.Graph == nil {
		log.Println(utils.RedColor("[!!] Nie wczytano grafu"))
		return
	}

	fmt.Println("Wybierz algorytm:")
	fmt.Println("0. Tabu Search")
	fmt.Println("1. Simulated Annealing")
	fmt.Print("> ")

	var choice int
	var tsp atsp.ATSP
	var prompt string

	if _, err := fmt.Scanln(&choice); err != nil {
		log.Println(utils.RedColor(err))
		return
	}
	switch choice {
	case 0:
		tsp, prompt = atsp.NewTabuSearchSolver(opts.Graph, opts.Tabu.Tenure, opts.Tabu.MaxIterations, opts.Tabu.MoveType), fmt.Sprintf("Tabu Search (%v)", opts.Tabu)
	case 1:
		tsp, prompt = atsp.NewSimulatedAnnealingSolver(opts.Graph, opts.Annealing.CoolingRate, opts.Annealing.MinimalTemperature, opts.Annealing.InitialTemperature, opts.Annealing.Epochs), fmt.Sprintf("Simulated Annealing (%v)", opts.Annealing)
	default:
		log.Println(utils.RedColor("[!!] Tylko opcje 0-2"))
		return
	}

	benchmark.MeasureSolveTime(tsp, prompt)
}

func SetTabuOptions() {
	fmt.Println("Ustawianie opcji dla Tabu Search")
	fmt.Print("Podaj kadencję: ")
	var tenure int
	if _, err := fmt.Scanln(&tenure); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	fmt.Print("Podaj liczbę iteracji: ")
	var maxIter int
	if _, err := fmt.Scanln(&maxIter); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	fmt.Println("Wybierz typ ruchu:")
	fmt.Println("0. swap")
	fmt.Println("1. insert")
	fmt.Print("> ")
	var choice int
	if _, err := fmt.Scanln(&choice); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	var moveType atsp.MovingMethod
	switch choice {
	case 0:
		moveType = atsp.MovingSwap
	case 1:
		moveType = atsp.MovingInsert
	default:
		log.Println(utils.RedColor("[!!] Tylko opcje 0-1"))
		return
	}

	opts.Tabu = TabuOpts{Tenure: tenure, MaxIterations: maxIter, MoveType: moveType}

}

func SetSimulatedAnnealingOptions() {
	fmt.Println("Ustawianie opcji dla Simulated Annealing")
	fmt.Print("Podaj współczynnik chłodzenia (alfa): ")
	var coolingRate float64
	if _, err := fmt.Scanln(&coolingRate); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	fmt.Print("Podaj liczbę epok: ")
	var epochs int
	if _, err := fmt.Scanln(&epochs); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	fmt.Print("Podaj minimalną temperaturę: ")
	var minTemp float64
	if _, err := fmt.Scanln(&minTemp); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	fmt.Print("Podaj początkową temperaturę: ")
	var initTemp float64
	if _, err := fmt.Scanln(&initTemp); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	opts.Annealing = AnnealingOpts{CoolingRate: coolingRate, Epochs: epochs, MinimalTemperature: minTemp, InitialTemperature: initTemp}
}

func GenerateGraph() {
	fmt.Print("Podaj liczbę miast: ")
	var cities int
	if _, err := fmt.Scanln(&cities); err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	g, err := generator.GenerateAdjacencyMatrix(cities)
	if err != nil {
		log.Println(utils.RedColor(err))
		return
	}

	opts.Graph = g
}
