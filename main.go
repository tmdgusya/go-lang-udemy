package main

import "fmt"

func main() {
	// create a map no value inside of it
	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	fmt.Println(colors["white"])
	delete(colors, "white")
	fmt.Println(colors) // map[]
}
