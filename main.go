package main

import (
	// "strings"
	"strings"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func CalculateNumberOfCharacters(path string) int {
	f, err := os.ReadFile(path)
	check(err)

	r := []rune(string(f))

	return len(r)
}

func formatToString(x int, largest int) string{
	s := strconv.Itoa(x)
	cap := len([]rune(strconv.Itoa(largest))) - len([]rune(s))
	for range cap {
		s = " " + s
	}
	return s
}



func main() {
	bytePtr := flag.Bool("c", false, "print the byte counts")
	linePtr := flag.Bool("l", false, "print the line counts")
	wordPtr := flag.Bool("w", false, "print the word counts")
	charPtr := flag.Bool("m", false, "print the char counts")

	flag.Parse()

	filePaths := flag.Args()

	flags := *bytePtr || *linePtr || *wordPtr || *charPtr
	
	commands := make(map[string][]int)
	// total := make([]int, 0)

	var l, c, w int

	var largest int
	for _, v := range filePaths {
		f, err := filepath.Abs(v)
		check(err)

		if !flags {
			l = CalculateNumberOfLines(f)
			w = CalculateNumberOfWords(f)
			c = CalculateNumberOfBytes(f)
			

			largest = max(l, c, w, largest)

			commands[v] = []int{l, w, c}
		}		
	}

	var s strings.Builder
	for k, v := range commands {
		for _, j  := range v {
			s .WriteString(formatToString(j, largest) + " ")
		}

		s .WriteString(k + "\n")
	}
	
	fmt.Printf("%s", s.String())
}
