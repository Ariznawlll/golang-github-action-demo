package main

import "fmt"

func Cat() string {
	return "Wang~~~"
}
func main() {
	saying := Cat()
	fmt.Print(saying)
}
