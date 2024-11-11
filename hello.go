package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	helloEnglishPrefix = "Hello, "
	helloSpanishPrefix = "Hola, "
	helloFrenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world!"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return prefix 
}

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}
