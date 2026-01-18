package go_reloaded

func CheckForArticle(s []string) []string{
	if len(s) < 2 {
		return s
	}
	
	mapArticles := map[string]struct{}{
	"a": {}, "e": {}, "i": {}, "o": {}, "u": {}, "h": {},
	"A": {}, "E": {}, "I": {}, "O": {}, "U": {}, "H": {},
}
		for i, char := range s {
			if (char == "a" || char == "A" && s[i+1] == " ") {
				if _, exists := mapArticles[s[i+2]]; exists {
				char = "an"
			}
		}
	}
	
	return s
}