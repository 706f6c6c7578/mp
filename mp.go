package main

import (
	"fmt"
	"io"
	"os"
)

func padText(text []byte) []byte {
	length := len(text)
	paddingNeeded := 5 - (length % 5)
	if paddingNeeded == 5 {
		return text
	}
	
	paddedText := make([]byte, length+paddingNeeded)
	copy(paddedText, text)
	for i := length; i < len(paddedText); i++ {
		paddedText[i] = 'X' // Padding with uppercase 'X'
	}
	return paddedText
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: mp < infile > outfile")
	fmt.Fprintln(os.Stderr, "\nThis program pads UTF-8 text to the nearest multiple of 5 bytes using 'X'.")
	fmt.Fprintln(os.Stderr, "It reads from stdin and writes to stdout, allowing use in pipelines.")
}

func main() {
	// Check for -h flag or no input
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		printUsage()
		return
	}

	// Check if there's input available
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		// No input provided
		printUsage()
		return
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	paddedContent := padText(input)

	_, err = os.Stdout.Write(paddedContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Data successfully padded.\n")
	fmt.Fprintf(os.Stderr, "Original size: %d bytes\n", len(input))
	fmt.Fprintf(os.Stderr, "Padded size: %d bytes\n", len(paddedContent))
}
