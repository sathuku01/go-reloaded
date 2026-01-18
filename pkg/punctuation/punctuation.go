package punctuate

import (
	"strings"
	"unicode"
)

func FormatPunctuation(words []string) []string {
	punct := "!.,?:;"
	result := make([]string, 0, len(words))

	for _, word := range words {
		if len(word) == 1 && strings.ContainsAny(word, punct) {
			if len(result) > 0 {
				result[len(result)-1] += word
			} else {
				result = append(result, word)
			}
		} else {
			result = append(result, word)
		}
	}

	text := strings.Join(result, " ")

	for _, p := range punct {
		char := string(p)
		text = strings.ReplaceAll(text, " "+char, char)
	}

	text = addSpacesAfterPunctuation(text)

	text = handlePunctuationGroups(text)

	text = smartSentenceCapitalization(text)

	return strings.Fields(text)
}

func addSpacesAfterPunctuation(text string) string {
	runes := []rune(text)
	result := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		result = append(result, r)

		if strings.ContainsRune(".,!?:;", r) {
			if i+1 < len(runes) {
				next := runes[i+1]

				if unicode.IsLetter(next) && next != '\'' &&
					!strings.ContainsRune(".,!?:;", next) {
					result = append(result, ' ')
				}
			}
		}
	}

	return string(result)
}

func handlePunctuationGroups(text string) string {
	runes := []rune(text)
	result := make([]rune, 0, len(runes))

	i := 0
	for i < len(runes) {
		r := runes[i]
		result = append(result, r)

		if strings.ContainsRune(".,!?:;", r) {
			j := i + 1
			for j < len(runes) && strings.ContainsRune(".,!?:;", runes[j]) {
				result = append(result, runes[j])
				j++
			}
			if j > i+1 {
				i = j - 1
			}
		}

		i++
	}

	return string(result)
}

func smartSentenceCapitalization(text string) string {
	runes := []rune(text)
	capitalizeNext := true

	for i, r := range runes {
		if capitalizeNext && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			capitalizeNext = false
		}

		if r == '.' || r == '!' || r == '?' {
			isEndOfSentence := false

			if i+1 < len(runes) {
				next := runes[i+1]
				if next == ' ' || i+1 == len(runes)-1 {
					isEndOfSentence = true
				}
			} else {
				isEndOfSentence = true
			}

			if isEndOfSentence {
				capitalizeNext = true
			}
		}

		if r == ':' || (i > 0 && runes[i-1] == '.' && i > 1 && runes[i-2] == '.' && i > 2 && runes[i-3] == '.') {
			capitalizeNext = false
		}
	}

	return string(runes)
}
