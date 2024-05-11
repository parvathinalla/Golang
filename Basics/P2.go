package main

import (
	"fmt"
	"time"
)

func printOutDate() {
	t := time.Now()
	fmt.Println(t)
}

func waitForIt(message string) {
	defer fmt.Println("Done!!!")
	fmt.Println("Waiting")
	fmt.Println("waiting")
	fmt.Println("waiting")
	fmt.Println(message)
}

func joinTwoStrings(prefix string, suffix string) string {
	return prefix + suffix
}

func main() {
	printOutDate()
	waitForIt(joinTwoStrings("Hi", " there "))
}
