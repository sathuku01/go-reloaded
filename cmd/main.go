package main

import (
	"bufio"
	"fmt"
	pc "go_reloaded/process"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Error: Empty field")
	}
	path := string(os.Args[1])

	// Open the file.
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// Read the contents of the file.
	text := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	output, Err := os.Create("result.txt")

	if Err != nil {
		log.Fatal(Err)
		return
	}

	text = pc.ProcessData(text)
	_, Err = output.WriteString(strings.Join(text, " "))

	if Err != nil {
		log.Fatal(Err)
		return
	}

}
