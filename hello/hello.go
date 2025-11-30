package hello

const englishHelloPrefix = "Hello, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
