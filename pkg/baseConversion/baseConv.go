package go_reloaded

import(
	"strconv"
)


func HexTodec(splitted []string) []string {

    for i := 0; i < len(splitted); i++ {
        if splitted[i] == "(hex)" && i > 0 {
            val, err := strconv.ParseInt(splitted[i-1], 16, 64)
            if err == nil {
                splitted[i-1] = strconv.FormatInt(val, 10)
            }
            splitted = append(splitted[:i], splitted[i+1:]...)
            i--
        }
    }

    return splitted

}


func BinaryToDecimal(bin []string) []string {
	
	for i := 0; i < len(bin); i++ {

		if bin[i] == "(bin)" {

			bins, err := strconv.ParseInt(bin[i-1], 2, 64)
			if err == nil {
				bin[i-1] = strconv.FormatInt(bins, 10)
			}

			bin = append(bin[:i], bin[i+1:]...)
			i--
		}

	}
	return bin
}