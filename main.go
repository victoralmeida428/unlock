package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unlockods/unlock"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Uso: ./prog <arquivo.ods>")
	}
	
	inputFile := os.Args[1]
	outputFile := getOutputFilename(inputFile)
	
	if err := unlock.UnlunkFile(inputFile, outputFile); err != nil {
		log.Fatalf("Erro ao processar arquivo: %v", err)
	}
	
	fmt.Printf("Arquivo processado com sucesso! Saída: %s\n", outputFile)
}

func getOutputFilename(inputPath string) string {
	ext := filepath.Ext(inputPath)
	base := inputPath[:len(inputPath)-len(ext)]
	return base + "_unprotected" + ext
}
