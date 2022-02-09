package nineteen

func PrintFrame(strings []string) {
	var maxLen int
	for _, a := range strings {
		if maxLen < len(a) {
			maxLen = len(a)
		}
	}
	for i := 0; i < maxLen+4; i++ {
		print("*")
	}
	println()
	for i := 0; i < len(strings); i++ {
		print("* ")
		print(strings[i])
		for j := 0; j < maxLen-len(strings[i]); j++ {
			print(" ")
		}
		println(" *")
	}
	for i := 0; i < maxLen+4; i++ {
		print("*")
	}
	println()
}
