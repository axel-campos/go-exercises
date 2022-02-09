package twenty

func PigLatinToNormal(text string) string {

	return text
}

func NormalToPigLatin(text string) string {
	text += " "
	var pig_latin_text string
	var length int = 0
	for i, char := range text {
		if char == 32 && length != 0 {
			pig_latin_text += text[i-length+1 : i]
			pig_latin_text += string(text[i-length])
			pig_latin_text += "ay "
			length = 0
		} else if char == 32 && length == 0 {
			pig_latin_text += " "
		} else {
			length++
		}

	}
	return pig_latin_text
}
