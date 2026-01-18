package process

import (
    
    "strings"
    aq "go_reloaded/pkg/alignquotes"
    ac "go_reloaded/pkg/article"
    hc "go_reloaded/pkg/baseConversion"
    sc "go_reloaded/pkg/capitalise"
    pc "go_reloaded/pkg/punctuation"
)

func ProcessData(text []string) []string {
	finalStuff := []string{}
  for _, item := range text {
        if strings.TrimSpace(item) == "" {
            finalStuff = append(finalStuff, "")
            continue
        }

        bits := strings.Fields(item)

        bits = hc.HexTodec(bits)
        bits = hc.BinaryToDecimal(bits)
        bits = sc.ApplyTransformations(bits)
        bits = ac.CheckForArticle(bits)
        bits = pc.FormatPunctuation(bits)
        bits = aq.AlignQuotes(bits)

        finalStuff = append(finalStuff, strings.Join(bits, " "))
    }
	return finalStuff
}