
# unlockods

Este é um utilitário escrito em Go para remover a proteção de planilhas `.ods|.xlsz|.xls`, útil para desbloquear planilhas protegidas por senha.

## Funcionalidade

O programa extrai o conteúdo do arquivo `.ods|.xlsz|.xls`, modifica o XML interno para remover as restrições de proteção de tabela, e gera um novo arquivo `.ods|.xlsz|.xls` desbloqueado.

#

## Exemplo de Código

```go
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
```

## Limitações

- Atualmente, o foco é em arquivos `.ods`. A lógica para `.xls` está parcialmente implementada, mas não plenamente funcional.
- Não remove senhas de outras partes do documento (como arquivos criptografados com senha no nível do ZIP).

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo `LICENSE` para mais informações.
# unlock
