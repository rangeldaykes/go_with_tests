package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello() string {
	return "Hello, world"
}

func HelloWithName(name string) string {

	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func HelloWithNameAndLanguage(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

//func main() {
//	fmt.Println(Hello())
//}
