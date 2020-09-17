package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "prac1.4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("%s\n", f.Name())
		}
	}
}
