package process

import (
	"strings"
	"log"
	"errors"
	sc "go_reloaded/pkg/capitalise"
	hc "go_reloaded/pkg/baseConversion"

)

func ProcessData(strSlice string) string{
	splitted := strings.Fields(strSlice)
		for i := 0; i < len(splitted); i++ {
		if splitted[i] == "(cap)" {
			splitted[i-1] = sc.CapsFirstLetter(splitted[i-1])
			splitted = append(splitted[:i], splitted[i+1:]...)
			strings.Join(splitted, " ")
		}else if splitted[i] == "(cap,"{
			iter := sc.ExrtractNumFromString(splitted[i+1])
			if iter > i{
				log.Fatal(errors.New("Lookback period is larger than start of sentence"))
				// return
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
				log.Fatal(errors.New("Lookback period is larger than start of sentence"))
				
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
				log.Fatal(errors.New("Lookback period is larger than start of sentence"))
				// return
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
			strings.Join(splitted[:], " ")
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
	return finalResult
	
}