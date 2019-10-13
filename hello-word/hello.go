package main

import "fmt"

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	const englishHelloPrefix = "Hello, "
	const brasilHelloPrefix = "Olá, "
	const frenchHelloPrefix = "Bonjour, "

	switch language {
	case "Brazil":
		prefix = brasilHelloPrefix
	case "French": 
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main()  {
	fmt.Println(Hello("", ""))
}