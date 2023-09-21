package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entrega3 struct {
	Horario int
	Destino string
	Bonus   int
}

type grafico3 struct {
	Conexoes map[string]map[string]int
}

func novoGrafo3() *grafico3 {
	return &grafico3{Conexoes: make(map[string]map[string]int)}
}

func (g *grafico3) AdicionarConexao(origem, destino string, tempo int) {
	if _, ok := g.Conexoes[origem]; !ok {
		g.Conexoes[origem] = make(map[string]int)
	}
	g.Conexoes[origem][destino] = tempo
}

func carregarDestinos3(arquivo string) *grafico3 {
	grafo := novoGrafo3()

	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de destinos:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := scanner.Text()
		parts := strings.Split(linha, ",")
		if len(parts) == 3 {
			origem := parts[0]
			destino := parts[1]
			tempo, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Erro ao analisar o tempo:", err)
				os.Exit(1)
			}
			grafo.AdicionarConexao(origem, destino, tempo)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo de destinos:", err)
		os.Exit(1)
	}

	return grafo
}

func carregarEntregas3(arquivo string) []entrega3 {
	var entregas []entrega3

	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de entregas:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := scanner.Text()
		parts := strings.Split(linha, ",")
		if len(parts) == 3 {
			horario, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Erro ao analisar o horário:", err)
				os.Exit(1)
			}
			destino := parts[1]
			bonus, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Erro ao analisar o bônus:", err)
				os.Exit(1)
			}
			entregas = append(entregas, entrega3{Horario: horario, Destino: destino, Bonus: bonus})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo de entregas:", err)
		os.Exit(1)
	}

	return entregas
}

func leilaoEntregas3(grafo *grafico3, entregas []entrega3) ([]entrega3, int) {
	sort.Slice(entregas, func(i, j int) bool {
		return entregas[i].Horario < entregas[j].Horario
	})

	sequenciaEntregas := []entrega3{}
	lucroTotal := 0
	tempoAtual := 0

	for _, entrega := range entregas {
		if entrega.Horario >= tempoAtual {
			tempoEntrega := grafo.Conexoes["A"][entrega.Destino]
			tempoRetorno := grafo.Conexoes[entrega.Destino]["A"]
			tempoTotal := tempoEntrega + tempoRetorno

			if tempoAtual+tempoTotal <= entrega.Horario {
				sequenciaEntregas = append(sequenciaEntregas, entrega)
				lucroTotal += entrega.Bonus
				tempoAtual += tempoTotal
			}
		}
	}

	return sequenciaEntregas, lucroTotal
}

func main() {
	arquivoDestinos := "destinos.txt"
	arquivoEntregas := "entregas.txt"

	grafo := carregarDestinos3(arquivoDestinos)
	entregas := carregarEntregas3(arquivoEntregas)

	sequenciaEntregas, lucroTotal := leilaoEntregas3(grafo, entregas)

	fmt.Println("Sequência de Entregas Programadas:")
	for _, entrega := range sequenciaEntregas {
		fmt.Printf("Entrega em %s - Tempo: %d minutos - Bônus: %d\n", entrega.Destino, grafo.Conexoes[entrega.Destino]["A"]+grafo.Conexoes["A"][entrega.Destino], entrega.Bonus)
	}
	fmt.Printf("Lucro Total Esperado: %d\n", lucroTotal)
}
