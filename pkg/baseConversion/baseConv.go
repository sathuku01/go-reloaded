package go_reloaded

import(
	"strconv"
	"log"
	// "fmt"
)

func HexTodec(strHex string) string {
	hex, err := strconv.ParseInt(strHex, 16, 64)

	if err != nil {
		log.Fatal("Error converting hex to dec") 
		
	}
	
	return strconv.FormatInt(hex, 10)
}

// converts a binary string to decimal.
func BinaryToDecimal(bin string) string {
	dec := 0
	for i := 0; i < len(bin); i++ {
		dec = dec*2 + int(bin[i]-'0')

	}
	return strconv.Itoa(dec)
}