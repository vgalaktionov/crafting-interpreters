package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/vgalaktionov/crafting-interpreters/glox/scanner"
)

func main() {
	fmt.Println("Welcome to glox!")

	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Failed to read source file at path: ", path)
	}
	run(string(bytes))

}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		line := fmt.Sprintln(scanner.Text())

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
		}
		run(line)
	}
}

func run(source string) {
	s := scanner.NewScanner()
}
