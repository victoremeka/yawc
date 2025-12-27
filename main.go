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
	total := make(map[int]int, 4)

	var l, c, w, m int

	var largest int
	for _, v := range filePaths {
		f, err := filepath.Abs(v)
		check(err)

		if !flags {
			l = CalculateNumberOfLines(f)
			w = CalculateNumberOfWords(f)
			c = CalculateNumberOfBytes(f)
			
			largest = max(l, c, w, largest)

			total[0] += l
			total[1] += w
			total[2] += c

			commands[v] = []int{l, w, c}
		} else {
			commands[v] = []int{}
			idx := 0

			if *linePtr {
				l = CalculateNumberOfLines(f)
				commands[v] = append(commands[v], l)
				total[idx] += l
				idx++
			}
			if *wordPtr {
				w = CalculateNumberOfWords(f)
				commands[v] = append(commands[v], w)
				total[idx] += w
				idx++
			}
			if *bytePtr {
				c = CalculateNumberOfBytes(f)
				commands[v] = append(commands[v], c)
				total[idx] += c
				idx++
			}
			if *charPtr {
				m = CalculateNumberOfCharacters(f)
				commands[v] = append(commands[v], m)
				total[idx] += m
				idx++
			}

			largest = max(l, w, c, m, largest)
		}
	}

	var s strings.Builder
	for _, path := range filePaths {
		v := commands[path]
		for _, j  := range v {
			s.WriteString(formatToString(j, largest) + " ")
		}

		s .WriteString(path + "\n")
	}
	
	if len(filePaths) > 1 {
		for i := 0; i < len(total); i++ {
			s.WriteString(formatToString(total[i], largest) + " ")
		}
		s.WriteString("total\n")
	}
	
	fmt.Printf("%s", s.String())
}
