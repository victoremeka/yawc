package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CalculateNumberOfBytes(path string) int {
	f, err := os.ReadFile(path)
	check(err)

	return len(f)
}

func CalculateNumberOfLines(path string) int {
	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}

func CalculateNumberOfWords(path string) int {
	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}

func main() {
	args := os.Args
	path := args[2]
	absPath, err := filepath.Abs(path)
	var res int
	switch cliOption := args[1]; cliOption {
	case "-c":
		res = CalculateNumberOfBytes(absPath)
		check(err)
	case "-l":
		res = CalculateNumberOfLines(absPath)
	case "-w":
		res = CalculateNumberOfWords(absPath)
	}
	fmt.Printf("%d %s\n", res, path)
}
