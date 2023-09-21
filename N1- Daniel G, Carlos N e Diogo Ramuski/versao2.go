package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entrega2 struct {
	Horario int
	Destino string
	Bonus   int
}

type grafico2 struct {
	Conexoes map[string]map[string]int
}

func novoGrafo2() *grafico2 {
	return &grafico2{Conexoes: make(map[string]map[string]int)}
}

func (g *grafico2) addConexao2(origem, destino string, tempo int) {
	if _, ok := g.Conexoes[origem]; !ok {
		g.Conexoes[origem] = make(map[string]int)
	}
	g.Conexoes[origem][destino] = tempo
}

func carregarDestinos2(filename string) *grafico2 {
	grafo := novoGrafo2()

	arquivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de destinos:", err)
		os.Exit(1)
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		partes := strings.Split(linha, ",")
		if len(partes) == 3 {
			origem := partes[0]
			destino := partes[1]
			tempo, err := strconv.Atoi(partes[2])
			if err != nil {
				fmt.Println("Erro ao analisar o tempo:", err)
				os.Exit(1)
			}
			grafo.addConexao2(origem, destino, tempo)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo de destinos:", err)
		os.Exit(1)
	}

	return grafo
}

func carregarEntregas2(filename string) []entrega2 {
	var entregas []entrega2

	arquivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de entregas:", err)
		os.Exit(1)
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		partes := strings.Split(linha, ",")
		if len(partes) == 3 {
			horario, err := strconv.Atoi(partes[0])
			if err != nil {
				fmt.Println("Erro ao analisar o horário:", err)
				os.Exit(1)
			}
			destino := partes[1]
			bonus, err := strconv.Atoi(partes[2])
			if err != nil {
				fmt.Println("Erro ao analisar o bônus:", err)
				os.Exit(1)
			}
			entregas = append(entregas, entrega2{Horario: horario, Destino: destino, Bonus: bonus})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo de entregas:", err)
		os.Exit(1)
	}

	return entregas
}

type estado2 struct {
	TempoAtual   int
	DestinoAtual string
	Visitados    map[string]bool
	Lucro        int
	Caminho      []entrega2
}

func calcularLucro2(entregas []entrega2, grafo *grafico2) (int, []entrega2) {
	sort.Slice(entregas, func(i, j int) bool {
		return entregas[i].Horario < entregas[j].Horario
	})

	n := len(entregas)
	dp := make([]int, n)
	anterior := make([]int, n)

	for i := range dp {
		dp[i] = entregas[i].Bonus
		anterior[i] = -1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			tempoViagem := grafo.Conexoes[entregas[j].Destino][entregas[i].Destino]
			tempoRetorno := grafo.Conexoes[entregas[i].Destino]["A"]
			tempoTotal := tempoViagem + tempoRetorno

			if entregas[i].Horario >= entregas[j].Horario+tempoTotal {
				if dp[i] < dp[j]+entregas[i].Bonus {
					dp[i] = dp[j] + entregas[i].Bonus
					anterior[i] = j
				}
			}
		}
	}

	lucroMaximo := 0
	indice := -1
	for i, lucro := range dp {
		if lucro > lucroMaximo {
			lucroMaximo = lucro
			indice = i
		}
	}

	sequencia := []entrega2{}
	for indice != -1 {
		sequencia = append([]entrega2{entregas[indice]}, sequencia...)
		indice = anterior[indice]
	}

	return lucroMaximo, sequencia
}

func main() {
	arquivoDestinos := "destinos.txt"
	arquivoEntregas := "entregas.txt"

	grafo := carregarDestinos2(arquivoDestinos)
	entregas := carregarEntregas2(arquivoEntregas)

	lucroMaximo, sequencia := calcularLucro2(entregas, grafo)

	fmt.Println("Sequência de Entregas Programadas:")
	for _, entrega := range sequencia {
		fmt.Printf("Entrega em %s - Tempo: %d minutos - Bônus: %d\n", entrega.Destino, grafo.Conexoes[entrega.Destino]["A"]+grafo.Conexoes["A"][entrega.Destino], entrega.Bonus)
	}
	fmt.Printf("Lucro Total Esperado: %d\n", lucroMaximo)
}
