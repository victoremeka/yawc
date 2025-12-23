package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func CalculateNumberOfBytes(path string) int {
	absPath, err := filepath.Abs(path)
	check(err)

	f, err := os.ReadFile(absPath)
	check(err)

	return len(f)
}

func CalculateNumberOfLines(path string) int {

}


func main() {
	args := os.Args

	path := args[2]
	switch cliOption := args[1]; cliOption {
		case "-c":
			nob := CalculateNumberOfBytes(path)
			fmt.Printf("%d %s\n", nob, path)
	}
}
