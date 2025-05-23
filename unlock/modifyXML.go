package unlock

import (
	"bytes"
	"regexp"
)

type FileType int

const (
	EXCELTYPE FileType = iota
	ODSTYPE
)

func modifyXMLContent(content []byte, filetp FileType) []byte {
	// Padrão para encontrar a tabela com proteção
	
	switch filetp {
	case EXCELTYPE:
		sheetProtectionPattern := regexp.MustCompile(
			`<sheetProtection.*?.>`)
		
		// Remover proteção de pasta de trabalho
		return sheetProtectionPattern.ReplaceAll(content, []byte{})
	
	case ODSTYPE:
		tablePattern := regexp.MustCompile(
			`
[(<table:table [^>]*table:name="[^"]+"[^>]*table:protected="true"[^>]*)(table:protection-key="[^"]+"\s*table:protection-key-digest-algorithm="[^"]+"[^>]*)?>]

	`)
		
		// Substituir o padrão encontrado
		return tablePattern.ReplaceAllFunc(content, func(match []byte) []byte {
			// Substituir protected="true" por protected="false"
			modified := bytes.Replace(match,
				[]byte(`table:protected="true"`),
				[]byte(`table:protected="false"`), 1)
			
			// Remover os atributos de proteção
			protectionPattern := regexp.MustCompile(
				`table:protection-key="[^"]+"\s*table:protection-key-digest-algorithm="[^"]+"`)
			return protectionPattern.ReplaceAll(modified, []byte{})
		})
	default:
		return content
	}
	
}
