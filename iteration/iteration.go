package iteration

// takes in a string and repeats it 5 times
func Repeat(char string, iterations int) (repeat string) {

	var repeated string

	for i := 0; i < iterations; i++ {
		repeated += char
	}

	return repeated

}
