package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Input and output file paths
	inputPath := "Resome golang developer.pdf"
	outputPath := "Resome golang developer123.pdf"

	// Reduce the size of the PDF file
	err := reducePDFSize(inputPath, outputPath)
	if err != nil {
		fmt.Println("Error reducing PDF size:", err)
		return
	}

	fmt.Println("PDF size reduced successfully!")
}

func reducePDFSize(inputPath, outputPath string) error {
	// Read the input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create a buffer to store the compressed data
	var compressedBuffer bytes.Buffer

	// Compress the PDF data
	zw := zlib.NewWriter(&compressedBuffer)
	_, err = io.Copy(zw, inputFile)
	if err != nil {
		return err
	}
	zw.Close()

	// Write the compressed data to the output file
	err = ioutil.WriteFile(outputPath, compressedBuffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}


