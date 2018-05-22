package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#ff0000",
		"yellow": "#FFFF00",
		"blue":   "#0000FF",
	}

	fmt.Println(colors)

	colors["green"] = "#00ff00"
	delete(colors, "blue")

	print(colors)
}

func print(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)
	}
}
