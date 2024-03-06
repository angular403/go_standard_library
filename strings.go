package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Andrew Wiliam", "Andrew"))
	fmt.Println(strings.Split("Andrew Wiliam", " "))
	fmt.Println(strings.ToLower("Andrew Wiliam"))
	fmt.Println(strings.ToUpper("Andrew Wiliam"))
	fmt.Println(strings.Trim("              Andrew Wiliam         ", " "))
	fmt.Println(strings.ReplaceAll("Andrew Wiliam Andrew Wiliam", "Andrew", "Budi"))
}
