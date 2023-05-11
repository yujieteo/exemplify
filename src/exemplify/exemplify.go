package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	// Read each line into a list
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Create a slice to hold the output strings
	var output []string

	// Loop through each line and substitute it into the string
	for _, line := range lines {
		output = append(output, fmt.Sprintf(`Write, in the form of LaTeX .tex source with beamer,
on the topic of %s with detailed technical jargon. 
All beamer frames have small font applied. These are the following beamer frames. 
Each frame should have appropriate descriptive mathematical formulae:
1. Important corollaries of %s
2. Counter-intuitive examples in %s
3. Special cases of %s
4. Non-examples of %s
5. Precise definition of %s
6. Concluding sentence.`, line, line, line, line, line, line))
	}

	// Open the output file for writing
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Write each output string to the file
	for _, line := range output {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	writer.Flush()

	fmt.Println("Output file generated successfully.")
}
