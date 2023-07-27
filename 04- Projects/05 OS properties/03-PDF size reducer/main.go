// https://pdfcpu.io/core/optimize

package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	// "os"
)

func main() {
	// Input and output file paths
	inputPath := "input.pdf"
	outputPath := "output.pdf"

	// Reduce the size of the PDF file
	err := reducePDFSize(inputPath, outputPath)
	if err != nil {
		fmt.Println("Error reducing PDF size:", err)
		return
	}

	fmt.Println("PDF size reduced successfully!")
}

func reducePDFSize(inputPath, outputPath string) error {
	// Configuration for reducing the size
	config := pdfcpu.NewDefaultConfiguration()
	config.ValidationMode = pdfcpu.ValidationNone

	// Optimize the PDF file
	err := api.OptimizeFile(inputPath, outputPath, config)
	if err != nil {
		return err
	}

	return nil
}