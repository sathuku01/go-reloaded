package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"strings"
	hc "go_reloaded/pkg/capitalise"
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

	splitted := strings.Fields(text)
	for i := 0; i < len(splitted); i++ {
		if splitted[i] == "(cap)" {
			splitted[i-1] = hc.CapsFirstLetter(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(cap,"{
			iter := hc.ExrtractNumFromString(splitted[i+1])
			for j := 1; j <= iter; j++ {
				splitted[i-j] = hc.CapsFirstLetter(splitted[i-j])
			}
			splitted = append(splitted[:i], splitted[i+2:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(up)"{

		}

	}
	fmt.Println(text)
	fmt.Println()
	fmt.Println(splitted)
}