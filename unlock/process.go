package unlock

import (
	"archive/zip"
	"fmt"
	"os"
	"strings"
)

func UnlunkFile(inputPath, outputPath string) error {
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
	
	return nil
}
