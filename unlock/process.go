package unlock

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func UnlunkFile(inputPath, outputPath string) error {
	isXls := false
	if strings.HasSuffix(inputPath, ".xls") {
		isXls = true
		switch runtime.GOOS {
		case "linux":
			cmd := exec.Command("libreoffice", "--headless", "--convert-to", "xlsx", inputPath, "--outdir", filepath.Dir(inputPath))
			output, err := cmd.Output()
			fmt.Println(string(output))
			if err != nil {
				log.Printf(err.Error())
			}
		case "windows":
			sofficePath := getLibreOfficePath()
			if sofficePath != "" {
				cmd := exec.Command(sofficePath, "--headless", "--convert-to", "xlsx", inputPath, "--outdir", filepath.Dir(inputPath))
				output, err := cmd.Output()
				fmt.Println(string(output))
				if err != nil {
					log.Printf(err.Error())
				}
			}

		}

		inputPath += "x"
	}
	// 1. Abrir arquivo ODS como ZIP
	r, err := zip.OpenReader(inputPath)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
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

	}

	return nil
}

func getLibreOfficePath() string {
	// Locais comuns de instalação no Windows
	possiblePaths := []string{
		`C:\Program Files\LibreOffice\program\soffice.exe`,
		`C:\Program Files (x86)\LibreOffice\program\soffice.exe`,
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Tentar encontrar no PATH
	if path, err := exec.LookPath("soffice.exe"); err == nil {
		return path
	}

	return ""
}
