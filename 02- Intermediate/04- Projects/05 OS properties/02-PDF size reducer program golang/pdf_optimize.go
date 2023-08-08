/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

 
 package main

 import (
	 "fmt"
	 "log"
	 "os"
	 "time"
 
	 "github.com/unidoc/unipdf/v3/model"
	 "github.com/unidoc/unipdf/v3/model/optimize"
 )
 
 const usage = "Usage: %s INPUT_PDF_PATH OUTPUT_PDF_PATH\n"
 func main() {
	 args := os.Args
	 if len(args) < 3 {
		fmt.Println("*****")
		fmt.Println()
		fmt.Printf(usage, os.Args[0])
		fmt.Println()
		fmt.Println("*****")
		return
	 }
	 inputPath := args[1]
	 outputPath := args[2]
 
	 // Initialize starting time.
	 start := time.Now()
 
	 // Get input file stat.
	 inputFileInfo, err0 := os.Stat(inputPath)
	 if err0 != nil {
		 log.Fatalf("Fail: %v\n", err0)
	 }
 
	 // Create reader.
	 inputFile, err1 := os.Open(inputPath)
	 if err1 != nil {
		 log.Fatalf("Fail: %v\n", err1)
	 }
	 defer inputFile.Close()
 
	 reader, err2 := model.NewPdfReader(inputFile)
	 if err2 != nil {
		 log.Fatalf("Fail: %v\n", err2)
	 }
 
	 // Get number of pages in the input file.
	 pages, err3 := reader.GetNumPages()
	 if err3 != nil {
		 log.Fatalf("Fail: %v\n", err3)
	 }
 
	 // Add input file pages to the writer.
	 writer := model.NewPdfWriter()
	 for i := 1; i <= pages; i++ {
		 page, err4 := reader.GetPage(i)
		 if err4 != nil {
			 log.Fatalf("Fail: %v\n", err4)
		 }
 
		 if err4 = writer.AddPage(page); err4 != nil {
			 log.Fatalf("Fail: %v\n", err4)
		 }
	 }
 
	 // Add reader AcroForm to the writer.
	 if reader.AcroForm != nil {
		 writer.SetForms(reader.AcroForm)
	 }
 
	 // Set optimizer.
	 writer.SetOptimizer(optimize.New(optimize.Options{
		 CombineDuplicateDirectObjects:   true,
		 CombineIdenticalIndirectObjects: true,
		 CombineDuplicateStreams:         true,
		 CompressStreams:                 true,
		 UseObjectStreams:                true,
		 ImageQuality:                    80,
		 ImageUpperPPI:                   100,
	 }))
 
	 // Create output file.
	 outputFile, err6 := os.Create(outputPath)
	 if err6 != nil {
		 log.Fatalf("Fail: %v\n", err6)
	 }
	 defer outputFile.Close()
 
	 // Write output file.
	 err6 = writer.Write(outputFile)
	 if err6 != nil {
		 log.Fatalf("Fail: %v\n", err6)
	 }
 
	 // Get output file stat.
	 outputFileInfo, err7 := os.Stat(outputPath)
	 if err7 != nil {
		 log.Fatalf("Fail: %v\n", err7)
	 }
 
	 // Print basic optimization statistics.
	 inputSize := inputFileInfo.Size()
	 outputSize := outputFileInfo.Size()
	 ratio := 100.0 - (float64(outputSize) / float64(inputSize) * 100.0)
	 duration := float64(time.Since(start)) / float64(time.Millisecond)
 
	 fmt.Printf("Original file: %s\n", inputPath)
	 fmt.Printf("Original size: %d bytes\n", inputSize)
	 fmt.Printf("Optimized file: %s\n", outputPath)
	 fmt.Printf("Optimized size: %d bytes\n", outputSize)
	 fmt.Printf("Compression ratio: %.2f%%\n", ratio)
	 fmt.Printf("Processing time: %.2f ms\n", duration)
 }





//  Resome golang developer.pdf



/*
 * PDF optimization (compression) example.
 *
 * Run as: go run pdf_optimize.go <input.pdf> <output.pdf>
 */