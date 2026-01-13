package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"strings"
	//"go_reloaded"
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

	fmt.Println(CapsFirstLetter(text))

 }

 func CapsFirstLetter(modString string ) string{

	splitted := strings.Fields(modString)
	for i := 0; i < len(splitted); i++ {
		if splitted[i] == "(cap)" {
			splitted[i-1] = Caps(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
		}

	}
	return strings.Join(splitted, " ")
 }

 func Caps(s string) string {
	// word := ""
	sByte := []byte(s)
	for i, char := range sByte{
		if i == 0 && (char >= 65 && char <= 90){
			continue
		} else if i == 0 && (char >= 97 && char <= 122) {
			sByte[i] = char - 32
		} else if char >= 65 && char <= 90 && i > 0{
			sByte[i] = char + 32
		} 
	}
	return string(sByte)
 }