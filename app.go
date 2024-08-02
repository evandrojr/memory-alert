package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"strconv"
	"strings"

	"os"

	"github.com/martinlindhe/notify"
)

const MinimaMemoriaDisponívelPadrao = 2000

func main() {

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Erro ao obter o caminho do executável:", err)
		return
	}
	exeDir := filepath.Dir(exePath)
	fmt.Println(exeDir)

	MinimaMemoriaDisponível := MinimaMemoriaDisponívelPadrao
	if len(os.Args) > 1 {
		MinimaMemoriaDisponível, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Printf("erro ao converter valor da memória mínima: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("Iniciando com mínima memória disponível de: " + strconv.Itoa(MinimaMemoriaDisponível) + "KiB")
		notify.Notify("memory-alert", "Memory Alert", "Mem min disp: "+strconv.Itoa(MinimaMemoriaDisponível)+"KiB", exeDir+"/ram.png")
	} else {
		fmt.Println("Iniciando com o padrão da mínima memória disponível de: " + strconv.Itoa(MinimaMemoriaDisponívelPadrao) + "KiB, passe como parâmetro para sobrescrevê-lo")
		notify.Notify("memory-alert", "Memory Alert", "Mem min disp padrão: "+strconv.Itoa(MinimaMemoriaDisponívelPadrao)+"KiB, passe como parâmetro para sobrescrevê-lo", exeDir+"/ram.png")
	}

	for {

		memAvailable, err := MemAvailableMB()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Memória disponível: %d MB\n", memAvailable)
		}

		if memAvailable < MinimaMemoriaDisponível {
			notify.Notify("memory-alert", "ALERTA!", "A memória tá lascada!", exeDir+"/ram.png")
		}
		time.Sleep(2 * time.Second)
	}
}

func MemAvailableMB() (int, error) {

	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("erro ao ler /proc/meminfo: %w", err)
	}
	meminfo := string(data)
	for _, line := range strings.Split(meminfo, "\n") {
		if strings.HasPrefix(line, "MemAvailable:") {
			parts := strings.Fields(line)
			memAvailable, err := strconv.Atoi(parts[1])
			if err != nil {
				return 0, fmt.Errorf("erro ao converter valor: %w", err)
			}
			memAvailableMB := memAvailable / 1024
			return memAvailableMB, nil
		}
	}
	return 0, fmt.Errorf("MemAvailable não encontrado em /proc/meminfo")
}
