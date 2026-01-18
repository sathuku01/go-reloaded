package fixquotes

import (
	"strings"
)

func AlignQuotes(words []string) []string {
	if len(words) == 0 {
		return words
	}
	
	text := strings.Join(words, " ")
	runes := []rune(text)
	result := make([]rune, 0, len(runes))
	inQuotes := false
	
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		
		if r == '\'' {
			if !inQuotes {
				inQuotes = true
				if len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, r)
			} else {
				inQuotes = false
				if len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, r)
				if i+1 < len(runes) {
					next := runes[i+1]
					if (next >= 'a' && next <= 'z') || (next >= 'A' && next <= 'Z') || next == '\'' {
						result = append(result, ' ')
					}
				}
			}
		} else if r == ',' && inQuotes {
			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}
			result = append(result, r)
		} else if r == ':' && i+1 < len(runes) && runes[i+1] == ' ' {
			result = append(result, r, ' ')
			i++ 
		} else {
			result = append(result, r)
		}
	}
	
	finalText := string(result)
	
	finalText = strings.ReplaceAll(finalText, " '", "'")
	finalText = strings.ReplaceAll(finalText, "' ", " '")
	
	return strings.Fields(finalText)
}