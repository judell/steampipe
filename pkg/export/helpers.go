package export

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func generateDefaultExportFileName(exporter Exporter, executionName string) string {
	now := time.Now()
	timeFormatted := fmt.Sprintf("%d%02d%02d-%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return fmt.Sprintf("%s-%s%s", executionName, timeFormatted, exporter.FileExtension())
}

func writeExport(filePath string, exportData *strings.Reader) error {
	// create the output file
	destination, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, exportData)
	return err
}