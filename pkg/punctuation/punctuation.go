package go_reloaded


func TextCorrected(input string) []rune {
	symbols := []rune{'.', ',', '!', '?', ':', ';'}
	runes := []rune(input)
for i := 0; i < len(runes); i++ {
    for _, sym := range symbols {
        if runes[i] == sym {
            // remove spaces before punctuation
            if i-1 >= 0 && runes[i-1] == ' ' {
                runes = append(runes[:i-1], runes[i:]...)
                i-- // adjust index
            }
            // ensure one space after punctuation
            if i+1 < len(runes) && runes[i+1] != ' ' {
                runes = append(runes[:i+1], append([]rune{' '}, runes[i+1:]...)...)
            }
            break
        }
    }
}

	runes = TrimEdges(runes)

	openQuote := false

	for j := 0; j < len(runes); j++ {
		if runes[j] == '\'' {
			openQuote = !openQuote
			if openQuote && j+1 < len(runes) && runes[j+1] == ' ' {
				runes = append(runes[:j+1], runes[j+2:]...)
			}

			if !openQuote && j-1 >= 0 && runes[j-1] == ' ' {
				runes = append(runes[:j-1], runes[j:]...)
				j--
			}
		}
	}

	runes = TrimEdges(runes)

	return runes
}


// func TrimExtraSpaces(chars []rune) []rune {
// 	for i := len(chars) - 1; i > 0; i-- {
// 		if chars[i] == ' ' && chars[i-1] == ' ' {
// 			chars = append(chars[:i], chars[i+1:]...)
// 		}
// 	}

// 	if len(chars) > 0 && chars[len(chars)-1] == ' ' {
// 		chars = chars[:len(chars)-1]
// 	}

// 	return chars
// }

func TrimEdges(s []rune) []rune {
	start := 0
	end := len(s) - 1
	punctuation := s[0]
	startFlag := false
	endFlag := false
	
	

	for start <= end {
	    c := s[start]
	    if punctuation == '\''{
	        startFlag = true
	    }
		
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' && !startFlag {
			start++
		} else {
			break
		}
	}

	for end >= start {
		c := s[end]
		
	    if punctuation == '\''{
	        startFlag = true
	    }
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' && !endFlag {
			end--
		} else {
			break
		}
	}

	if start > end {
		return nil
	}

	return s[start : end+1]

}