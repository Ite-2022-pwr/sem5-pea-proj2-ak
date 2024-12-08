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

type Options struct {
	Graph graph.Graph
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
		default:
			log.Println(utils.RedColor("[!!] Tylko opcje 0-3"))
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
		tsp, prompt = atsp.NewTabuSearchSolver(opts.Graph, 7, opts.Graph.GetVerticesCount()*100, atsp.MovingInsert), "Tabu Search"
	default:
		log.Println(utils.RedColor("[!!] Tylko opcje 0-2"))
		return
	}

	benchmark.MeasureSolveTime(tsp, prompt)
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
