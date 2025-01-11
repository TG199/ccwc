package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func countFileStats(fname string) (int, int, int, error) {
	var lc, wc, cc int

	file, err := os.Open(fname)

	if err != nil {
		return 0, 0, 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lc++
		wc += len(strings.Fields(line))
		cc += len(line)
	}

	if scanner.Err() != nil {
		return 0, 0, 0, nil
	}

	return lc, wc, cc, nil
}

func lineCount(fname string) {
	lc, _, _, err := countFileStats(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("%4d %s\n", lc, fname)
}

func wordCount(fname string) {
	_, wc, _, err := countFileStats(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%4d %s\n", wc, fname)
}

func characterCount(fname string) {
	_, _, cc, err := countFileStats(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%4d %s\n", cc, fname)
}
func processFile(fname string) {
	lc, wc, cc, err := countFileStats(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%4d %4d %4d %s\n", lc, wc, cc, fname)
}


func main(){
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: wc [-l|-w|-c <file> or ccwc <file1> <file2> ...")
		os.Exit(1)
	}

	argType := os.Args[1]
	switch argType {
	case "-l":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Missing filename after -l")
			os.Exit(1)
		}
		lineCount(os.Args[2])
	case "-w":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Missing filename after -w")
			os.Exit(1)
		}
		wordCount(os.Args[2])
	case "-c":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Missing filename after -c")
		}
		characterCount(os.Args[2])
	default:
		for _, fname := range os.Args[1:] {
			processFile(fname)
		}
	}
}
