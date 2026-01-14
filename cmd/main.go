package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"strings"
	sc "go_reloaded/pkg/capitalise"
	hc "go_reloaded/pkg/baseConversion"
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
			splitted[i-1] = sc.CapsFirstLetter(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(cap,"{
			iter := sc.ExrtractNumFromString(splitted[i+1])
			if iter > i{
				log.Fatal(err)
				return
			}
			for j := 1; j <= iter; j++ {
				splitted[i-j] = sc.CapsFirstLetter(splitted[i-j])
			}
			splitted = append(splitted[:i], splitted[i+2:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(up)"{
			splitted[i-1] = sc.ToUpper(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(up,"{
			iter := sc.ExrtractNumFromString(splitted[i+1])
			if iter > i{
				log.Fatal(err)
				return
			}
			for j := 1; j <= iter; j++ {
				splitted[i-j] = sc.ToUpper(splitted[i-j])
			}
			splitted = append(splitted[:i], splitted[i+2:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(low)"{
			splitted[i-1] = sc.ToLower(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(low,"{
			iter := sc.ExrtractNumFromString(splitted[i+1])
			if iter > i{
				log.Fatal(err)
				return
			}
			for j := 1; j <= iter; j++ {
				splitted[i-j] = sc.ToLower(splitted[i-j])
			}
			splitted = append(splitted[:i], splitted[i+2:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(hex)" && i > 0 {
			splitted[i-1] = hc.HexTodec(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(bin)" && i > 0 {
			splitted[i-1] = hc.BinaryToDecimal(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
			strings.Join(splitted, " ")
		}

	}
	
	finalResult := ""
	for i, char := range splitted {
		if i == 0{
			finalResult += char
		}else {
			finalResult += " " + char
		}
	}
	fmt.Println(text)
	fmt.Println()
	fmt.Println(finalResult)
}