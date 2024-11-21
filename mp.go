package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var paddingSize int

func getRandomLetter() byte {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return letters[rand.Intn(len(letters))]
}

func padText(text []byte, paddingSize int) []byte {
	length := len(text)
	paddingNeeded := paddingSize - (length % paddingSize)
	if paddingNeeded == paddingSize {
		return text
	}
	
	paddedText := make([]byte, length+paddingNeeded)
	copy(paddedText, text)
	for i := length; i < len(paddedText); i++ {
		paddedText[i] = getRandomLetter() // Padding with random uppercase letters
	}
	return paddedText
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: mp [-p 4] < infile > outfile")
	fmt.Fprintln(os.Stderr, "\nThis program pads UTF-8 text to the nearest multiple of 5 (default) or 4 bytes using random uppercase letters.")
	fmt.Fprintln(os.Stderr, "It reads from stdin and writes to stdout, allowing use in pipelines.")
	fmt.Fprintln(os.Stderr, "\nFlags:")
	fmt.Fprintln(os.Stderr, "  -p 4\tPad to multiples of 4 instead of 5")
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed
	
	flag.IntVar(&paddingSize, "p", 5, "Padding size (4 or 5)")
	flag.Parse()

	if paddingSize != 4 && paddingSize != 5 {
		fmt.Fprintln(os.Stderr, "Invalid padding size. Use 5 (default) or 4.")
		printUsage()
		os.Exit(1)
	}

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

	paddedContent := padText(input, paddingSize)

	_, err = os.Stdout.Write(paddedContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Data successfully padded.\n")
	fmt.Fprintf(os.Stderr, "Original size: %d bytes\n", len(input))
	fmt.Fprintf(os.Stderr, "Padded size: %d bytes\n", len(paddedContent))
}
