package go_reloaded


func CheckForArticle(s []string) []string {
	if len(s) < 2 {
		return s
	}

	mapArticles := map[string]bool{
		"a": true, "e": true, "i": true, "o": true, "u": true, "h": true,
		"A": true, "E": true, "I": true, "O": true, "U": true, "H": true,
	}

	for i, char := range s {
		if char == "a" || char == "A" {
			if _, exists := mapArticles[string(s[i+1][0])]; exists {
				s[i] = "an"
			}
		}
	}
	return s
}
