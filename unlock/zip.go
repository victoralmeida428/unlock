package unlock

import (
	"archive/zip"
	"fmt"
	"io"
	"strings"
)

func processFileInZIP(file *zip.File, zipWriter *zip.Writer, fileType FileType) error {
	// Abrir arquivo dentro do ZIP
	
	reader, err := file.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	
	// Criar novo arquivo no ZIP de saída
	writer, err := zipWriter.Create(file.Name)
	if err != nil {
		return err
	}
	
	// Se não for content.xml, apenas copie
	if !strings.HasSuffix(file.Name, ".xml") {
		_, err = io.Copy(writer, reader)
		return err
	}
	
	// Processar content.xml
	content, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("erro ao ler content.xml: %w", err)
	}
	
	modifiedContent := modifyXMLContent(content, fileType)
	_, err = writer.Write(modifiedContent)
	return err
}
