package go_reloaded

func CheckForArticle(s string) string{
	mapArticles := map[byte]struct{}{
			'a' : {}, 'e' : {},	'i' : {}, 'o' : {}, 'u' : {}, 'h' : {},
			'A' : {}, 'E' : {}, 'I' : {}, 'O' : {}, 'U' : {}, 'H' : {},
			}
			if _, exists := mapArticles[s[0]]; exists {
				s = "an"
			}
			return s
}