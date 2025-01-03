package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filename, err)
	}
	return string(content), nil
}

func convertMarkdownToHTML(markdownContent string) string {
	html := blackfriday.Run([]byte(markdownContent))
	return string(html)
}

func writeHTMLFile(htmlContent string, outputFile string) error {
	err := ioutil.WriteFile(outputFile, []byte(htmlContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing HTML file %s: %w", outputFile, err)
	}
	fmt.Println("\nResume saved to", outputFile)
	return nil

}
func convertHTMLToPDF(inputHTML string, outputPDF string) error {

	// Create a temporary HTML file to pass to wkhtmltopdf
	tempHTMLFile, err := ioutil.TempFile("", "temp-html-*.html")
	if err != nil {
		return fmt.Errorf("error creating temporary HTML file: %w", err)
	}
	defer os.Remove(tempHTMLFile.Name()) // Clean up the temp file

	if _, err := tempHTMLFile.Write([]byte(inputHTML)); err != nil {
		return fmt.Errorf("error writing to temporary file: %w", err)
	}
	if err := tempHTMLFile.Close(); err != nil {
		return fmt.Errorf("error closing temp file: %w", err)
	}

	cmd := exec.Command("wkhtmltopdf", tempHTMLFile.Name(), outputPDF)
	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("error converting to PDF: %w", err)
	}

	fmt.Println("\nResume saved to", outputPDF)
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <resume in markdown>")
		os.Exit(1)
	}

	inputMarkdownFile := os.Args[1]

	if _, err := os.Stat(inputMarkdownFile); os.IsNotExist(err) {
		log.Fatalf("The file %s does not exist", inputMarkdownFile)
	}

	resumeContent, err := readFile(inputMarkdownFile)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}
	baseName := strings.TrimSuffix(inputMarkdownFile, filepath.Ext(inputMarkdownFile))
	resumeHTML := baseName + ".html"
	resumePDF := baseName + ".pdf"

	htmlContent := convertMarkdownToHTML(resumeContent)

	if err := writeHTMLFile(htmlContent, resumeHTML); err != nil {
		log.Fatalf("Error writing HTML file: %v", err)
	}

	if err := convertHTMLToPDF(htmlContent, resumePDF); err != nil {
		log.Fatalf("Error converting to PDF: %v", err)
	}

}
