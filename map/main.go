package main

import "fmt"

func main(){
	colors := map[string]string{
		"red" : "#ff0000",
		"green" : "#4bf745",
		"white" : "#fffff",
	}
	printMap(colors)
}

func printMap(c map[string]string){
	for color, hex := range c {
		fmt.Println("Hex color for", color, "is", hex)
	}
}