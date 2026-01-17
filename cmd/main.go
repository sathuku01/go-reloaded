package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	// "strings"
	// sc "go_reloaded/pkg/capitalise"
	// hc "go_reloaded/pkg/baseConversion"
	pc "go_reloaded/process"
)

func main(){
	if len(os.Args) < 1{
		// log.Fatal()
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
	text := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lineRead := scanner.Text()
		text = lineRead + " "
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
	_, Err = output.WriteString(text)

	if Err != nil {
		log.Fatal(Err)
		return
	}
	
}