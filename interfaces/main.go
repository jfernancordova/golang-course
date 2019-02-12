package main

import "fmt"

//Custom Type
type englishBot struct {}
type spanishBot struct {}

func main(){
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(sb spanishBot){
	fmt.Println(sb.getGreeting())
}

func printGreeting(eb englishBot){
	fmt.Println(eb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}