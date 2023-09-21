package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entrega1 struct {
	Horario int
	Destino string
	Bonus   int
}

type grafico1 struct {
	Conexoes map[string]map[string]int
}

func novoGrafo1() *grafico1 {
	return &grafico1{Conexoes: make(map[string]map[string]int)}
}

func (g *grafico1) addConexao1(origem, destino string, tempo int) {
	if _, ok := g.Conexoes[origem]; !ok {
		g.Conexoes[origem] = make(map[string]int)
	}
	g.Conexoes[origem][destino] = tempo
}

func carregarDestinos1(arquivo string) *grafico1 {
	grafo := novoGrafo1()

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
			grafo.addConexao1(origem, destino, tempo)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo de destinos:", err)
		os.Exit(1)
	}

	return grafo
}

func carregarEntregas1(arquivo string) []entrega1 {
	var entregas []entrega1

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
			entregas = append(entregas, entrega1{Horario: horario, Destino: destino, Bonus: bonus})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo de entregas:", err)
		os.Exit(1)
	}

	return entregas
}

type estado1 struct {
	TempoAtual         int
	DestinoAtual       string
	EntregasRealizadas []entrega1
}

func leilaoEntregas1(grafo *grafico1, entregas []entrega1) ([]entrega1, int) {
	sort.Slice(entregas, func(i, j int) bool {
		return entregas[i].Horario < entregas[j].Horario
	})

	sequenciaEntregas := []entrega1{}
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

func leitura1() {
	arquivoDestinos := "destinos.txt"
	arquivoEntregas := "entregas.txt"

	grafo := carregarDestinos1(arquivoDestinos)
	entregas := carregarEntregas1(arquivoEntregas)

	sequenciaEntregas, lucroTotal := leilaoEntregas1(grafo, entregas)

	fmt.Println("Sequência de Entregas Programadas:")
	for _, entrega := range sequenciaEntregas {
		fmt.Printf("Entrega em %s - Tempo: %d minutos - Bônus: %d\n", entrega.Destino, grafo.Conexoes[entrega.Destino]["A"]+grafo.Conexoes["A"][entrega.Destino], entrega.Bonus)
	}
	fmt.Printf("Lucro Total Esperado: %d\n", lucroTotal)
}

func main() {
	leitura1()
}
