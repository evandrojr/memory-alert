package main

import (
	"fmt"
	"path/filepath"
	"time"

	"strconv"
	"strings"

	"os"

	"github.com/martinlindhe/notify"
)

func main() {

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Erro ao obter o caminho do executável:", err)
		return
	}

	// Obtém o diretório do executável
	exeDir := filepath.Dir(exePath)

	for {

		memAvailable, err := MemAvailableMB()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Memória disponível: %d MB\n", memAvailable)
		}

		if memAvailable < 15000 {
			notify.Notify("memory-alert", "ALERT!", "Low memory!", exeDir+"/ram.png")
		}
		time.Sleep(2 * time.Second)
	}
}

func MemAvailableMB() (int, error) {
	// Lê o arquivo /proc/meminfo
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("erro ao ler /proc/meminfo: %w", err)
	}

	// Converte o conteúdo do arquivo para string
	meminfo := string(data)

	// Procura pela linha que começa com "MemAvailable"
	for _, line := range strings.Split(meminfo, "\n") {
		if strings.HasPrefix(line, "MemAvailable:") {
			// Divide a linha em partes
			parts := strings.Fields(line)
			// Converte o valor para um inteiro
			memAvailable, err := strconv.Atoi(parts[1])
			if err != nil {
				return 0, fmt.Errorf("erro ao converter valor: %w", err)
			}
			// O valor está em KB, converte para MB
			memAvailableMB := memAvailable / 1024
			return memAvailableMB, nil
		}
	}

	return 0, fmt.Errorf("MemAvailable não encontrado em /proc/meminfo")
}
