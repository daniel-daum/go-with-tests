package helloworld

import "fmt"

func Hello() string {
	return "Hello, world!"
}
<<<<<<< Updated upstream:hello.go
=======

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return prefix
}

>>>>>>> Stashed changes:hello_world/hello.go
func main() {
	fmt.Println(Hello())
}