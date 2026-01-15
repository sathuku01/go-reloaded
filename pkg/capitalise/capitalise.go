 package capitalise
 
 import(
	"strconv"

 )

 func CapsFirstLetter(s string) string {
	sByte := []byte(s)
	for i, char := range sByte{
		if i == 0 && (char >= 65 && char <= 90){
			continue
		} else if i == 0 && (char >= 97 && char <= 122) {
			sByte[i] = char - 32
		} else if char >= 65 && char <= 90 && i > 0{
			sByte[i] = char + 32
		} 
	}
	return string(sByte)
 }

func ToUpper(s string) string{
	sByte := []byte(s)

	convertedString := ""
	for _, char := range sByte{
		if (char >= 97 && char <= 122){
			convertedString += string(char - 32)
		} else {
			convertedString += string(char)
		}
	}
	return convertedString
}

func ToLower (s string) string{
sByte := []byte(s)

	convertedString := ""
	for _, char := range sByte{
		if (char >= 65 && char <= 90){
			convertedString += string(char + 32)
		} else {
			convertedString += string(char)
		}
	}
	return convertedString

}

 func ExrtractNumFromString(numString string) int {
	numStringByte := []byte(numString)
	extracted := ""
	for _, num := range numStringByte{
		if num >= 48 && num <= 57{
			extracted += string(num)
		}
	}
	intNum, err := strconv.Atoi(extracted) 
	if err != nil {
		return -1
	}
	return intNum
 }