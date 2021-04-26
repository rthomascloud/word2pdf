package main

import (
	"fmt"
	"os"

	officel "github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/document/convert"
	pdfl "github.com/unidoc/unipdf/v3/common/license"
)

func setLicenseKey() error {
	apiKey := os.Getenv("UNICLOUD_METERED_KEY")
	if apiKey == "" {
		return fmt.Errorf("Missing UNICLOUD_METERED_KEY environment variable")
	}

	err := officel.SetMeteredKey(apiKey)
	if err != nil {
		return err
	}

	err = pdfl.SetMeteredKey(apiKey)
	if err != nil {
		return err
	}

	return nil
}

func wordToPDF(inputPath string, outputPath string) error {
	doc, err := document.Open(inputPath)
	if err != nil {
		return err
	}
	defer doc.Close()
	c := convert.ConvertToPdf(doc)

	err = c.WriteToFile(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <inputfile.docx> <outputfile.pdf>\n", os.Args[0])
		os.Exit(1)
	}

	err := setLicenseKey()
	if err != nil {
		fmt.Printf("Failed to load license: %v\n", err)
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	err = wordToPDF(inputPath, outputPath)
	if err != nil {
		fmt.Printf("Error extracting text from %s : %v\n", inputPath, err)
		os.Exit(1)
	}
}
