package iteration

const defaultRepCnt = 5

// Repeat takes a character and repeats it 5 times
func Repeat(char string, rep int) string {
	var repeated string
	if rep < 0 {
		rep = defaultRepCnt
	}

	for i := 0; i < rep; i++ {
		repeated += char
	}
	return repeated
}