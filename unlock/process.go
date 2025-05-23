package unlock

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func UnlunkFile(inputPath, outputPath string) error {
	isXls := false
	if strings.HasSuffix(inputPath, ".xls") {
		isXls = true
		cmd := exec.Command("libreoffice", "--headless", "--convert-to", "xlsx", inputPath)
		output, err := cmd.Output()
		fmt.Println(string(output))
		if err != nil {
			log.Fatal(err)
		}
		inputPath += "x"
	}
	// 1. Abrir arquivo ODS como ZIP
	r, err := zip.OpenReader(inputPath)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo ODS: %w", err)
	}
	defer r.Close()
	
	// 2. Criar novo arquivo ODS
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de saída: %w", err)
	}
	defer outFile.Close()
	
	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()
	var typeFile FileType
	if strings.Contains(".xls", inputPath) {
		typeFile = EXCELTYPE
	} else if strings.Contains(".ods", inputPath) {
		typeFile = ODSTYPE
	}
	// 3. Processar cada arquivo dentro do ODS
	
	for _, file := range r.File {
		
		if err := processFileInZIP(file, zipWriter, typeFile); err != nil {
			return fmt.Errorf("erro ao processar arquivo %s: %w", file.Name, err)
		}
	}
	
	if isXls {
		_ = os.Remove(inputPath)
		cmd := exec.Command("libreoffice", "--headless", "--convert-to", "xls", outputPath)
		output, err := cmd.Output()
		fmt.Println(string(output))
		if err != nil {
			log.Fatal(err)
		}
		
	}
	
	return nil
}
