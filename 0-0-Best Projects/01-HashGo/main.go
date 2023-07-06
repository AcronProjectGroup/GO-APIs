package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	hashcatPath := "/path/to/hashcat" // Replace with the actual path to the Hashcat executable

	// Check if Hashcat exists
	if _, err := os.Stat(hashcatPath); os.IsNotExist(err) {
		fmt.Println("Hashcat not found at the specified path.")
		return
	}

	hashType := "0" // Replace with the desired hash type
	hashFile := "/path/to/hash/file" // Replace with the path to the file containing the hashes
	wordlist := "/path/to/wordlist" // Replace with the path to the wordlist file

	// Run Hashcat command
	cmd := exec.Command(hashcatPath, "-m", hashType, "-a", "0", hashFile, wordlist)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running Hashcat:", err)
		return
	}

	// Print the output
	fmt.Println(string(output))
}
