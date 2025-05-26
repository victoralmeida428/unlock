
# unlockods

Este é um utilitário escrito em Go para remover a proteção de planilhas `.ods|.xlsz|.xls`, útil para desbloquear planilhas protegidas por senha.

## Funcionalidade

O programa extrai o conteúdo do arquivo `.ods|.xlsz|.xls`, modifica o XML interno para remover as restrições de proteção de tabela, e gera um novo arquivo `.ods|.xlsz|.xls` desbloqueado.

#

## Exemplo de Código

```go
package main

import (
	"bufio"
	"fmt"
	"github.com/victoralmeida428/unlock/unlock"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	systemPath, inputFolder, outputFolder string
)

func init() {
	var err error
	systemPath, err = os.Getwd()
	if err != nil {
		log.Fatalf("Erro ao obter diretório atual: %v", err)
	}
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Pasta com os arquivos (default: input): ")
	inputFolder, _ = reader.ReadString('\n')
	inputFolder = strings.TrimSpace(inputFolder)
	
	fmt.Print("Pasta de saída (default: output): ")
	outputFolder, _ = reader.ReadString('\n')
	outputFolder = strings.TrimSpace(outputFolder)
	
	// Valores padrão
	if inputFolder == "" {
		inputFolder = "input"
	}
	if outputFolder == "" {
		outputFolder = "output"
	}
	
	// Criar pastas se não existirem
	if err = os.MkdirAll(inputFolder, 0755); err != nil {
		log.Fatalf("Erro ao criar pasta de entrada: %v", err)
	}
	if err = os.MkdirAll(outputFolder, 0755); err != nil {
		log.Fatalf("Erro ao criar pasta de saída: %v", err)
	}
	
	fmt.Printf("\nConfigurações:\n- Pasta de entrada: %s\n- Pasta de saída: %s\n\n",
		filepath.Join(systemPath, inputFolder),
		filepath.Join(systemPath, outputFolder))
}

func main() {
	
	if err := os.MkdirAll(filepath.Join(systemPath, outputFolder), 0755); err != nil {
		log.Printf("Erro ao criar pasta de saída: %v", err)
	}
	
	// Get all files in input folder
	files, err := os.ReadDir(filepath.Join(systemPath, inputFolder))
	if err != nil {
		log.Fatalf("Erro ao ler pasta de entrada: %v", err)
	}
	
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		
		ext := filepath.Ext(file.Name())
		if isSupportedFile(ext) {
			inputPath := filepath.Join(systemPath, inputFolder, file.Name())
			outputPath := filepath.Join(systemPath, outputFolder, getOutputFilename(file.Name()))
			
			fmt.Println("Processando arquivo: " + inputPath)
			if err = unlock.UnlunkFile(inputPath, outputPath); err != nil {
				log.Printf("Erro ao processar arquivo %s: %v", inputPath, err)
				time.Sleep(5 * time.Second)
				continue
			}
			
			fmt.Printf("Arquivo processado com sucesso: %s\n", outputPath)
		}
	}
	
	fmt.Println("Processamento concluído!")
	time.Sleep(5 * time.Second)
}

func isSupportedFile(extension string) bool {
	switch extension {
	case ".xls", ".xlsx", ".ods":
		return true
	default:
		return false
	}
}

func getOutputFilename(inputPath string) string {
	ext := filepath.Ext(inputPath)
	base := inputPath[:len(inputPath)-len(ext)]
	return base + "_unprotected" + ext
}

```

## Limitações

- Atualmente, o foco é em arquivos `.ods`. A lógica para `.xls` está parcialmente implementada, mas não plenamente funcional.
- Não remove senhas de outras partes do documento (como arquivos criptografados com senha no nível do ZIP).

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo `LICENSE` para mais informações.
# unlock
