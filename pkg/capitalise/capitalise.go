package capitalise

import (
	"strconv"
	"strings"
	"unicode"
)

func ApplyTransformations(tokens []string) []string {
	operations := map[string]func(string) string{
		"cap": capitalizeSingle,
		"up":  uppercaseSingle,
		"low": lowercaseSingle,
	}

	index := 0
	for index < len(tokens) {
		if !strings.Contains(tokens[index], "(") {
			index++
			continue
		}

		startIndex := index
		rawCommand := tokens[index]
		endIndex := index

		for !strings.Contains(tokens[endIndex], ")") {
			endIndex++
			if endIndex >= len(tokens) {
				break
			}
			rawCommand += tokens[endIndex]
		}

		processedCommand := rawCommand

		closingBracketPos := strings.Index(processedCommand, ")")
		trailingPunctuation := ""
		if closingBracketPos != -1 && closingBracketPos+1 < len(processedCommand) {
			trailingPunctuation = processedCommand[closingBracketPos+1:]
			processedCommand = processedCommand[:closingBracketPos+1]
		}

		processedCommand = strings.TrimPrefix(processedCommand, "(")
		processedCommand = strings.TrimSuffix(processedCommand, ")")

		commandParts := strings.Split(processedCommand, ",")

		actionKey := strings.TrimSpace(commandParts[0])
		transformFunc, exists := operations[actionKey]
		if !exists {
			index++
			continue
		}

		if len(commandParts) == 1 {
			if startIndex-1 >= 0 {
				tokens[startIndex-1] = transformFunc(tokens[startIndex-1])
			}
		}

		if len(commandParts) == 2 {
			lookbackCount, parseErr := strconv.Atoi(strings.TrimSpace(commandParts[1]))
			if parseErr == nil && lookbackCount > 0 {
				applyFrom := startIndex - lookbackCount
				if applyFrom < 0 {
					applyFrom = 0
				}
				for j := applyFrom; j < startIndex; j++ {
					tokens[j] = transformFunc(tokens[j])
				}
			}
		}

		if trailingPunctuation != "" {
			tokens[startIndex] = trailingPunctuation
			copy(tokens[startIndex+1:], tokens[endIndex+1:])
			tokens = tokens[:len(tokens)-(endIndex-startIndex)]
			index = startIndex + 1
		} else {
			copy(tokens[startIndex:], tokens[endIndex+1:])
			tokens = tokens[:len(tokens)-(endIndex-startIndex+1)]
			index = startIndex
		}
	}

	return tokens
}

func capitalizeSingle(input string) string {
	for i, character := range input {
		if unicode.IsLetter(character) {
			return input[:i] + strings.ToUpper(string(character)) + strings.ToLower(input[i+1:])
		}
	}
	return input
}

func uppercaseSingle(input string) string {
	return strings.ToUpper(input)
}

func lowercaseSingle(input string) string {
	return strings.ToLower(input)
}
